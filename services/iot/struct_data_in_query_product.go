package iot

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

// DataInQueryProduct is a nested struct in iot response
type DataInQueryProduct struct {
	GmtCreate           int64  `json:"GmtCreate" xml:"GmtCreate"`
	DataFormat          int    `json:"DataFormat" xml:"DataFormat"`
	Description         string `json:"Description" xml:"Description"`
	DeviceCount         int    `json:"DeviceCount" xml:"DeviceCount"`
	NodeType            int    `json:"NodeType" xml:"NodeType"`
	ProductKey          string `json:"ProductKey" xml:"ProductKey"`
	ProductName         string `json:"ProductName" xml:"ProductName"`
	ProductSecret       string `json:"ProductSecret" xml:"ProductSecret"`
	CategoryName        string `json:"CategoryName" xml:"CategoryName"`
	CategoryKey         string `json:"CategoryKey" xml:"CategoryKey"`
	AliyunCommodityCode string `json:"AliyunCommodityCode" xml:"AliyunCommodityCode"`
	Id2                 bool   `json:"Id2" xml:"Id2"`
	ProtocolType        string `json:"ProtocolType" xml:"ProtocolType"`
	ProductStatus       string `json:"ProductStatus" xml:"ProductStatus"`
	Owner               bool   `json:"Owner" xml:"Owner"`
	NetType             int    `json:"NetType" xml:"NetType"`
	AuthType            string `json:"AuthType" xml:"AuthType"`
}
