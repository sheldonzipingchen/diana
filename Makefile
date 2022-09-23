BINARY="diana"
VERSION="2022092301"

default: help

help: ## 显示帮助信息
	@echo 'usage: make [targets]'
	@echo ''
	@echo 'targets:'
	@egrep '^(.+)\:\ .*##\ (.+)' ${MAKEFILE_LIST} | sed 's/:.*##/#/' | column -t -c 2 -s '#'

build: ## 构建应用程序
	go build -o bin/${BINARY}-${VERSION}-local diana/cmd
	@echo '本地应用程序构建完毕' 

build-linux: ## 构建Linux应用程序
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/${BINARY}-${VERSION}-linux-amd64 diana/cmd
	@echo 'linux应用程序构建完毕'

build-mac: ## 构建mac应用程序
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/${BINARY}-${VERSION}-darwin-amd64 diana/cmd/main
	@echo 'mac应用程序构建完毕'

build-windows: ## 构建windows应用程序
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/${BINARY}-${VERSION}-windows-amd64.exe diana/cmd/main
	@echo 'windows应用程序构建完毕'

deploy-data: clean build-linux ## 发布到生产环境
	scp bin/${BINARY}-${VERSION}-linux-amd64 FacialServer:/home/neo/apps/facial/bin/facial
	@echo '发布完毕'

deploy-api: clean build-linux ## 发布服务到生产环境
	scp bin/${BINARY}-api-${VERSION}-linux-amd64 FacialServer:/home/neo/apps/facial/bin/facial_api
	@echo '发布完毕'

deploy-bot: clean build-linux ## 发布服务到生产环境
	scp bin/${BINARY}-bot-${VERSION}-linux-amd64 FacialServer:/home/neo/apps/facial/bin/facial_bot
	@echo '发布完毕'

clean: ## 清除构建文件
	rm bin/*
	@echo '过去构建文件清理完毕'

doc: ## 生成 godocs 文档并开启本地文档服务
	godoc -http=:8085 -index

db-migrate: ## 运行数据库迁移命令
	@echo '还没实现'

swagger-ui: ## 运行本地 swagger-ui
	@echo '还没实现'

