-- GRANT ALL PRIVILEGES ON *.* TO lgh@"%" IDENTIFIED BY "lgh" ;
-- FLUSH PRIVILEGES ;

-- 建库
CREATE DATABASE IF NOT EXISTS hui_iot DEFAULT CHARSET utf8 COLLATE utf8_general_ci; #创建数据库
-- 选择库
USE hui_iot;

DROP TABLE iot_device;;/*SkipError*/
CREATE TABLE iot_device(
    id BIGINT NOT NULL   COMMENT '主键 分布式唯一ID' ,
    device_uid BIGINT NOT NULL   COMMENT '设备UID' ,
    device_name VARCHAR(1024)    COMMENT '设备名称' ,
    device_type_id VARCHAR(32) NOT NULL   COMMENT '设备类型ID-外键 设备类型ID-外键' ,
    online_status INT NOT NULL   COMMENT '联网状态 在线（1）、离线（0）、未激活（-1）' ,
    PRIMARY KEY (ID)
) COMMENT = '物联网设备表';;

ALTER TABLE iot_device COMMENT '物联网设备表';;
DROP TABLE iot_device_config_attr;;/*SkipError*/
CREATE TABLE iot_device_config_attr(
    id BIGINT NOT NULL   COMMENT 'ID 主键' ,
    device_id BIGINT NOT NULL   COMMENT '设备ID-外键 设备ID-外键' ,
    device_type_config_attr_id VARCHAR(32) NOT NULL   COMMENT '配置属性ID 配置属性ID' ,
    VALUE VARCHAR(3072)    COMMENT '配置属性-值 配置属性-值' ,
    PRIMARY KEY (ID)
) COMMENT = '设备配置属性表 该表允许读和写。配合影子机制更新设备配置属性。';;

ALTER TABLE iot_device_config_attr COMMENT '设备配置属性表';;
DROP TABLE iot_device_invariant_attr;;/*SkipError*/
CREATE TABLE iot_device_invariant_attr(
    ID BIGINT NOT NULL   COMMENT 'ID 主键' ,
    device_id BIGINT NOT NULL   COMMENT '设备ID-外键 设备ID-外键' ,
    device_type_invariant_attr_id VARCHAR(32) NOT NULL   COMMENT '固定属性ID 固定属性ID' ,
    VALUE VARCHAR(3072)    COMMENT '固定属性-值 固定属性-值' ,
    PRIMARY KEY (ID)
) COMMENT = '设备固定属性表 该表关联iot_device_type_invariant_attr表。如果本表没有对应的值，设备取该iot_device_type_invariant_attr表默认值。';;

ALTER TABLE iot_device_invariant_attr COMMENT '设备固定属性表';;
DROP TABLE iot_device_event;;/*SkipError*/
CREATE TABLE iot_device_event(
    TIME BIGINT NOT NULL   COMMENT '落库时间 时序数据库的主键' ,
    event_time DATETIME NOT NULL   COMMENT '事件发送时间' ,
    device_type_event_id BIGINT NOT NULL   COMMENT '设备类型事件定义ID-外键 设备类型事件定义ID-外键' ,
    event_info_map TEXT NOT NULL   COMMENT '事件信息 事件信息，json，{key:value}结构' ,
    event_level INT NOT NULL   COMMENT '事件级别 事件级别(信息INFO、告警ALARM、故障ERROR)' ,
    event_status VARCHAR(1) NOT NULL   COMMENT '事件状态 事件状态（告警1，结束0）' ,
    PRIMARY KEY (TIME)
) COMMENT = '设备事件表 设备事件表。用于记录设备历史数据。使用时序数据库。';;

ALTER TABLE iot_device_event COMMENT '设备事件表';;
DROP TABLE iot_device_event_status;;/*SkipError*/
CREATE TABLE iot_device_event_status(
    TIME BIGINT NOT NULL   COMMENT '落库时间 时序数据库的主键' ,
    event_time DATETIME NOT NULL   COMMENT '事件发生时间 事件发生时间，如果上报为空，取服务器时间。' ,
    device_type_event_id BIGINT NOT NULL   COMMENT '设备类型事件定义ID-外键 设备类型事件定义ID-外键' ,
    event_info_map TEXT NOT NULL   COMMENT '事件信息 事件信息，json，{key:value}结构' ,
    event_level INT NOT NULL   COMMENT '事件级别 事件级别(信息INFO、告警ALARM、故障ERROR)' ,
    event_status VARCHAR(1) NOT NULL   COMMENT '事件状态 事件状态（告警1，结束0）' ,
    PRIMARY KEY (TIME)
) COMMENT = '设备事件状态表 本表表示该设备对应各种事件的当前状态（持续中1、结束0）。便于批量查询某种事件处于告警的设备列表。因此，设备的事件上报需要更新该表。更新逻辑是，告警以上级别的事件上报持续中时，状态改为1，否则改为0。当仅当状态不一致时修改数据库。';;

