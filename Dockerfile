FROM daocloud.io/golang:1.3

MAINTAINER Mike Chen <stupidchen@foxmail.com>

COPY . /go/src/blog

RUN go get -v -d .
RUN go install blog 

CMD [ "blog" ]
