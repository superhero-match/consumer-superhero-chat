prepare:
	go get -u github.com/golang/dep/cmd/dep
	go get -u github.com/gin-gonic/gin
	go get -u golang.org/x/sys/unix
	go get -u github.com/jinzhu/configor
	go get -u github.com/go-sql-driver/mysql
	go get -u go.uber.org/zap
	go get -u github.com/segmentio/kafka-go
	go get -u golang.org/x/net/context
    go get -u github.com/go-redis/redis

run:
	go build -o bin/main cmd/consumer/main.go
	./bin/main

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o bin/main cmd/consumer/main.go
	chmod +x bin/main

init:
	dep init

deps:
	dep ensure -v

dkb:
	docker build -t consumer .

dkr:
	docker run consumer

launch: dkb dkr

cr-log:
	docker logs consumer -f

db-log:
	docker logs db -f

es-log:
	docker logs es -f

rmc:
	docker rm -f $$(docker ps -a -q)

rmi:
	docker rmi -f $$(docker images -a -q)

clear: rmc rmi

cr-ssh:
	docker exec -it consumer /bin/bash

db-ssh:
	docker exec -it db /bin/bash

es-ssh:
	docker exec -it es /bin/bash

PHONY: prepare build dkb dkr launch cr-log db-log es-log cr-ssh db-ssh es-ssh rmc rmi clear