FROM golang:alpine as base

FROM base as dev

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o demo-ddd

ENTRYPOINT ["/app/demo-ddd"]