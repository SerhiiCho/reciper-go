FROM nginx:1.17-alpine

COPY . /app
COPY default.conf /etc/nginx/conf.d/default.conf

RUN apk add --update nodejs nodejs-npm \
    && cd /app \
    && npm i \
    && npm run build

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]

