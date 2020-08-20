FROM quay.io/giantswarm/alpine:3.11-giantswarm

USER giantswarm

ADD ./awscnfm /usr/local/bin/awscnfm

ENTRYPOINT ["awscnfm"]
