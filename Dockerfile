FROM debian:bullseye

WORKDIR /app

COPY *.* ./

ADD go1.21.0.linux-amd64.tar.gz /usr/local

CMD /usr/local/go/bin/go run main.go