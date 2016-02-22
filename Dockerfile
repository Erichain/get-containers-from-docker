FROM golang

ADD . /go/src/ ~/Developer/Workspace/get-containers-from-docker
RUN go install ~/Developer/Workspace/get-containers-from-docker
ENTRYPOINT /go/bin/basic_web_server

EXPOSE 8080