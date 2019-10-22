package push

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// MassPush invokes the push.MassPush API synchronously
// api document: https://help.aliyun.com/api/push/masspush.html
func (client *Client) MassPush(request *MassPushRequest) (response *MassPushResponse, err error) {
	response = CreateMassPushResponse()
	err = client.DoAction(request, response)
	return
}

// MassPushWithChan invokes the push.MassPush API asynchronously
// api document: https://help.aliyun.com/api/push/masspush.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) MassPushWithChan(request *MassPushRequest) (<-chan *MassPushResponse, <-chan error) {
	responseChan := make(chan *MassPushResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.MassPush(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// MassPushWithCallback invokes the push.MassPush API asynchronously
// api document: https://help.aliyun.com/api/push/masspush.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) MassPushWithCallback(request *MassPushRequest, callback func(response *MassPushResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *MassPushResponse
		var err error
		defer close(result)
		response, err = client.MassPush(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// MassPushRequest is the request struct for api MassPush
type MassPushRequest struct {
	*requests.RpcRequest
	PushTask *[]MassPushPushTask `position:"Query" name:"PushTask"  type:"Repeated"`
	AppKey   requests.Integer    `position:"Query" name:"AppKey"`
}

// MassPushPushTask is a repeated param struct in MassPushRequest
type MassPushPushTask struct {
	AndroidNotificationBarType     string `name:"AndroidNotificationBarType"`
	AndroidExtParameters           string `name:"AndroidExtParameters"`
	IOSBadge                       string `name:"iOSBadge"`
	IOSBadgeAutoIncrement          string `name:"iOSBadgeAutoIncrement"`
	AndroidOpenType                string `name:"AndroidOpenType"`
	Title                          string `name:"Title"`
	Body                           string `name:"Body"`
	DeviceType                     string `name:"DeviceType"`
	PushTime                       string `name:"PushTime"`
	SendSpeed                      string `name:"SendSpeed"`
	AndroidPopupActivity           string `name:"AndroidPopupActivity"`
	IOSRemindBody                  string `name:"iOSRemindBody"`
	IOSExtParameters               string `name:"iOSExtParameters"`
	AndroidNotifyType              string `name:"AndroidNotifyType"`
	AndroidPopupTitle              string `name:"AndroidPopupTitle"`
	IOSMusic                       string `name:"iOSMusic"`
	IOSApnsEnv                     string `name:"iOSApnsEnv"`
	IOSMutableContent              string `name:"iOSMutableContent"`
	AndroidNotificationBarPriority string `name:"AndroidNotificationBarPriority"`
	ExpireTime                     string `name:"ExpireTime"`
	AndroidPopupBody               string `name:"AndroidPopupBody"`
	IOSNotificationCategory        string `name:"iOSNotificationCategory"`
	StoreOffline                   string `name:"StoreOffline"`
	IOSSilentNotification          string `name:"iOSSilentNotification"`
	JobKey                         string `name:"JobKey"`
	Target                         string `name:"Target"`
	AndroidOpenUrl                 string `name:"AndroidOpenUrl"`
	AndroidNotificationChannel     string `name:"AndroidNotificationChannel"`
	AndroidRemind                  string `name:"AndroidRemind"`
	AndroidActivity                string `name:"AndroidActivity"`
	AndroidXiaoMiNotifyBody        string `name:"AndroidXiaoMiNotifyBody"`
	IOSSubtitle                    string `name:"iOSSubtitle"`
	IOSRemind                      string `name:"iOSRemind"`
	TargetValue                    string `name:"TargetValue"`
	AndroidMusic                   string `name:"AndroidMusic"`
	AndroidXiaoMiActivity          string `name:"AndroidXiaoMiActivity"`
	AndroidXiaoMiNotifyTitle       string `name:"AndroidXiaoMiNotifyTitle"`
	PushType                       string `name:"PushType"`
}

// MassPushResponse is the response struct for api MassPush
type MassPushResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	MessageIds MessageIds `json:"MessageIds" xml:"MessageIds"`
}

// CreateMassPushRequest creates a request to invoke MassPush API
func CreateMassPushRequest() (request *MassPushRequest) {
	request = &MassPushRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Push", "2016-08-01", "MassPush", "", "")
	return
}

// CreateMassPushResponse creates a response to parse from MassPush response
func CreateMassPushResponse() (response *MassPushResponse) {
	response = &MassPushResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}