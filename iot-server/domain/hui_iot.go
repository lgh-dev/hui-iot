package domain

import (
	"time"
)

/******sql******
CREATE TABLE `iot_device_type_read_attr` (
  `id` bigint(19) DEFAULT NULL COMMENT 'ID 主键',
  `device_type_id` bigint(19) DEFAULT NULL COMMENT '设备类型ID-外键 设备类型ID-外键',
  `attr_p_key` varchar(32) NOT NULL COMMENT '只读属性字段名 只读属性字段名',
  `data_type` varchar(32) NOT NULL COMMENT '只读属性-数据类型 只读属性-数据类型',
  `defalut_value` varchar(3072) DEFAULT NULL COMMENT '只读属性-默认值 只读属性-默认值',
  `unit` varchar(32) DEFAULT NULL COMMENT '只读属性-单位 只读属性-单位'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备类型只读属性定义表 '
******sql******/
// IotDeviceTypeReadAttr 设备类型只读属性定义表
type IotDeviceTypeReadAttr struct {
	ID           int64  `gorm:"column:id;type:bigint(19)"`                   // ID 主键
	DeviceTypeID int64  `gorm:"column:device_type_id;type:bigint(19)"`       // 设备类型ID-外键 设备类型ID-外键
	AttrPKey     string `gorm:"column:attr_p_key;type:varchar(32);not null"` // 只读属性字段名 只读属性字段名
	DataType     string `gorm:"column:data_type;type:varchar(32);not null"`  // 只读属性-数据类型 只读属性-数据类型
	DefalutValue string `gorm:"column:defalut_value;type:varchar(3072)"`     // 只读属性-默认值 只读属性-默认值
	Unit         string `gorm:"column:unit;type:varchar(32)"`                // 只读属性-单位 只读属性-单位
}

/******sql******
CREATE TABLE `PDMAN_DB_VERSION` (
  `DB_VERSION` varchar(256) DEFAULT NULL,
  `VERSION_DESC` varchar(1024) DEFAULT NULL,
  `CREATED_TIME` varchar(32) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8
******sql******/

/******sql******
CREATE TABLE `iot_device` (
  `id` bigint(19) NOT NULL COMMENT '主键 分布式唯一ID',
  `device_uid` varchar(32) NOT NULL COMMENT '设备UID',
  `device_name` varchar(1024) DEFAULT NULL COMMENT '设备名称',
  `device_type_id` bigint(20) NOT NULL COMMENT '设备类型ID-外键 设备类型ID-外键',
  `online_status` int(10) NOT NULL COMMENT '联网状态 在线（1）、离线（0）、未激活（-1）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='物联网设备表;'
******sql******/
// IotDevice 物联网设备表;
type IotDevice struct {
	ID           int64  `gorm:"primary_key;column:id;type:bigint(19);not null"` // 主键 分布式唯一ID
	DeviceUId    string `gorm:"column:device_uid;type:varchar(32);not null"`    // 设备UID
	DeviceName   string `gorm:"column:device_name;type:varchar(1024)"`          // 设备名称
	DeviceTypeID int64  `gorm:"column:device_type_id;type:bigint(20);not null"` // 设备类型ID-外键 设备类型ID-外键
	OnlineStatus int    `gorm:"column:online_status;type:int(10);not null"`     // 联网状态 在线（1）、离线（0）、未激活（-1）
}

