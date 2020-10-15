FROM golang

EXPOSE 8080

RUN mkdir /go/src/gourlshort
ADD . /go/src/gourlshort/

RUN go get github.com/gorilla/mux
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/joho/godotenv

WORKDIR /go/src/gourlshort

RUN go build gourlshort
RUN cp gourlshort /
RUN chmod +x /gourlshort
ENTRYPOINT ["/gourlshort"]