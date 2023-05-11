FROM --platform=linux/amd64 node:16-alpine AS build

ENV PORT=8081
ENV NODE_ENV production

WORKDIR /usr/src/app/

COPY ./package*.json ./

RUN apk add --no-cache python3 make g++ && npm i --omit=dev

COPY ./ ./

EXPOSE ${PORT}
ENTRYPOINT [ "npm", "run", "start" ]