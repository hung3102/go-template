FROM node:20.16.0-alpine

RUN apk --no-cache add openjdk11-jre bash curl openssl && \
    npm cache clean --force && \
    npm i -g firebase-tools
RUN npm install -g firebase-tools