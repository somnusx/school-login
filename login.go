package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	data := make(url.Values)
	data["username"] = []string{"root"}
	data["domain"] = []string{"studentphone"}
	data["password"] = []string{"password"}
	data["enablemacauth"] = []string{"0"}
			
	
	res, err := http.PostForm("http://n.njcit.cn/index.php/index/login", data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer res.Body.Close()

	fmt.Println("post send success")
}
