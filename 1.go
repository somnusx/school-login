package main

import (
	"fmt"
	"os"
	"time"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"encoding/json"
)



func pan() {

		response,err3:=http.Get("http://www.baidu.com")
		if err3 != nil {
			return
		}
		defer response.Body.Close()

		body,err4:=ioutil.ReadAll(response.Body)
		if err4 != nil {
			return
		}
		
		reg := regexp.MustCompile(`baidu`)
		if reg.FindAllString(string(body), -1)[0]=="baidu"{
			fmt.Println("\u767b\u5f55\u6210\u529f")
			time.Sleep(1200 * time.Second)
			os.Exit(1)
		}
}



func main() {
	//s := "KETX15132"
	//pw := "OTcwNTAx"
	s := "71441P20"
	pw := "NzE0NDEy"
	time1 := 0.5
	time2 := 0.6
	for i := 1; i < 82; i++ {
	if i == 80{
	fmt.Println("\u767b\u5f55\u5931\u8d25\u3002\u3002\u3002")
	time.Sleep(1200 * time.Second)
	os.Exit(1)	
	}
	if i == 40{
	time1 = time1 + 0.4
	time2 = time2 + 0.4
	fmt.Println("\u5207\u6362\u65f6\u95f4")
	fmt.Println("time=0.9")
	fmt.Println("time=1")
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
	time.Sleep(1200 * time.Second)
	os.Exit(1)	
	}

	if user["info"] == "\u8bbe\u5907\u672a\u54cd\u5e94"{
			pan()
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
		pan()
}
}