FROM golang:1.19-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o /docker-gs-ping
CMD ["/docker-gs-ping"]