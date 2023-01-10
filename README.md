# Notebook


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

<img src="https://github.com/magdyismail88/notebook/blob/main/assets/note02.png" />


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
├── assets
│   ├── note00.png
│   ├── note01.png
│   └── note02.png
├── bin
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
│   ├── build
│   │   ├── build.js
│   │   ├── check-versions.js
│   │   ├── logo.png
│   │   ├── utils.js
│   │   ├── vue-loader.conf.js
│   │   ├── webpack.base.conf.js
│   │   ├── webpack.dev.conf.js
│   │   └── webpack.prod.conf.js
│   ├── config
│   │   ├── dev.env.js
│   │   ├── index.js
│   │   └── prod.env.js
│   ├── index.html
│   ├── package.json
│   ├── package-lock.json
│   ├── src
│   │   ├── App.vue
│   │   ├── assets
│   │   │   └── logo.png
│   │   ├── components
│   │   │   ├── Auth
│   │   │   │   └── Login.vue
│   │   │   ├── Container
│   │   │   │   ├── ContainerChangeForm.vue
│   │   │   │   ├── ContainerEditForm.vue
│   │   │   │   ├── ContainerForm.vue
│   │   │   │   └── Container.vue
│   │   │   ├── Flash.vue
│   │   │   ├── Home.vue
│   │   │   ├── Note
│   │   │   │   ├── NoteEditForm.vue
│   │   │   │   ├── NoteForm.vue
│   │   │   │   ├── NoteShow.vue
│   │   │   │   └── Note.vue
│   │   │   ├── partials
│   │   │   │   └── Header.vue
│   │   │   └── Tab
│   │   │       ├── TabEditForm.vue
│   │   │       ├── TabForm.vue
│   │   │       ├── Tabs.vue
│   │   │       └── Tab.vue
│   │   ├── helpers
│   │   │   └── utils.js
│   │   ├── main.js
│   │   ├── router
│   │   │   └── index.js
│   │   └── store
│   │       ├── index.js
│   │       └── modules
│   │           ├── container.js
│   │           ├── note.js
│   │           └── tab.js
│   └── static
│       ├── css
│       │   └── style.css
│       ├── img
│       │   └── logo.png
│       └── js
│           └── script.js
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