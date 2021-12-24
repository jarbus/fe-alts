FROM nimlang/nim:alpine as build
RUN apk --no-cache add git libsass-dev libffi-dev openssl-dev openssh-client
WORKDIR /nitter
RUN git clone --depth 1 https://github.com/zedeus/nitter.git .
RUN nimble build -y -d:release --passC:"-flto" --passL:"-flto" \
	&& strip -s nitter \
	&& nimble scss

FROM alpine as app
RUN apk --no-cache add pcre-dev sqlite-dev
RUN addgroup -g 1000 nitter \
	&& adduser -u 1000 -G nitter -s /bin/sh -D nitter
USER nitter
WORKDIR /nitter
COPY --chown=nitter --from=build /nitter/nitter ./
COPY --chown=nitter --from=build /nitter/public ./public
COPY ./nitter.conf ./nitter.conf
CMD ["./nitter"]
