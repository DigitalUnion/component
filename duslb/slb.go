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
	"encoding/json"
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	slb20140515 "github.com/alibabacloud-go/slb-20140515/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"time"
)

type server struct {
	ServerId    string `json:"ServerId,omitempty"`
	Weight      string `json:"Weight,omitempty"`
	Type        string `json:"Type,omitempty"`
	Port        string `json:"Port,omitempty"`
	Description string `json:"Description,omitempty"`
}

type Request struct {
	AccessKeyId     string // AccessKeyId （必须）
	AccessKeySecret string // AccessKeySecret （必须）
	RegionId        string // 地域 id （必须）
	LoadBalancerId  string // 负载均衡 id （操作 lb 默认服务器组时必须）
	VServerGroupId  string // 虚拟组 id （操作 lb 虚拟服务器组时必须）
	ServerGroupId   string // alb 服务器组 id （操作 lb 虚拟服务器组时必须）
	Port            string // 端口号（操作 lb 虚拟服务器组时必须）
	Weight          string // 权重（添加操作时必须）
	ServerId        string // ecs_instanceID（非必需，默认获取当前服务器 instanceId）
}

func (r *Request) setServerId() error {
	if r.ServerId != "" {
		return nil
	}
	var err error
	r.ServerId, err = getInstanceID(r.AccessKeyId, r.AccessKeySecret, r.RegionId)
	return err
}

func (r *Request) addVServerGroupCheck() (bool, error) {
	err := r.setServerId()
	if err != nil {
		return false, err
	}
	if r.AccessKeySecret == "" || r.AccessKeyId == "" {
		return false, nil
	}
	if r.RegionId == "" || r.VServerGroupId == "" || r.Port == "" || r.Weight == "" || r.ServerId == "" {
		return false, nil
	}
	return true, nil
}

func (r *Request) addAlbBackendCheck() (bool, error) {
	err := r.setServerId()
	if err != nil {
		return false, err
	}
	if r.AccessKeySecret == "" || r.AccessKeyId == "" {
		return false, err
	}
	if r.Port == "" || r.ServerGroupId == "" || r.Weight == "" || r.ServerId == "" {
		return false, err
	}
	return true, err
}

func (r *Request) addBackendCheck() (bool, error) {
	err := r.setServerId()
	if err != nil {
		return false, err
	}
	if r.AccessKeySecret == "" || r.AccessKeyId == "" {
		return false, err
	}
	if r.RegionId == "" || r.LoadBalancerId == "" || r.Weight == "" || r.ServerId == "" {
		return false, err
	}
	return true, err
}

func (r *Request) delVServerGroupCheck() (bool, error) {
	err := r.setServerId()
	if err != nil {
		return false, err
	}
	if r.AccessKeySecret == "" || r.AccessKeyId == "" {
		return false, nil
	}
	if r.RegionId == "" || r.VServerGroupId == "" || r.Port == "" || r.ServerId == "" {
		return false, nil
	}
	return true, nil
}

func (r *Request) delBackendCheck() (bool, error) {
	err := r.setServerId()
	if err != nil {
		return false, err
	}
	if r.AccessKeySecret == "" || r.AccessKeyId == "" {
		return false, nil
	}
	if r.RegionId == "" || r.LoadBalancerId == "" || r.ServerId == "" {
		return false, nil
	}
	return true, nil
}

func (r *Request) delAlbBackendCheck() (bool, error) {
	err := r.setServerId()
	if err != nil {
		return false, err
	}
	if r.AccessKeySecret == "" || r.AccessKeyId == "" {
		return false, nil
	}
	if r.Port == "" || r.ServerGroupId == "" || r.ServerId == "" {
		return false, nil
	}
	return true, nil
}

func createClient(accessKeyId, accessKeySecret string) (*slb20140515.Client, error) {
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: &accessKeyId,
		// 您的AccessKey Secret
		AccessKeySecret: &accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("slb.aliyuncs.com")
	return slb20140515.NewClient(config)
}

