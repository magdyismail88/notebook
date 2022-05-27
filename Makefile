run:
	./bin/notebook &
	cd ./frontend && npm run dev
	echo "Server running, for stop: make stop"
stop:
	killall -9 notebook
	sleep 1
	echo "Server stopped"
build:
	# Build js
	cd ./frontend && npm install --save --legacy-peer-deps
	# Build go
	go mod tidy # install go models requirements (pre-requesties)
	go build .
	go run init/migration.go
	sleep 2
	mv notebook ./bin/notebook
	echo "Successfully Built!"
	sleep 1
