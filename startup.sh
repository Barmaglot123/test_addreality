#!/bin/sh
gosu postgres pg_ctl -D "$PGDATA" -w start
redis-server $INSTALL_DIR/redis.conf
./test_task