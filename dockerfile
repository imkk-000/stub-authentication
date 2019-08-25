FROM golang:1.12 as builder
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/login/main.go

FROM alpine:3.10.2
COPY --from=builder /build/app .
EXPOSE 8888
CMD ./app
