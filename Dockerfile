FROM golang:latest

LABEL maintainer="SpartanerSpaten <revol-xut@protonmail.com"

WORKDIR /app

RUN apt-get install git

RUN go get -u github.com/go-sql-driver/mysql

COPY . .
COPY ./app /app

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
