#!/bin/bash
set -e
# скрипт на создание бд и пользователя при запуске
psql -v ON_ERROR_STOP=1 --username "$DB_USER" --dbname "$DB_NAME" <<-EOSQL
    DO
    \$\$
    BEGIN
        IF NOT EXISTS (SELECT FROM pg_catalog.pg_user WHERE usename = '$DB_USER') THEN
            CREATE USER "$DB_USER" WITH PASSWORD '$DB_PASSWORD';
        END IF;
    END
    \$\$;
    DO
    \$\$
    BEGIN
        IF NOT EXISTS (SELECT FROM pg_catalog.pg_database WHERE datname = '$DB_NAME') THEN
            CREATE DATABASE "$DB_NAME";
        END IF;
    END
    \$\$;
    GRANT ALL PRIVILEGES ON DATABASE "$DB_NAME" TO "$DB_USER";
EOSQL
