-- public.consumer_group definition

-- Drop table

-- DROP TABLE public.consumer_group;

CREATE TABLE public.consumer_group (
	id varchar NOT NULL, -- id
	group_id varchar NULL, -- 消费组ID
	auth_type varchar NULL, -- 权限类型 ，aksign
	cloud_commodity_type varchar NULL, -- 云商品类型iothub_senior
	iot_id varchar NULL, -- 疑似租户id待确定
	group_name varchar NULL, -- 消费组名称
	CONSTRAINT consumer_group_pk PRIMARY KEY (id),
	CONSTRAINT consumer_group_un UNIQUE (iot_id)
);
COMMENT ON TABLE public.consumer_group IS '消费组';

-- Column comments

COMMENT ON COLUMN public.consumer_group.id IS 'id';
COMMENT ON COLUMN public.consumer_group.group_id IS '消费组ID';
COMMENT ON COLUMN public.consumer_group.auth_type IS '权限类型 ，aksign';
COMMENT ON COLUMN public.consumer_group.cloud_commodity_type IS '云商品类型iothub_senior';
COMMENT ON COLUMN public.consumer_group.iot_id IS '疑似租户id待确定';
COMMENT ON COLUMN public.consumer_group.group_name IS '消费组名称';


-- public.product definition

-- Drop table

-- DROP TABLE public.product;

CREATE TABLE public.product (
	product_name varchar(30) NOT NULL, -- 支持中文、英文字母、日文、数字、和特殊字符_-@()，长度限制 4~30个字符，中文及日文算 2 个字符
	product_status varchar(30) NULL, -- 开发中DEVELOPMENT_STATUS，RELEASE_STATUS已发布
	node_type int4 NOT NULL, -- 节点类型：0直连，1网关子设备，2网关
	net_type int4 NULL, -- 连网方式:0-wifi,1-蜂窝（2G/3G/4G/5G）,2-以太网,3-LoRaWAN,-1其他
	data_format int4 NULL, -- 1-Alink JSON ，0-透传/自定义
	category_id int8 NULL, -- 产品品类。数据来源于内置的标准产品。其他/自定义为-1
	auth_type varchar(30) NULL, -- 认证方式。设备密钥，ID2,X.509证书
	product_key varchar(15) NOT NULL, -- productKey主键
	product_secret varchar(32) NULL, -- 产品密钥
	CONSTRAINT product_pk PRIMARY KEY (product_key)
);
COMMENT ON TABLE public.product IS '产品表';

-- Column comments

COMMENT ON COLUMN public.product.product_name IS '支持中文、英文字母、日文、数字、和特殊字符_-@()，长度限制 4~30个字符，中文及日文算 2 个字符';
COMMENT ON COLUMN public.product.product_status IS '开发中DEVELOPMENT_STATUS，RELEASE_STATUS已发布';
COMMENT ON COLUMN public.product.node_type IS '节点类型：0直连，1网关子设备，2网关';
COMMENT ON COLUMN public.product.net_type IS '连网方式:0-wifi,1-蜂窝（2G/3G/4G/5G）,2-以太网,3-LoRaWAN,-1其他';
COMMENT ON COLUMN public.product.data_format IS '1-Alink JSON ，0-透传/自定义';
COMMENT ON COLUMN public.product.category_id IS '产品品类。数据来源于内置的标准产品。其他/自定义为-1';
COMMENT ON COLUMN public.product.auth_type IS '认证方式。设备密钥，ID2,X.509证书';
COMMENT ON COLUMN public.product.product_key IS 'productKey主键';
COMMENT ON COLUMN public.product.product_secret IS '产品密钥';


-- public.data_script definition

-- Drop table

-- DROP TABLE public.data_script;

CREATE TABLE public.data_script (
	product_key varchar NULL, -- 产品Key
	id int8 NOT NULL, -- 主键
	script_type varchar NULL, -- 脚步类型，目前支持js
	script_content text NULL, -- 脚本内容
	script_code varchar(32) NULL, -- 脚本标识
	CONSTRAINT data_script_pk PRIMARY KEY (id),
	CONSTRAINT data_script_un UNIQUE (script_code),
	CONSTRAINT data_script_fk FOREIGN KEY (product_key) REFERENCES product(product_key)
);
COMMENT ON TABLE public.data_script IS '数据解释脚本，支持js，ECMAScript 5语法。';

-- Column comments

COMMENT ON COLUMN public.data_script.product_key IS '产品Key';
COMMENT ON COLUMN public.data_script.id IS '主键';
COMMENT ON COLUMN public.data_script.script_type IS '脚步类型，目前支持js';
COMMENT ON COLUMN public.data_script.script_content IS '脚本内容';
COMMENT ON COLUMN public.data_script.script_code IS '脚本标识';


