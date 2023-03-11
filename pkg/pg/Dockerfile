FROM golang:1.20-alpine

RUN go install github.com/rubenv/sql-migrate/...@latest
ADD migrations /migrations
ADD migrate.sh /migrate.sh
ADD dbconfig.yml /dbconfig.yml
RUN chmod +x /migrate.sh

ENTRYPOINT []
