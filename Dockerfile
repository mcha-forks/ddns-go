FROM scratch

ENV UPDATE_INTERVAL="300"

COPY dnsd /usr/bin/dnsd

ENTRYPOINT ["/usr/bin/dnsd"]
