package tutorial

import (
	"fmt"
	// "io/ioutil"
	"encoding/json"
	"net/http"
)

type IpInfo struct {
	Aaa string
	Bbb string
	Ccc string
}

func HttpServer() {
	fmt.Println("http server")
	http.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {
		ip := r.URL.Query().Get("ip")
		if ip == "" {
			ip = "0.0.0.0"
		}
		fmt.Println(ip)
		ipInfo := IpInfo{"aaa", "bbb", "ccc"}
		b, err := json.Marshal(ipInfo)
		if err != nil {
			b = []byte("[]")
		}
		fmt.Fprintf(w, string(b))
	})
	http.ListenAndServe(":8082", nil)
}