/******sql******
CREATE TABLE `iot_device_event` (
  `time` bigint(19) NOT NULL COMMENT '落库时间 时序数据库的主键',
  `event_time` datetime NOT NULL COMMENT '事件发送时间',
  `device_type_event_id` bigint(19) NOT NULL COMMENT '设备类型事件定义ID-外键 设备类型事件定义ID-外键',
  `event_info_map` text NOT NULL COMMENT '事件信息 事件信息，json，{key:value}结构',
  `event_level` int(10) NOT NULL COMMENT '事件级别 事件级别(信息INFO、告警ALARM、故障ERROR)',
  `event_status` varchar(1) NOT NULL COMMENT '事件状态 事件状态（告警1，结束0）',
  PRIMARY KEY (`time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备事件表 '
******sql******/
// IotDeviceEvent 设备事件表
type IotDeviceEvent struct {
	Time              int64     `gorm:"primary_key;column:time;type:bigint(19);not null"`     // 落库时间 时序数据库的主键
	EventTime         time.Time `gorm:"column:event_time;type:datetime;not null"`             // 事件发送时间
	DeviceTypeEventID int64     `gorm:"column:device_type_event_id;type:bigint(19);not null"` // 设备类型事件定义ID-外键 设备类型事件定义ID-外键
	EventInfoMap      string    `gorm:"column:event_info_map;type:text;not null"`             // 事件信息 事件信息，json，{key:value}结构
	EventLevel        int       `gorm:"column:event_level;type:int(10);not null"`             // 事件级别 事件级别(信息INFO、告警ALARM、故障ERROR)
	EventStatus       string    `gorm:"column:event_status;type:varchar(1);not null"`         // 事件状态 事件状态（告警1，结束0）
}

/******sql******
CREATE TABLE `iot_device_invariant_attr` (
  `id` bigint(19) NOT NULL COMMENT 'ID 主键',
  `device_id` bigint(19) NOT NULL COMMENT '设备ID-外键 设备ID-外键',
  `device_type_invariant_attr_id` varchar(32) NOT NULL COMMENT '固定属性ID 固定属性ID',
  `value` varchar(3072) DEFAULT NULL COMMENT '固定属性-值 固定属性-值',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备固定属性表 '
******sql******/
// IotDeviceInvariantAttr 设备固定属性表
type IotDeviceInvariantAttr struct {
	ID                        int64  `gorm:"primary_key;column:id;type:bigint(19);not null"`                 // ID 主键
	DeviceID                  int64  `gorm:"column:device_id;type:bigint(19);not null"`                      // 设备ID-外键 设备ID-外键
	DeviceTypeInvariantAttrID string `gorm:"column:device_type_invariant_attr_id;type:varchar(32);not null"` // 固定属性ID 固定属性ID
	Value                     string `gorm:"column:value;type:varchar(3072)"`                                // 固定属性-值 固定属性-值
}

/******sql******
CREATE TABLE `iot_device_type` (
  `id` bigint(19) NOT NULL COMMENT '主键 分布式唯一ID',
  `device_type_uid` varchar(128) NOT NULL COMMENT '设备类型UID',
  `device_type_name` varchar(512) DEFAULT NULL COMMENT '设备类型名称',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='物联网设备类型表 '
******sql******/
// IotDeviceType 物联网设备类型表
type IotDeviceType struct {
	ID             int64  `gorm:"primary_key;column:id;type:bigint(19);not null"`    // 主键 分布式唯一ID
	DeviceTypeUId  string `gorm:"column:device_type_uid;type:varchar(128);not null"` // 设备类型UID
	DeviceTypeName string `gorm:"column:device_type_name;type:varchar(512)"`         // 设备类型名称
}

