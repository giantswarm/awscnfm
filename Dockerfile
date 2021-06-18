FROM alpine:3.14.0

ADD ./awscnfm /usr/local/bin/awscnfm

ENTRYPOINT ["/usr/local/bin/awscnfm"]
