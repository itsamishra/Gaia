package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := "Screenshots/screenshot.png"

	buf := bytes.NewBuffer(nil)
	f, _ := os.Open(fileName)
	io.Copy(buf, f)
	f.Close()
	fmt.Println(buf.Bytes())

	// for _, filename := range filenames {
	// 	f, _ := os.Open(filename) // Error handling elided for brevity.
	// 	io.Copy(buf, f)           // Error handling elided for brevity.
	// 	f.Close()
	// }
	// s := string(buf.Bytes())
}
