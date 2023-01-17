# Build stage
FROM golang:1.19 as builder

ENV GO111MODULE=on

WORKDIR /go/src/go_crud_app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o ./app ./cmd/main

RUN ls

# Final stage
FROM alpine

WORKDIR /root/

COPY --from=builder /go/src/go_crud_app/app ./app
RUN ls /root/

EXPOSE 8080
CMD ["./app"]