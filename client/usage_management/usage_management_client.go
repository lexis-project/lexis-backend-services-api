// Code generated by go-swagger; DO NOT EDIT.

package usage_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery -name API -inpkg

// API is the interface of the usage management client
type API interface {
	/*
	   AddConsumption adds a consumption to the system*/
	AddConsumption(ctx context.Context, params *AddConsumptionParams) (*AddConsumptionOK, error)
	/*
	   DecreaseCredit inserts a new reseller in the system*/
	DecreaseCredit(ctx context.Context, params *DecreaseCreditParams) (*DecreaseCreditOK, error)
	/*
	   GetCredit credits status of the account with the provided id*/
	GetCredit(ctx context.Context, params *GetCreditParams) (*GetCreditOK, error)
	/*
	   GetHistory credits history of the customer with id*/
	GetHistory(ctx context.Context, params *GetHistoryParams) (*GetHistoryOK, error)
	/*
	   GetUsage gets account usage for specific project

	   get account usage of project with given id*/
	GetUsage(ctx context.Context, params *GetUsageParams) (*GetUsageOK, error)
	/*
	   IncreaseCredit inserts a new reseller in the system*/
	IncreaseCredit(ctx context.Context, params *IncreaseCreditParams) (*IncreaseCreditOK, error)
}

// New creates a new usage management API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for usage management API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
AddConsumption adds a consumption to the system
*/
func (a *Client) AddConsumption(ctx context.Context, params *AddConsumptionParams) (*AddConsumptionOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addConsumption",
		Method:             "POST",
		PathPattern:        "/accounting/{id}/manage/{medium}/consume/{amount}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddConsumptionReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddConsumptionOK), nil

}

/*
DecreaseCredit inserts a new reseller in the system
*/
func (a *Client) DecreaseCredit(ctx context.Context, params *DecreaseCreditParams) (*DecreaseCreditOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "decreaseCredit",
		Method:             "POST",
		PathPattern:        "/accounting/{id}/manage/{medium}/decrease/{amount}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DecreaseCreditReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DecreaseCreditOK), nil

}

/*
GetCredit credits status of the account with the provided id
*/
func (a *Client) GetCredit(ctx context.Context, params *GetCreditParams) (*GetCreditOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCredit",
		Method:             "GET",
		PathPattern:        "/accounting/{id}/available",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetCreditReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCreditOK), nil

}

/*
GetHistory credits history of the customer with id
*/
func (a *Client) GetHistory(ctx context.Context, params *GetHistoryParams) (*GetHistoryOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getHistory",
		Method:             "GET",
		PathPattern:        "/accounting/{id}/history",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetHistoryReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetHistoryOK), nil

}

/*
GetUsage gets account usage for specific project

get account usage of project with given id
*/
func (a *Client) GetUsage(ctx context.Context, params *GetUsageParams) (*GetUsageOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUsage",
		Method:             "GET",
		PathPattern:        "/accounting/{id}/usage",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetUsageReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUsageOK), nil

}

/*
IncreaseCredit inserts a new reseller in the system
*/
func (a *Client) IncreaseCredit(ctx context.Context, params *IncreaseCreditParams) (*IncreaseCreditOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "increaseCredit",
		Method:             "POST",
		PathPattern:        "/accounting/{id}/manage/{medium}/increase/{amount}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IncreaseCreditReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*IncreaseCreditOK), nil

}