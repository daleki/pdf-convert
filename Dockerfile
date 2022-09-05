FROM golang:alpine
RUN apk add poppler-utils --no-cache
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY *.go ./
RUN go build -o /usr/bin/pdf-convert
RUN chmod +x /usr/bin/pdf-convert
CMD ["/usr/bin/pdf-convert"]
