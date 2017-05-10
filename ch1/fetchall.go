package main

import (
	"time"
	"fmt"
	"net/http"
	"os"
)

func main() {

	var ch1 chan string = make(chan string)
	var ch2 chan string = make(chan string)
	var ch3 chan string = make(chan string)
	var ch4 chan string = make(chan string)

	var url1 string = "https://www.baidu.com"
	var url2 string = "https://www.sohu.com"
	var url3 string = "https://www.so.com"
	var url4 string = "https://cn.bing.com"

	go fetch(url1, ch1)
	go fetch(url2, ch2)
	go fetch(url3, ch3)
	go fetch(url4, ch4)
	<-ch1
	<-ch2
	<-ch3
	<-ch4

}

func fetch(url string, ch chan string) {
	fmt.Println("start fetching url...")
	start := time.Now()
	resp, err := http.Get(url)
	checkError(err)
	fmt.Printf("%s\n", resp.Header)

	secs := time.Since(start).Seconds()
	fmt.Printf("%.2fs", secs)
	ch <- url
}

func checkError(err error) {
	if err != nil {
		fmt.Printf("error:%v", err)
		os.Exit(-1)
	}

}
