FROM alpine:3.6 as builder

COPY deploy/_output/api/api /usr/local/bin/api

RUN apk --no-cache add curl ca-certificates

EXPOSE 3200
ENTRYPOINT ["api"]
