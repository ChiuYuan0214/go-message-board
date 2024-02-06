FROM mysql:latest

WORKDIR /db

COPY init_script.sql /docker-entrypoint-initdb.d/ddl_script.sql

EXPOSE 3306