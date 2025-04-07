FROM golang:1.22-alpine

RUN apk add --update --no-cache make tzdata busybox sqlite

WORKDIR /usr/local/src
RUN mkdir -p ./backup && chmod -R 777 ./backup
RUN mkdir -p ./logs && chmod -R 777 ./logs
RUN mkdir -p ./.bin && chmod -R 777 ./.bin
COPY --chown=1000:1000 . .

RUN make build_clean
RUN echo "0 3 * * * /usr/local/src/.bin/cleaner" >> /etc/crontabs/root

RUN make build_backup
RUN echo "0 4 * * * /usr/local/src/.bin/backup" >> /etc/crontabs/root

RUN make build_import
RUN echo "0 2,14 * * * /usr/local/src/.bin/importer" >> /etc/crontabs/root

RUN make build_sender
RUN echo "0,30 * * * * /usr/local/src/.bin/sender" >> /etc/crontabs/root

CMD ["sh", "-c", "crond -f & make run_bot"]
