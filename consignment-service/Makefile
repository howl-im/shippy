all: build run

build:
	@echo "Building~~~~~~~~~~~~~~~~~~~~~~~~"
# for grpc	
# protoc -I. --go_out=plugins=grpc:$(GOPATH)/src/shippy/consignment-service proto/consignment/consignment.proto

# for go-micro
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/shippy/consignment-service proto/consignment/consignment.proto

# 编译 consignment-service CGO_ENABLED=0
	GOOS=linux GOARCH=amd64 go build

# 根据当前目录下的 Dockerfile 生成名为 consignment-service 的镜像
	docker build -t consignment-service .

run:
# 在 Docker alpine 容器的 50001 端口上运行 consignment-service 服务
# 可添加 -d 参数将微服务放到后台运行
# 使用 -p 参数，将容器的 50051 端口映射到本机的 50051 端口
# -p 本机端口:容器端口
# for grpc
# docker run -p 50051:50051 consignment-service
	
# -e 选项用于设置镜像中的环境变量, 其中	-e MICRO_REGISTRY=mdns 会使 
#  go-micro 在本地使用 mdns 多播作为服务发现的中间层。
#  在生产环境中一般会使用 Consul 或 Etcd 代替 mdns 做服务发现。
	@echo "Runing~~~~~~~~~~~~~~~~~~~~~~~~"
	docker run -p 50051:50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns \
		consignment-service

.PHONY: run build all 
