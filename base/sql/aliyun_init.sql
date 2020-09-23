-- public.consumer_group definition

-- Drop table

-- DROP TABLE public.consumer_group;

CREATE TABLE public.consumer_group (
	id varchar NOT NULL, -- id
	group_id varchar NULL, -- ������ID
	auth_type varchar NULL, -- Ȩ������ ��aksign
	cloud_commodity_type varchar NULL, -- ����Ʒ����iothub_senior
	iot_id varchar NULL, -- �����⻧id��ȷ��
	group_name varchar NULL, -- ����������
	CONSTRAINT consumer_group_pk PRIMARY KEY (id),
	CONSTRAINT consumer_group_un UNIQUE (iot_id)
);
COMMENT ON TABLE public.consumer_group IS '������';

-- Column comments

COMMENT ON COLUMN public.consumer_group.id IS 'id';
COMMENT ON COLUMN public.consumer_group.group_id IS '������ID';
COMMENT ON COLUMN public.consumer_group.auth_type IS 'Ȩ������ ��aksign';
COMMENT ON COLUMN public.consumer_group.cloud_commodity_type IS '����Ʒ����iothub_senior';
COMMENT ON COLUMN public.consumer_group.iot_id IS '�����⻧id��ȷ��';
COMMENT ON COLUMN public.consumer_group.group_name IS '����������';


-- public.product definition

-- Drop table

-- DROP TABLE public.product;

CREATE TABLE public.product (
	product_name varchar(30) NOT NULL, -- ֧�����ġ�Ӣ����ĸ�����ġ����֡��������ַ�_-@()���������� 4~30���ַ������ļ������� 2 ���ַ�
	product_status varchar(30) NULL, -- ������DEVELOPMENT_STATUS��RELEASE_STATUS�ѷ���
	node_type int4 NOT NULL, -- �ڵ����ͣ�0ֱ����1�������豸��2����
	net_type int4 NULL, -- ������ʽ:0-wifi,1-���ѣ�2G/3G/4G/5G��,2-��̫��,3-LoRaWAN,-1����
	data_format int4 NULL, -- 1-Alink JSON ��0-͸��/�Զ���
	category_id int8 NULL, -- ��ƷƷ�ࡣ������Դ�����õı�׼��Ʒ������/�Զ���Ϊ-1
	auth_type varchar(30) NULL, -- ��֤��ʽ���豸��Կ��ID2,X.509֤��
	product_key varchar(15) NOT NULL, -- productKey����
	product_secret varchar(32) NULL, -- ��Ʒ��Կ
	CONSTRAINT product_pk PRIMARY KEY (product_key)
);
COMMENT ON TABLE public.product IS '��Ʒ��';

-- Column comments

COMMENT ON COLUMN public.product.product_name IS '֧�����ġ�Ӣ����ĸ�����ġ����֡��������ַ�_-@()���������� 4~30���ַ������ļ������� 2 ���ַ�';
COMMENT ON COLUMN public.product.product_status IS '������DEVELOPMENT_STATUS��RELEASE_STATUS�ѷ���';
COMMENT ON COLUMN public.product.node_type IS '�ڵ����ͣ�0ֱ����1�������豸��2����';
COMMENT ON COLUMN public.product.net_type IS '������ʽ:0-wifi,1-���ѣ�2G/3G/4G/5G��,2-��̫��,3-LoRaWAN,-1����';
COMMENT ON COLUMN public.product.data_format IS '1-Alink JSON ��0-͸��/�Զ���';
COMMENT ON COLUMN public.product.category_id IS '��ƷƷ�ࡣ������Դ�����õı�׼��Ʒ������/�Զ���Ϊ-1';
COMMENT ON COLUMN public.product.auth_type IS '��֤��ʽ���豸��Կ��ID2,X.509֤��';
COMMENT ON COLUMN public.product.product_key IS 'productKey����';
COMMENT ON COLUMN public.product.product_secret IS '��Ʒ��Կ';


-- public.data_script definition

-- Drop table

-- DROP TABLE public.data_script;

CREATE TABLE public.data_script (
	product_key varchar NULL, -- ��ƷKey
	id int8 NOT NULL, -- ����
	script_type varchar NULL, -- �Ų����ͣ�Ŀǰ֧��js
	script_content text NULL, -- �ű�����
	script_code varchar(32) NULL, -- �ű���ʶ
	CONSTRAINT data_script_pk PRIMARY KEY (id),
	CONSTRAINT data_script_un UNIQUE (script_code),
	CONSTRAINT data_script_fk FOREIGN KEY (product_key) REFERENCES product(product_key)
);
COMMENT ON TABLE public.data_script IS '���ݽ��ͽű���֧��js��ECMAScript 5�﷨��';

-- Column comments

COMMENT ON COLUMN public.data_script.product_key IS '��ƷKey';
COMMENT ON COLUMN public.data_script.id IS '����';
COMMENT ON COLUMN public.data_script.script_type IS '�Ų����ͣ�Ŀǰ֧��js';
COMMENT ON COLUMN public.data_script.script_content IS '�ű�����';
COMMENT ON COLUMN public.data_script.script_code IS '�ű���ʶ';


