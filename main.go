package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"os"
)


func Compress(src []byte) ([]byte, error) {
	var buf bytes.Buffer    //bytesbuffer is a struct
	w := zlib.NewWriter(&buf)

	l, err := w.Write(src)
	w.Close()    //dont use defer...The zlib.Writer isn't flushed until after the function exits
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


func test(){
	var input = []byte("data:image/png;base64,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")
	//buf := bytes.NewBuffer(input)
	c , err:=Compress(input)
	if err != nil {

	} else {
		//写文件，使用c++zip尝试解压
		fileObj,err := os.OpenFile("dem0.zip",os.O_RDWR|os.O_CREATE|os.O_TRUNC,0644)
		if err != nil {
			fmt.Println("Failed to open the file",err.Error())
			os.Exit(2)
		}
		defer fileObj.Close()
		if _,err := fileObj.Write(c);err == nil {
			fmt.Println("Successful writing to thr file with os.OpenFile and *File.Write method.",c)
		}


		//if _,err := fileObj.WriteString(content);err == nil {
		//	fmt.Println("Successful writing to the file with os.OpenFile and *File.WriteString method.",content)
		//}
		//contents := []byte(content)
		//if _,err := fileObj.Write(contents);err == nil {
		//	fmt.Println("Successful writing to thr file with os.OpenFile and *File.Write method.",content)
		//}
	}


	o , err1 := Decompress(c)
	if err1 != nil {

	}

	res:=bytes.Compare(input, o)
	if res == 0 {
		fmt.Println("!..Slices are equal..!")
	} else {
		fmt.Println("!..Slice are not equal..!")
	}



	_, err2 := Decompress(input)
	if err2 != nil {
		println("check decom error ok")
		
	}
}


func nromalcompress(){
	//byte array init
	var input = []byte("data:image/png;base64,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")
	var buf bytes.Buffer    //bytesbuffer is a struct
	w := zlib.NewWriter(&buf)
	w.Write(input)
	w.Close()
	fmt.Println(buf.Bytes())    //压缩到buf
	fmt.Println(len(buf.Bytes()))
	fmt.Println(len(input))

	//decompress
	b := buf.Bytes()
	z:=bytes.NewBuffer(b)  // []byte to bytes.buffer
	var output bytes.Buffer
	r, err1 :=zlib.NewReader(z)
	if err1 != nil {
		fmt.Println("jie压缩失败")
		return
	}

	//io
	io.Copy(&output, r)
	res:=bytes.Compare(input, output.Bytes())
	if res == 0 {
		fmt.Println("!..Slices are equal..!")
	} else {
		fmt.Println("!..Slice are not equal..!")
	}
}



func bestcompress(){
	//byte array init
	var input = []byte("data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAKsAAAAgCAYAAABtn4gCAAAI9klEQVR4Xu2cfYxcVRmHn3Pu3Nm6lXa2ta0AYlsgFSGgThMU0W5hGkqApNGARk3cGm3BRN3FYGuCISrBXaNphcSkG2I3hn9gY2gSNWq32AoE1C6KRihFti2QSgvsTmm783HvOa+zM69cSS8Tacsq7n2SX87cj5P7ZvLs2ztnbseICG8HMjIsbxMyMnIoxhja8fJnriuYMNdjOsIVNhd0Y0wBRbxH6nHZ16OdUot2iXNDC3/2mzJtOImOnpHJ2p4Xr13ZkDTYJLHrwRgw4L1gAptIFzt87AoSuzXi/BqJ3KZ/XL1iSGLXd9aOh8u81WRksh5cefmUeFuBgmBABHEeEziwBrUVvCCxayWaSow419MY17zw4eVrz3ls9zbeKjIyWV+4bPkmiVyvCtnqps6Bta2uahJZxQs4jziHxP7fxS1I7B947tJLN5/7xBN9tIeN37yNN0EXME6LecAEGfR/746ZJeuBSy9RUVsiGucxsYPAYuxUzOtlFQEniG8JS+xRaTW+d/+y97P46Sf7yMg4XbLuPW/ZGuqu1+cE4y0m8OAs0uyopjmqqArJrYDz4FVU1wrJ6959S5ftWjL29DZOlYxM1ut/eU3h7tEjW2/+yWE683kIPBJYTGzBGkwjaFdVXRFIZPXSktWJ3hbotheUrQ1hdzaELXMqZGSyCmzaV5xT+PNDR/jg3hqzcjmM9aCSijHQDLBwISxYAAjUD0PtUEtKmYoHWq9bUwQDYKRgZ8kmxljLyZKRyXrdz1cXgB5jDb/4xLtY8oPnmRc78tYiLdMwBrjicuxN6xg/YwGxWIwxzM4bOuUQPH8X9thvsYFvRJpj8NrYSOjIhdJTv/2dfUCZ9mwAbgCKJIwAw8Dgf3j+gGYi5YPZKLBcz+8HSrTYqHPQY1qHAqNawwAn0q/nrwLGdFvngta9UetZqueuI2FEj4+Shp6vc9FrDAIDM0pWxPRgaPLq4k5+V5zNqkeOQi5sCguCWb2aV756Kzv+XuXQAU8+pClrFDsWzZ7PVRcMcOaR2wmPPdCUMxc6gtC35LVCC4P3pgfYTDpdwHYVReVgTIUpqSgKqIhb9PhYIhEbEtFYnrJa0KXH7kdRxpL59Cf7GFZJ9A+CkkqZRkml6gIGdFynKamQW3T/sNatx9iu9Y6lvydMqJxJjVrLjJFVYIUh4fer53Px46+ycLLOnDDPrO6PM/61bzD48DgeQxgGRM5iAOeFvcfrPHu4wvorvsv5s1+mI/o1b4xZ0UbW+1WGEeDGlK6YJl1Jzx1GUSF2A0UVYSBlXr9eZ1WKbP0qxnoVSqFLayzpeGN6B2Q4OZbUo8Lfn3LdjSpkSeevR1GxiyrpxrQ5Wu/GmfFsgEg3CURnvoMdl5/B8Thmol5j1i293Dc6wZHJiFrdUa3GTFYjjjdSaaRWiykfjxh+fIL6/G/TFm+6SaekmUgRFd1OYxAYTt+v8ikp4q9Kla3FQHon19qSLkvaOSn7BklYz4kMpNRb1OuMpsqY7Fs3Yx5kEaGA8Dr2XLOIPXMFd/5SDs85kyeeO8Zk3TFZixuJmKw0UtXUYir1mD8dOMrByjm48CLSEAHBFEinqOPwm1zkHyGdUVTMNpK3ryOdieQYpdTrpjOGjpoTJE+pt9SuFpV4QueUZsg9KycQFPLsum4h73tyNvvHIyqRw3khCiyBNVhrVEBp7o+dEDlP1VkI5kJECga8ob0kjNGW9nKcouTFpOtqHe3FK55EPWOk0f496de0o2smyJqsNgGGhKMrFvLXox0sNjAZO+rOE1qV1QAYBJW1GU+lVsOGljREDN4b/keYeBtddwQYpT2jM0JWnCDeYKzaqpiOgCdLc1i1aDa1eg0XBETWYrAqKwjaXcXjnaMjEKx/FUnrqmLwLiDjpGQdyB6+BsRRFiekGMYLcw/RVZhkXidMVitUqlUq9anUWqlVmaxN7a9w9lzLsvmvINW/pHdVZ6dSbn9Px1Kml/R7wPZ1LJ2mbjaqYxElkzWWnRKDOJAUYe/ZM8imz36AwFc4PnmsmWP/GiutMfBVbrv+AvLj30oVVZzFxcFUdra/h+SG//L910j7T9i6RqvnTlMtJf0DyWT1EbskEiQW8CAoymOHH+Wl3OMM9hRZkI+pHzvSTE3HhR2Oe76wnIL/KbnKr1AUg4iKGuUaCXa1vy/Ttcx0Srz1DCSyckPqWrAu6E9TZx1O3pNUYbtm1tJVLEOuJvg2wt71t81srw5x761F7uvrpvfaC+lrZLhvJffeUuTB/T9iQaX/BFG9t/g4II5ComqIi+0Qb8yNwJhKKbro3Q+IpjRNtwLrE0HYrTVsSWpgRM+ZDtbr9YrAs8C41rMdEN0uzZhvsEa+tL181ZZVQyagx+jDVSZnEAuGhN2v/KGZ985eTOeiTsQJWw8c45k/jrHxrKeYm49f/0+/qKj1lqhRPTf07msPlEXaLuucl3yvTkmTdN7pYVClXacppjwbMF3olxdsAEqaDShayyj/5xhRa4wxdP+4VLAB+3KdthDMMti8aQqLBWNIxdWFuOK5UMa5c9lTdAQeMC1RX+uoOaJqnupkWPaxXXL2mn1lEeGkycgeEdz55ZHyx+6+ai34B8AiAtbTEjYATCuGBPFCB46b3rOfvBVELMmHqVxT1LgaUq2ExPVg7bmffLbMKZGRyao89JUd2z66+crN4unNOQsdYEIwOTCBwVgQk4gqHq4+40XO66wiLkC8xamoLgqIaiH1SiPVYPOSTz+zjdNCRiar8kjvg30f+eGVhLH0hpEl6DDY0GBy0hQWCxiQWChInU/NfZG4lsd700iAjwNcPaBez1Gv5IhqdvP5n9vbx2klI5NVefTrD/Zd9v2Vu+JItoY1U8jlbSKsdlgXCT1dBwmjkFrN4p3BxQFxs6NOxZadM2uXfT77f1cZp/kDVhrL7+wuGMOmIGd6cqEhaMSqrBfnKtw6/yAWdHnKEEe2ETP1esgLfRd9cU95Wn6RJSOTVeFD31lRMJYea80KY2n+fNAd817i3FyMCIinLMJO780u8QxdcvNTp//ngzIyWU8jGRnZrwhmZPwT5I+Pd2qC5IkAAAAASUVORK5CYII")
	var buf bytes.Buffer    //bytesbuffer is a struct
	w, err := zlib.NewWriterLevelDict(&buf, zlib.BestCompression, input)
	if err != nil  {
		fmt.Println("压缩失败")
		return
	}

	w.Write(input)
	w.Close()
	fmt.Println(buf.Bytes())
	fmt.Println(len(buf.Bytes()))
	fmt.Println(len(input))


	//decompress
	var output bytes.Buffer
	r, err1 :=zlib.NewReaderDict(&buf, input)
	if err1 != nil {
		fmt.Println("jie压缩失败")
		return
	}

	//io
	io.Copy(&output, r)

	res:=bytes.Compare(input, output.Bytes())

	if res == 0 {
		fmt.Println("!..Slices are equal..!")
	} else {
		fmt.Println("!..Slice are not equal..!")
	}


	//for i, v := range input{
	//	output.Bytes()
	//}
	//if output.String() == input {
	//
	//}
}

func main() {
	test()

	//nromalcompress()
	//bestcompress()



}
