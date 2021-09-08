# rest-skeleton

Golang编写的一个rest API骨架项目

## 配置

目前的配置项从环境变量中读取，程序运行会导入`.env`文件中的环境变量

已使用的配置项

```yaml
# APP
APP_NAME
APP_DEBUG

  # GIN
GIN_ADDR
GIN_MODE

  #DATABASE
DATABASE_DSN

  # REDIS
REDIS_ADDR
REDIS_PASSWORD
REDIS_DATABASE
REDIS_DIAL_TIMEOUT
```

## 鸣谢

Mix Go [https://github.com/mix-go/mix](https://github.com/mix-go/mix)