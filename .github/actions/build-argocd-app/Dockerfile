FROM golang:alpine as builder
WORKDIR /app
COPY . .
RUN go build



FROM alpine
COPY --from=builder /app/daedalus /daedalus
ENTRYPOINT [ "/daedalus" ]