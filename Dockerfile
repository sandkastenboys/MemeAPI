FROM golang:latest

LABEL maintainer="SpartanerSpaten <revol-xut@protonmail.com"

WORKDIR /app

#COPY go.mod go.sum ./

#Downloads Dependecy
RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
