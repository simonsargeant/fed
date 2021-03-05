FROM alpine:latest

ENTRYPOINT ["fed"]

COPY bin/fed /

RUN install /fed /usr/local/bin


