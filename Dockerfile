FROM golang:1.17-alpine AS build
ADD . /crawler
ENV CGO_ENABLED=0
WORKDIR /crawler
RUN go build -o crawler.bin ./cmd/crawler

FROM alpine:latest
COPY --from=build /crawler/crawler.bin /crawler/crawler.bin
EXPOSE 9999
ENTRYPOINT ["/crawler/crawler.bin"]
