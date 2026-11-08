# build
FROM golang:latest as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/air-verse/air@v1.61.1
COPY . .
ARG CGO_ENABLED=0
ARG GOOS=linux
RUN go build -ldflags "-w -s" -o ./bin/ghostbin ./cmd/webapp/main.go

# deploy
FROM gcr.io/distroless/static-debian13 as worker
WORKDIR /app
COPY --from=builder /app .
EXPOSE 8080
CMD ["/app/bin/ghostbin"]
