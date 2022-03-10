FROM golang as builder

WORKDIR /app

# Put mod an sum here in order to decrease build time then only code was changed and all deps are up to date
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

# we do not need all golang tools in order to run service and we can use smaller image
FROM scratch

COPY --from=builder /app/gohttpserver /app/
EXPOSE 8080
ENTRYPOINT [ "app/gohttpserver" ]