APP := gateway
GO := GOOS=linux go
IMAGE := microservicedemo/${APP}

build:
	${GO} build -o bundles/${APP} .

image: build
	docker build -t ${IMAGE} .

run:
	-docker service rm ${APP}
	docker service create --name ${APP} --network demo -p 9091:9091 ${IMAGE}
