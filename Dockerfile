FROM golang:1.21 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server ./cmd/server
FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY ./internal/assets ./internal/assets
COPY --from=builder /app/server .
EXPOSE 3070
CMD ["./server"]
