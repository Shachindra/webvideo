# Maintainer Shachindra

FROM golang:alpine AS build-app
RUN apk update
WORKDIR /app
COPY . .
RUN go build -o webvideo .

FROM alpine
WORKDIR /app
COPY --from=build-app /app/webvideo .
COPY --from=build-app /app/ui ./ui
RUN chmod +x ./webvideo
RUN apk update && apk add --no-cache bash openresolv bind-tools gettext
RUN ./webvideo
EXPOSE 9070