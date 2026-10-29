package api

import (
	"github.com/netbox-community/go-netbox/v3/netbox/client/ipam"
	"github.com/netbox-community/netbox-operator/pkg/netbox/models"
	"github.com/netbox-community/netbox-operator/pkg/netbox/utils"
)

func (c *NetboxCompositeClient) getIPAMRoleDetails(name string) (*models.IpamRole, error) {
	request := ipam.NewIpamRolesListParams().WithName(&name)
	response, err := c.clientV3.Ipam.IpamRolesList(request, nil)
	if err != nil {
		return nil, utils.NetboxError("failed to fetch IPAM Role details", err)
	}
	if len(response.Payload.Results) == 0 {
		return nil, utils.NetboxNotFoundError("ipam role '" + name + "'")
	}

	return &models.IpamRole{
		Id:   response.Payload.Results[0].ID,
		Slug: *response.Payload.Results[0].Slug,
		Name: *response.Payload.Results[0].Name,
	}, nil
}
