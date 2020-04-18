# 价值定位
1. 低消耗、安全、高性能、易接入、一键部署、快速应用开发。

# 需求文档
## 需求要点：
1. 兼容和适配不同设备。
1. 向应用提供统一设备交互方式。
1. 收集数据并存储。
1. 把数据推送出去。
1. 可以控制设备。
1. 支持超高的性能。
1. 超低的资源占用。
1. 支持插件化应用。
1. 对数据做简单的分析。
1. ui简单清晰友好。

## 架构设计：
### 业务架构：
1. iot-server服务。
    1. 设备模型：
        1. 定义（配置文件）：
            1. 设备配置属性。
            1. 设备传感属性。
            1. 告警事件。
            1. 指令下发。
        1. 后端API:
            1. 查询。
    1. 编解码配置：
        1. 定义（配置文件）：
            1. 配置类型：
                1. emqx。
            1. 属性上报。
            1. 指令下发。
            1. 下行指令。
        1. emqx：
            1. 编解码规则定义（配置文件）:
                1. 
    1. 应用管理：
        1. 定义（配置文件）：
            1. 应用名称。
            1. 应用账号密码。
            1. 设备业务属性（跟应用1对1）。
        1. 后端API:
            1. 应用accessToken获取。
            1. 权限控制。
    1. 设备管理：
        1. 后端API:
            1. 查询。
            1. 新增。
            1. 更新。
            1. 删除。
            1. 指令下发。
            1. 数据查询。

1. iot-worker服务。
    1. 设备数据采集和处理。
    1. 在线和离线策略。
    1. 定时任务。

### API设计
参考`api.go`。
1. iot-server服务。
    1. 设备模型：
        1. 初始化。
        1. 批量查询。
        1. 单个查询。
    1. 编解码配置：
        1. 初始化。
        1. 批量查询。
        1. 单个查询。
    1. 应用管理：
        1. 初始化。
        1. 数据库地址查询。
        1. 应用accessToken获取。
        1. 权限控制。
    1. 设备管理：
        1. 后端API:
            1. 批量查询。
            1. 单个查询。
            1. 新增（包括业务属性）。
            1. 更新（包括业务属性）。
            1. 删除。
            1. 指令下发。
            1. 数据查询。
                1. 传感数据。
                1. 告警事件。
                1. 指令下发。

### UI设计。
1. 使用element-ui。主颜色`#170399`。


### mqtt的topic设计。

#### 设备相关TOPIC

设备上报固定属性
```
请求Topic：/sys/{deviceModelID}/{deviceID}/thing/deviceinfo/update
响应Topic：/sys/{deviceModelID}/{deviceID}/thing/deviceinfo/update_reply
```
设备上报只读属性
```
上行（透传）
请求Topic：/sys/{deviceModelID}/{deviceID}/thing/model/up_raw
响应Topic：/sys/{deviceModelID}/{deviceID}/thing/model/up_raw_reply

上行（Alink JSON）

请求Topic： /sys/{deviceModelID}/{deviceID}/thing/event/property/post
响应Topic： /sys/{deviceModelID}/{deviceID}/thing/event/property/post_reply
上报格式
{
  "id": "123",
  "version": "1.0",
  "params": {
    "Power": {
      "value": "on",
      "time": 1524448722000
    },
    "WF": {
      "value": 23.6,
      "time": 1524448722000
    }
  },
  "method": "thing.event.property.post"
}
```
设置设备配置属性
```
下行（透传）
请求Topic：/sys/{deviceModelID}/{deviceID}/thing/model/down_raw
响应Topic：/sys/{deviceModelID}/{deviceID}/thing/model/down_raw_reply

下行（Alink JSON）
请求Topic：/sys/{deviceModelID}/{deviceID}/thing/service/property/set
响应Topic：/sys/{deviceModelID}/{deviceID}/thing/service/property/set_reply
数据格式
{
  "id": "123",
  "version": "1.0",
  "params": {
    "temperature": "30.5"
  },
  "method": "thing.service.property.set"
}
```
设备事件上报
```
上行（透传）

请求Topic：/sys/{deviceModelID}/{deviceID}/thing/model/up_raw
响应Topic：/sys/{deviceModelID}/{deviceID}/thing/model/up_raw_reply

上行（Alink JSON）

请求Topic：/sys/{deviceModelID}/{deviceID}/thing/event/{tsl.event.identifier}/post
响应Topic：/sys/{deviceModelID}/{deviceID}/thing/event/{tsl.event.identifier}/post_reply
```


设备影子
```
/shadow/update/${YourProductKey}/${YourDeviceName}
/shadow/get/${YourProductKey}/${YourDeviceName}
```


### 名词解释


# 数据驱动型IOT体系架构设计
原文
[https://www.ibm.com/developerworks/cn/iot/library/iot-lo-architecture-design-from-application-layer/index.html][从应用层表象出发进行IoT架构抽象设计]: 


## 进度
### 原型设计-设备管理
还差父子设备结构。

## 疑问集合

###　物联网设备的一些告警信息如何体现在五元组上？是只读属性么？
答：只是设备的五元组。可以扩充事件。

### 父子设备如何体现？
答：

## 脚步记录
代码生成器命令。
```
D:\GoProject\代码生成器\go-mygen.exe -h 192.168.20.101 -P 3306 -u root -p lgh -d hui_iot
```