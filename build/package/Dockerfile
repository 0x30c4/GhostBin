FROM golang:1.18-alpine3.15 AS builder

ENV GO111MODULE=on 

WORKDIR /go/src/github.com/0x30c4/GoPasteBin/
 
COPY . .

ENV TZ Asia/Dhaka

#dependencies
RUN go mod tidy 

WORKDIR /go/src/github.com/javiersoto15/skeleton-tutorial/cmd/tutorial
 
RUN GOOS=linux go build -a -installsuffix cgo -o app .


#Second stage for a smaller image
FROM alpine:3.15
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/javiersoto15/skeleton-tutorial/cmd/tutorial/app .
CMD ["./app"]
