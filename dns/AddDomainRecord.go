package dns

import (
	"github.com/denverdino/aliyungo/common"
	"log"
)

type AddDomainRecordArgs struct {
	DomainName string
	RR         string
	Type       string
	Value      string

	//optional
	TTL      int32
	Priority int
	Line     string
}

type AddDomainRecordResponse struct {
	common.Response
	InstanceId string
	RecordId   string
}

// AddDomainRecord
//
// You can read doc at https://docs.aliyun.com/?spm=5176.100054.201.106.OeZ3dN#/pub/dns/api-reference/record-related&AddDomainRecord
func (client *Client) AddDomainRecord(args *AddDomainRecordArgs) (response *AddDomainRecordResponse, err error) {
	action := "AddDomainRecord"
	response = &AddDomainRecordResponse{}
	err = client.Invoke(action, args, response)
	if err == nil {
		return response, nil
	} else {
		log.Fatalf("%s error, %v", action, err)
		return response, err
	}
}
