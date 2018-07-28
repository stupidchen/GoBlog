FROM daocloud.io/golang:1.9

MAINTAINER Mike Chen <stupidchen@foxmail.com>

COPY . /go/src/blog

RUN go get -v -t github.com/jinzhu/gorm
RUN go get -v -t github.com/go-sql-driver/mysql
RUN go get -v -t github.com/google/uuid
RUN go install blog

CMD [ "blog" ]
