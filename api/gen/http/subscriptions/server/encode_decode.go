// Code generated by goa v3.5.4, DO NOT EDIT.
//
// Subscriptions HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/lawrencejones/pgsink/api/design -o api

package server

import (
	"context"
	"io"
	"net/http"

	subscriptions "github.com/lawrencejones/pgsink/api/gen/subscriptions"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeGetResponse returns an encoder for responses returned by the
// Subscriptions Get endpoint.
func EncodeGetResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*subscriptions.Subscription)
		enc := encoder(ctx, w)
		body := NewGetResponseBody(res)
		w.WriteHeader(http.StatusCreated)
		return enc.Encode(body)
	}
}

// EncodeAddTableResponse returns an encoder for responses returned by the
// Subscriptions AddTable endpoint.
func EncodeAddTableResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*subscriptions.Subscription)
		enc := encoder(ctx, w)
		body := NewAddTableResponseBody(res)
		w.WriteHeader(http.StatusCreated)
		return enc.Encode(body)
	}
}

// DecodeAddTableRequest returns a decoder for requests sent to the
// Subscriptions AddTable endpoint.
func DecodeAddTableRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body AddTableRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateAddTableRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewAddTableSubscriptionPublishedTable(&body)

		return payload, nil
	}
}

// EncodeStopTableResponse returns an encoder for responses returned by the
// Subscriptions StopTable endpoint.
func EncodeStopTableResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*subscriptions.Subscription)
		enc := encoder(ctx, w)
		body := NewStopTableResponseBody(res)
		w.WriteHeader(http.StatusAccepted)
		return enc.Encode(body)
	}
}

// DecodeStopTableRequest returns a decoder for requests sent to the
// Subscriptions StopTable endpoint.
func DecodeStopTableRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body StopTableRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateStopTableRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewStopTableSubscriptionPublishedTable(&body)

		return payload, nil
	}
}

// marshalSubscriptionsSubscriptionPublishedTableToSubscriptionPublishedTableResponseBody
// builds a value of type *SubscriptionPublishedTableResponseBody from a value
// of type *subscriptions.SubscriptionPublishedTable.
func marshalSubscriptionsSubscriptionPublishedTableToSubscriptionPublishedTableResponseBody(v *subscriptions.SubscriptionPublishedTable) *SubscriptionPublishedTableResponseBody {
	res := &SubscriptionPublishedTableResponseBody{
		Schema: v.Schema,
		Name:   v.Name,
	}

	return res
}
