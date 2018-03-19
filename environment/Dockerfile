FROM golang:1.10-stretch

COPY requirements.apt /tmp/requirements.apt

RUN apt-get update && apt-get install -y aptitude
RUN aptitude full-upgrade -y && aptitude install -y $(tr '\n' ' ' < /tmp/requirements.apt)

WORKDIR /code

# Install golang dep manager
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

ENV GOPATH /godeps

# Create the user/group for the running stuff
RUN useradd -s /bin/bash -mk /etc/skel savsgio

USER savsgio

CMD ["ls -las"]
