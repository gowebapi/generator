// Go WASM output
package gowasm

import (
	"bytes"
	"fmt"
	"go/format"
	"go/parser"
	"go/token"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"github.com/gowebapi/webidlgenerator/types"
)

const fileTemplInput = `
{{define "header"}}
// Code generated by webidlgenerator. DO NOT EDIT.

package {{.Package}}

import "syscall/js"

// @IMPORT@

// ReleasableApiResource is used to release underlaying
// allocated resources.
type ReleasableApiResource interface {
	Release()
}

type releasableApiResourceList []ReleasableApiResource

func (a releasableApiResourceList) Release() {
	for _, v := range a {
		v.Release()
	}
}

// workaround for compiler error
func unused(value interface{}) {
	// TODO remove this method
}

type Union struct {
	Value js.Value
}

func (u *Union) JSValue() js.Value {
	return u.Value
}

func UnionFromJS(value js.Value) * Union {
	return & Union{Value: value}
}

{{end}}
`

var fileTempl = template.Must(template.New("file").Parse(fileTemplInput))

// Data in header evaluation
type fileData struct {
	Package string
}

type writeFn func(dst io.Writer, in types.Type) error

type Source struct {
	Package string
	name    string
	Content []byte
}

var reservedGoKeywords = map[string]bool{
	"make":   true,
	"nil":    true,
	"int":    true,
	"uint":   true,
	"string": true,
	"len":    true,
	"iota":   true,
}

// WriteSource is create source code files.
// returns map["path/filename"]"file content"
func WriteSource(conv *types.Convert) ([]*Source, error) {
	oldTB := types.TransformBasic
	restoreTB := func() { types.TransformBasic = oldTB }
	defer restoreTB()
	types.TransformBasic = pkgMgr.transformPackageName
	target := make(map[string]*bytes.Buffer)
	var err error
	for _, e := range conv.Enums {
		if e.InUse() {
			err = writeType(e, target, writeEnum, err)
		}
	}
	for _, v := range conv.Callbacks {
		if v.InUse() {
			err = writeType(v, target, writeCallback, err)
		}
	}
	for _, v := range conv.Dictionary {
		if v.InUse() {
			err = writeType(v, target, writeDictionary, err)
		}
	}
	for _, v := range conv.Interface {
		if v.InUse() {
			err = writeType(v, target, writeInterface, err)
		}
	}
	if err != nil {
		return nil, err
	}
	ret := make([]*Source, 0)
	for pkg, buffer := range target {
		content := buffer.Bytes()
		content = sourceCodeRemoveEmptyLines(content)
		if content, err = insertImportLines(pkg, content); err != nil {
			fmt.Fprintf(os.Stderr, "error:%s:unable to remove unused imports from source code: %s\n", pkg, err)

		}
		if source, err := format.Source(content); err == nil {
			content = source
		} else {
			// we just print this error to get an output file that we
			// later can correct and fix the bug
			fmt.Fprintf(os.Stderr, "error:%s:unable to format output source code: %s\n", pkg, err)
		}
		wasm, desktop := createMultieOSLib(content)
		base := strings.ToLower(shortPackageName(pkg))
		ret = append(ret, &Source{
			Package: pkg,
			name:    base + ".go",
			Content: desktop,
		})
		ret = append(ret, &Source{
			Package: pkg,
			name:    base + "_js_wasm.go",
			Content: wasm,
		})
	}
	sort.Slice(ret, func(i, j int) bool {
		if ret[i].Package == ret[j].Package {
			return ret[i].name < ret[j].name
		}
		return ret[i].Package < ret[j].Package
	})
	return ret, nil
}

func writeType(value types.Type, target map[string]*bytes.Buffer, conv writeFn, err error) error {
	if err != nil {
		return err
	}
	pkgMgr.setPackageName(value)
	dst, err := getTarget(value, target)
	if err != nil {
		return err
	}
	if err := conv(dst, value); err != nil {
		return err
	}
	return nil
}

func getTarget(value types.Type, target map[string]*bytes.Buffer) (*bytes.Buffer, error) {
	pkg := value.Basic().Package
	dst, ok := target[pkg]
	if ok {
		return dst, nil
	}
	dst = &bytes.Buffer{}
	target[pkg] = dst
	data := fileData{
		Package: shortPackageName(pkg),
	}
	if err := fileTempl.ExecuteTemplate(dst, "header", data); err != nil {
		return nil, err
	}
	return dst, nil
}

// sourceCodeRemoveEmptyLines will remove empty lines
func sourceCodeRemoveEmptyLines(code []byte) []byte {
	add := []string{"//", "func", "type", "const", "var"}
	in := bytes.NewBuffer(code)
	var out bytes.Buffer
	ignore := false
	for {
		s, err := in.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		if err == io.EOF {
			break
		}
		if len(strings.TrimSpace(s)) == 0 {
			continue
		}
		if strings.HasPrefix(s, "package") {
			out.WriteByte('\n')
		}
		found := false
		for _, prefix := range add {
			if strings.HasPrefix(s, prefix) {
				found = true
				if !ignore {
					out.WriteByte('\n')
				}
				ignore = true
			}
		}
		if !found {
			ignore = false
		}
		out.WriteString(s)
	}
	return out.Bytes()
}

func createMultieOSLib(content []byte) (wasm, others []byte) {
	oldImport := []byte("import \"syscall/js\"")
	newImport := []byte("import js \"github.com/gowebapi/webapi/core/failjs\"")
	oldTag := []byte("package")
	newTag := []byte("// +build !js\n\npackage")
	wasm = content
	others = bytes.Replace(content, oldImport, newImport, 1)
	others = bytes.Replace(others, oldTag, newTag, 1)
	return
}

func insertImportLines(pkg string, content []byte) ([]byte, error) {
	// first we extract all unresolved symbols to find out
	// what import lines that is actually is used. there is
	// logic bug in the current import type detector. for
	// some types that is changes to by the type system to
	// a more generic, like callback -> js.Func, generate
	// an import line to that callback type. this is very
	// hard to write a proper detector as we need to do
	// template inspection if returned TypeInfo.Input is used
	// etc. the code below is simply trying to figure out
	// what imports are really used and just include those.
	removeImports := true
	names := make(map[string]struct{})
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", content, parser.DeclarationErrors)
	if err == nil {
		for _, unres := range f.Unresolved {
			if _, f := reservedGoKeywords[unres.Name]; !f {
				names[unres.Name] = struct{}{}
			}
		}
	} else {
		removeImports = false
	}

	// then we are writing the import lines
	file := pkgMgr.packages[pkg]
	lines := file.importLines(names, removeImports)
	lines = lines + "\n" + file.importInfo()
	content = bytes.Replace(content, []byte("// @IMPORT@"), []byte(lines), 1)
	return content, err
}

func (src *Source) Filename(insidePkg string) (string, bool) {
	full := filepath.Join(src.Package, src.name)
	if insidePkg == "" {
		return full, true
	}
	limit := insidePkg
	if !strings.HasSuffix(limit, "/") {
		limit = limit + "/"
	} else {
		insidePkg = insidePkg[0 : len(insidePkg)-1]
	}
	if src.Package == insidePkg {
		return src.name, true
	}
	if strings.HasPrefix(src.Package, limit) {
		return full[len(limit):], true
	}
	return full, false
}
