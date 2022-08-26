FROM golang:1.19.0

WORKDIR /src

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -v ./...

CMD ["go", "test", "-v", "./..."]
