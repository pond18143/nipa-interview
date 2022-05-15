FROM golang:1.13.3-alpine

ENV GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

WORKDIR /app
COPY . .

EXPOSE 8080
RUN go build
CMD ["./nipa-interview"]
