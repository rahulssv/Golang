package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://stackoverflow.com/",
		"https://www.youtube.com/",
	}
	c:=make(chan string)

	for _,link := range(links){
		go getUrl(link,c)
	}
	for data:= range(c){
		go func(d string){
			time.Sleep(5*time.Second)
			getUrl(d,c)
		}(data)
	}
}

func getUrl(link string,c chan string){
	_, err := http.Get(link)

	if err!=nil{
		fmt.Println(link,"server may be down")
		c<-link
	}
	fmt.Println(link,"server is up")
	c<-link
}