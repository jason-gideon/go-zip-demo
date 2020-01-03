package czip

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
)

func Compress(src []byte) ([]byte, error) {
	var buf bytes.Buffer    //bytesbuffer is a struct
	w := zlib.NewWriter(&buf)

	defer w.Close()
	l, err := w.Write(src)
	if err != nil {
		return nil, err
	}

	fmt.Println(buf.Bytes())
	fmt.Println(len(buf.Bytes()))
	fmt.Println(len(src), l)

	return buf.Bytes(), nil
}

func Decompress(src []byte) ([]byte, error){

	buf := bytes.NewBuffer(src)
	r, err :=zlib.NewReader(buf)
	if err != nil {
		fmt.Println("jie压缩失败")
		return nil, err
	}

	io.Copy(buf, r)

	return buf.Bytes(), nil
	//res:=bytes.Compare(src, output.Bytes())
	//if res == 0 {
	//	fmt.Println("!..Slices are equal..!")
	//} else {
	//	fmt.Println("!..Slice are not equal..!")
	//}
}