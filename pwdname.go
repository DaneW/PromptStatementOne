package main

import (
	"bytes"
	"fmt"
	"os"
	s "strings"
)

func main() {
	// If these env vars don't exist...
	// maybe you should consider fixing them before pimping out your bash profile
	home, _ := os.LookupEnv("HOME")
	pwd, _ := os.LookupEnv("PWD")
	formattedPwd := s.Replace(pwd, home, "", 1)
	dirs := s.SplitAfter(formattedPwd, "/")
	dirs = dirs[1:]
	length := len(dirs)
	switch {
	case length > 2:
		var output bytes.Buffer
		output.WriteString("~/")
		for idx, dir := range dirs {
			if idx > (length - 3) {
				output.WriteString(dir)
			} else {
				output.WriteString("../")
			}
		}
		fmt.Println(output.String())
	default:
		fmt.Println("~" + formattedPwd)
	}
}
