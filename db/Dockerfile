FROM mysql:8.0
COPY ./sql/schema_v2.sql /docker-entrypoint-initdb.d/
COPY ./sql/test_seed_v2.sql /docker-entrypoint-initdb.d/
EXPOSE 3306
