FROM golang:1.24-alpine as builder

RUN mkdir /app
WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 go build -o brokerapp ./cmd/api

RUN chmod +x brokerapp


FROM alpine:latest
RUN mkdir /app

COPY --from=builder app/brokerapp /app/brokerapp

CMD ["/app/brokerapp"]