# npm dockerfile

FROM node:alpine as web-builder
WORKDIR /app
COPY . .
RUN npm install -g pnpm
COPY . .

WORKDIR /app/ui
RUN pnpm install && pnpm run build

FROM golang:alpine as builder

WORKDIR /app

COPY --from=web-builder /app ./
RUN go build



FROM alpine
COPY --from=builder /app/daedalus /daedalus
ENTRYPOINT [ "/daedalus" ]
EXPOSE 80
CMD [ "server", "--port", "80"]