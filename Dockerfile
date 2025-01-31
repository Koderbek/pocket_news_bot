FROM golang:1.17-alpine

ENV TZ Europe/Moscow
RUN apk add --update --no-cache make tzdata

WORKDIR /usr/local/src
RUN mkdir -p ./.bin && chmod -R 777 ./.bin
COPY --chown=1000:1000 . .

CMD ["sh", "-c", "make run_import & make run_bot"]
