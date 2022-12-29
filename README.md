# Golang Tips

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

## Install all third party packages required

From the root folder of the project (where is located the go.mod file):

```
go mod tidy
```