ALTER TABLE iot_device_event_status COMMENT '设备事件状态表';;
DROP TABLE iot_device_type;;/*SkipError*/
CREATE TABLE iot_device_type(
    ID BIGINT NOT NULL   COMMENT '主键 分布式唯一ID' ,
    device_type_uid VARCHAR(128) NOT NULL   COMMENT '设备类型UID' ,
    device_type_name VARCHAR(512)    COMMENT '设备类型名称' ,
    PRIMARY KEY (ID)
) COMMENT = '物联网设备类型表 ';;

ALTER TABLE iot_device_type COMMENT '物联网设备类型表';;
DROP TABLE iot_device_type_invariant_attr;;/*SkipError*/
CREATE TABLE iot_device_type_invariant_attr(
    id BIGINT    COMMENT 'ID 主键' ,
    device_type_id BIGINT    COMMENT '设备类型ID-外键 设备类型ID-外键' ,
    attr_p_key VARCHAR(32)    COMMENT '固定属性-字段名 固定属性-字段名' ,
    data_type VARCHAR(32)    COMMENT '固定属性-数据类型 固定属性-数据类型' ,
    defalut_value VARCHAR(3072)    COMMENT '固定属性-默认值 固定属性-默认值' ,
    unit VARCHAR(32)    COMMENT '固定属性-单位 固定属性-单位' 
) COMMENT = '设备类型固定属性定义表 ';;

ALTER TABLE iot_device_type_invariant_attr COMMENT '设备类型固定属性定义表';;
DROP TABLE iot_device_type_read_attr;;/*SkipError*/
CREATE TABLE iot_device_type_read_attr(
    id BIGINT    COMMENT 'ID 主键' ,
    device_type_id BIGINT    COMMENT '设备类型ID-外键 设备类型ID-外键' ,
    attr_p_key VARCHAR(32) NOT NULL   COMMENT '只读属性字段名 只读属性字段名' ,
    data_type VARCHAR(32) NOT NULL   COMMENT '只读属性-数据类型 只读属性-数据类型' ,
    defalut_value VARCHAR(3072)    COMMENT '只读属性-默认值 只读属性-默认值' ,
    unit VARCHAR(32)    COMMENT '只读属性-单位 只读属性-单位' 
) COMMENT = '设备类型只读属性定义表 ';;

ALTER TABLE iot_device_type_read_attr COMMENT '设备类型只读属性定义表';;
DROP TABLE iot_device_type_config_attr;;/*SkipError*/
CREATE TABLE iot_device_type_config_attr(
    id BIGINT    COMMENT 'ID 主键' ,
    device_type_id BIGINT    COMMENT '设备类型ID-外键 设备类型ID-外键' ,
    attr_p_key VARCHAR(32) NOT NULL   COMMENT '配置属性字段名 配置属性字段名' ,
    data_type VARCHAR(32) NOT NULL   COMMENT '配置属性-数据类型 配置属性-数据类型' ,
    defalut_value VARCHAR(3072)    COMMENT '配置属性-默认值 配置属性-默认值' ,
    unit VARCHAR(32)    COMMENT '配置属性-单位 配置属性-单位' ,
    MAX_VALUE VARCHAR(32)    COMMENT '最大值 最大值' ,
    min_value VARCHAR(32)    COMMENT '最小值 最小值' 
) COMMENT = '设备类型配置属性定义表 ';;

ALTER TABLE iot_device_type_config_attr COMMENT '设备类型配置属性定义表';;
DROP TABLE iot_device_type_func;;/*SkipError*/
CREATE TABLE iot_device_type_func(
    id BIGINT    COMMENT 'ID 主键' ,
    device_type_id BIGINT    COMMENT '设备类型ID-外键 设备类型ID-外键' ,
    func_name VARCHAR(32) NOT NULL   COMMENT '功能名称 功能名称' ,
    in_para VARCHAR(3072)    COMMENT '输入参数信息 输入参数信息[key:value:data_type,key:value:data_type]' 
) COMMENT = '设备类型功能函数定义表 ';;

ALTER TABLE iot_device_type_func COMMENT '设备类型功能函数定义表';;
DROP TABLE iot_device_type_event;;/*SkipError*/
CREATE TABLE iot_device_type_event(
    id BIGINT NOT NULL   COMMENT 'ID 主键' ,
    device_type_id BIGINT NOT NULL   COMMENT '设备类型ID-外键 设备类型ID-外键' ,
    event_name VARCHAR(32) NOT NULL   COMMENT '事件名称 事件名称' ,
    event_level INT NOT NULL   COMMENT '事件级别 事件级别(信息INFO、告警ALARM、故障ERROR)' ,
    PRIMARY KEY (ID)
) COMMENT = '设备类型事件定义表 ';;

ALTER TABLE iot_device_type_event COMMENT '设备类型事件定义表';;
