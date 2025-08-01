FROM alpine:3.22.1

WORKDIR /var/lib/effective-dollop

RUN apk add gcompat

RUN addgroup -S evdp && adduser evdp -S evdp -G evdp


COPY ./build/effective-dollop /effective-dollop
COPY ./dev/cert ./dev/cert
COPY ./config.toml ./

RUN chown -R evdp:evdp /var/lib/effective-dollop

USER evdp
EXPOSE 8643
CMD ["/effective-dollop"]
