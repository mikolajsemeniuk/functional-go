# Functional Go

Go Web API using functional paradigm

```sh
docker-compose up -d
air
go test -coverprofile=coverage.out `go list ./... | grep -v docs`

go tool cover -func=coverage.out # cli
go tool cover -html=coverage.out # browser
```
