-- GRANT ALL PRIVILEGES ON *.* TO lgh@"%" IDENTIFIED BY "lgh" ;
-- FLUSH PRIVILEGES ;

-- 建库
CREATE DATABASE IF NOT EXISTS hui_iot DEFAULT CHARSET utf8 COLLATE utf8_general_ci; #创建数据库
-- 选择库
USE hui_iot;
DROP TABLE device;;/*SkipError*/
CREATE TABLE device(
                       id VARCHAR(32) NOT NULL   COMMENT '主键 分布式唯一ID' ,
                       device_uid VARCHAR(32) NOT NULL   COMMENT '设备UID' ,
                       device_name VARCHAR(1024)    COMMENT '设备名称' ,
                       gateway_id VARCHAR(32)    COMMENT '网关ID' ,
                       device_model_id BIGINT NOT NULL   COMMENT '设备模型ID-外键 设备模型ID-外键' ,
                       online_status INT(10) NOT NULL   COMMENT '联网状态 在线（1）、离线（0）、未激活（-1）' ,
                       PRIMARY KEY (id)
) COMMENT = '物联网设备表 ';;

ALTER TABLE device COMMENT '物联网设备表';;
DROP TABLE device_read_attr;;/*SkipError*/
CREATE TABLE device_read_attr(
                                 REVISION INT    COMMENT '乐观锁' ,
                                 CREATED_BY VARCHAR(32)    COMMENT '创建人' ,
                                 CREATED_TIME DATETIME    COMMENT '创建时间' ,
                                 UPDATED_BY VARCHAR(32)    COMMENT '更新人' ,
                                 UPDATED_TIME DATETIME    COMMENT '更新时间'
) COMMENT = '设备只读属性表 ';;

ALTER TABLE device_read_attr COMMENT '设备只读属性表';;
DROP TABLE device_config_attr;;/*SkipError*/
CREATE TABLE device_config_attr(
                                   id BIGINT(19) NOT NULL   COMMENT 'ID 主键' ,
                                   device_id BIGINT(19) NOT NULL   COMMENT '设备ID-外键' ,
                                   ukey VARCHAR(32) NOT NULL   COMMENT '配置属性标识' ,
                                   value VARCHAR(3072)    COMMENT '配置属性-值' ,
                                   PRIMARY KEY (id)
) COMMENT = '设备配置属性表 ';;

ALTER TABLE device_config_attr COMMENT '设备配置属性表';;
DROP TABLE device_invariant_attr;;/*SkipError*/
CREATE TABLE device_invariant_attr(
                                      id BIGINT(19) NOT NULL   COMMENT 'ID 主键' ,
                                      device_id BIGINT(19) NOT NULL   COMMENT '设备ID-外键 设备ID-外键' ,
                                      ukey VARCHAR(32) NOT NULL   COMMENT '固定属性标识' ,
                                      value VARCHAR(3072)    COMMENT '固定属性-值' ,
                                      PRIMARY KEY (id)
) COMMENT = '设备固定属性表 ';;

ALTER TABLE device_invariant_attr COMMENT '设备固定属性表';;
DROP TABLE device_event;;/*SkipError*/
CREATE TABLE device_event(
                             time BIGINT(19) NOT NULL   COMMENT '落库时间 时序数据库的主键' ,
                             send_time DATETIME NOT NULL   COMMENT '事件发送时间' ,
                             ukey BIGINT(19) NOT NULL   COMMENT '事件定义标识' ,
                             info TEXT NOT NULL   COMMENT '事件信息' ,
                             level VARCHAR(32) NOT NULL   COMMENT '事件级别(信息INFO、告警ALARM、故障ERROR)' ,
                             status VARCHAR(32) NOT NULL   COMMENT '事件状态 事件状态（告警ALARM，结束OVER）' ,
                             PRIMARY KEY (time)
) COMMENT = '设备事件表 ';;

ALTER TABLE device_event COMMENT '设备事件表';;
DROP TABLE device_command;;/*SkipError*/
CREATE TABLE device_command(
                               cmd_id VARCHAR(32)    COMMENT '指令ID' ,
                               device_id VARCHAR(32)    COMMENT '设备ID' ,
                               send_data VARCHAR(32)    COMMENT '发送数据' ,
                               return_data VARCHAR(32)    COMMENT '响应数据（同步或异步）' ,
                               status VARCHAR(32)    COMMENT '指令状态(send【下发】、ack【设备已接收】、succ【执行成功】、error【异常】)' ,
                               msg VARCHAR(32)    COMMENT '执行信息，包括成功或失败的信息' ,
                               is_retry VARCHAR(1)    COMMENT '是否需要重发' ,
                               retry_interval VARCHAR(32)    COMMENT '重发间隔(s)' ,
                               retry_number VARCHAR(32)    COMMENT '重发次数' ,
                               retry_completed_number VARCHAR(32)    COMMENT '已重发次数'
) COMMENT = '指令下发表 ';;

ALTER TABLE device_command COMMENT '指令下发表';;