/******sql******
CREATE TABLE `iot_device_type_config_attr` (
  `id` bigint(19) DEFAULT NULL COMMENT 'ID 主键',
  `device_type_id` bigint(19) DEFAULT NULL COMMENT '设备类型ID-外键 设备类型ID-外键',
  `attr_p_key` varchar(32) NOT NULL COMMENT '配置属性字段名 配置属性字段名',
  `data_type` varchar(32) NOT NULL COMMENT '配置属性-数据类型 配置属性-数据类型',
  `defalut_value` varchar(3072) DEFAULT NULL COMMENT '配置属性-默认值 配置属性-默认值',
  `unit` varchar(32) DEFAULT NULL COMMENT '配置属性-单位 配置属性-单位',
  `max_value` varchar(32) DEFAULT NULL COMMENT '最大值 最大值',
  `min_value` varchar(32) DEFAULT NULL COMMENT '最小值 最小值'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备类型配置属性定义表 '
******sql******/
// IotDeviceTypeConfigAttr 设备类型配置属性定义表
type IotDeviceTypeConfigAttr struct {
	ID           int64  `gorm:"column:id;type:bigint(19)"`                   // ID 主键
	DeviceTypeID int64  `gorm:"column:device_type_id;type:bigint(19)"`       // 设备类型ID-外键 设备类型ID-外键
	AttrPKey     string `gorm:"column:attr_p_key;type:varchar(32);not null"` // 配置属性字段名 配置属性字段名
	DataType     string `gorm:"column:data_type;type:varchar(32);not null"`  // 配置属性-数据类型 配置属性-数据类型
	DefalutValue string `gorm:"column:defalut_value;type:varchar(3072)"`     // 配置属性-默认值 配置属性-默认值
	Unit         string `gorm:"column:unit;type:varchar(32)"`                // 配置属性-单位 配置属性-单位
	MaxValue     string `gorm:"column:max_value;type:varchar(32)"`           // 最大值 最大值
	MinValue     string `gorm:"column:min_value;type:varchar(32)"`           // 最小值 最小值
}

/******sql******
CREATE TABLE `iot_device_type_func` (
  `id` bigint(19) DEFAULT NULL COMMENT 'ID 主键',
  `device_type_id` bigint(19) DEFAULT NULL COMMENT '设备类型ID-外键 设备类型ID-外键',
  `func_name` varchar(32) NOT NULL COMMENT '功能名称 功能名称',
  `in_para` varchar(3072) DEFAULT NULL COMMENT '输入参数信息 输入参数信息[key:value:data_type,key:value:data_type]'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备类型功能函数定义表 '
******sql******/
// IotDeviceTypeFunc 设备类型功能函数定义表
type IotDeviceTypeFunc struct {
	ID           int64  `gorm:"column:id;type:bigint(19)"`                  // ID 主键
	DeviceTypeID int64  `gorm:"column:device_type_id;type:bigint(19)"`      // 设备类型ID-外键 设备类型ID-外键
	FuncName     string `gorm:"column:func_name;type:varchar(32);not null"` // 功能名称 功能名称
	InPara       string `gorm:"column:in_para;type:varchar(3072)"`          // 输入参数信息 输入参数信息[key:value:data_type,key:value:data_type]
}

/******sql******
CREATE TABLE `iot_device_config_attr` (
  `id` bigint(19) NOT NULL COMMENT 'ID 主键',
  `device_id` bigint(19) NOT NULL COMMENT '设备ID-外键 设备ID-外键',
  `device_type_config_attr_id` varchar(32) NOT NULL COMMENT '配置属性ID 配置属性ID',
  `value` varchar(3072) DEFAULT NULL COMMENT '配置属性-值 配置属性-值',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备配置属性表 '
******sql******/
// IotDeviceConfigAttr 设备配置属性表
type IotDeviceConfigAttr struct {
	ID                     int64  `gorm:"primary_key;column:id;type:bigint(19);not null"`              // ID 主键
	DeviceID               int64  `gorm:"column:device_id;type:bigint(19);not null"`                   // 设备ID-外键 设备ID-外键
	DeviceTypeConfigAttrID string `gorm:"column:device_type_config_attr_id;type:varchar(32);not null"` // 配置属性ID 配置属性ID
	Value                  string `gorm:"column:value;type:varchar(3072)"`                             // 配置属性-值 配置属性-值
}

