package hackerrank

import (
	"bytes"
	"io"
	"testing"
)

func Checker(solver func(io.Reader, io.Writer), t *testing.T, src_, dst_ string) {
	var out_ bytes.Buffer

	src := bytes.TrimSpace([]byte(src_))
	dst := bytes.TrimSpace([]byte(dst_))

	solver(bytes.NewBuffer(src), &out_)
	out := bytes.TrimSpace(out_.Bytes())

	if !bytes.Equal(dst, out) {
		t.Fatalf("Input:\n%s\nExp:\n[%s]\nGot:\n[%s]\n", src, dst, out)
	}
}
