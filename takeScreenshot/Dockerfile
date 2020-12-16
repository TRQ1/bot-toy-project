FROM node:10-alpine

ENV CHROME_BIN=/usr/bin/chromium-browser
RUN echo @edge http://nl.alpinelinux.org/alpine/edge/community >> /etc/apk/repositories && \
    echo @edge http://nl.alpinelinux.org/alpine/edge/main >> /etc/apk/repositories && \
    apk add --no-cache \
    chromium@edge \
    terminus-font \
    nss@edge

WORKDIR /usr/src/app

COPY . .

RUN npm install

EXPOSE 8080
CMD [ "npx", "babel-node", "init.js" ]
