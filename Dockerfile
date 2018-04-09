FROM golang:1.9-alpine3.6

ENV LANG en_US.utf8

# Install Redis
RUN apk --update add redis
RUN adduser -D red && \
    passwd -u red

# Install PGSQL
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
    psql --command "CREATE USER test_user WITH PASSWORD 'rUeIZVWr';" &&\
    psql --command "CREATE DATABASE test_task ;" &&\
    pg_ctl stop

# Install
ENV INSTALL_DIR="/opt/test_task"
USER root
WORKDIR $INSTALL_DIR

ADD ./.build/ $INSTALL_DIR
RUN chmod +x test_task

# Run
CMD ["/bin/sh", "/opt/test_task/startup.sh"]