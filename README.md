# 看房网

基于Golang+Kratos+MySQL+Redis+Kafka+elk+Opentracing实现的微服务项目


## 技术点

- [x] 微服务架构（BFF、Service、Admin、Job、Task 分模块）
- [x] API 设计（包括 API 定义、错误码规范、Error 的使用）
- [x] gRPC 的使用
- [x] Go 项目工程化（项目结构、DI、代码分层、ORM 框架）
- [ ] 并发的使用（errgroup 的并行链路请求）
- [ ] 微服务中间件的使用（ELK、Opentracing、Prometheus、Kafka）
- [ ] 缓存的使用优化（一致性处理、Pipeline 优化）

## 开发命令

1. 生成proto代码

    ```shell
    kratos proto client api/account/service/v1/account.proto
    ```

## docker部署

1. 创建volume

```shell
docker volume create kfan-db
```

## 调试

1. SaveHistory
```shell

grpcurl -d '{"obj_id": 2,"obj_type": "house","user_id": 1}' -plaintext 127.0.0.1:9002 api.history.service.v1.History/SaveHistory
```

2. GetHistory

```shell
grpcurl -d '{"user_id": 1}' -plaintext 127.0.0.1:9002 api.history.service.v1.History/GetHistory
```