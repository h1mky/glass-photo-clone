FROM golang:1.24.5-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . . 


RUN go build -o app ./cmd

FROM alpine


WORKDIR /root/

COPY --from=builder /app/app . 


CMD ["./app"]
