package main

import (
    "fmt"
    "net/http"
    "log"
    "os"
    "path/filepath"
    "time"
    "strings"
)


func atag(href string) string {
    splice := strings.Split(href,"/")
    endpoint := strings.Join(splice[1:],"/")
    var tag = "<a href="+endpoint+">"+href+"</a>"
    return tag
}

func main(){
    args:=os.Args
    var dir, port string
    if len(args) <= 2 {
        fmt.Println("Missing arguments. The correct format is:\n    fileserver port directory")
        os.Exit(1)
    }
    if len(args) == 3 {
        port = args[1]
        dir = args[2]
    }
    if len(args) > 3 {
        fmt.Println("Too many arguments. The correct format is:\n    fileserver port directory")
        os.Exit(1)
    }
    myServer := &http.Server{
        MaxHeaderBytes : 1 << 20,
        ReadTimeout : 1 * time.Second,
        WriteTimeout : 10 * time.Second,
        Addr : ":" + port,
    }
    http.HandleFunc("/index", func(w http.ResponseWriter, r * http.Request){
        var response string = `<!DOCTYPE html>
<head>
<title>Files</title>
<meta name="viewport" content="width=device-width initial-scale=1.0">
</head>
<body>
<h1>Files</h1>
`
        var files []string
        err := filepath.Walk(dir, func (path string, info os.FileInfo, err error) error {
            if !info.IsDir() {
                files = append(files, path)
            }
            return nil
        })
        if err == nil {
        response+="<ol type='I'>"
        for _, path := range files {
            response += "<li>"+atag(path)+"</li>"
        }
        response += "</ol>"
        }
        fmt.Fprintf(w,"%s",response+"<p>&copy fileserver by Paul Caron</p></body>")
    })
    log.Println("Starting file server in directory : ",dir+":"+port)
    fileServer := http.FileServer(http.Dir(dir))
    http.Handle("/", fileServer)
    err := myServer.ListenAndServe()
    if err != nil {
        log.Println(err);
    }
}

