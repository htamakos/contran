# contran
contran is a small utility to transform various docker container formats to one another.

The following are currently supported by contran

|From|To|
|:---|:-|
|Elastic Beanstalk|Docker-compose configuration|

# Usage

```sh
$ cat Dockerrun.aws.json
{
  "AWSEBDockerrunVersion": 2,
  "volumes": [
    {
      "name": "php-app",
      "host": {
        "sourcePath": "/var/app/current/php-app"
      }
    },
    {
      "name": "nginx-proxy-conf",
      "host": {
        "sourcePath": "/var/app/current/proxy/conf.d"
      }
    }
  ],
  "containerDefinitions": [
    {
      "name": "php-app",
      "image": "php:fpm",
      "environment": [
        {
          "name": "Container",
          "value": "PHP"
        }
      ],
      "essential": true,
      "memory": 128,
      "memoryReservation": 128,
      "privileged": true,
      "mountPoints": [
        {
          "sourceVolume": "php-app",
          "containerPath": "/var/www/html",
          "readOnly": true
        }
      ],
      "command": [
        "--path.procfs",
        "/host/proc",
        "--path.sysfs",
        "/host/sys",
        "--collector.filesystem.ignored-mount-points",
        "\"^/(sys|proc|dev|host|etc|rootfs/var/lib/docker|rootfs/var/run/docker)($|/)\""
      ],
      "ulimits": [
        {
          "name": "nofile",
          "hardLimit": 65535,
          "softLimit": 65535
        }
      ]
    }
  ]
}

$ cat Dockerrun.aws.json | contran
version: "2.4"
services:
  php-app:
    image: php:fpm
    environment:
      Container: PHP
    command:
    - --path.procfs
    - /host/proc
    - --path.sysfs
    - /host/sys
    - --collector.filesystem.ignored-mount-points
    - '"^/(sys|proc|dev|host|etc|rootfs/var/lib/docker|rootfs/var/run/docker)($|/)"'
    mem_limit: 128m
    mem_reservation: 128m
    privileged: true
    volumes:
    - /var/app/current/php-app:/var/www/html:ro
    ulimits:
      nofile:
        soft: 65535
        hard: 65535
```

## Installation

```sh
go get github.com/htamakos/contran
```
