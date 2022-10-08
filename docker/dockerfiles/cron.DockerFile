FROM golang:1.19.1-alpine3.16

WORKDIR $GOPATH/src/youtube_search_engine

COPY . .

RUN go mod download
RUN go build -o $GOPATH/bin/cron cmd/cron/main.go 

CMD [ "cron" ]