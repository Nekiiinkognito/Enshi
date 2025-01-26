FROM node:18-alpine as builder

WORKDIR /app

ENV VITE_ENV=docker
COPY package*.json ./
RUN npm ci
COPY . .
RUN npm run build


FROM nginx:alpine

COPY --from=builder /app/dist /usr/share/nginx/html
COPY ./nginx/nginx.conf /etc/nginx/conf.d/default.conf

# Certificates
COPY ./nginx/fullchain.pem /etc/nginx/ssl/
COPY ./nginx/privkey.pem /etc/nginx/ssl/

EXPOSE 80
EXPOSE 443