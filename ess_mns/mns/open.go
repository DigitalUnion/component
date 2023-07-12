package mns

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v2/client"
	"github.com/alibabacloud-go/tea/tea"
)

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *ecs20140526.Client, _err error) {
	configApi := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	configApi.Endpoint = tea.String(EcsInfo.Endpoint)
	_result = &ecs20140526.Client{}
	_result, _err = ecs20140526.NewClient(configApi)
	return _result, _err
}

// 通过实例 id 获取实例信息，用来区分当家机器 ip
func GetInstanceIdInfo(instanceIds []string) (res string, _err error) {
	instanceIdsJson, err := json_iterator.MarshalToString(instanceIds)
	if err != nil {
		return
	}

	client, _err := CreateClient(tea.String(MNS.AccessKeyId), tea.String(MNS.AccessKeySecret))
	if _err != nil {
		return "", _err
	}

	describeInstancesRequest := &ecs20140526.DescribeInstancesRequest{
		RegionId:    tea.String(EcsInfo.RegionId),
		InstanceIds: tea.String(instanceIdsJson),
	}

	tmp, _err := client.DescribeInstances(describeInstancesRequest)
	if _err != nil {
		return "", _err
	}
	return tmp.String(), _err
}
