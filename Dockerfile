FROM golang:1.13.8-alpine

RUN set -x \
  && apk update \
  && apk add --update --no-cache vim git make musl-dev curl \
  && apk add --update --no-cache clang llvm9 binutils gcc

ADD . /visket-playground
WORKDIR /visket-playground

ENV PORT 8080
EXPOSE 8080

ENV GIN_MODE release

# compile visket

RUN git clone https://github.com/visket-lang/visket ./visket \
  && cd ./visket && make && cd .. \
  && cp ./visket/bin/visket ./static/ 

# compile playground

RUN go build -o main

CMD ["./main"]

