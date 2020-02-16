# 价值定位
1. 低消耗、高性能、易接入、一键部署、快速应用开发。

# 需求文档
## 需求要点：
1. 兼容不同设备。
1. 收集数据并存储。
1. 把数据推送出去。
1. 可以控制设备。
1. 支持超高的性能。
1. 超低的资源占用。
1. 对数据做简单的分析。
1. 通用API支持上层应用开发。移动端、PC端。

## 架构设计：
服务关联功能：
1. 设备服务。（支持设备CURDF，影子机制）
1. 数据服务。（支持数据的存储、事件流分析、数据推送）
1. 计算服务。（支持通用报表统计）
1. API服务。（代理内部接口，对外统一支持移动端、PC端、三方平台）


### mqtt的topic设计。

#### 设备相关TOPIC

设备上报固定属性
```
请求Topic：/sys/{productKey}/{deviceName}/thing/deviceinfo/update
响应Topic：/sys/{productKey}/{deviceName}/thing/deviceinfo/update_reply
```
设备上报只读属性
```
上行（透传）
请求Topic：/sys/{productKey}/{deviceName}/thing/model/up_raw
响应Topic：/sys/{productKey}/{deviceName}/thing/model/up_raw_reply

上行（Alink JSON）

请求Topic： /sys/{productKey}/{deviceName}/thing/event/property/post
响应Topic： /sys/{productKey}/{deviceName}/thing/event/property/post_reply
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
请求Topic：/sys/{productKey}/{deviceName}/thing/model/down_raw
响应Topic：/sys/{productKey}/{deviceName}/thing/model/down_raw_reply

下行（Alink JSON）
请求Topic：/sys/{productKey}/{deviceName}/thing/service/property/set
响应Topic：/sys/{productKey}/{deviceName}/thing/service/property/set_reply
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

请求Topic：/sys/{productKey}/{deviceName}/thing/model/up_raw
响应Topic：/sys/{productKey}/{deviceName}/thing/model/up_raw_reply

上行（Alink JSON）

请求Topic：/sys/{productKey}/{deviceName}/thing/event/{tsl.event.identifier}/post
响应Topic：/sys/{productKey}/{deviceName}/thing/event/{tsl.event.identifier}/post_reply
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