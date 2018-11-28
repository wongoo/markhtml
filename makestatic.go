//+build dev

package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"unicode/utf8"
)

var files = []string{
	"index.min.html",
	"markhtml.min.js",
}

func main() {
	if err := makestatic(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func makestatic() error {
	f, err := os.Create("static.go")
	if err != nil {
		return err
	}
	defer f.Close()
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "package main\n\n")
	fmt.Fprintf(buf, "var Files = map[string]string{\n")
	for _, fn := range files {
		b, err := ioutil.ReadFile(fn)
		if err != nil {
			return err
		}
		fmt.Fprintf(buf, "\t%q: ", fn)
		if utf8.Valid(b) {
			fmt.Fprintf(buf, "`%s`", sanitize(b))
		} else {
			fmt.Fprintf(buf, "%q", b)
		}
		fmt.Fprintln(buf, ",\n")
	}
	fmt.Fprintln(buf, "}")
	fmtbuf, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}
	return ioutil.WriteFile("static.go", fmtbuf, 0666)
}

// sanitize prepares a valid UTF-8 string as a raw string constant.
func sanitize(b []byte) []byte {
	// Replace ` with `+"`"+`
	b = bytes.Replace(b, []byte("`"), []byte("`+\"`\"+`"), -1)

	// Replace BOM with `+"\xEF\xBB\xBF"+`
	// (A BOM is valid UTF-8 but not permitted in Go source files.
	// I wouldn't bother handling this, but for some insane reason
	// jquery.js has a BOM somewhere in the middle.)
	return bytes.Replace(b, []byte("\xEF\xBB\xBF"), []byte("`+\"\\xEF\\xBB\\xBF\"+`"), -1)
}
