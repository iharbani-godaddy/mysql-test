linux:
	 env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -a -installsuffix cgo -o ./build/linux/mysql-test -v main.go
