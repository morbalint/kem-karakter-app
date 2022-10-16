FROM alpine as deploy

WORKDIR /app

FROM golang:1.19.0-alpine as builder

WORKDIR /src

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o /server ./cmd/api

FROM deploy

COPY --from=builder /server /app/server

ENTRYPOINT [ "/app/server" ]
