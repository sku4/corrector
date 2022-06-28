FROM golang:1.18.1-alpine3.15 AS builder

RUN go version

COPY . /github.com/sku4/corrector/
WORKDIR /github.com/sku4/corrector/

RUN go mod download
RUN GOOS=linux go build -o ./.bin/app ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/sku4/corrector/.bin/app .
COPY --from=0 /github.com/sku4/corrector/configs/config.yml configs/config.yml

EXPOSE 8000

CMD ["./app"]