/******sql******
CREATE TABLE `iot_device_event_status` (
  `time` bigint(19) NOT NULL COMMENT '落库时间 时序数据库的主键',
  `event_time` datetime NOT NULL COMMENT '事件发生时间 事件发生时间，如果上报为空，取服务器时间。',
  `device_type_event_id` bigint(19) NOT NULL COMMENT '设备类型事件定义ID-外键 设备类型事件定义ID-外键',
  `event_info_map` text NOT NULL COMMENT '事件信息 事件信息，json，{key:value}结构',
  `event_level` int(10) NOT NULL COMMENT '事件级别 事件级别(信息INFO、告警ALARM、故障ERROR)',
  `event_status` varchar(1) NOT NULL COMMENT '事件状态 事件状态（告警1，结束0）',
  PRIMARY KEY (`time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备事件状态表 '
******sql******/
// IotDeviceEventStatus 设备事件状态表
type IotDeviceEventStatus struct {
	Time              int64     `gorm:"primary_key;column:time;type:bigint(19);not null"`     // 落库时间 时序数据库的主键
	EventTime         time.Time `gorm:"column:event_time;type:datetime;not null"`             // 事件发生时间 事件发生时间，如果上报为空，取服务器时间。
	DeviceTypeEventID int64     `gorm:"column:device_type_event_id;type:bigint(19);not null"` // 设备类型事件定义ID-外键 设备类型事件定义ID-外键
	EventInfoMap      string    `gorm:"column:event_info_map;type:text;not null"`             // 事件信息 事件信息，json，{key:value}结构
	EventLevel        int       `gorm:"column:event_level;type:int(10);not null"`             // 事件级别 事件级别(信息INFO、告警ALARM、故障ERROR)
	EventStatus       string    `gorm:"column:event_status;type:varchar(1);not null"`         // 事件状态 事件状态（告警1，结束0）
}

/******sql******
CREATE TABLE `iot_device_type_event` (
  `id` bigint(19) NOT NULL COMMENT 'ID 主键',
  `device_type_id` bigint(19) NOT NULL COMMENT '设备类型ID-外键 设备类型ID-外键',
  `event_name` varchar(32) NOT NULL COMMENT '事件名称 事件名称',
  `event_level` int(10) NOT NULL COMMENT '事件级别 事件级别(信息INFO、告警ALARM、故障ERROR)',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备类型事件定义表 '
******sql******/
// IotDeviceTypeEvent 设备类型事件定义表
type IotDeviceTypeEvent struct {
	ID           int64  `gorm:"primary_key;column:id;type:bigint(19);not null"` // ID 主键
	DeviceTypeID int64  `gorm:"column:device_type_id;type:bigint(19);not null"` // 设备类型ID-外键 设备类型ID-外键
	EventName    string `gorm:"column:event_name;type:varchar(32);not null"`    // 事件名称 事件名称
	EventLevel   int    `gorm:"column:event_level;type:int(10);not null"`       // 事件级别 事件级别(信息INFO、告警ALARM、故障ERROR)
}

/******sql******
CREATE TABLE `iot_device_type_invariant_attr` (
  `id` bigint(19) DEFAULT NULL COMMENT 'ID 主键',
  `device_type_id` bigint(19) DEFAULT NULL COMMENT '设备类型ID-外键 设备类型ID-外键',
  `attr_p_key` varchar(32) DEFAULT NULL COMMENT '固定属性-字段名 固定属性-字段名',
  `data_type` varchar(32) DEFAULT NULL COMMENT '固定属性-数据类型 固定属性-数据类型',
  `defalut_value` varchar(3072) DEFAULT NULL COMMENT '固定属性-默认值 固定属性-默认值',
  `unit` varchar(32) DEFAULT NULL COMMENT '固定属性-单位 固定属性-单位'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备类型固定属性定义表 '
******sql******/
// IotDeviceTypeInvariantAttr 设备类型固定属性定义表
type IotDeviceTypeInvariantAttr struct {
	ID           int64  `gorm:"column:id;type:bigint(19)"`               // ID 主键
	DeviceTypeID int64  `gorm:"column:device_type_id;type:bigint(19)"`   // 设备类型ID-外键 设备类型ID-外键
	AttrPKey     string `gorm:"column:attr_p_key;type:varchar(32)"`      // 固定属性-字段名 固定属性-字段名
	DataType     string `gorm:"column:data_type;type:varchar(32)"`       // 固定属性-数据类型 固定属性-数据类型
	DefalutValue string `gorm:"column:defalut_value;type:varchar(3072)"` // 固定属性-默认值 固定属性-默认值
	Unit         string `gorm:"column:unit;type:varchar(32)"`            // 固定属性-单位 固定属性-单位
}
