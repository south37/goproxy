FROM golang:1.14.6

COPY . /app

WORKDIR /app

RUN make

EXPOSE 8080

CMD ["/app/goproxy"]
