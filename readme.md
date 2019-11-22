# go_file_server

[![issues x](https://img.shields.io/github/issues/paul-caron/go_file_server)](https://img.shields.io/github/issues/paul-caron/go_file_server) [![license x](https://img.shields.io/github/license/paul-caron/go_file_server)](https://img.shields.io/github/license/paul-caron/go_file_server)

This file server lists the files that are recursively present inside the static directory.

### Configuration

All configuration is done inside the "config.json"

Configuration is done as follows :
```json
{
    "Port": 8080,
    "TLSEnabled": true,
    "SSLCertificate": "cert.pem",
    "SSLKey": "key.pem"
}
```

### Start
Execute a compiled binary of server.go or run it uncompiled.
```bash
go run server.go
```
All files inside the static directory are then listed at url :
```
https://domain:8080/
```
or
```
http://domain:8080/
```
### Dependencies

The go file itself has no dependencies but the html template uses the following two: 

- [JQuery](https://jquery.com/)
- [DataTables](https://datatables.net/)

DataTables is a Jquery plugin that makes HTML tables beautiful.
