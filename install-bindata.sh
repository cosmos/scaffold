go get -u github.com/go-bindata/go-bindata/... 
cd ${GOPATH}/src/github.com/go-bindata/go-bindata/
go test
cd go-bindata
go build 
mv go-bindata /usr/local/bin/