FROM golang:1.19-alpine

WORKDIR /build

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o ./main cmd/app/main.go

EXPOSE 8080

CMD ["./main"]

