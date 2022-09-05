FROM golang:alpine
RUN apk add poppler-utils --no-cache
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY pdf-convert.go ./
RUN  go build
COPY pdf-convert /usr/bin/pdf-convert
RUN chmod +x /usr/bin/pdf-convert
CMD ["/usr/bin/pdf-convert"]
