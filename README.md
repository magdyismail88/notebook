# Notebook

Simple notebook application written with golang, vue and Sqlite3 for more portable and simple



### Requirments
- Golang
- npm



```bash
$ go run app.go --help
```

```bash
Usage:

	go run app.go [--argument]

The commands are:
	--build          build dependencies
	--server	     start server
```

```bash
$ go run app.go --build
```

```bash
$ go run app.go --server
```
or

```bash
$ ./bin/notebook-bin
```

### Example

![Screenshot](https://github.com/magdyismail88/notebook/blob/651a9dccc2776d49b47fe0ef07d3b720f2e53038/assets/screenshot01.png?raw=true")


### Application Structure

```
.
├── api
│   ├── constants.go
│   ├── container.go
│   ├── main.go
│   ├── note.go
│   ├── router.go
│   ├── tab.go
│   └── uploader.go
├── app.go
├── bin
│   └── notebook-bin
├── bootstrap
│   └── install.go
├── data
│   └── notebook.db
├── database
│   ├── constants.go
│   ├── container.go
│   ├── migration.go
│   ├── note.go
│   └── tab.go
├── data.txt
├── frontend
├── go.mod
├── go.sum
├── LICENSE
├── README.md
├── storage
├── tests
│   └── demo_test.go
└── util
    └── move_file.go
```



Magdy Ismail