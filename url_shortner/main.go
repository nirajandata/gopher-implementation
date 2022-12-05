package main

import (
	"fmt"
	"net/http"
	"html/template"
	"github.com/gorilla/mux"
	"crypto/md5"
    "encoding/hex"
)
var data map[string] string

// snippet from a md5 related gist
func md5Hasher(text string) string {
    hasher := md5.New()
    hasher.Write([]byte(text))
    return hex.EncodeToString(hasher.Sum(nil))
}

func formProcess(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		temp:=template.Must(template.ParseFiles("home.html"))
		temp.Execute(w,nil)

	} 	else {
		var text string =r.FormValue("value")
		encode:=md5Hasher(text)
		data[encode]=text
		fmt.Fprintf(w,"your new short link is http://localhost:8080/short/"+encode)
	}
}

func shorter(w http.ResponseWriter, r *http.Request){
	vars:=mux.Vars(r)
	url:=vars["url"]
	if value,ok :=data[url]; ok {
//		http.RedirectHandler(value,http.StatusMovedPermanently)
	http.Redirect(w,r,value,http.StatusSeeOther)
	}
	fmt.Fprintf(w,"sorry, your link is invalid \n")


}
func main(){
	data=make(map[string]string)
	r:=mux.NewRouter()
	r.HandleFunc("/",formProcess)	
	r.HandleFunc("/short/{url}",shorter)
	http.ListenAndServe(":8080",r)
}