-- public.product_tag definition

-- Drop table

-- DROP TABLE public.product_tag;

CREATE TABLE public.product_tag (
	id varchar NOT NULL, -- id
	tag_key varchar NULL, -- ��ǩkey
	tag_value varchar NULL, -- ��ǩvalue
	product_key varchar NULL, -- ��Ʒkey
	CONSTRAINT product_label_pk PRIMARY KEY (id),
	CONSTRAINT product_tag_fk FOREIGN KEY (product_key) REFERENCES product(product_key)
);
COMMENT ON TABLE public.product_tag IS '��Ʒ��ǩ��';

-- Column comments

COMMENT ON COLUMN public.product_tag.id IS 'id';
COMMENT ON COLUMN public.product_tag.tag_key IS '��ǩkey';
COMMENT ON COLUMN public.product_tag.tag_value IS '��ǩvalue';
COMMENT ON COLUMN public.product_tag.product_key IS '��Ʒkey';


-- public.subscribe definition

-- Drop table

-- DROP TABLE public.subscribe;

CREATE TABLE public.subscribe (
	id int8 NOT NULL, -- id
	product_key varchar NULL,
	subscribe_type varchar NULL, -- �������͡�AMQP
	callback_type int4 NULL, -- �ص����͡�206
	device_life_cycle_flag bool NULL, -- �豸�������ڱ��
	device_data_flag bool NULL, -- �豸�ϱ���Ϣ
	device_status_change_flag bool NULL, -- �豸״̬�仯
	thing_history_flag bool NULL, -- ��ģ����ʷ�����ϱ�
	ota_event_flag bool NULL, -- �̼�����״̬֪ͨ
	ota_job_flag bool NULL, -- ota��������֪ͨ
	ota_version_flag bool NULL, -- otaģ�Ͱ汾���ϱ�
	consumer_group varchar NULL,
	CONSTRAINT subscribe_pk PRIMARY KEY (id),
	CONSTRAINT subscribe_fk FOREIGN KEY (product_key) REFERENCES product(product_key),
	CONSTRAINT subscribe_fk_consumer_group FOREIGN KEY (consumer_group) REFERENCES consumer_group(iot_id)
);
COMMENT ON TABLE public.subscribe IS '���ı�';

-- Column comments

COMMENT ON COLUMN public.subscribe.id IS 'id';
COMMENT ON COLUMN public.subscribe.subscribe_type IS '�������͡�AMQP';
COMMENT ON COLUMN public.subscribe.callback_type IS '�ص����͡�206';
COMMENT ON COLUMN public.subscribe.device_life_cycle_flag IS '�豸�������ڱ��';
COMMENT ON COLUMN public.subscribe.device_data_flag IS '�豸�ϱ���Ϣ';
COMMENT ON COLUMN public.subscribe.device_status_change_flag IS '�豸״̬�仯';
COMMENT ON COLUMN public.subscribe.thing_history_flag IS '��ģ����ʷ�����ϱ�';
COMMENT ON COLUMN public.subscribe.ota_event_flag IS '�̼�����״̬֪ͨ';
COMMENT ON COLUMN public.subscribe.ota_job_flag IS 'ota��������֪ͨ';
COMMENT ON COLUMN public.subscribe.ota_version_flag IS 'otaģ�Ͱ汾���ϱ�';


-- public.topic definition

-- Drop table

-- DROP TABLE public.topic;

CREATE TABLE public.topic (
	id int8 NOT NULL, -- ����
	operation int2 NULL, -- ����ȫ�£�0-���ģ�1-������
	product_key varchar(15) NOT NULL, -- ��ƷKey
	topic_short_name varchar(255) NOT NULL, -- ������
	CONSTRAINT topic_pk PRIMARY KEY (id),
	CONSTRAINT topic_fk FOREIGN KEY (product_key) REFERENCES product(product_key)
);
COMMENT ON TABLE public.topic IS '��Ϣ�����';

-- Column comments

COMMENT ON COLUMN public.topic.id IS '����';
COMMENT ON COLUMN public.topic.operation IS '����ȫ�£�0-���ģ�1-������';
COMMENT ON COLUMN public.topic.product_key IS '��ƷKey';
COMMENT ON COLUMN public.topic.topic_short_name IS '������';


-- public.tsl definition

-- Drop table

-- DROP TABLE public.tsl;

CREATE TABLE public.tsl (
	id int8 NOT NULL, -- ID
	product_key varchar NULL, -- ��ƷKey
	tsl text NULL, -- ��ģ��
	"version" int8 NULL, -- �汾��ʱ���
	CONSTRAINT tsl_fk FOREIGN KEY (product_key) REFERENCES product(product_key)
);
COMMENT ON TABLE public.tsl IS '��ģ�ͱ�';

-- Column comments

COMMENT ON COLUMN public.tsl.id IS 'ID';
COMMENT ON COLUMN public.tsl.product_key IS '��ƷKey';
COMMENT ON COLUMN public.tsl.tsl IS '��ģ��';
COMMENT ON COLUMN public.tsl."version" IS '�汾��ʱ���';