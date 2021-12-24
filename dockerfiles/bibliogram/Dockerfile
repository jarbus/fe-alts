FROM node:14-alpine as build
RUN apk --no-cache add git python3 make g++
WORKDIR /app
RUN git clone --depth 1 https://git.sr.ht/~cadence/bibliogram .
RUN npm install --production --no-optional

FROM node:14-alpine as app
RUN apk add --no-cache graphicsmagick
# Non-root user
USER node
COPY --chown=node --from=build /app /home/node/app
WORKDIR /home/node/app
RUN mkdir -p ./db
COPY --chown=node ./config.js ./config.js
CMD ["npm", "start"]
