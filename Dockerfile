FROM alpine:3.3
RUN apk add poppler-utils --no-cache
COPY pdf-convert /usr/bin/pdf-convert
RUN chmod +x /usr/bin/pdf-convert
CMD ["/usr/bin/pdf-convert"]
