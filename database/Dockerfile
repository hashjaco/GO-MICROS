FROM mysql:latest
COPY migrations/user_schema.up.sql /docker-entrypoint-initdb.d/user_schema.up.sql
COPY migrations/user_schema.down.sql /docker-entrypoint-initdb.d/user_schema.down.sql