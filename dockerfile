From golang:1.23
LABEL org.opencontainers.image.source="https://github.com/kilianp07/cassandracrud"
WORKDIR /go/src/app

COPY . .

RUN go build -o main main.go

CMD ["./main"]
