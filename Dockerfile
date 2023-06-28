FROM ubuntu:latest
# Install supervisor
RUN apt-get update && apt-get install -y supervisor golang-go ca-certificates git
RUN apt-get update && apt-get install -y python3 python3-pip
#RUN python3 -m pip install --upgrade pip
RUN pip install notebook
# Add supervisor configuration file
COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf

# Run supervisor
# CMD ["/usr/bin/supervisord"]
COPY . /goproject
WORKDIR /goproject
RUN go build .
CMD ["/goproject/goproject"]
