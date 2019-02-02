package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/login", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
	fmt.Println("run API ok", "http://localhost:8080/login")
}

/*
 * 处理表单
 *
 */
func disposeForm(w http.ResponseWriter, r *http.Request) {

	// if r.ParseForm() != nil {
	// 	fmt.Println(r.ParseForm())
	// 	w.WriteHeader(http.StatusInternalServerError) // Proper HTTP response
	// 	return
	// }
	fmt.Println("接收到的值", r.PostFormValue(""))
	var bodyAll = r.PostForm
	var user = r.PostForm["username"][0]

	fmt.Println("接收到bodyAll", bodyAll)
	fmt.Println("接收到urser", user)
	fmt.Println("接收到Form", r.Form["username"])

	fmt.Println("method:", r.Method)
	fmt.Println("path:", r.URL.Path)
	fmt.Println("scheme:", r.URL.Scheme)
	fmt.Println("url", r.URL)
	jsondata(w, user)
}

/*
 *	设置跨域的头
 */
func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")                //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")    //header的类型
	w.Header().Set("Content-Type", "application/json; charset=utf-8") //返回数据格式是json

	disposeForm(w, r) //处理表单

}

/*
 *	设置返回的值
 */
func jsondata(w http.ResponseWriter, str string) {
	type Info struct {
		Name string
		Age  int
		Sex  string
	}
	infos := []Info{
		{
			Name: str,
			Age:  23,
			Sex:  "female",
		},
		{
			Name: "Benjie",
			Age:  24,
			Sex:  "male",
		},
	}
	data, _ := json.Marshal(infos)

	fmt.Fprintf(w, "%s", data)
}
