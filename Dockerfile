FROM golang:1.8

COPY . /go/src/github.com/turkenh/play-with-ansible

WORKDIR /go/src/github.com/turkenh/play-with-ansible

RUN go get -v -d ./...

RUN ssh-keygen -N "" -t rsa -f /etc/ssh/ssh_host_rsa_key >/dev/null

RUN CGO_ENABLED=0 go build -a -installsuffix nocgo -o /go/bin/play-with-ansible .


FROM alpine

RUN apk --update add ca-certificates
RUN mkdir -p /app/pwd

COPY --from=0 /go/bin/play-with-ansible /app/play-with-ansible
COPY --from=0 /etc/ssh/ssh_host_rsa_key /etc/ssh/ssh_host_rsa_key
COPY ./www /app/www

WORKDIR /app
CMD ["./play-with-ansible"]

EXPOSE 3000
