package goid

import (
	"runtime"
	"bytes"
	"strconv"
)

// ID get an uniq id, not as the same goroutine id
func ID() int64 {
	return int64(uintptr(G()))
}

// GoID get goroutine id, this method is very slow
func IDSlow() int64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseInt(string(b), 10, 64)
	return n
}
