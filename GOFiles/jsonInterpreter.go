package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	d := json.NewDecoder(os.Stdin)				//accepts from standard input
	var err error 								//deff error
	for err == nil {
		var v interface {}						//deff interface
		if err = d.Decode(&v); err != nil { 		//witchcraft...
			break
		}

		var b [] byte

		if b, err = json.MarshalIndent(v, "", " "); err != nil {
			break
		}
		_, err = os.Stdout.Write(b)
	}
	if err != io.EOF {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
}
