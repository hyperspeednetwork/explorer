#FROM golang:1.13.0
#
#RUN mkdir -p /go/src/github.com
#WORKDIR /go/src/github.com
#COPY . /go/src/github.com
#RUN go get github.com/astaxie/beego && go get github.com/beego/bee
#EXPOSE 8080
#CMD ["bee","run"]
