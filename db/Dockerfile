FROM mysql:8.0
COPY ./sql/schema_v2.sql /docker-entrypoint-initdb.d/
COPY ./sql/test_seed_v2.sql /docker-entrypoint-initdb.d/
VOLUME ["/var/lib/mysql"]
EXPOSE 3306
