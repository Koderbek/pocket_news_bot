FROM golang:1.22-alpine

RUN apk add --update --no-cache make tzdata busybox sqlite3

WORKDIR /usr/local/src
RUN mkdir -p ./.bin && chmod -R 777 ./.bin
COPY --chown=1000:1000 . .

RUN make build_clean
RUN echo "0 0 * * * /usr/local/src/.bin/clean_sent_news" >> /etc/crontabs/root

RUN make build_import
RUN echo "0 0,12 * * * /usr/local/src/.bin/import_blocked_resources" >> /etc/crontabs/root

RUN make build_sender
RUN echo "0,30 * * * * /usr/local/src/.bin/message_sender" >> /etc/crontabs/root

CMD ["sh", "-c", "crond -f && make run_bot"]
