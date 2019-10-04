FROM nginx:1.17-alpine

WORKDIR /usr/share/nginx/html

COPY . /app

RUN apk add --update nodejs nodejs-npm \
    && cd /app \
    && npm i \
    && npm run build \
    && mv dist/* /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]

