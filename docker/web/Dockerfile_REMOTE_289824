FROM golang:latest

EXPOSE 8080

ADD . /go/src/gourlshort
WORKDIR /go/src/gourlshort

RUN go get github.com/gorilla/mux
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/joho/godotenv

RUN go build
RUN cp ./gourlshort /

RUN chmod +x /gourlshort
ENTRYPOINT ["/gourlshort"]
