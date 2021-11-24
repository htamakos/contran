package translater

import (
	"bytes"
	"testing"
)

func TestTraslate(t *testing.T) {
	input := []byte(`
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
      "cpu": 73,
      "memory": 128,
      "memoryReservation": 128,
      "user": "php",
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
    },
    {
      "name": "nginx-proxy",
      "image": "nginx",
      "essential": true,
      "memory": 128,
      "portMappings": [
        {
          "hostPort": 80,
          "containerPort": 80
        }
      ],
      "links": ["php-app"],
      "mountPoints": [
        {
          "sourceVolume": "php-app",
          "containerPath": "/var/www/html",
          "readOnly": true
        },
        {
          "sourceVolume": "nginx-proxy-conf",
          "containerPath": "/etc/nginx/conf.d",
          "readOnly": true
        },
        {
          "sourceVolume": "awseb-logs-nginx-proxy",
          "containerPath": "/var/log/nginx"
        }
      ]
    }
  ]
}`)

	expect := `version: "2.4"
services:
  nginx-proxy:
    image: nginx
    mem_limit: 128m
    links:
    - php-app
    ports:
    - 80:80
    volumes:
    - /var/app/current/php-app:/var/www/html:ro
    - /var/app/current/proxy/conf.d:/etc/nginx/conf.d:ro
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
    cpu_shares: 73
    mem_limit: 128m
    mem_reservation: 128m
    user: php
    privileged: true
    volumes:
    - /var/app/current/php-app:/var/www/html:ro
    ulimits:
      nofile:
        soft: 65535
        hard: 65535
`
	sourceTranslater := &Eb{}
	targetTranslater := &Composer{}

	buf := new(bytes.Buffer)
	tran := NewTranslateService(sourceTranslater, targetTranslater, buf)
	tran.Translate(input)
	if buf.String() != expect {
		t.Fatal()
	}
}
