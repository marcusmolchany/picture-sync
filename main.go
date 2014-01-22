package main

import (
    "html/template"
    "net/http"
    "os"
    "fmt"
    "picture-sync/src/gplus"
)

func main() {
    http.HandleFunc("/", indexPage)

    http.HandleFunc("/redirect", redirectPage)
    http.HandleFunc("/redirect/", redirectPage)

    fileServer := http.StripPrefix("/css/", http.FileServer(http.Dir("css")))
    http.Handle("/css/", fileServer)

    fileServer = http.StripPrefix("/js/", http.FileServer(http.Dir("js")))
    http.Handle("/js/", fileServer)

    fileServer = http.StripPrefix("/img/", http.FileServer(http.Dir("img")))
    http.Handle("/img/", fileServer)

    err := http.ListenAndServe(":80", nil)
    checkError(err)
}

func indexPage(rw http.ResponseWriter, req *http.Request) {
    type Page struct {
        Title string
    }

    p := Page{
        Title: "index",
    }

    tmpl := make(map[string]*template.Template)
    tmpl["index.html"] = template.Must(template.ParseFiles("index.html"))
    tmpl["index.html"].ExecuteTemplate(rw, "base", p)
}

func redirectPage(rw http.ResponseWriter, req *http.Request) {
    type Page struct {
        Title string
        Url   string
    }

    url := gplus.Connect()
    // gplus.GetRequest("https://www.googleapis.com/plus/v1/activities?query=hello&key=AIzaSyATrgLzeMVYdQemwyzhmbTrE4oYB2-sQp0")

    p := Page{
        Title: "index",
        Url: url,
    }

    tmpl := make(map[string]*template.Template)
    tmpl["redirect.html"] = template.Must(template.ParseFiles("redirect.html"))
    tmpl["redirect.html"].ExecuteTemplate(rw, "base", p)
}

func checkError(err error) {
    if err != nil {
        fmt.Println("Fatal error ", err.Error())
        os.Exit(1)
    }
}
