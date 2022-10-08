FROM golang:1.19.1-alpine3.16

COPY . $GOPATH/src/youtube_search_engine

WORKDIR $GOPATH/src/youtube_search_engine

RUN go mod download
RUN go build -o $GOPATH/bin/youtube_search_engine cmd/youtube/main.go 

CMD [ "youtube_search_engine" ]