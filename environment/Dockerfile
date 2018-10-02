FROM golang:1.11-stretch

ENV GOPATH /usr/local

COPY requirements.apt /tmp/requirements.apt

RUN apt-get update && apt-get install -y aptitude
RUN aptitude full-upgrade -y && aptitude install -y $(tr '\n' ' ' < /tmp/requirements.apt)

# Install golang dep manager
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# Create the user/group for the running stuff
RUN useradd -s /bin/bash -mk /etc/skel savsgio

WORKDIR /code

USER savsgio

CMD ["ls -las"]
