package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
)

// Writes file to hard drive
func writeNewFile(newFileName string, newFileBytes []byte) {
	err := ioutil.WriteFile(newFileName, newFileBytes, 0777)
	if err != nil {
		panic(err)
	}
}

// Reads file from hard drive
func readFile(fileName string) []byte {
	imageFileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	return imageFileBytes
}

func main() {
	imageName := "Screenshots/screenshot.png"
	imageFileBytes := readFile(imageName)

	fmt.Println("File Size:")
	fmt.Println(strconv.Itoa(len(imageFileBytes)) + " bytes")

	// Encode []byte -> base64
	encoded := base64.StdEncoding.EncodeToString(imageFileBytes)
	fmt.Println(len(encoded))

	decoded, _ := base64.StdEncoding.DecodeString(encoded)
	fmt.Println(reflect.TypeOf(decoded))
	fmt.Println(len([]byte(decoded)))

	fmt.Println("Writing new file...")
	writeNewFile("newscreenshot.png", decoded)
}