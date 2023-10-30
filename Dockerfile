FROM node:18-alpine as base

RUN npm install -g pnpm

FROM base as dependencies

WORKDIR /app
COPY package.json pnpm-lock.yaml ./
RUN pnpm install

FROM base as builder

WORKDIR /app
COPY . .
COPY --from=dependencies /app/node_modules ./mode_modules
RUN pnpm build

FROM caddy
COPY Caddyfile /etc/caddy/Caddyfile
COPY --from=builder /app/dist /usr/share/caddy/html