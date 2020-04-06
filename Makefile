prepare:
	go mod download

run:
	go build -o bin/main cmd/consumer/main.go
	./bin/main

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o bin/main cmd/consumer/main.go
	chmod +x bin/main

dkb:
	docker build -t consumer-superhero-chat .

dkr:
	docker run consumer-superhero-chat

launch: dkb dkr

cr-log:
	docker logs consumer-superhero-chat -f

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
	docker exec -it consumer-superhero-chat /bin/bash

db-ssh:
	docker exec -it db /bin/bash

es-ssh:
	docker exec -it es /bin/bash

PHONY: prepare build dkb dkr launch cr-log db-log es-log cr-ssh db-ssh es-ssh rmc rmi clear