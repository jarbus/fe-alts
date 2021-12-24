FROM alpine/git as repo
WORKDIR /app
RUN git clone --depth 1 https://codeberg.org/teddit/teddit.git .

FROM node:14-alpine
RUN apk --no-cache add ffmpeg
USER node
COPY --chown=node --from=repo /app /teddit
WORKDIR /teddit
RUN npm install --production --no-optional
# TODO: Move into docker-compose
COPY --chown=node ./config.js ./config.js
CMD ["npm", "start"]
