# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

ENV PORT=8080


RUN  mkdir -p /go/src \
&& mkdir -p /go/bin \
&& mkdir -p /go/pkg
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH

RUN mkdir -p $GOPATH/src/app 
ADD . $GOPATH/src/app

WORKDIR $GOPATH/src/app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY *.go ./
COPY .env ./

RUN go build -v -o /calendar-server

EXPOSE $PORT

CMD [ "/calendar-server" ]