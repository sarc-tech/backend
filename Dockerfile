FROM golang:1.23.4

WORKDIR /go/src/back
COPY . /go/src/back
RUN go mod download
RUN make build

EXPOSE 80/tcp

CMD [ "/go/src/back/bin/back" ]