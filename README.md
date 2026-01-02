# sheepim-API-gateway

## 初始化项目
```bash
hz update -idl idl/user.thrift   
hz update -idl idl/auth.thrift   
hz update -idl idl/project.thrift   
hz update -idl idl/feedback.thrift   
hz update -idl idl/room.thrift
hz update -idl idl/file.thrift
hz update -idl idl/aichat.thrift



cd biz/infra/container
wire
```
## 配置文件示例

```yml


server:
  listen-address: 0.0.0.0:9100
  service-name: sheepim-api-gateway
etcd:
  endpoint:
    - localhost:2379

open-telemetry:
  endpoint: localhost:4417

rpc:
  user-service-name: sheepim-user-service
  auth-service-name: sheepim-auth-service
  project-service-name: personal-project-service
  room-service-name: sheepim-room-service
  feedback-service-name: personal-feedback-service
  file-service-name: personal-file-service

```

## 开发环境

```bash
export ENV=development
```

## 生产环境

```bash
export ENV=production
```