package mns

type MnsTopic struct {
	Url             string `yaml:"url"`
	AccessKeyId     string `yaml:"AccessKeyId"`
	AccessKeySecret string `yaml:"AccessKeySecret"`
	Topic           string `yaml:"Topic"`
}

type InstanceInfo struct {
	RegionId string `yaml:"RegionId"`
	Endpoint string `yaml:"Endpoint"`
}

var MNS = MnsTopic{}

var EcsInfo = InstanceInfo{}
