# go_file_server

[![dependencies 0](https://img.shields.io/github/issues/paul-caron/go_file_server)](https://img.shields.io/github/issues/paul-caron/go_file_server)

This file server lists the files that are recursively present inside the static directory.

To use this file server, the command goes like:
```
fileserver port
```

For example, we could start the server like:
```
fileserver 8080
```
This would start the server at port 8080

All files inside the static directory are all listed at url endpoint "/".
Like:
```
http://domain:8080/
```



