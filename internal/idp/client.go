package idp

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	ServerUrl string
}

func NewClient(
	protocol string,
	host string,
	port int,
) *Client {
	return &Client{
		ServerUrl: fmt.Sprintf(
			"%s://%s:%d",
			protocol,
			host,
			port,
		),
	}
}

func (c *Client) GetSuperiors(
	ehid string,
) (SuperiorResponseDto, error) {
	response, err := http.Get(c.ServerUrl + "/account/" + ehid + "/superiors")
	if err != nil {
		return SuperiorResponseDto{}, err
	}
	defer response.Body.Close()

	var superiorResponseDto SuperiorResponseDto
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return SuperiorResponseDto{}, err
	}

	json.Unmarshal(body, &superiorResponseDto)
	return superiorResponseDto, nil
}

func (c *Client) GetOrganization(
	id string,
) (OrganizationSingleResponseDto, error) {
	response, err := http.Get(c.ServerUrl + "/organizations/" + id)
	if err != nil {
		return OrganizationSingleResponseDto{}, err
	}
	defer response.Body.Close()

	var organizationDto OrganizationSingleResponseDto
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return OrganizationSingleResponseDto{}, err
	}

	json.Unmarshal(body, &organizationDto)
	return organizationDto, nil
}

func (c *Client) GetOrganizationMembers(
	id string,
) (OrganizationMemberResponseDto, error) {
	response, err := http.Get(c.ServerUrl + "/organizations/" + id + "/members")
	if err != nil {
		return OrganizationMemberResponseDto{}, err
	}
	defer response.Body.Close()

	var organizationMemberDto OrganizationMemberResponseDto
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return OrganizationMemberResponseDto{}, err
	}

	json.Unmarshal(body, &organizationMemberDto)
	return organizationMemberDto, nil
}
