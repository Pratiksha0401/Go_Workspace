package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error")
		os.Exit(1)
	}

	//fmt.Println(resp)
	//to read body by byte size
	/*s := make([]byte , 99999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))
	*/
	//use of writer and reader to get body
	//io.Copy(os.Stdout , resp.Body)

	//custom writer
	lw := logWriter{}
	io.Copy(lw, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
