package mylib

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

func Zip() {
	files := []string{"aa", "1.txt"}
	output := "out.zip"

	newZipFile, err := os.Create(output)
	checkErr(err)

	defer newZipFile.Close()

	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	for _, file := range files {
		fileToZip, err := os.Open(file)
		checkErr(err)

		info, err := fileToZip.Stat()
		checkErr(err)

		header, err := zip.FileInfoHeader(info)
		checkErr(err)

		header.Name = file
		header.Method = zip.Deflate

		writer, err := zipWriter.CreateHeader(header)
		checkErr(err)

		_, err = io.Copy(writer, fileToZip)
		checkErr(err)
	}

	fmt.Println("zip done!")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
