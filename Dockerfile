FROM golang:1.22.2 AS builder

WORKDIR /build

COPY go.mod ./

RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

RUN go build -o /server cmd/cli/main.go

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=builder /server /server

EXPOSE 4173

ENTRYPOINT ["/server"]
