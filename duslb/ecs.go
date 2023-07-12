/*
 * The MIT License (MIT)
 *
 * Copyright (c) 2014 Manuel Martínez-Almeida
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 */

package duslb

import (
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v3/client"
	"github.com/alibabacloud-go/tea/tea"
	"net"
)

func createEcsClient(accessKeyId, accessKeySecret string) (*ecs20140526.Client, error) {
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: &accessKeyId,
		// 您的AccessKey Secret
		AccessKeySecret: &accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("ecs.cn-beijing.aliyuncs.com")
	return ecs20140526.NewClient(config)
}

func getInstanceID(accessKeyId, accessKeySecret, RegionId string) (string, error) {
	client, err := createEcsClient(accessKeyId, accessKeySecret)
	if err != nil {
		return "", err
	}
	privateIp := getPrivateIp()
	describeInstancesRequest := &ecs20140526.DescribeInstancesRequest{
		RegionId:           tea.String(RegionId),
		PrivateIpAddresses: tea.String(fmt.Sprintf("[\"%s\"]", privateIp)),
	}
	result, err := client.DescribeInstances(describeInstancesRequest)
	if err != nil {
		return "", err
	}
	if result != nil && result.Body != nil && result.Body.Instances != nil && result.Body.Instances.Instance != nil && len(result.Body.Instances.Instance) > 0 {
		return *result.Body.Instances.Instance[0].InstanceId, nil
	}
	return "", err

}

func getPrivateIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
