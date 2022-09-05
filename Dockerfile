## Build
FROM golang:alpine AS build
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY *.go ./
RUN go build -o /usr/bin/pdf-convert

## Deploy
FROM alpine
WORKDIR /
RUN apk add poppler-utils --no-cache
COPY --from=build /usr/bin/pdf-convert /usr/bin/pdf-convert
CMD ["/usr/bin/pdf-convert"]
