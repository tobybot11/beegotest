FROM golang:1.7.3

# COPY . /go/src/github.com/tobybot11/beegotest
#VOLUME ["$(pwd)", "/go/src/github.com/tobybot11/beegotest"]
WORKDIR /go/src/github.com/tobybot11/beegotest

RUN go get github.com/astaxie/beego && go get github.com/beego/bee

EXPOSE 8080

CMD ["bee", "run"]

