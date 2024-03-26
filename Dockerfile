FROM golang:latest
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/cosmtrek/air@latest
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/ghostbin ./cmd/ghostbin/main.go
EXPOSE 8080
ENTRYPOINT ["/app/bin/ghostbin"]
