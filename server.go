package main

import (
_    "fmt"
    "net/http"
    "log"
    "os"
    "path/filepath"
    "time"
    "html/template"
    "io/ioutil"
)

const dir string = "static"

type infoStruct struct {
    Path string
    Info os.FileInfo
}

func main(){
    var port string
    if len(os.Args) != 2 {
        log.Fatal("The command is incorrect. The correct format is:\n    fileserver port ")
    }
    port = os.Args[1]
    myServer := &http.Server{
        MaxHeaderBytes : 1 << 20,
        ReadTimeout : 10 * time.Second,
        WriteTimeout : 10 * time.Second,
        Addr : ":" + port,
    }
    http.HandleFunc("/", root)
    http.HandleFunc("/delete", deleteFile)
    http.HandleFunc("/upload", uploadFile)
    log.Println("Starting file server in directory : ",dir+":"+port)
    fileServer := http.FileServer(http.Dir("./"))
    http.Handle("/static/", fileServer)
    log.Fatal(myServer.ListenAndServeTLS("cert.pem","key.pem"))
}


func root(w http.ResponseWriter, r *http.Request) {
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
}

func deleteFile(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    if len(r.Form["filepath"]) > 0 {
        log.Println(os.Remove(r.Form["filepath"][0]))
        templ, _ := template.ParseFiles("templateDelete.html")
        templ.Execute(w, "")
    }
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
    r.ParseMultipartForm(1<<20)
    file, handler, err := r.FormFile("upload")
    if err != nil {
        log.Println("Error Retrieving the File")
        log.Println(err)
    }
    defer file.Close()
    b, _ := ioutil.ReadAll(file)
    ioutil.WriteFile("static/"+handler.Filename, b, 0666)
    templ, _ := template.ParseFiles("templateUpload.html")
    templ.Execute(w, nil)
}
