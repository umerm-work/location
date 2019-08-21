FROM golang

RUN mkdir -p /go/src/github.com/fb-investments

ADD . /go/src/github.com/fb-investments
WORKDIR /go/src/github.com/fb-investments/aone-location-service

RUN go get  -t -v ./...
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run github.com/fb-investments/aone-location-service/cmd  -watch github.com/fb-investments/aone-location-service
