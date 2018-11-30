package main

import (
	"io"
	"io/ioutil"

	"github.com/lapsang-boys/galvin/typeless/ast"
	"github.com/pkg/errors"
)

// ParseFile parses the given LLVM IR assembly file into an LLVM IR module.
func ParseFile(path string) (*ast.File, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	content := string(buf)
	return ParseString(path, content)
}

// Parse parses the given LLVM IR assembly file into an LLVM IR module, reading
// from r. An optional path to the source file may be specified for error
// reporting.
func Parse(path string, r io.Reader) (*ast.File, error) {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	content := string(buf)
	return ParseString(path, content)
}

// ParseBytes parses the given LLVM IR assembly file into an LLVM IR module,
// reading from b. An optional path to the source file may be specified for
// error reporting.
func ParseBytes(path string, b []byte) (*ast.File, error) {
	content := string(b)
	return ParseString(path, content)
}

// ParseString parses the given LLVM IR assembly file into an LLVM IR module,
// reading from content. An optional path to the source file may be specified
// for error reporting.
func ParseString(path, content string) (*ast.File, error) {
	tree, err := ast.Parse(path, content)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to parse %q into AST", path)
	}
	root := ast.ToTypelessNode(tree.Root())
	return root.(*ast.File), nil
}
