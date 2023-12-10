FROM golang:1.17-alpine

ENV TZ Europe/Moscow
RUN apk add --update --no-cache make tzdata

WORKDIR /usr/local/src
COPY . .

CMD ["make", "run"]
