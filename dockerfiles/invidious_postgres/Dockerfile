FROM alpine/git as config
WORKDIR /invidious
RUN git clone --depth 1 https://github.com/iv-org/invidious.git .

FROM postgres:10 as app
# config/sql is a folder.
COPY --from=config /invidious/config/sql /config/sql
COPY --from=config /invidious/docker/init-invidious-db.sh /docker-entrypoint-initdb.d/init-invidious-db.sh
