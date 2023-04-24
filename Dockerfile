# npm dockerfile 
FROM node:alpine as builder
WORKDIR /app
COPY package.json .
RUN npm install
COPY . .
RUN npm run build


FROM caddy:2-alpine
COPY --from=builder /app/build /var/www/html
COPY Caddyfile /etc/caddy/Caddyfile
ENTRYPOINT [ "caddy" ]
EXPOSE 80
CMD [ "run", "--config", "/etc/caddy/Caddyfile" , "--adapter", "caddyfile"]