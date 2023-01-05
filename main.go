package main

import (
	"fmt"
)

func main() {
	// for update
	fileList := [2]string{"steps_demo.xlsx", "README.md"}
	fileUrl := "https://raw.githubusercontent.com/shouaya/sg/main/"
	
	for i, v := range fileList {
		fmt.Println(i, v)
		err := DownloadFile( v, fileUrl + v)
		if err != nil {
			panic(err)
		}
		fmt.Println("Downloaded: " + fileUrl + v)
    }
	// run script
	run()
}
