package config

type Oss struct {
	Endpoint        string `mapstructure:"endpoint" yaml:"endpoint"`
	AccessKeyId     string `mapstructure:"accessKeyId" yaml:"accessKeyId"`
	AccessKeySecret string `mapstructure:"accessKeySecret" yaml:"accessKeySecret"`
	BucketName      string `mapstructure:"bucketName" yaml:"bucketName"`
	Host            string `mapstructure:"host" yaml:"host"`
}
