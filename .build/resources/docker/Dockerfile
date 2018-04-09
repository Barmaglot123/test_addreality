FROM golang:1.9-alpine3.6

ENV LANG en_US.utf8
ENV GOOS linux
ENV GOARCH amd64

# Install Redis
RUN apk --update add redis

# Install PGSQL
COPY ./1_migration.sql /db/1_migration.sql
ENV PGDATA="/var/lib/postgresql/data"
ADD https://github.com/tianon/gosu/releases/download/1.2/gosu-amd64 /usr/local/bin/gosu
RUN echo "@edge http://nl.alpinelinux.org/alpine/edge/main" >> /etc/apk/repositories && \
    apk update && \
    apk add "libpq@edge<9.7" "postgresql-client@edge<9.7" "postgresql@edge<9.7" "postgresql-contrib@edge<9.7"&& \
    mkdir -p $PGDATA && \
    chmod +x /usr/local/bin/gosu && \
    rm -rf /var/cache/apk/*

RUN chown -R postgres:postgres $PGDATA
RUN mkdir -p /run/postgresql
RUN chown -R postgres:postgres /run/postgresql

USER postgres
RUN initdb && \
    sed -ri "s/^#(listen_addresses\s*=\s*)\S+/\1'*'/" $PGDATA/postgresql.conf && \
    pg_ctl -w start && \
    psql --command "CREATE USER test_user WITH SUPERUSER PASSWORD 'rUeIZVWr';" &&\
    psql --command "CREATE DATABASE test_task ;" &&\
    psql -h localhost -p 5432 -U postgres -d test_task -a -f /db/1_migration.sql &&\
    pg_ctl stop

RUN echo "host all all 172.17.0.1/32 md5" >> /var/lib/postgresql/data/pg_hba.conf

ENV INSTALL_DIR=$GOPATH/src/untitled3

USER root
WORKDIR $INSTALL_DIR

ADD . $INSTALL_DIR

EXPOSE 5432

CMD ["/bin/sh", "/go/src/untitled3/startup.sh"]