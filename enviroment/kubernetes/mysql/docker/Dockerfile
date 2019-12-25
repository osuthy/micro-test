FROM mysql:8

EXPOSE 3306
ENV MYSQL_ALLOW_EMPTY_PASSWORD=yes
ENV MYSQL_ROOT_USER=root
COPY ./init.sql /docker-entrypoint-initdb.d/
