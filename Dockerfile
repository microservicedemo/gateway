FROM alpine:edge

ADD bundles/gateway /usr/bin/gateway

EXPOSE 9091
CMD /usr/bin/gateway
