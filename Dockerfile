FROM golang:1.27rc3-alpine3.24

WORKDIR /golangApp

COPY . .
RUN go build -o app_bin .

ENTRYPOINT [ "./app_bin" ]
