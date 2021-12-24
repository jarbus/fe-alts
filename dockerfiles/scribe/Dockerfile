FROM alpine/git as repo
WORKDIR /scribe
RUN git clone --depth 1 https://git.sr.ht/~edwardloveall/scribe .

FROM node:16-alpine as node_build
WORKDIR /tmp_build
COPY --from=repo /scribe/package.json .
COPY --from=repo /scribe/yarn.lock .
RUN yarn install --no-progress --frozen-lockfile
COPY --from=repo /scribe/webpack.mix.js .
COPY --from=repo /scribe/src ./src
RUN yarn prod

FROM crystallang/crystal:1.2.1-alpine as lucky_build
ENV SKIP_LUCKY_TASK_PRECOMPILATION="1"
RUN apk add yaml-static
WORKDIR /tmp_build
COPY --from=repo /scribe/shard.* ./
RUN  shards install --production
COPY --from=repo /scribe .
COPY --from=node_build /tmp_build/public/mix-manifest.json public/mix-manifest.json
RUN crystal build --static src/start_server.cr
RUN crystal build --static tasks.cr -o run_task

FROM alpine
ARG PUID=1000
ARG PGID=1000
RUN addgroup -g ${PGID} -S lucky && \
    adduser -u ${PUID} -S lucky -G lucky
WORKDIR /home/lucky/app
COPY --chown=lucky:lucky --from=node_build /tmp_build/public public
COPY --chown=lucky:lucky --from=lucky_build /tmp_build/start_server start_server
COPY --chown=lucky:lucky --from=lucky_build /tmp_build/run_task run_task
COPY --chown=lucky:lucky --from=repo  /scribe/script/docker_entrypoint ./
RUN mkdir ./config
RUN chown -R lucky /home/lucky
USER lucky
# TODO: Replace with mount in docker-compose
COPY --chown=lucky:lucky ./watch.yml ./config/watch.yml
CMD ["./start_server"]

