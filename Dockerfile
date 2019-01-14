FROM golang:1.11-stretch

RUN mkdir /migrate
ADD cmd/migrations /migrate/
WORKDIR /migrate

RUN go build -o migrate .

CMD ["/migrate/migrate"]
