package main

import (
	"bufio"
	"fmt"
	"github.com/upyun/go-sdk/v3/upyun"
	"mime/multipart"
	"os"
)

const (
	Bucket   = ""
	Operator = ""
	Password = ""
)

func main() {
	up := upyun.NewUpYun(&upyun.UpYunConfig{
		Bucket:   Bucket,
		Operator: Operator,
		Password: Password,
	})

	/**
		args : folder || filepath
	 */
	args := "tests001"
	_, err := up.GetInfo(args)
	if err == nil {
		fmt.Println("folder exist")
	}
	// the folder isn't exist
	// create folder
	up.Mkdir(args)

	// upload file
	file, e := os.Open("README.md")
	if e != nil {
		panic("not file or folder ")
	}
	defer file.Close()
	isSuccess := up.Put(&upyun.PutObjectConfig{
		Path: args + `/README.md`,
		//LocalPath: "README.md",
		Reader: bufio.NewReader(file),
	})
	fmt.Println(isSuccess) // nil


	// 上传 form文件
	var header *multipart.FileHeader
	//filename := header.Filename
	out, _ := header.Open()
	defer out.Close()
	up.Put(&upyun.PutObjectConfig{
		Path: args + `/file_random_name.jpeg` ,
		//LocalPath: "README.md",
		Reader: bufio.NewReader(out),
	})
}
