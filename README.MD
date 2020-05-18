# 价值定位
1. 低消耗、安全、高性能、易接入、一键部署、快速应用开发。

# 需求文档
## 需求要点：
1. 兼容和适配不同设备。
1. 向应用提供统一设备交互方式。
1. 屏蔽底层设备和协议的差异。
1. 收集数据并存储。
1. 把数据推送出去。
1. 可以控制设备。
1. 支持超高的性能。
1. 超低的资源占用。
1. 支持插件化应用。
1. 对数据做简单的分析。
1. ui简单清晰友好。

## 架构设计：

1. 上行事件编码和下行指令编码不再考虑使用emqx的规则引擎，原因有二
    1. 不支持api添加动作。
    1. 不够强大，而且不好维护相关逻辑。
转而采用go语言自己封装基于mqtt协议的驱动编解码组件，只需要编写转码方法即可。具体代码参考：
temp-over-event.go
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

### 技术架构
1. 后端技术。
    1. Go、gin、emqx、mongodb、TDengine。
2. 前端技术。
    1. vue、element-ui。

### UI设计。
1. 使用element-ui。主颜色`#170399`。


### mqtt的topic设计。

#### 设备相关TOPIC

设备上报只读属性
```
上行
请求Topic：/hiot/sys/{deviceModelID}/{deviceID}/dv/r/up
响应Topic：/hiot/sys/{deviceModelID}/{deviceID}/dv/r/up_reponse
```
设置设备配置属性
```
下行
请求Topic：/hiot/sys/{deviceModelID}/{deviceID}/dv/c/set
响应Topic：/hiot/sys/{deviceModelID}/{deviceID}/dv/c/set_reponse
```
设备上报配置属性
```
上行
请求Topic：/hiot/sys/{deviceModelID}/{deviceID}/dv/c/up
响应Topic：/hiot/sys/{deviceModelID}/{deviceID}/dv/c/up_reponse
```
设备事件上报
```
上行
请求Topic：/hiot/sys/{deviceModelID}/{deviceID}/dv/e/up
响应Topic：/hiot/sys/{deviceModelID}/{deviceID}/dv/e/up_reponse
```
设备指令下发
```
上行
请求Topic：/hiot/sys/{deviceModelID}/{deviceID}/dv/cmd/down
响应Topic：/hiot/sys/{deviceModelID}/{deviceID}/dv/cmd/down_reponse
```

设备影子
```
/hiot/shadow/update/${YourProductKey}/${YourDeviceName}
/hiot/get/${YourProductKey}/${YourDeviceName}
```
智能汽车设备信息例子
```
 {
        "k-sc_id": "0001",
        "k-engine_id": "111111",
        "k-licence_plate_number": "粤AF88888",
        "wheel": [{
            "k-w_id": "00",
            "p-model": "奥迪",
            "p-size": "21寸",
            "r-rpm": "sensor_0001_wheel_00_rpm",    
            "r-pressure": "sensor_0001_wheel_00_pressure"
        }, {
            "k-w_id": "10",
            "p-model": "米其林",
            "p-size": "21寸",
            "r-rpm": "sensor_0001_wheel_10_rpm",
            "r-pressure": "sensor_0001_wheel_10_pressure"
        },],
        "head_light": {
            "c-brightness":100
        },
        "p-model": "奥迪",
        "p-tank_capacity": 60,
        "r-driving_speed": "sensor_0001_driving_speed",
        "r-remaining_voyage": "sensor_0001_remaining_voyage",
        "r-latitude": "sensor_0001_latitude",
        "r-longitude": "sensor_0001_longitude",
        "c-speed_lock": 70,
        "e-crash_event": "alarm_0001_crash_event",
        "f-unlock": true,
        "f-autopilot": true
    }
```

### 名词解释


