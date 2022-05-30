package ch03

import (
	"net"
	"testing"
)

func TestListener(t *testing.T) {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = listener.Close() }()
	t.Logf("bound to %q", listener.Addr())

}
