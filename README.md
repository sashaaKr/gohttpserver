# Setup

1. For a new golang project, first run

```bash
got mod init NAME
```

`NAME` - should be something like github.com/sashaaKr/gohttpserver
This will create `go.mod` file

2. Now you can install packages with

```bash
go get github.com/sirupsen/logrus
```

It will add dependency to `go.mod` file and crate `go.sum` file with all fixed versions

# Getting started

```bash
go build
./gohttpserver
curl localhost:8080
```

# Build with docker

```bash
docker build -t gohttpserver .
docker run -p 8080:8080 gohttpserver .
```
