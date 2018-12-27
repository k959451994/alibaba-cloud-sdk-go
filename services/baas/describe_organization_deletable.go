package baas

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

// DescribeOrganizationDeletable invokes the baas.DescribeOrganizationDeletable API synchronously
// api document: https://help.aliyun.com/api/baas/describeorganizationdeletable.html
func (client *Client) DescribeOrganizationDeletable(request *DescribeOrganizationDeletableRequest) (response *DescribeOrganizationDeletableResponse, err error) {
	response = CreateDescribeOrganizationDeletableResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeOrganizationDeletableWithChan invokes the baas.DescribeOrganizationDeletable API asynchronously
// api document: https://help.aliyun.com/api/baas/describeorganizationdeletable.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeOrganizationDeletableWithChan(request *DescribeOrganizationDeletableRequest) (<-chan *DescribeOrganizationDeletableResponse, <-chan error) {
	responseChan := make(chan *DescribeOrganizationDeletableResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeOrganizationDeletable(request)
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

// DescribeOrganizationDeletableWithCallback invokes the baas.DescribeOrganizationDeletable API asynchronously
// api document: https://help.aliyun.com/api/baas/describeorganizationdeletable.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeOrganizationDeletableWithCallback(request *DescribeOrganizationDeletableRequest, callback func(response *DescribeOrganizationDeletableResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeOrganizationDeletableResponse
		var err error
		defer close(result)
		response, err = client.DescribeOrganizationDeletable(request)
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

// DescribeOrganizationDeletableRequest is the request struct for api DescribeOrganizationDeletable
type DescribeOrganizationDeletableRequest struct {
	*requests.RpcRequest
	OrganizationId string `position:"Query" name:"OrganizationId"`
	Location       string `position:"Body" name:"Location"`
}

// DescribeOrganizationDeletableResponse is the response struct for api DescribeOrganizationDeletable
type DescribeOrganizationDeletableResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateDescribeOrganizationDeletableRequest creates a request to invoke DescribeOrganizationDeletable API
func CreateDescribeOrganizationDeletableRequest() (request *DescribeOrganizationDeletableRequest) {
	request = &DescribeOrganizationDeletableRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Baas", "2018-07-31", "DescribeOrganizationDeletable", "", "")
	return
}

// CreateDescribeOrganizationDeletableResponse creates a response to parse from DescribeOrganizationDeletable response
func CreateDescribeOrganizationDeletableResponse() (response *DescribeOrganizationDeletableResponse) {
	response = &DescribeOrganizationDeletableResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}