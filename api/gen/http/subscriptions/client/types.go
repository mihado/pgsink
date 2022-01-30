// Code generated by goa v3.5.4, DO NOT EDIT.
//
// Subscriptions HTTP client types
//
// Command:
// $ goa gen github.com/lawrencejones/pgsink/api/design -o api

package client

import (
	subscriptions "github.com/lawrencejones/pgsink/api/gen/subscriptions"
	goa "goa.design/goa/v3/pkg"
)

// AddTableRequestBody is the type of the "Subscriptions" service "AddTable"
// endpoint HTTP request body.
type AddTableRequestBody struct {
	// Postgres table schema
	Schema string `form:"schema" json:"schema" xml:"schema"`
	// Postgres table name
	Name string `form:"name" json:"name" xml:"name"`
}

// StopTableRequestBody is the type of the "Subscriptions" service "StopTable"
// endpoint HTTP request body.
type StopTableRequestBody struct {
	// Postgres table schema
	Schema string `form:"schema" json:"schema" xml:"schema"`
	// Postgres table name
	Name string `form:"name" json:"name" xml:"name"`
}

// GetResponseBody is the type of the "Subscriptions" service "Get" endpoint
// HTTP response body.
type GetResponseBody struct {
	// ID of subscription
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// List of published tables
	PublishedTables []*SubscriptionPublishedTableResponseBody `form:"published_tables,omitempty" json:"published_tables,omitempty" xml:"published_tables,omitempty"`
}

// AddTableResponseBody is the type of the "Subscriptions" service "AddTable"
// endpoint HTTP response body.
type AddTableResponseBody struct {
	// ID of subscription
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// List of published tables
	PublishedTables []*SubscriptionPublishedTableResponseBody `form:"published_tables,omitempty" json:"published_tables,omitempty" xml:"published_tables,omitempty"`
}

// StopTableResponseBody is the type of the "Subscriptions" service "StopTable"
// endpoint HTTP response body.
type StopTableResponseBody struct {
	// ID of subscription
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// List of published tables
	PublishedTables []*SubscriptionPublishedTableResponseBody `form:"published_tables,omitempty" json:"published_tables,omitempty" xml:"published_tables,omitempty"`
}

// SubscriptionPublishedTableResponseBody is used to define fields on response
// body types.
type SubscriptionPublishedTableResponseBody struct {
	// Postgres table schema
	Schema *string `form:"schema,omitempty" json:"schema,omitempty" xml:"schema,omitempty"`
	// Postgres table name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// NewAddTableRequestBody builds the HTTP request body from the payload of the
// "AddTable" endpoint of the "Subscriptions" service.
func NewAddTableRequestBody(p *subscriptions.SubscriptionPublishedTable) *AddTableRequestBody {
	body := &AddTableRequestBody{
		Schema: p.Schema,
		Name:   p.Name,
	}
	return body
}

// NewStopTableRequestBody builds the HTTP request body from the payload of the
// "StopTable" endpoint of the "Subscriptions" service.
func NewStopTableRequestBody(p *subscriptions.SubscriptionPublishedTable) *StopTableRequestBody {
	body := &StopTableRequestBody{
		Schema: p.Schema,
		Name:   p.Name,
	}
	return body
}

// NewGetSubscriptionCreated builds a "Subscriptions" service "Get" endpoint
// result from a HTTP "Created" response.
func NewGetSubscriptionCreated(body *GetResponseBody) *subscriptions.Subscription {
	v := &subscriptions.Subscription{
		ID: *body.ID,
	}
	v.PublishedTables = make([]*subscriptions.SubscriptionPublishedTable, len(body.PublishedTables))
	for i, val := range body.PublishedTables {
		v.PublishedTables[i] = unmarshalSubscriptionPublishedTableResponseBodyToSubscriptionsSubscriptionPublishedTable(val)
	}

	return v
}

// NewAddTableSubscriptionCreated builds a "Subscriptions" service "AddTable"
// endpoint result from a HTTP "Created" response.
func NewAddTableSubscriptionCreated(body *AddTableResponseBody) *subscriptions.Subscription {
	v := &subscriptions.Subscription{
		ID: *body.ID,
	}
	v.PublishedTables = make([]*subscriptions.SubscriptionPublishedTable, len(body.PublishedTables))
	for i, val := range body.PublishedTables {
		v.PublishedTables[i] = unmarshalSubscriptionPublishedTableResponseBodyToSubscriptionsSubscriptionPublishedTable(val)
	}

	return v
}

// NewStopTableSubscriptionAccepted builds a "Subscriptions" service
// "StopTable" endpoint result from a HTTP "Accepted" response.
func NewStopTableSubscriptionAccepted(body *StopTableResponseBody) *subscriptions.Subscription {
	v := &subscriptions.Subscription{
		ID: *body.ID,
	}
	v.PublishedTables = make([]*subscriptions.SubscriptionPublishedTable, len(body.PublishedTables))
	for i, val := range body.PublishedTables {
		v.PublishedTables[i] = unmarshalSubscriptionPublishedTableResponseBodyToSubscriptionsSubscriptionPublishedTable(val)
	}

	return v
}

// ValidateGetResponseBody runs the validations defined on GetResponseBody
func ValidateGetResponseBody(body *GetResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.PublishedTables == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("published_tables", "body"))
	}
	for _, e := range body.PublishedTables {
		if e != nil {
			if err2 := ValidateSubscriptionPublishedTableResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateAddTableResponseBody runs the validations defined on
// AddTableResponseBody
func ValidateAddTableResponseBody(body *AddTableResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.PublishedTables == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("published_tables", "body"))
	}
	for _, e := range body.PublishedTables {
		if e != nil {
			if err2 := ValidateSubscriptionPublishedTableResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateStopTableResponseBody runs the validations defined on
// StopTableResponseBody
func ValidateStopTableResponseBody(body *StopTableResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.PublishedTables == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("published_tables", "body"))
	}
	for _, e := range body.PublishedTables {
		if e != nil {
			if err2 := ValidateSubscriptionPublishedTableResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateSubscriptionPublishedTableResponseBody runs the validations defined
// on SubscriptionPublishedTableResponseBody
func ValidateSubscriptionPublishedTableResponseBody(body *SubscriptionPublishedTableResponseBody) (err error) {
	if body.Schema == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("schema", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	return
}
