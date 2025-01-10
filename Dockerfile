FROM golang:1.23.4

WORKDIR /go/src/back
COPY . /go/src/back
RUN go mod download
RUN apt-get update && \
    apt-get install wget postgresql-client --yes && \
    mkdir --parents ~/.postgresql && \
    wget "https://storage.yandexcloud.net/cloud-certs/CA.pem" \
         --output-document ~/.postgresql/root.crt && \
    chmod 0655 ~/.postgresql/root.crt
RUN make build

EXPOSE 80/tcp

CMD [ "/go/src/back/bin/back" ]