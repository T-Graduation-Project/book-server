# book-server
图书服务
## HOW TO RUN
~~~bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
./borrow-server --gf.gcfg.file=config.toml

docker build -t csighub.tencentyun.com/urftian/tendoc-oms-server-go .
docker run --rm -it -p 8199:8199 csighub.tencentyun.com/urftian/tendoc-oms-server-go
~~~

## 开发常用命令
~~~bash
# proto 协议生成代码
protoc --gofast_out=plugins=grpc:. protocol/borrow.proto

# model 默认使用 config.toml的[database.default]
gf gen dao
gf gen dao -l "mysql:username:password@tcp(ip:port)/tablename"

# docker
docker build -t image_name .
docker run --rm -it -p 58080:58080 image_name
docker login 
docker push image_name
~~~

protoc 命令来自于 https://github.com/google/protobuf，可以产生序列化和反序列化的代码，无go相关代码。
protoc-gen-go插件则来自于https://github.com/golang/protobuf/protoc-gen-go， 可以产生go相关代码， 除上述序列化和反序列化代码之外， 还增加了一些通信公共库。
编译方法分别为：
~~~bash
protoc --go_out=. my.proto  
protoc --go_out=plugins=grpc:.  my.proto 
~~~
而grpc又来源于https://github.com/grpc/grpc-go, protoc和protoc-gen-go这两个工具都不在其中，所以玩grpc的时候，少不了protoc和protoc-gen-go.
