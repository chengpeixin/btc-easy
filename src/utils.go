package src

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

func IntToByte(num int64)[]byte{
	var buffer bytes.Buffer
	err :=binary.Write(&buffer,binary.BigEndian,num)
	CheckErr("IntToByte",err)
	return buffer.Bytes()
}

func CheckErr(data string,err error){
	if err != nil{
		fmt.Println("err occur",err,"pos :",pos)
		os.Exit(1)
	}
}