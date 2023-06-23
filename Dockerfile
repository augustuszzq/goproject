FROM ubuntu:latest
RUN apt-get update && apt-get install -y golang-go git
COPY . /goproject
WORKDIR /goproject
RUN go build .
CMD ["/goproject/goproject"]
