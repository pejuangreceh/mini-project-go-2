FROM golang

#Set working directory didalam docker
WORKDIR /app

#copy go.mod dan go.sum
COPY go.mod .
COPY go.sum .

#Download dependensi Go
RUN go mod download

#copy semua file aplikasi kedalam container
COPY . .

#Build aplikasi Go
RUN go build -o crud_api

EXPOSE 8080

CMD ["./crud_api"]