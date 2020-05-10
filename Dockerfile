FROM golang:1.14 as build-env

WORKDIR /go/src/github.com/adelowo/queryapp
ADD . /go/src/github.com/adelowo/queryapp

ENV GO111MODULE=on

RUN go mod download
RUN go mod verify
RUN go install ./cmd

FROM gcr.io/distroless/base
COPY --from=build-env /go/bin/cmd /
CMD ["/cmd"]
