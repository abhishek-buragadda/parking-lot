#/bin/bash
go get github.com/stretchr/testify/assert
go test ./... -v
GOOS=linux go build -o parking_lot