type SlbConfig struct {
	AccessKey     string
	SecretKey     string
	RegionId      string
	LoadBalancers []string
	VServerGroups []string
	ServerGroups  []string
	Port          string
}

// OnStartup 服务启动时执行,将此节点挂载到SLB上
func OnStartup(cfg SlbConfig) {
	for _, e := range cfg.LoadBalancers {
		r := &Request{
			AccessKeyId:     cfg.AccessKey,
			AccessKeySecret: cfg.SecretKey,
			RegionId:        cfg.RegionId,
			LoadBalancerId:  e,
			Weight:          "50",
		}
		start := time.Now()
		fmt.Printf("AddBackendServers:%+v", r)
		err := AddBackendServers(r)
		if err != nil {
			fmt.Println(err.Error())
		}
		Notice("AddBackendServers", e, r.ServerId, start, err)
	}
	for _, e := range cfg.VServerGroups {
		r := &Request{
			AccessKeyId:     cfg.AccessKey,
			AccessKeySecret: cfg.SecretKey,
			RegionId:        cfg.RegionId,
			VServerGroupId:  e,
			Weight:          "50",
			Port:            cfg.Port,
		}
		start := time.Now()
		fmt.Printf("AddVServerGroupBackendServers:%+v", r)
		err := AddVServerGroupBackendServers(r)
		if err != nil {
			fmt.Println(err.Error())
		}
		Notice("AddVServerGroupBackendServers", e, r.ServerId, start, err)
	}

	for _, e := range cfg.ServerGroups {
		r := &Request{
			AccessKeyId:     cfg.AccessKey,
			AccessKeySecret: cfg.SecretKey,
			RegionId:        cfg.RegionId,
			ServerGroupId:   e,
			Port:            cfg.Port,
			Weight:          "50",
		}
		start := time.Now()
		fmt.Printf("AddAlbBackendServers:%+v", r)
		err := AddAlbBackendServers(r)
		if err != nil {
			fmt.Println(err.Error())
		}
		Notice("AddAlbBackendServers", e, r.ServerId, start, err)
	}
}

// OnShutdown 服务关闭时执行,将此节点从SLB摘除
func OnShutdown(cfg SlbConfig) {
	for _, e := range cfg.LoadBalancers {
		r := &Request{
			AccessKeyId:     cfg.AccessKey,
			AccessKeySecret: cfg.SecretKey,
			RegionId:        cfg.RegionId,
			LoadBalancerId:  e,
			Weight:          "50",
		}
		start := time.Now()
		fmt.Printf("DelBackendServers:%+v", r)
		err := DelBackendServers(r)
		if err != nil {
			fmt.Println(err.Error())
		}
		Notice("DelBackendServers", e, r.ServerId, start, err)
	}
	for _, e := range cfg.VServerGroups {
		r := &Request{
			AccessKeyId:     cfg.AccessKey,
			AccessKeySecret: cfg.SecretKey,
			RegionId:        cfg.RegionId,
			VServerGroupId:  e,
			Weight:          "50",
			Port:            cfg.Port,
		}
		start := time.Now()
		fmt.Printf("DelVServerGroupBackendServers:%+v", r)
		err := DelVServerGroupBackendServers(r)
		if err != nil {
			fmt.Println(err.Error())
		}
		Notice("DelVServerGroupBackendServers", e, r.ServerId, start, err)
	}
	for _, e := range cfg.ServerGroups {
		r := &Request{
			AccessKeyId:     cfg.AccessKey,
			AccessKeySecret: cfg.SecretKey,
			RegionId:        cfg.RegionId,
			ServerGroupId:   e,
			Port:            cfg.Port,
		}
		start := time.Now()
		fmt.Printf("DelAlbBackendServers:%+v", r)
		err := DelAlbBackendServers(r)
		if err != nil {
			fmt.Println(err.Error())
		}
		Notice("DelAlbBackendServers", e, r.ServerId, start, err)
	}
}
func AddEndServer(r *Request) error {
	if r.VServerGroupId != "" {
		return AddVServerGroupBackendServers(r)
	} else if r.ServerGroupId != "" {
		return AddAlbBackendServers(r)
	} else {
		return AddBackendServers(r)
	}
}

