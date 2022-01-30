// Code generated by goa v3.5.4, DO NOT EDIT.
//
// Tables client
//
// Command:
// $ goa gen github.com/lawrencejones/pgsink/api/design -o api

package tables

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "Tables" service client.
type Client struct {
	ListEndpoint goa.Endpoint
}

// NewClient initializes a "Tables" service client given the endpoints.
func NewClient(list goa.Endpoint) *Client {
	return &Client{
		ListEndpoint: list,
	}
}

// List calls the "List" endpoint of the "Tables" service.
func (c *Client) List(ctx context.Context, p *ListPayload) (res []*Table, err error) {
	var ires interface{}
	ires, err = c.ListEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.([]*Table), nil
}