# 数据驱动型IOT体系架构设计
原文
[https://www.ibm.com/developerworks/cn/iot/library/iot-lo-architecture-design-from-application-layer/index.html][从应用层表象出发进行IoT架构抽象设计]: 


## 进度
### 实体定义

## 疑问集合

### 定义一类设备的时候，以什么原则判断一个设备是否属于该类设备，如汽车、电动车，如通过集中控制器和单灯节点的路灯与nb通讯的单灯？
答：定义一类物联网设备的时候，以网关ID为唯一标识，同一网关下的设备共同组成一个物联网模型。如果一个模型里面的子设备独立联网应当划分为另外一个物联网模型。
正例：
    1. 气象站模型=网关ID+气象站。
    2. 物联网自行车模型=通讯卡唯一标识+自行车+智能胎压监测的车轮。
    3. 智能汽车模型=通讯卡唯一标识+智能汽车。
反例：
    1. 一体化灯杆模型=灯杆出厂ID+摄像头+户外广告屏+广播音箱+单灯控制器。
正确的应该是分开为：摄像头模型、户外广告屏模型、广播音箱模型、单灯控制器模型，然后在应用层组合使用。
原因是：
    不具有通讯唯一标识的模型，当数据上来的时候因数据包中不存在设备的唯一标识，无法映射到具体的设备上。
    如：`topic=/read/{网关ID}  {"temp":30}`
    
###　物联网设备的一些告警信息如何体现在五元组上？是只读属性么？
答：只是设备的五元组。可以扩充事件。  

### 待解决问题：
 1. tdengine的超级表标签只能6个？
 
 2. 自动建表性能是不是差点？
 
 3. 驱动程序是否需要查询设备信息？
 答：不需要。数据上报直接落库，指令下发由服务层将指令数据完整带下来。
 4. 驱动程序需要查询设备模型数据么？
 答：由于驱动和设备模型是多对一的关系，程序启动初始化应该就有。
 5. 应该是新增设备的时候，同时在tdengine中新增子表么？
 答：是的。`ALTER TABLE <table_name> SET TAG <tag_name>=<new_tag_value>`。
 6. 是否支持扩展多套应用属性？
 答：只支持扩展一套应用属性，支持多应用使用同一套应用属性。使用多套应用属性的话，请部署多套系统。

## 脚步记录

## 基准测试命令
测试ID
```
go test -bench=BenchmarkNextId -benchmem  ./iot-server/service/
```

## TDengine建表语句

```
# 状态超级表

CREATE TABLE iot_online_status (ts timestamp, status binary(20)) TAGS(device_model binary(20),brand binary(20),model binary(20),device_id binary(20));

INSERT INTO hui_iot.status_0001_0001_onlin1e_status USING hui_iot.iot_online_status TAGS ('smart_car_camera','海康威视', 'SZ-0001','0001_1000') VALUES ('2018-01-01 00:00:00.000','online');

SELECT * FROM hui_iot.iot_status_smart_car_camera ;

# 传感值超级表
CREATE TABLE iot_read_smart_car_camera (ts timestamp, value double) TAGS(device_model binary(20),device_id binary(20), read_id binary(20));

# 插入值
INSERT INTO hui_iot.read_0001_1000_temp USING hui_iot.iot_read_smart_car_camera TAGS ('smart_car_camera','0001_1000', 'temp') VALUES ('2018-01-04 00:00:00.000',95);
INSERT INTO hui_iot.read_0001_1000_wind_direction USING hui_iot.iot_sensor_smart_car_camera TAGS ('smart_car_camera','0001_1000', 'wind_direction') VALUES ('2018-01-01 00:00:00.000',44);
INSERT INTO hui_iot.read_0001_1000_wind_direction USING hui_iot.iot_sensor_smart_car_camera TAGS ('smart_car_camera','0001_1000', 'wind_direction') VALUES ('2018-01-02 00:00:00.000',123);

# 测试表

CREATE TABLE test (ts timestamp, value double) TAGS(device_model binary(20),device_id binary(20), read_id binary(20), area_id binary(20), device_type binary(20), manufacture_id binary(20), supplie_id binary(20));
CREATE TABLE test (ts timestamp, value double) TAGS(device_model binary(20),device_id binary(20), read_id binary(20), area_id binary(20), device_type binary(20), manufacture_id binary(20));

```

