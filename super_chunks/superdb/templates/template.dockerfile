FROM redis:alpine

COPY ./conf/redis.conf /usr/local/etc/redis/redis.conf

CMD [ "redis-server", "/usr/local/etc/redis/redis.conf" ]

EXPOSE ${redis_port}

VOLUME [ "/data" ]