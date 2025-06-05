package main

// import (
// 	"flag"
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	// Determine the initial directories.
// 	flag.Parse()
// 	roots := flag.Args()
// 	if len(roots) == 0 {
// 		roots = []string{"."}
// 	}
// 	// Traverse the file tree.
// 	fileSizes := make(chan int64)
// 	go func() {
// 		for _, root := range roots {
// 			walkDir(root, fileSizes)
// 		}
// 		close(fileSizes)
// 	}()
// 	// Print the results.
// 	var nfiles, nbytes int64
// 	for size := range fileSizes {
// 		nfiles++
// 		nbytes += size
// 	}
// 	printDiskUsage(nfiles, nbytes)
// }
// func printDiskUsage(nfiles, nbytes int64) {
// 	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
// }

// func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
// 	defer n.Done()
// 	for _, entry := range dirents(dir) {
// 		if entry.IsDir() {
// 			n.Add(1)
// 			subdir := filepath.Join(dir, entry.Name())
// 			go walkDir(subdir, n, fileSizes)
// 		} else {
// 			fileSizes <- entry.Size()
// 		}
// 	}
// }
