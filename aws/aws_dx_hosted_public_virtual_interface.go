// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/directconnect"
)

func ListDxHostedPublicVirtualInterface(client *Client) ([]Resource, error) {
	req := client.directconnectconn.DescribeVirtualInterfacesRequest(&directconnect.DescribeVirtualInterfacesInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.VirtualInterfaces) > 0 {
		for _, r := range resp.VirtualInterfaces {
			tags := map[string]string{}
			for _, t := range r.Tags {
				tags[*t.Key] = *t.Value
			}

			result = append(result, Resource{
				Type: "aws_dx_hosted_public_virtual_interface",
				ID:   *r.VirtualInterfaceId,
				Tags: tags,
			})
		}
	}

	return result, nil
}
