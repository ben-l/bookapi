FROM golang:alpine
COPY . /bookapi
WORKDIR /bookapi
EXPOSE 8855
RUN apk update && apk add git
RUN go get -d github.com/gorilla/mux
CMD ["go", "run", "main.go"]
