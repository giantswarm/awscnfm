FROM alpine:3.16.0

ADD ./awscnfm /usr/local/bin/awscnfm

ENTRYPOINT ["/usr/local/bin/awscnfm"]
