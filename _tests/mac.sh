go run '_tests/main.go'
go test -v "sync_tree/calc"
go test -v "sync_tree/data"
go test -v "sync_tree/lock"
go test -v "sync_tree/logs"
go test -v "sync_tree/market"
go test -v "sync_tree/user"
go run '_tests/temp/main.go'
go test ./...
go run '_tests/temp/main.go'