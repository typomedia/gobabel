build:
	go mod tidy
	go build -ldflags "-s -w" -o dist/ .

run:
	go mod tidy
	go run main.go

cross:
	go mod tidy
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o dist/gobabel-linux-amd64 .
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o dist/gobabel-macos-amd64 .
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o dist/gobabel-windows-amd64.exe .
