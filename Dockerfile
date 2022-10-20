FROM golang:latest
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .
EXPOSE 4000
CMD ["/app/main"]