git pull

go mod download

CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o server

./server 
