# Golang Course (and Tips)

## Course

https://www.udemy.com/course/curso-golang

## Run code

```
go run mainFile.go
```

Examples:

```
go run ./topics/27-with-concurrency/main.go
go run ./challenges/06-geometry-modules/main.go
```

## New Project

Create new module

```
go mod init repo-or-project-name
```

## Install third party packages

From the root folder of the project (where is located the go.mod file):

```
go get package_repo_url
```

Example:

```
go get github.com/donvito/hellomod
```

## Install all required third party packages

From the root folder of the project (where is located the go.mod file):

```
go mod tidy
```

## Test

```
go test ./folder_with_tests
```

Example:

```
go test ./topics/28-basic-testing/
```

## Test Coverage

```
go test -cover ./folder_with_tests
```

Example:

```
go test -cover ./topics/28-basic-testing/
```

### More coverage info

```
go test ./folder_with_tests -coverprofile=coverage
go tool cover -func=coverage
```

Example:

```
go test ./topics/28-basic-testing/ -coverprofile=coverage
go tool cover -func=coverage
```

### HTML coverage info

```
go test ./folder_with_tests -coverprofile=coverage
go tool cover -html=coverage
```

Example:

```
go test ./topics/28-basic-testing/ -coverprofile=coverage
go tool cover -html=coverage
```

## Test Profile

Install [graphviz](https://graphviz.org/)

```
go test ./folder_with_tests -cpuprofile=profile
go tool pprof profile
top
list functionName
web
pdf
quit
```

Example:

```
go test ./topics/28-basic-testing/ -cpuprofile=profile
go tool pprof profile
top
list Fibonacci
web
pdf
quit
```
