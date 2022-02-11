FROM golang:1.13 as build
COPY forward/ /fwd
WORKDIR /fwd
RUN go build -ldflags "-linkmode external -extldflags -static"

FROM scratch
COPY --from=build /fwd/forward forward
COPY --from=build /fwd/index.html index.html
CMD ["./forward"]
