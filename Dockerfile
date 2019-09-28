FROM golang:1.10
RUN go get -u github.com/golang/dep/cmd/dep

WORKDIR /go/src/github.com/myriadeinc/pickaxe/

COPY . .

RUN dep ensure

RUN go build -o build/pickaxe

ENTRYPOINT [ "./build/pickaxe" ]