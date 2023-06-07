FROM golang

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o crud_api

CMD ["./crud_api"]