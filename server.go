package main

import (
    "net/http"
    "log"
    "os"
    "path/filepath"
    "time"
    "html/template"
)

func main(){
    var dir string = "static"
    var port string
    if len(os.Args) != 2 {
        log.Fatal("The command is incorrect. The correct format is:\n    fileserver port ")
    }
    port = os.Args[1]
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

