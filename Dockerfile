FROM golang:1.17-alpine AS go_builder

WORKDIR /restService

COPY ./src .

RUN apk add gcc musl-dev && \
    go mod tidy && \
    go get -d && \
    go test ./... -cover && \
    go build .


FROM alpine:3.14

COPY --from=go_builder /restService/tkRestService .

EXPOSE 8084

ENTRYPOINT [ "./tkRestService" ]
