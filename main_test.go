package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestCluserWideScope(t *testing.T) {
	os.Args = []string{"./seal-cli", "-h", "http://host", "-s", "MySecret", "-dry-run"}

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	w.Close()
	os.Stdout = old
	out := <-outC
	if !strings.Contains(out, "cluster-wide") {
		t.Errorf("expected kubeseal CMD to be called with cluster-wide scope, but got %s", out)
	}
}

// func TestNamespaceWideScope(t *testing.T) {
// 	os.Args = []string{"./seal-cli", "-h", "http://host", "-s", "MySecret", "-n", "myspace", "-dry-run"}

// 	old := os.Stdout
// 	r, w, _ := os.Pipe()
// 	os.Stdout = w

// 	main()

// 	outC := make(chan string)
// 	go func() {
// 		var buf bytes.Buffer
// 		io.Copy(&buf, r)
// 		outC <- buf.String()
// 	}()

// 	w.Close()
// 	os.Stdout = old
// 	out := <-outC
// 	if !strings.Contains(out, "namespace-wide") {
// 		t.Errorf("expected kubeseal CMD to be called with namespace-wide scope, but got %s", out)
// 	}
// }
