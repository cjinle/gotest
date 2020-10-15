package tutorial

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func DownloadPic() {
	picUrl := "http://pic-bucket.ws.126.net/photo/0001/2020-10-15/FOVJECMS4T8E0001NOS.jpg"
	outFile := "aa.jpg"
	fmt.Println("Pic URL: ", picUrl)
	res, err := http.Get(picUrl)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	bytes, _ := ioutil.ReadAll(res.Body)
	err = ioutil.WriteFile(outFile, bytes, 0644)
	if err != nil {
		panic(err)
	}
}