# Notebook

Simple notebook application written with golang, vue and Sqlite3 for more portable and simple



### Requirments
- Golang
- npm



```bash
$ go run run.go --help
```

```bash
Usage:

	go run app.go [--argument]

The commands are:
	--build        build dependencies
	--server	     start server
```

```bash
$ go run run.go --build
```

```bash
$ go run run.go --server
```
or

```bash
$ ./bin/notebook-bin
```

```
localhost:8888
```

### Example

![Screenshot](https://github.com/magdyismail88/notebook/blob/651a9dccc2776d49b47fe0ef07d3b720f2e53038/assets/screenshot01.png?raw=true")


### Application Structure

```
.
├── app
│   ├── controllers
│   │   ├── application.go
│   │   ├── container.go
│   │   ├── note.go
│   │   ├── tab.go
│   │   └── uploader.go
│   ├── models
│   │   ├── container.go
│   │   ├── note.go
│   │   └── tab.go
│   └── services
│       ├── container.go
│       ├── note.go
│       └── tab.go
├── bin
├── bootstrap
│   ├── app.go
│   └── env.go
├── cmd
│   └── install.go
├── config
│   └── routes
│       ├── api.go
│       ├── kernel.go
│       └── web.go
├── data
├── database
│   ├── constants.go
│   ├── container.go
│   ├── migration.go
│   ├── note.go
│   └── tab.go
├── go.mod
├── go.sum
├── LICENSE
├── README.md
├── run.go
├── server
│   └── main.go
├── storage
│   └── upload_image.png
├── tests
└── util
    └── move_file.go
```
