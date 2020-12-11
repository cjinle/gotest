package regex

import (
	"log"
	"net/http"
	"net/url"
	"io/ioutil"
	"regexp"
)

func GetVideoUrl() {
	link := "http://v.baidu.com/watch/01345452608823380694.html"
	res, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	// log.Println(string(bytes))
	reg := `videoFlashPlayUrl\s=\s\'(.*)\'`
	submatch := regexp.MustCompile(reg).FindAllSubmatch(bytes, -1)
	if len(submatch) > 0 && len(submatch[0]) > 0 {
		vv := submatch[0][1]
		urlInfo, _ := url.Parse(string(vv))

		log.Println(urlInfo.Query()["video"])
	}
	
	
}