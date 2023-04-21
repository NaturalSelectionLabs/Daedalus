# npm dockerfile 
FROM node:alpine as builder
WORKDIR /app
COPY package.json .
RUN npm install
COPY . .
RUN npm run build


FROM caddy:2-alpine
COPY --from=builder /app/build /usr/share/caddy
ENTRYPOINT [ "caddy" ]
EXPOSE 80
CMD [ "file-server", "--root", "/usr/share/caddy" ]