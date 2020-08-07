IMG_NAME = alw_golang_redis_geoloc_img
NET_NAME = alw_golang_redis_geoloc_net
REDIS_NAME = alw_golang_redis_geoloc_redis
SERVER_NAME = alw_golang_redis_geoloc_test
MYGOPATH = ${GOPATH}

#for first time setup use make all
all: create-network run-redis connect-redis build run

build:
	cd "$(MYGOPATH)\src\albert02lowis\alw-golang-redis-geoloc"
	docker build \
	-t $(IMG_NAME) \
	.

rerun: stoprm-server run

run:
	docker run \
	--net $(NET_NAME) \
	--name $(SERVER_NAME) \
	-p 8080:8080 \
	-d \
	$(IMG_NAME) alw-golang-redis-geoloc -redisAddr=172.19.0.22:6379

stop:
	docker stop $(SERVER_NAME)

start:
	docker start $(SERVER_NAME)

stoprm-server: stop
	docker rm $(SERVER_NAME)

run-redis:
	docker run \
	--name $(REDIS_NAME) \
	-d redis

connect-redis:
	docker network connect \
	--ip 172.19.0.22 \
	$(NET_NAME) \
	$(REDIS_NAME)

stop-redis:
	docker stop $(REDIS_NAME)

stoprm-redis: stop-redis
	docker rm $(REDIS_NAME)

start-redis:
	docker start $(REDIS_NAME)

create-network:
	docker network create \
	--attachable \
	-d bridge \
	--subnet 172.19.0.0/16 \
	$(NET_NAME)

test:
	cd "$(MYGOPATH)\src\albert02lowis\alw-golang-redis-geoloc"
	go test

cleanall: stoprm-redis stoprm-server
	docker network rm $(NET_NAME)
	docker rmi $(IMG_NAME)
