FROM postgres:11

ADD ./1_init_product.sql /docker-entrypoint-initdb.d/
ADD ./2_insert_dummy.sql /docker-entrypoint-initdb.d/
