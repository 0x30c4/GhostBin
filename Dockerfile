# build
FROM golang:latest as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/air-verse/air@latest
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/ghostbin ./cmd/webapp/main.go

# deploy
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app .
EXPOSE 8080
CMD ["/app/bin/ghostbin"]
