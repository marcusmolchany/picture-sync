package main

import (
    "html/template"
    "net/http"
    "os"
    "fmt"
)

type Page struct {
    Title string
}

func main() {
    http.HandleFunc("/", indexPage)

    fileServer := http.StripPrefix("/css/", http.FileServer(http.Dir("css")))
    http.Handle("/css/", fileServer)

    fileServer = http.StripPrefix("/js/", http.FileServer(http.Dir("js")))
    http.Handle("/js/", fileServer)

    fileServer = http.StripPrefix("/img/", http.FileServer(http.Dir("img")))
    http.Handle("/img/", fileServer)

    err := http.ListenAndServe(":8080", nil)
    checkError(err)
}

func indexPage(rw http.ResponseWriter, req *http.Request) {
    p := Page{
        Title: "index",
    }

    tmpl := make(map[string]*template.Template)
    tmpl["index.html"] = template.Must(template.ParseFiles("index.html"))
    tmpl["index.html"].ExecuteTemplate(rw, "base", p)
}

func checkError(err error) {
    if err != nil {
        fmt.Println("Fatal error ", err.Error())
        os.Exit(1)
    }
}