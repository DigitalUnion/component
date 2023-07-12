// This file is auto-generated, don't edit it. Thanks.
package duslb

import (
	"fmt"
	alb20200616 "github.com/alibabacloud-go/alb-20200616/v2/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"strconv"
	"time"
)

func AddAlbBackendServers(r *Request) error {
	ok, err := r.addAlbBackendCheck()
	if err != nil {
		return err
	}
	if !ok {
		// 缺少必传参数
		return fmt.Errorf("miss required parameter")
	}
	client, err := createAlbClient(r.AccessKeyId, r.AccessKeySecret)
	if err != nil {
		return err
	}
	port, err := strconv.Atoi(r.Port)
	if err != nil {
		return err
	}
	weight, err := strconv.Atoi(r.Weight)
	if err != nil {
		return err
	}
	servers0 := &alb20200616.AddServersToServerGroupRequestServers{
		Port:       tea.Int32(int32(port)),
		ServerId:   tea.String(r.ServerId),
		ServerType: tea.String("Ecs"),
		Weight:     tea.Int32(int32(weight)),
	}
	addServersToServerGroupRequest := &alb20200616.AddServersToServerGroupRequest{
		Servers:       []*alb20200616.AddServersToServerGroupRequestServers{servers0},
		ServerGroupId: tea.String(r.ServerGroupId),
	}
	_, err = client.AddServersToServerGroup(addServersToServerGroupRequest)
	return err
}

func DelAlbBackendServers(r *Request) error {
	ok, err := r.delAlbBackendCheck()
	if err != nil {
		return err
	}
	if !ok {
		// 缺少必传参数
		return fmt.Errorf("miss required parameter")
	}
	client, err := createAlbClient(r.AccessKeyId, r.AccessKeySecret)
	if err != nil {
		return err
	}
	port, err := strconv.Atoi(r.Port)
	if err != nil {
		return err
	}
	servers0 := &alb20200616.RemoveServersFromServerGroupRequestServers{
		Port:       tea.Int32(int32(port)),
		ServerId:   tea.String(r.ServerId),
		ServerType: tea.String("Ecs"),
	}
	removeServersFromServerGroupRequest := &alb20200616.RemoveServersFromServerGroupRequest{
		ServerGroupId: tea.String(r.ServerGroupId),
		Servers:       []*alb20200616.RemoveServersFromServerGroupRequestServers{servers0},
	}
	_, err = client.RemoveServersFromServerGroup(removeServersFromServerGroupRequest)
	
	// 由于ALB的RemoveServersFromServerGroup接口属于异步接口，即系统返回一个请求ID，但该后端服务器尚未移除成功，系统后台的移除任务仍在进行
	// 故调用删除接口后，需要等待一段时间，避免后续创建同样的后端服务器时，系统报错
	time.Sleep(5 * time.Second)
	return err
}

func createAlbClient(accessKeyId, accessKeySecret string) (*alb20200616.Client, error) {
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: &accessKeyId,
		// 您的AccessKey Secret
		AccessKeySecret: &accessKeySecret,
	}
	
	// 访问的域名
	config.Endpoint = tea.String("alb.cn-beijing.aliyuncs.com")
	return alb20200616.NewClient(config)
}
