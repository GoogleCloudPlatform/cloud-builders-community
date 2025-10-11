FROM gcr.io/cloud-builders/go AS build-env

WORKDIR /build
COPY . .

RUN go mod download
RUN go build -v -o ./main .

FROM alpine:latest

RUN apk add --no-cache ca-certificates

COPY --from=build-env /build/main /cancelot

ENTRYPOINT [ "/cancelot" ]
