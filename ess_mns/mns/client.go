package mns

import (
	ali_mns "github.com/aliyun/aliyun-mns-go-sdk"
	jsoniter "github.com/json-iterator/go"
	"log"
	"strings"
	"sync/atomic"
	"time"
)

var (
	tc                  = time.NewTicker(30 * time.Second)
	json_iterator       = jsoniter.ConfigCompatibleWithStandardLibrary
	sendNow       int64 = 0
	ready         int64 = 0
	sending       int64 = 1
)

const (
	MNS_QUEUE_PRE = "MnsQueue"
	MNS_SUB_PRE   = "MnsSubscriptionName"
)

type MnsClient struct {
	client       ali_mns.MNSClient
	topic        ali_mns.AliMNSTopic
	queueManager ali_mns.AliQueueManager
	queue        ali_mns.AliMNSQueue
}

var respChan = make(chan ali_mns.MessageReceiveResponse)
var errChan = make(chan error)

func MnsInit() (*MnsClient, error) {
	mnsCient := new(MnsClient)
	mnsCient.client = ali_mns.NewAliMNSClient(MNS.Url,
		MNS.AccessKeyId,
		MNS.AccessKeySecret)

	// 创建队列，队列名称为testQueue。 队列名称只能是数字和字母
	mnsCient.queueManager = ali_mns.NewMNSQueueManager(mnsCient.client)
	err := mnsCient.queueManager.CreateSimpleQueue(getQueueName())
	if err != nil && !ali_mns.ERR_MNS_QUEUE_ALREADY_EXIST_AND_HAVE_SAME_ATTR.IsEqual(err) {
		return nil, err
	}

	// 订阅主题，本示例中的endpoint设置为队列名称。
	mnsCient.topic = ali_mns.NewMNSTopic(MNS.Topic, mnsCient.client)
	sub := ali_mns.MessageSubsribeRequest{
		Endpoint:            mnsCient.topic.GenerateQueueEndpoint(getQueueName()),
		NotifyContentFormat: ali_mns.SIMPLIFIED,
	}

	err = mnsCient.topic.Subscribe(getSubscriptionName(), sub)
	if err != nil && !ali_mns.ERR_MNS_SUBSCRIPTION_ALREADY_EXIST_AND_HAVE_SAME_ATTR.IsEqual(err) {
		return nil, err
	}

	mnsCient.queue = ali_mns.NewMNSQueue(getQueueName(), mnsCient.client)
	go mnsCient.startReceiveMessage()
	return mnsCient, nil
}

func (mns *MnsClient) ReceiveMsg() {
	go func() {
		for {
			select {
			case resp := <-respChan:
				{
					err := mns.parseBody(resp.MessageBody)
					if err != nil {
						log.Println("parseBody err:", err)
					}
					// 报错忽略，删除消息
					if ret, e := mns.queue.ChangeMessageVisibility(resp.ReceiptHandle, 5); e != nil {
						log.Println("ChangeMessageVisibility err:", e)
					} else {
						log.Println("visibility changed", ret)
						log.Println("delete it now: ", ret.ReceiptHandle)
						if e := mns.queue.DeleteMessage(ret.ReceiptHandle); e != nil {
							log.Println("DeleteMessage err:", e)
						}
					}
				}
			case err := <-errChan:
				{
					if !strings.Contains(err.Error(), "State not exist") && !strings.Contains(err.Error(), "Message not exist") {
						log.Println("ReceiveMsg err:", err)
					}
				}
			}
		}
	}()
}

func (mns *MnsClient) startReceiveMessage() {
	for range tc.C {
		flag := atomic.CompareAndSwapInt64(&sendNow, ready, sending)
		if !flag {
			continue
		}
		mns.queue.ReceiveMessage(respChan, errChan, 2)
		atomic.CompareAndSwapInt64(&sendNow, sending, ready)
	}
}

func (mns *MnsClient) Close() {
	mns.topic.Unsubscribe(getSubscriptionName())
	mns.queueManager.DeleteQueue(getQueueName())
}

type Content struct {
	DefaultResult        string   `json:"defaultResult"`
	InstanceIds          []string `json:"instanceIds"`
	LifecycleActionToken string   `json:"lifecycleActionToken"`
	LifecycleHookId      string   `json:"lifecycleHookId"`
	LifecycleHookName    string   `json:"lifecycleHookName"`
	LifecycleTransition  string   `json:"lifecycleTransition"`
	NotificationMetadata string   `json:"notificationMetadata"`
	RequestId            string   `json:"requestId"`
	ScalingActivityId    string   `json:"scalingActivityId"`
	ScalingGroupId       string   `json:"scalingGroupId"`
	ScalingGroupName     string   `json:"scalingGroupName"`
	ScalingRuleId        string   `json:"scalingRuleId"`
}
type MsgBody struct {
	Content     Content   `json:"content"`
	Product     string    `json:"product"`
	RegionId    string    `json:"regionId"`
	ResourceArn string    `json:"resourceArn"`
	Time        time.Time `json:"time"`
	UserId      string    `json:"userId"`
}

func (mns *MnsClient) parseBody(content string) error {
	content, err := Base64Decode(content)
	if err != nil {
		return err
	}
	log.Println("msg:", content)

	msg := new(MsgBody)
	err = json_iterator.UnmarshalFromString(content, &msg)
	if err != nil || len(msg.Content.InstanceIds) == 0 {
		return err
	}

	if msg.Content.LifecycleTransition != "SCALE_IN" {
		return nil
	}

	// 获取通知实例的 ip
	instandInfo, err := GetInstanceIdInfo(msg.Content.InstanceIds)
	if err != nil {
		return err
	}

	ip := "\"" + GetIp() + "\""
	if !strings.Contains(instandInfo, ip) {
		return nil
	}

	log.Println("current node scalein")
	State = "MovingOut"

	// 当前机器缩容，做善后处理
	mns.Close()
	return err
}

func getQueueName() string {
	return MNS_QUEUE_PRE + strings.Replace(GetIp(), ".", "s", -1)
}

func getSubscriptionName() string {
	return MNS_SUB_PRE + strings.Replace(GetIp(), ".", "s", -1)
}
