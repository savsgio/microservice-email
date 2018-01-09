FROM golang:1.9-stretch

COPY requirements.apt /tmp/requirements.apt

RUN apt-get update && apt-get install -y aptitude
RUN aptitude full-upgrade -y && aptitude install -y $(tr '\n' ' ' < /tmp/requirements.apt)

WORKDIR /code

ENV GOPATH /code

CMD ["ls -las"]
