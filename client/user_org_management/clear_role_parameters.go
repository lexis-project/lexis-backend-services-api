// Code generated by go-swagger; DO NOT EDIT.

package user_org_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewClearRoleParams creates a new ClearRoleParams object
// with the default values initialized.
func NewClearRoleParams() *ClearRoleParams {
	var ()
	return &ClearRoleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewClearRoleParamsWithTimeout creates a new ClearRoleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewClearRoleParamsWithTimeout(timeout time.Duration) *ClearRoleParams {
	var ()
	return &ClearRoleParams{

		timeout: timeout,
	}
}

// NewClearRoleParamsWithContext creates a new ClearRoleParams object
// with the default values initialized, and the ability to set a context for a request
func NewClearRoleParamsWithContext(ctx context.Context) *ClearRoleParams {
	var ()
	return &ClearRoleParams{

		Context: ctx,
	}
}

// NewClearRoleParamsWithHTTPClient creates a new ClearRoleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewClearRoleParamsWithHTTPClient(client *http.Client) *ClearRoleParams {
	var ()
	return &ClearRoleParams{
		HTTPClient: client,
	}
}

/*ClearRoleParams contains all the parameters to send to the API endpoint
for the clear role operation typically these are written to a http.Request
*/
type ClearRoleParams struct {

	/*OrganizationID
	  Id of the organization linked

	*/
	OrganizationID strfmt.UUID
	/*ProjectID
	  Id of the project linked

	*/
	ProjectID *strfmt.UUID
	/*ProjectShortName
	  (REQUIRED when ProjectID is provided!) Short name of the project linked

	*/
	ProjectShortName *string
	/*UserID
	  Id of the user to be modified

	*/
	UserID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the clear role params
func (o *ClearRoleParams) WithTimeout(timeout time.Duration) *ClearRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the clear role params
func (o *ClearRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the clear role params
func (o *ClearRoleParams) WithContext(ctx context.Context) *ClearRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the clear role params
func (o *ClearRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the clear role params
func (o *ClearRoleParams) WithHTTPClient(client *http.Client) *ClearRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the clear role params
func (o *ClearRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the clear role params
func (o *ClearRoleParams) WithOrganizationID(organizationID strfmt.UUID) *ClearRoleParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the clear role params
func (o *ClearRoleParams) SetOrganizationID(organizationID strfmt.UUID) {
	o.OrganizationID = organizationID
}

// WithProjectID adds the projectID to the clear role params
func (o *ClearRoleParams) WithProjectID(projectID *strfmt.UUID) *ClearRoleParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the clear role params
func (o *ClearRoleParams) SetProjectID(projectID *strfmt.UUID) {
	o.ProjectID = projectID
}

// WithProjectShortName adds the projectShortName to the clear role params
func (o *ClearRoleParams) WithProjectShortName(projectShortName *string) *ClearRoleParams {
	o.SetProjectShortName(projectShortName)
	return o
}

// SetProjectShortName adds the projectShortName to the clear role params
func (o *ClearRoleParams) SetProjectShortName(projectShortName *string) {
	o.ProjectShortName = projectShortName
}

// WithUserID adds the userID to the clear role params
func (o *ClearRoleParams) WithUserID(userID strfmt.UUID) *ClearRoleParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the clear role params
func (o *ClearRoleParams) SetUserID(userID strfmt.UUID) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *ClearRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param organizationID
	qrOrganizationID := o.OrganizationID
	qOrganizationID := qrOrganizationID.String()
	if qOrganizationID != "" {
		if err := r.SetQueryParam("organizationID", qOrganizationID); err != nil {
			return err
		}
	}

	if o.ProjectID != nil {

		// query param projectID
		var qrProjectID strfmt.UUID
		if o.ProjectID != nil {
			qrProjectID = *o.ProjectID
		}
		qProjectID := qrProjectID.String()
		if qProjectID != "" {
			if err := r.SetQueryParam("projectID", qProjectID); err != nil {
				return err
			}
		}

	}

	if o.ProjectShortName != nil {

		// query param projectShortName
		var qrProjectShortName string
		if o.ProjectShortName != nil {
			qrProjectShortName = *o.ProjectShortName
		}
		qProjectShortName := qrProjectShortName
		if qProjectShortName != "" {
			if err := r.SetQueryParam("projectShortName", qProjectShortName); err != nil {
				return err
			}
		}

	}

	// path param userID
	if err := r.SetPathParam("userID", o.UserID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
