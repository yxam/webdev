export PORT=8080

go build webdev/cmd/webdev/modelutil
go install webdev/cmd/webdev

cp ../../bin/webdev .

./webdev
