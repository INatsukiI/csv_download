FROM node:18.17.0-bullseye-slim

WORKDIR /app

RUN apt-get update  && \
    npm install -g npm && \
    npm install -g @vue/cli

EXPOSE 3000