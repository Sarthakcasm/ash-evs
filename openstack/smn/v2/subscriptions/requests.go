package subscriptions

import (
	"github.com/gophercloud/gophercloud"
)

var RequestOpts gophercloud.RequestOpts = gophercloud.RequestOpts{
	MoreHeaders: map[string]string{"Content-Type": "application/json", "X-Language": "en-us"},
}

//CreateOpsBuilder is used for creating subscription parameters.
//any struct providing the parameters should implement this interface
type CreateOpsBuilder interface {
	ToSubscriptionCreateMap() (map[string]interface{}, error)
}

//CreateOps is a struct that contains all the parameters.
type CreateOps struct {
	//Message endpoint
	Endpoint            string         `json:"endpoint" required:"true"`
    //Protocol of the message endpoint
	Protocol            string         `json:"protocol" required:"true"`
	//Description of the subscription
	Remark              string         `json:"remark,omitempty"`
}


func (ops CreateOps) ToSubscriptionCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(ops, "")
}


//Create a subscription with given parameters.
func Create(client *gophercloud.ServiceClient, ops CreateOpsBuilder, topicUrn string) (r CreateResult) {
	b, err := ops.ToSubscriptionCreateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(createURL(client, topicUrn), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes:     []int{201, 200},
		MoreHeaders: RequestOpts.MoreHeaders,
	})

	return
}

//delete a subscription via subscription urn
func Delete(client *gophercloud.ServiceClient, subscriptionUrn string) (r DeleteResult) {
	_, r.Err = client.Delete(deleteURL(client, subscriptionUrn), &RequestOpts)
	return
}

//get a subscription with detailed information by subscription urn
//func Get(client *gophercloud.ServiceClient, subscriptionUrn string) (r GetResult) {
//	_, r.Err = client.Get(getURL(client, subscriptionUrn), &r.Body, &RequestOpts)
//	return
//}

//list all the subscriptions
func List(client *gophercloud.ServiceClient) (r ListResult) {
	_, r.Err = client.Get(listURL(client), &r.Body, &RequestOpts)
	return
}

//list all the subscriptions
func ListFromTopic(client *gophercloud.ServiceClient, subscriptionUrn string) (r ListResult) {
	_, r.Err = client.Get(listFromTopicURL(client, subscriptionUrn), &r.Body, &RequestOpts)
	return
}