-- public.product_tag definition

-- Drop table

-- DROP TABLE public.product_tag;

CREATE TABLE public.product_tag (
	id varchar NOT NULL, -- id
	tag_key varchar NULL, -- 标签key
	tag_value varchar NULL, -- 标签value
	product_key varchar NULL, -- 产品key
	CONSTRAINT product_label_pk PRIMARY KEY (id),
	CONSTRAINT product_tag_fk FOREIGN KEY (product_key) REFERENCES product(product_key)
);
COMMENT ON TABLE public.product_tag IS '产品标签表';

-- Column comments

COMMENT ON COLUMN public.product_tag.id IS 'id';
COMMENT ON COLUMN public.product_tag.tag_key IS '标签key';
COMMENT ON COLUMN public.product_tag.tag_value IS '标签value';
COMMENT ON COLUMN public.product_tag.product_key IS '产品key';


-- public.subscribe definition

-- Drop table

-- DROP TABLE public.subscribe;

CREATE TABLE public.subscribe (
	id int8 NOT NULL, -- id
	product_key varchar NULL,
	subscribe_type varchar NULL, -- 订阅类型。AMQP
	callback_type int4 NULL, -- 回调类型。206
	device_life_cycle_flag bool NULL, -- 设备生命周期变更
	device_data_flag bool NULL, -- 设备上报消息
	device_status_change_flag bool NULL, -- 设备状态变化
	thing_history_flag bool NULL, -- 物模型历史数据上报
	ota_event_flag bool NULL, -- 固件升级状态通知
	ota_job_flag bool NULL, -- ota升级批次通知
	ota_version_flag bool NULL, -- ota模型版本号上报
	consumer_group varchar NULL,
	CONSTRAINT subscribe_pk PRIMARY KEY (id),
	CONSTRAINT subscribe_fk FOREIGN KEY (product_key) REFERENCES product(product_key),
	CONSTRAINT subscribe_fk_consumer_group FOREIGN KEY (consumer_group) REFERENCES consumer_group(iot_id)
);
COMMENT ON TABLE public.subscribe IS '订阅表';

-- Column comments

COMMENT ON COLUMN public.subscribe.id IS 'id';
COMMENT ON COLUMN public.subscribe.subscribe_type IS '订阅类型。AMQP';
COMMENT ON COLUMN public.subscribe.callback_type IS '回调类型。206';
COMMENT ON COLUMN public.subscribe.device_life_cycle_flag IS '设备生命周期变更';
COMMENT ON COLUMN public.subscribe.device_data_flag IS '设备上报消息';
COMMENT ON COLUMN public.subscribe.device_status_change_flag IS '设备状态变化';
COMMENT ON COLUMN public.subscribe.thing_history_flag IS '物模型历史数据上报';
COMMENT ON COLUMN public.subscribe.ota_event_flag IS '固件升级状态通知';
COMMENT ON COLUMN public.subscribe.ota_job_flag IS 'ota升级批次通知';
COMMENT ON COLUMN public.subscribe.ota_version_flag IS 'ota模型版本号上报';


-- public.topic definition

-- Drop table

-- DROP TABLE public.topic;

CREATE TABLE public.topic (
	id int8 NOT NULL, -- 主键
	operation int2 NULL, -- 操作全新：0-订阅，1-发布。
	product_key varchar(15) NOT NULL, -- 产品Key
	topic_short_name varchar(255) NOT NULL, -- 主题简称
	CONSTRAINT topic_pk PRIMARY KEY (id),
	CONSTRAINT topic_fk FOREIGN KEY (product_key) REFERENCES product(product_key)
);
COMMENT ON TABLE public.topic IS '消息主题表';

-- Column comments

COMMENT ON COLUMN public.topic.id IS '主键';
COMMENT ON COLUMN public.topic.operation IS '操作全新：0-订阅，1-发布。';
COMMENT ON COLUMN public.topic.product_key IS '产品Key';
COMMENT ON COLUMN public.topic.topic_short_name IS '主题简称';


-- public.tsl definition

-- Drop table

-- DROP TABLE public.tsl;

CREATE TABLE public.tsl (
	id int8 NOT NULL, -- ID
	product_key varchar NULL, -- 产品Key
	tsl text NULL, -- 物模型
	"version" int8 NULL, -- 版本号时间戳
	CONSTRAINT tsl_fk FOREIGN KEY (product_key) REFERENCES product(product_key)
);
COMMENT ON TABLE public.tsl IS '物模型表';

-- Column comments

COMMENT ON COLUMN public.tsl.id IS 'ID';
COMMENT ON COLUMN public.tsl.product_key IS '产品Key';
COMMENT ON COLUMN public.tsl.tsl IS '物模型';
COMMENT ON COLUMN public.tsl."version" IS '版本号时间戳';