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
│   │   ├── container.go
│   │   ├── note.go
│   │   ├── tab.go
│   │   └── uploader.go
│   └── models
│       ├── constants.go
│       ├── container.go
│       ├── note.go
│       └── tab.go
├── app.go
├── bin
├── bootstrap
│   ├── app.go
│   ├── database.go
│   └── env.go
├── cmd
│   └── main.go
├── config
│   └── routes.go
├── data
│   └── notebook.db
├── database
│   ├── constants.go
│   ├── container.go
│   ├── migration.go
│   ├── note.go
│   └── tab.go
├── frontend
├── go.mod
├── go.sum
├── install.go
├── LICENSE
├── README.md
├── storage
├── tests
└── util
    └── move_file.go
```
