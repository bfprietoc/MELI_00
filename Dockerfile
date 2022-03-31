FROM golang:1.18

LABEL mainteiner="Fabian Prieto <bfprietoc@gmail.com>"

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

ENV PORT 1234

RUN go build

CMD ["./meli_00"]