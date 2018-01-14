Microservice for send emails
============================

#### System Dependencies

- [Go v1.9.2](https://golang.org/dl/)
- [RabbitMQ](https://www.rabbitmq.com/)
- make

#### Go dependencies

- [fasthttp](https://github.com/valyala/fasthttp)
- [fasthttprouter](https://github.com/buaazp/fasthttprouter)
- [amqp](https://github.com/streadway/amqp)
- [gomail](https://github.com/go-gomail/gomail)
- [yaml](https://github.com/go-yaml/yaml)
- [go-logger](https://github.com/savsgio/go-logger)


## General

Optimize Go code:
```bash
make simplify
```

Download Go dependencies and build:
```bash
make
```

Run:

Copy `etc/config.yml` to `etc/config-dev.yml` *(this file not tracked in Git)*, modify each config and exec:
```bash
make run
```

***Note:*** If you want to use with Docker, make sure you have this rabbitmq host in `etc/config-dev.yml`:
```yaml
...
rabbitmq:
  host: rabbitmq_server
...
```

Install:
```bash
make install
```

Command line arguments:
- `-log-level`: Level of log that you want to show (default: *info*)
- `-config-file`:  Path of configuration file (default: */etc/microservice-email.yml*)

#### Run in Docker

#### Dependencies

- [Docker](https://www.docker.com/)
- [Docker-compose](https://docs.docker.com/compose/) *Recommended to install with `pip3` (python3)*

Run:
```bash
make docker_run
```

#### Others
**Feel free to contribute it or fork me...** :wink: