FROM alpine:3.14.2

ADD ./awscnfm /usr/local/bin/awscnfm

ENTRYPOINT ["/usr/local/bin/awscnfm"]
