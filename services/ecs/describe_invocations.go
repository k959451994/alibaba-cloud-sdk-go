package ecs

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

func (client *Client) DescribeInvocations(request *DescribeInvocationsRequest) (response *DescribeInvocationsResponse, err error) {
	response = CreateDescribeInvocationsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeInvocationsWithChan(request *DescribeInvocationsRequest) (<-chan *DescribeInvocationsResponse, <-chan error) {
	responseChan := make(chan *DescribeInvocationsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInvocations(request)
		responseChan <- response
		errChan <- err
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) DescribeInvocationsWithCallback(request *DescribeInvocationsRequest, callback func(response *DescribeInvocationsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInvocationsResponse
		var err error
		defer close(result)
		response, err = client.DescribeInvocations(request)
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

type DescribeInvocationsRequest struct {
	*requests.RpcRequest
	PageSize             string `position:"Query" name:"PageSize"`
	InvokeId             string `position:"Query" name:"InvokeId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	CommandName          string `position:"Query" name:"CommandName"`
	Timed                string `position:"Query" name:"Timed"`
	PageNumber           string `position:"Query" name:"PageNumber"`
	CommandId            string `position:"Query" name:"CommandId"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	InvokeStatus         string `position:"Query" name:"InvokeStatus"`
	CommandType          string `position:"Query" name:"CommandType"`
	InstanceId           string `position:"Query" name:"InstanceId"`
}

type DescribeInvocationsResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	TotalCount int64  `json:"TotalCount" xml:"TotalCount"`
	PageNumber int64  `json:"PageNumber" xml:"PageNumber"`
	PageSize   int64  `json:"PageSize" xml:"PageSize"`
	Invocation struct {
		InvocationItem []struct {
			InvokeId     string `json:"InvokeId" xml:"InvokeId"`
			CommandId    string `json:"CommandId" xml:"CommandId"`
			CommandType  string `json:"CommandType" xml:"CommandType"`
			CommandName  string `json:"CommandName" xml:"CommandName"`
			Timed        bool   `json:"Timed" xml:"Timed"`
			InvokeStatus string `json:"InvokeStatus" xml:"InvokeStatus"`
			InvokeItem   struct {
				InvokeItemItem []struct {
					InstanceId string `json:"InstanceId" xml:"InstanceId"`
					Status     string `json:"Status" xml:"Status"`
				} `json:"InvokeItemItem" xml:"InvokeItemItem"`
			} `json:"InvokeItem" xml:"InvokeItem"`
		} `json:"InvocationItem" xml:"InvocationItem"`
	} `json:"Invocation" xml:"Invocation"`
}

func CreateDescribeInvocationsRequest() (request *DescribeInvocationsRequest) {
	request = &DescribeInvocationsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeInvocations", "", "")
	return
}

func CreateDescribeInvocationsResponse() (response *DescribeInvocationsResponse) {
	response = &DescribeInvocationsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}