FROM golang:1.16.5-buster
WORKDIR /pickaxe
COPY . /pickaxe
RUN useradd gouser --home /pickaxe --uid 1000 
# Home directory for managing go imports
RUN chown -R gouser:gouser /pickaxe

USER gouser
RUN go mod download
RUN go build -o ./pickaxe cmd/pickaxe/pickaxe.go

CMD ["/pickaxe/pickaxe"]
