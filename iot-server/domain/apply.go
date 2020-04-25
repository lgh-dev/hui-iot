package domain

type Apply struct {
	Name    string `yaml:"id"`      //唯一标识
	Version string `yaml:"version"` //版本
	author  string `yaml:"author"`  //作者
	Desc    string `yaml:"desc"`    //描述
	Auth    Auth   `yaml:"auth"`    //权限信息
}

type Auth struct {
	AppKey string `yaml:"appKey"` // 应用账号
	Secret string `yaml:"secret"` // 应用密钥
}
