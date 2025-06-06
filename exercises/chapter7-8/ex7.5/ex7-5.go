package main

import (
	"io"
	"strings"
	"fmt"
)

func main() {
  // Create a data source
  data := "Hello, Gopher! This is a test string."
  src := strings.NewReader(data) // turns string into an io.Reader

  // Wrap it in a LimitReader with a 12-byte limit
  limited := LimitReader(src, 12)

  // read in 5-byte chunks
  buf := make([]byte, 5) 
  for {
	  n, err := limited.Read(buf)
	  if err == io.EOF {
		  fmt.Println("\nReached limit / end of data.")
		  break
	  }
	  if err != nil {
		  fmt.Println("Read error:", err)
		  break
	  }

	  fmt.Print(string(buf[:n]))
  }
}

type limitReader struct {
	reader io.Reader
	numberOfBytes int64
}


func (lr *limitReader) Read(p []byte) (int, error) {
	if lr.numberOfBytes <= 0 {
		return 0, io.EOF
	}

	if int64(len(p)) > lr.numberOfBytes {
		p = p[:lr.numberOfBytes]
	}

    n, err := lr.reader.Read(p)
    lr.numberOfBytes -= int64(n)
    return n, err
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitReader{r, n}
}

// func LimitReader(r io.Reader, n int64) io.Reader {
// 	buf := make([]byte, n)

// 	for {
// 		res, err := r.Read(buf)

// 		if err == io.EOF {
// 			fmt.Println("End of file!!!")
// 			break
// 		}

// 		if err != nil {
// 			break
// 		}

// 		fmt.Println(res)
// 	}
		
// 	// if err != nil {
// 	// 	return nil, err
// 	// }

// 	return nil
// }