FROM node:18.20.2-alpine

WORKDIR /bot

COPY package*.json ./

RUN npm install --only=production

COPY ./ ./

ENTRYPOINT [ "node","./bot.js" ]