package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	w, count := CountingWriter(os.Stdout)

    fmt.Fprintln(w, "Hello, world!") // Write 14 bytes
    fmt.Fprintln(w, "Another line!") // Write 14 more bytes

    fmt.Println("Total bytes written:", *count)
}

type countingWriter struct {
	writer io.Writer
	count *int64
}

func (cw *countingWriter) Write(p []byte) (int, error) {
	n, err := cw.writer.Write(p)
	*cw.count += int64(n)

	return n, err
}



// func CountingWriter(w io.Writer) (io.Writer, *int64) {
// 	c := countingWriter(w, )
// }

func CountingWriter(w io.Writer) (io.Writer, *int64) {
    var count int64
    return &countingWriter{writer: w, count: &count}, &count
}


