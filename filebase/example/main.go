package main

import (
	_ "fmt"
	"xyuTools/filebase"
)

func main() {

	s1 := []byte{0x31, 0x32}
	s2 := []byte{0x33, 0x34}

	filebase.AppendDataByte("./data.txt", s1)
	filebase.AppendDataByte("./data.txt", s2)
}
