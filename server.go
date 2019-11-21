package main

import (
    "fmt"
    "net/http"
    "log"
    "os"
    "path/filepath"
    "time"
    "html/template"
    "io/ioutil"
)

type infoStruct struct {
    Path string
    Info os.FileInfo
}

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
        if r.Method == "POST" {
            r.ParseMultipartForm(1<<20)
            file, handler, err := r.FormFile("upload")
            if err != nil {
                log.Println("Error Retrieving the File")
                log.Println(err)
            }
            defer file.Close()
            b, _ := ioutil.ReadAll(file)
            ioutil.WriteFile("static/"+handler.Filename, b, 0666)
        }else if r.Method == "DELETE" {
             r.ParseForm()
             os.Remove(r.Form["filepath"][0])
             fmt.Fprint(w,"deleted")
             return
        }
        var files []infoStruct
        filepath.Walk(dir, func (fpath string, finfo os.FileInfo, err error) error {
            if !finfo.IsDir() {
                fileInfos := infoStruct {
                    Path : string(fpath),
                    Info : finfo,
                }
                files = append(files, fileInfos)
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

