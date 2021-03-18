# go-zero框架的脚手架


## 环境准备

```bash
# 安装 go-zero
go get -u github.com/tal-tech/go-zero
 
# 安装 go-ctl工具
go get -u github.com/tal-tech/go-zero/tools/goctl

# 安装完成后，在命令行执行，确保指令执行成功
goctl  -v

```


### 项目初始化

```bash
#创建项目
goctl api new zero

cd zero
go mod init
go mod tidy


go run zero.go -f etc/zero-api.yaml

```

访问 [http://localhost:8888/from/you](http://localhost:8888/from/you) 可看到返回值

```json
{
  message: ""
}
```

### 生成swagger

```bash
go get -u github.com/zeromicro/goctl-swagger
```

### 添加grpc支持

