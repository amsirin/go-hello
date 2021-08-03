FROM golang:1.9.2 as builder
WORKDIR /go/src/go-hello
COPY . .
RUN go test -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-hello .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
ENV PORT 8080
COPY --from=builder /go/src/go-hello .
CMD ["./go-hello"]