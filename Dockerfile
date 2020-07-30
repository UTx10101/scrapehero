FROM node:alpine AS JS_BUILD
COPY web /web
WORKDIR web
RUN npm install && npm run build

FROM golang:1.14-alpine AS GO_BUILD
RUN apk update && apk add build-base
COPY api /api
WORKDIR /api
RUN go build -o /go/bin/api

FROM alpine:latest
COPY --from=JS_BUILD /web/build* ./web/
COPY --from=GO_BUILD /go/bin/api ./
CMD ./api
