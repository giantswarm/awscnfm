FROM alpine:3.13.5

ADD ./awscnfm /usr/local/bin/awscnfm

ENTRYPOINT ["/usr/local/bin/awscnfm"]
