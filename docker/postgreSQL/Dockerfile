FROM postgres:15.2-alpine

ENV POSTGRES_USER=user
ENV POSTGRES_PASSWORD=password
ENV POSTGRES_DB=postgres
ENV TZ=Asia/Tokyo

COPY ./sql/init.sql /docker-entrypoint-initdb.d/