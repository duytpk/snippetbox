### Run on CentOS 7 ###
cd snippetbox
go run $(ls -1 cmd/web/*.go | grep -v _test.go)