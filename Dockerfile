FROM golang:alpine3.11 as build-step

RUN apk add --update --no-cache ca-certificates git
COPY . /workdir
WORKDIR /workdir/app
RUN CGO_ENABLED=0 GOOS=linux go build -o main

FROM alpine:3.12
COPY --from=build-step /workdir/app/main /app/main
ENTRYPOINT ["/app/main"]
