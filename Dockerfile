FROM daocloud.io/golang:1.3

MAINTAINER Mike Chen <stupidchen@foxmail.com>

COPY . /go/src/blog

RUN go get -v -d .
RUN go get github.com/jinzhu/gorm
RUN go get github.com/go-sql-driver/mysql
RUN go install blog

CMD [ "blog" ]
