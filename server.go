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
    "encoding/json"
    "strconv"
)

const dir string = "static"

type infoStruct struct {
    Path string
    Info os.FileInfo
}

type Config struct {
    Port int `json:"port"`
    TLSEnabled bool `json:"TLSEnabled"`
    SSLCertificate string `json:"SSLCertificate"`
    SSLKey string `json:"SSLKey"`
}

func main(){
    var conf Config
    configFile, _ := os.Open("config.json")
    defer configFile.Close()
    byteConfig, _ := ioutil.ReadAll(configFile)
    json.Unmarshal(byteConfig, &conf)
    fmt.Println(conf.Port,conf.TLSEnabled)
    myServer := &http.Server{
        MaxHeaderBytes : 1 << 20,
        ReadTimeout : 10 * time.Second,
        WriteTimeout : 10 * time.Second,
        Addr : ":" + strconv.Itoa(conf.Port),
    }
    http.HandleFunc("/", root)
    http.HandleFunc("/delete", deleteFile)
    http.HandleFunc("/upload", uploadFile)
    fileServer := http.FileServer(http.Dir("./"))
    http.Handle("/static/", fileServer)
    if conf.TLSEnabled {
        log.Println("https")
        log.Fatal(myServer.ListenAndServeTLS(conf.SSLCertificate, conf.SSLKey))
    }else{
        log.Println("http")
        log.Fatal(myServer.ListenAndServe())
    }
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
        os.Remove(r.Form["filepath"][0])
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
