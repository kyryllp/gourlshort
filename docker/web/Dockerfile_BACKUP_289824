<<<<<<< HEAD
FROM golang

EXPOSE 8080

RUN mkdir /go/src/gourlshort
ADD . /go/src/gourlshort/
=======
FROM golang:latest

EXPOSE 8080

ADD . /go/src/gourlshort
WORKDIR /go/src/gourlshort
>>>>>>> feature/docker-deployment

RUN go get github.com/gorilla/mux
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/joho/godotenv

<<<<<<< HEAD
WORKDIR /go/src/gourlshort

RUN go build gourlshort
RUN cp gourlshort /
RUN chmod +x /gourlshort
ENTRYPOINT ["/gourlshort"]
=======
RUN go build
RUN cp ./gourlshort /

RUN chmod +x /gourlshort
ENTRYPOINT ["/gourlshort"]
>>>>>>> feature/docker-deployment
