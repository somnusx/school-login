package main

import (
	"fmt"
	"os"
	"time"
	"io/ioutil"
	"net/http"
	"net/url"
	"encoding/json"
)

func main() {
	s := "KETX15132"
	pw := "OTcwNTAx"

//	var i int = 0
	time1 := 0.9
	time2 := 1
	for i := 0; i < 81; i++ {
	if i == 80{
	fmt.Println("\u767b\u5f55\u5931\u8d25\u3002\u3002\u3002")
	os.Exit(1)	
	}
	fmt.Printf("\u7b2c%d\u6b21\u5c1d\u8bd5......\n",i)
	data := make(url.Values)
	data["username"] = []string{s}
	data["domain"] = []string{"student"}
	data["password"] = []string{pw}
	data["enablemacauth"] = []string{"0"}
			
	
	res, err := http.PostForm("http://n.njcit.cn/index.php/index/login", data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer res.Body.Close()
	
	var user map[string]interface{}  
	body, _ := ioutil.ReadAll(res.Body)  
	json.Unmarshal(body, &user)  
	if (user["info"] == "\u8ba4\u8bc1\u5931\u8d25\u002c\u0020\u5e76\u53d1\u767b\u5f55\u8d85\u8fc7\u6700\u5927\u9650\u5236") || (user["info"] == "\u8ba4\u8bc1\u5931\u8d25\u002c\u0020\u8bf7\u68c0\u67e5\u5bc6\u7801\u53ca\u8d26\u6237\u72b6\u6001"){
	fmt.Println("Local account is not available")
	os.Exit(1)	
	}

	if user["info"] == "\u8bbe\u5907\u672a\u54cd\u5e94"{
		fmt.Println("\u767b\u5f55\u6210\u529f")
		break
		os.Exit(1)
	}
	
	time.Sleep(time.Duration(time1) * time.Second)
	logout()
	time.Sleep(time.Duration(time2) * time.Second)
	}
}

func logout() {
	data := make(url.Values)
	data[""] = []string{""}
	
	res, err := http.PostForm("http://n.njcit.cn/index.php/index/logout", data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer res.Body.Close()

	var use map[string]interface{}  
	body, _ := ioutil.ReadAll(res.Body)  
	json.Unmarshal(body, &use)  
	if use["info"] == "\u8bbe\u5907\u672a\u54cd\u5e94"{
		fmt.Println("\u767b\u5f55\u6210\u529f")
		os.Exit(1)
	}
}
