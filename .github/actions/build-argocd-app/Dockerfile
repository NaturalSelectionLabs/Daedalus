FROM golang:alpine as builder
WORKDIR /app
COPY . .
RUN go build



FROM alpine
COPY --from=builder /app/daedalus /daedalus
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]