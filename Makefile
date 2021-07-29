composeup:
	docker-compose up
composedown:
	docker-compose down
	
# composebuildgoaccountsrv:
# 	docker-compose build go-account
# composerestartpromocli:
# 	docker-compose rm -fsv promotioncli
# 	docker-compose up promotioncli
# docbuildgoaccountsrv:
# 	docker build -t go-account -f /go-account/Dockerfile .

# docrungoaccountsrv:
# 	docker run -p 7070:7070 --name goacountsrvcont go-account
# # docbuildpromocli:
# # 	docker build -t promocli -f /promotion/DockerfileCli .
# # docrunpromocli:
# # 	docker run -p 50051:50051 --name promoclicont promocli

# rungoaccountsrv:
# 	go run go-account/server/app.go
# # runpromocli:
# # 	go run promotion/client/promotionClient.go

# # genpromotionproto:
# # 	protoc --proto_path=$$GOPATH/src:. --micro_out=source_relative:.. --go_out=. --go_opt=paths=source_relative promotion/promotion.proto

startpgdb:
	docker-compose up gcref-pgdb
