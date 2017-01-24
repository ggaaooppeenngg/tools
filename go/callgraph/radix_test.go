package callgraph

import (
	"testing"
)

func TestRadix(t *testing.T) {
	var r = RadixTree{Root: new(RadixNode)}
	r.Add("hash")
	r.Add("hash/crc32")
	r.Add("hash/crc64")
	r.Add("html")
	r.Add("html/template")
	r.Add("io")
	r.Add("io/ioutil")
	if l := r.LongestPrefix("html/template"); l != "html/template" {
		t.Fatalf("failed get %s\n", l)
	}
	if l := r.LongestPrefix("html/another"); l != "html" {
		t.Fatalf("failed get %s\n", l)
	}

}
