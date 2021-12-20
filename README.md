# Testing in Go

go test -cover
go test -cover ./...
go test -coverprofile coverFile.out
go test -coverprofile coverFile.out ./...

go tool cover -func coverFile.out 
go tool cover -html coverFile.out 

go test -json
go test -json ./...
go test -json ./...  > testResults.json