func AddVServerGroupBackendServers(r *Request) error {
	ok, err := r.addVServerGroupCheck()
	if err != nil {
		return err
	}
	if !ok {
		// 缺少必传参数
		return fmt.Errorf("miss required parameter")
	}
	client, err := createClient(r.AccessKeyId, r.AccessKeySecret)
	if err != nil {
		return err
	}
	backendServers := []server{
		{
			ServerId: r.ServerId,
			Port:     r.Port,
			Weight:   r.Weight,
		},
	}
	server, _ := json.Marshal(backendServers)
	req := &slb20140515.AddVServerGroupBackendServersRequest{
		RegionId:       tea.String(r.RegionId),
		VServerGroupId: tea.String(r.VServerGroupId),
		BackendServers: tea.String(string(server)),
	}
	_, err = client.AddVServerGroupBackendServers(req)
	if err != nil {
		return err
	}
	return err
}

func AddBackendServers(r *Request) error {
	ok, err := r.addBackendCheck()
	if err != nil {
		return err
	}
	if !ok {
		// 缺少必传参数
		return fmt.Errorf("miss required parameter")
	}
	client, err := createClient(r.AccessKeyId, r.AccessKeySecret)
	if err != nil {
		return err
	}
	backendServers := []server{
		{
			ServerId: r.ServerId,
			Weight:   r.Weight,
		},
	}
	server, _ := json.Marshal(backendServers)
	req := &slb20140515.AddBackendServersRequest{
		RegionId:       tea.String(r.RegionId),
		LoadBalancerId: tea.String(r.LoadBalancerId),
		BackendServers: tea.String(string(server)),
	}
	_, err = client.AddBackendServers(req)
	return err
}

func DelVServerGroupBackendServers(r *Request) error {
	ok, err := r.delVServerGroupCheck()
	if err != nil {
		return err
	}
	if !ok {
		// 缺少必传参数
		return fmt.Errorf("miss required parameter")
	}
	client, err := createClient(r.AccessKeyId, r.AccessKeySecret)
	if err != nil {
		return err
	}
	backendServers := []server{
		{
			ServerId: r.ServerId,
			Port:     r.Port,
		},
	}
	server, _ := json.Marshal(backendServers)

	req := &slb20140515.RemoveVServerGroupBackendServersRequest{
		RegionId:       tea.String(r.RegionId),
		VServerGroupId: tea.String(r.VServerGroupId),
		BackendServers: tea.String(string(server)),
	}
	_, err = client.RemoveVServerGroupBackendServers(req)
	return err
}

func DelEndServer(r *Request) error {
	if r.VServerGroupId != "" {
		return DelVServerGroupBackendServers(r)
	} else if r.ServerGroupId != "" {
		return DelAlbBackendServers(r)
	} else {
		return DelBackendServers(r)
	}
}

func DelBackendServers(r *Request) error {
	ok, err := r.delBackendCheck()
	if err != nil {
		return err
	}
	if !ok {
		// 缺少必传参数
		return fmt.Errorf("miss required parameter")
	}
	client, err := createClient(r.AccessKeyId, r.AccessKeySecret)
	if err != nil {
		return err
	}
	backendServers := []server{
		{
			ServerId: r.ServerId,
		},
	}
	server, _ := json.Marshal(backendServers)
	req := &slb20140515.RemoveBackendServersRequest{
		RegionId:       tea.String(r.RegionId),
		LoadBalancerId: tea.String(r.LoadBalancerId),
		BackendServers: tea.String(string(server)),
	}
	_, err = client.RemoveBackendServers(req)
	return err
}
