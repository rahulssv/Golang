package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)
type customReader struct{
}
// type customWriter struct{

// }
func main() {
	cr:= customReader{}
	bs, err := ioutil.ReadFile(os.Args[1])
	if err!=nil{
		fmt.Println(err)
	}
	cr.Read(bs)
	io.Copy(os.Stdout,cr)
}

// func (customWriter) Write(p []byte) (n int, err error){
// 	fmt.Println(string(p))
// 	return len(p),nil
// }
func (customReader) Read(p []byte) (n int, err error){
	fmt.Println(string(p))
	return len(p),nil
}