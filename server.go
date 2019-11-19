package main

import (
    "fmt"
    "net/http"
    "log"
    "os"
    "path/filepath"
    "time"
    "html/template"
)

func main(){
    args:=os.Args
    var dir string = "static"
    var port string
    if len(args) < 2 {
        fmt.Println("Missing arguments. The correct format is:\n    fileserver port ")
        os.Exit(1)
    }
    if len(args) == 2 {
        port = args[1]
    }
    if len(args) >=3 {
        fmt.Println("Too many arguments. The correct format is:\n    fileserver port ")
        os.Exit(1)
    }
    myServer := &http.Server{
        MaxHeaderBytes : 1 << 20,
        ReadTimeout : 1 * time.Second,
        WriteTimeout : 10 * time.Second,
        Addr : ":" + port,
    }
    http.HandleFunc("/", func(w http.ResponseWriter, r * http.Request){
        var files []string
        filepath.Walk(dir, func (path string, info os.FileInfo, err error) error {
            if !info.IsDir() {
                files = append(files, path)
            }
            return nil
        })
        templ, _ := template.ParseFiles("template.html")
        templ.Execute(w, files)
    })
    log.Println("Starting file server in directory : ",dir+":"+port)
    fileServer := http.FileServer(http.Dir("./"))
    http.Handle("/static/", fileServer)
    log.Fatal(myServer.ListenAndServe())
}

