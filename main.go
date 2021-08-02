package main

import (
	"html/template"
	"log"	// 來輸出程式目前執行的狀態
	"net/http"	// 伺服器監聽與運行的方法與 request handler
)

// 定義 template 所需的參數
type IndexData struct {
	Title   string
	Content string
}

func test(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))	//將data 定義到 index.html
	data := new(IndexData)
	data.Title = "首頁"
	data.Content = "Hello! World"
	tmpl.Execute(w, data)	// 執行程式
}

func main() {
	// Routing
	http.HandleFunc("/", test)	//綁定 test() 到這個 Routing 讓 server 知道 Routing 的 " / " 要執行 test()
	err := http.ListenAndServe(":8888", nil)
	/*
		ListenAndServe 有兩個參數 address 與 handler
		address 為存取的 url 與 port 因為目前沒有指定 url 所以只單純填寫 port
		handler 為所需處理的程序 目前也沒有 所以填寫 nil
	*/
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}