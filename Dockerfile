FROM debian:bullseye

WORKDIR /app

COPY main ./

CMD /app/main