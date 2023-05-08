FROM golang:1.20
RUN mkdir /app
WORKDIR /app
ADD . /app
RUN go build -o main ./server.go
# FROM ubuntu
EXPOSE 8080
CMD /app/main