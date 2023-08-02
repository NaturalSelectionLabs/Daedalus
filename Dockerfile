# npm dockerfile 
FROM golang:alpine as builder

WORKDIR /app

RUN apk add --no-cache nodejs npm
RUN npm install -g pnpm
COPY . .

WORKDIR /app/ui
RUN pnpm install && pnpm run build


#FROM caddy:2-alpine
#COPY --from=builder /app/build /var/www/html
#COPY Caddyfile /etc/caddy/Caddyfile
#ENTRYPOINT [ "caddy" ]
#EXPOSE 80
#CMD [ "run", "--config", "/etc/caddy/Caddyfile" , "--adapter", "caddyfile"]