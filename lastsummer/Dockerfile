FROM golang:latest

WORKDIR $GOPATH/src/lastsummer
COPY . $GOPATH/src/lastsummer
RUN go build main.go

EXPOSE 8012
ENTRYPOINT ["./lastsummer"]