# sheepim-API-gateway

## 初始化项目
```bash
hz update -idl idl/user.thrift   
hz update -idl idl/auth.thrift   
hz update -idl idl/project.thrift   

cd biz/infra/container
wire
```
## 配置文件示例

```yml

server:
  listen-address: 0.0.0.0:8888
  service-name: sheepim-user-controller

etcd:
  endpoint: 127.0.0.1:2379

open-telemetry:
  endpoint: 127.0.0.1:4317

database:
  username: xxx
  password: xxx
  database: xxx
  address: xxx
  port: xxx

```

## 开发环境

```bash
export ENV=development
```

## 生产环境

```bash
export ENV=production
```