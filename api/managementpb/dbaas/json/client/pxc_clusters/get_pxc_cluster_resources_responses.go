// Code generated by go-swagger; DO NOT EDIT.

package pxc_clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetPXCClusterResourcesReader is a Reader for the GetPXCClusterResources structure.
type GetPXCClusterResourcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPXCClusterResourcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPXCClusterResourcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetPXCClusterResourcesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPXCClusterResourcesOK creates a GetPXCClusterResourcesOK with default headers values
func NewGetPXCClusterResourcesOK() *GetPXCClusterResourcesOK {
	return &GetPXCClusterResourcesOK{}
}

/*
GetPXCClusterResourcesOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetPXCClusterResourcesOK struct {
	Payload *GetPXCClusterResourcesOKBody
}

func (o *GetPXCClusterResourcesOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/PXCCluster/Resources/Get][%d] getPxcClusterResourcesOk  %+v", 200, o.Payload)
}

func (o *GetPXCClusterResourcesOK) GetPayload() *GetPXCClusterResourcesOKBody {
	return o.Payload
}

func (o *GetPXCClusterResourcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(GetPXCClusterResourcesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPXCClusterResourcesDefault creates a GetPXCClusterResourcesDefault with default headers values
func NewGetPXCClusterResourcesDefault(code int) *GetPXCClusterResourcesDefault {
	return &GetPXCClusterResourcesDefault{
		_statusCode: code,
	}
}

/*
GetPXCClusterResourcesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetPXCClusterResourcesDefault struct {
	_statusCode int

	Payload *GetPXCClusterResourcesDefaultBody
}

// Code gets the status code for the get PXC cluster resources default response
func (o *GetPXCClusterResourcesDefault) Code() int {
	return o._statusCode
}

func (o *GetPXCClusterResourcesDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/PXCCluster/Resources/Get][%d] GetPXCClusterResources default  %+v", o._statusCode, o.Payload)
}

func (o *GetPXCClusterResourcesDefault) GetPayload() *GetPXCClusterResourcesDefaultBody {
	return o.Payload
}

func (o *GetPXCClusterResourcesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(GetPXCClusterResourcesDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetPXCClusterResourcesBody get PXC cluster resources body
swagger:model GetPXCClusterResourcesBody
*/
type GetPXCClusterResourcesBody struct {
	// params
	Params *GetPXCClusterResourcesParamsBodyParams `json:"params,omitempty"`
}

// Validate validates this get PXC cluster resources body
func (o *GetPXCClusterResourcesBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPXCClusterResourcesBody) validateParams(formats strfmt.Registry) error {
	if swag.IsZero(o.Params) { // not required
		return nil
	}

	if o.Params != nil {
		if err := o.Params.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "params")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get PXC cluster resources body based on the context it is used
func (o *GetPXCClusterResourcesBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPXCClusterResourcesBody) contextValidateParams(ctx context.Context, formats strfmt.Registry) error {
	if o.Params != nil {
		if err := o.Params.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "params")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPXCClusterResourcesBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPXCClusterResourcesBody) UnmarshalBinary(b []byte) error {
	var res GetPXCClusterResourcesBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetPXCClusterResourcesDefaultBody get PXC cluster resources default body
swagger:model GetPXCClusterResourcesDefaultBody
*/
type GetPXCClusterResourcesDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*GetPXCClusterResourcesDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this get PXC cluster resources default body
func (o *GetPXCClusterResourcesDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPXCClusterResourcesDefaultBody) validateDetails(formats strfmt.Registry) error {
	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("GetPXCClusterResources default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("GetPXCClusterResources default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get PXC cluster resources default body based on the context it is used
func (o *GetPXCClusterResourcesDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPXCClusterResourcesDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("GetPXCClusterResources default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("GetPXCClusterResources default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPXCClusterResourcesDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPXCClusterResourcesDefaultBody) UnmarshalBinary(b []byte) error {
	var res GetPXCClusterResourcesDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetPXCClusterResourcesDefaultBodyDetailsItems0 get PXC cluster resources default body details items0
swagger:model GetPXCClusterResourcesDefaultBodyDetailsItems0
*/
type GetPXCClusterResourcesDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this get PXC cluster resources default body details items0
func (o *GetPXCClusterResourcesDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get PXC cluster resources default body details items0 based on context it is used
func (o *GetPXCClusterResourcesDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetPXCClusterResourcesDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPXCClusterResourcesDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res GetPXCClusterResourcesDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetPXCClusterResourcesOKBody get PXC cluster resources OK body
swagger:model GetPXCClusterResourcesOKBody
*/
type GetPXCClusterResourcesOKBody struct {
	// expected
	Expected *GetPXCClusterResourcesOKBodyExpected `json:"expected,omitempty"`
}

// Validate validates this get PXC cluster resources OK body
func (o *GetPXCClusterResourcesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateExpected(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPXCClusterResourcesOKBody) validateExpected(formats strfmt.Registry) error {
	if swag.IsZero(o.Expected) { // not required
		return nil
	}

	if o.Expected != nil {
		if err := o.Expected.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getPxcClusterResourcesOk" + "." + "expected")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getPxcClusterResourcesOk" + "." + "expected")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get PXC cluster resources OK body based on the context it is used
func (o *GetPXCClusterResourcesOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateExpected(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPXCClusterResourcesOKBody) contextValidateExpected(ctx context.Context, formats strfmt.Registry) error {
	if o.Expected != nil {
		if err := o.Expected.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getPxcClusterResourcesOk" + "." + "expected")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getPxcClusterResourcesOk" + "." + "expected")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPXCClusterResourcesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPXCClusterResourcesOKBody) UnmarshalBinary(b []byte) error {
	var res GetPXCClusterResourcesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetPXCClusterResourcesOKBodyExpected Resources contains Kubernetes cluster resources.
swagger:model GetPXCClusterResourcesOKBodyExpected
*/
type GetPXCClusterResourcesOKBodyExpected struct {
	// Memory in bytes.
	MemoryBytes string `json:"memory_bytes,omitempty"`

	// CPU in millicpus. For example 0.1 of CPU is equivalent to 100 millicpus.
	// See https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#meaning-of-cpu.
	CPUm string `json:"cpu_m,omitempty"`

	// Disk size in bytes.
	DiskSize string `json:"disk_size,omitempty"`
}

// Validate validates this get PXC cluster resources OK body expected
func (o *GetPXCClusterResourcesOKBodyExpected) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get PXC cluster resources OK body expected based on context it is used
func (o *GetPXCClusterResourcesOKBodyExpected) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetPXCClusterResourcesOKBodyExpected) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPXCClusterResourcesOKBodyExpected) UnmarshalBinary(b []byte) error {
	var res GetPXCClusterResourcesOKBodyExpected
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetPXCClusterResourcesParamsBodyParams PXCClusterParams represents PXC cluster parameters that can be updated.
swagger:model GetPXCClusterResourcesParamsBodyParams
*/
type GetPXCClusterResourcesParamsBodyParams struct {
	// Cluster size.
	ClusterSize int32 `json:"cluster_size,omitempty"`

	// backup
	Backup *GetPXCClusterResourcesParamsBodyParamsBackup `json:"backup,omitempty"`

	// haproxy
	Haproxy *GetPXCClusterResourcesParamsBodyParamsHaproxy `json:"haproxy,omitempty"`

	// proxysql
	Proxysql *GetPXCClusterResourcesParamsBodyParamsProxysql `json:"proxysql,omitempty"`

	// pxc
	PXC *GetPXCClusterResourcesParamsBodyParamsPXC `json:"pxc,omitempty"`

	// restore
	Restore *GetPXCClusterResourcesParamsBodyParamsRestore `json:"restore,omitempty"`
}

// Validate validates this get PXC cluster resources params body params
func (o *GetPXCClusterResourcesParamsBodyParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBackup(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateHaproxy(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateProxysql(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePXC(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRestore(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPXCClusterResourcesParamsBodyParams) validateBackup(formats strfmt.Registry) error {
	if swag.IsZero(o.Backup) { // not required
		return nil
	}

	if o.Backup != nil {
		if err := o.Backup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "backup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "params" + "." + "backup")
			}
			return err
		}
	}

	return nil
}

func (o *GetPXCClusterResourcesParamsBodyParams) validateHaproxy(formats strfmt.Registry) error {
	if swag.IsZero(o.Haproxy) { // not required
		return nil
	}

	if o.Haproxy != nil {
		if err := o.Haproxy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "haproxy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "params" + "." + "haproxy")
			}
			return err
		}
	}

	return nil
}

func (o *GetPXCClusterResourcesParamsBodyParams) validateProxysql(formats strfmt.Registry) error {
	if swag.IsZero(o.Proxysql) { // not required
		return nil
	}

	if o.Proxysql != nil {
		if err := o.Proxysql.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "proxysql")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "params" + "." + "proxysql")
			}
			return err
		}
	}

	return nil
}

func (o *GetPXCClusterResourcesParamsBodyParams) validatePXC(formats strfmt.Registry) error {
	if swag.IsZero(o.PXC) { // not required
		return nil
	}

	if o.PXC != nil {
		if err := o.PXC.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "pxc")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "params" + "." + "pxc")
			}
			return err
		}
	}

	return nil
}

func (o *GetPXCClusterResourcesParamsBodyParams) validateRestore(formats strfmt.Registry) error {
	if swag.IsZero(o.Restore) { // not required
		return nil
	}

	if o.Restore != nil {
		if err := o.Restore.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "restore")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "params" + "." + "restore")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get PXC cluster resources params body params based on the context it is used
func (o *GetPXCClusterResourcesParamsBodyParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateBackup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateHaproxy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateProxysql(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidatePXC(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateRestore(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPXCClusterResourcesParamsBodyParams) contextValidateBackup(ctx context.Context, formats strfmt.Registry) error {
	if o.Backup != nil {
		if err := o.Backup.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "backup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "params" + "." + "backup")
			}
			return err
		}
	}

	return nil
}

func (o *GetPXCClusterResourcesParamsBodyParams) contextValidateHaproxy(ctx context.Context, formats strfmt.Registry) error {
	if o.Haproxy != nil {
		if err := o.Haproxy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "haproxy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "params" + "." + "haproxy")
			}
			return err
		}
	}

	return nil
}

func (o *GetPXCClusterResourcesParamsBodyParams) contextValidateProxysql(ctx context.Context, formats strfmt.Registry) error {
	if o.Proxysql != nil {
		if err := o.Proxysql.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "proxysql")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "params" + "." + "proxysql")
			}
			return err
		}
	}

	return nil
}

func (o *GetPXCClusterResourcesParamsBodyParams) contextValidatePXC(ctx context.Context, formats strfmt.Registry) error {
	if o.PXC != nil {
		if err := o.PXC.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "pxc")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "params" + "." + "pxc")
			}
			return err
		}
	}

	return nil
}

func (o *GetPXCClusterResourcesParamsBodyParams) contextValidateRestore(ctx context.Context, formats strfmt.Registry) error {
	if o.Restore != nil {
		if err := o.Restore.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "restore")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "params" + "." + "restore")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPXCClusterResourcesParamsBodyParams) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPXCClusterResourcesParamsBodyParams) UnmarshalBinary(b []byte) error {
	var res GetPXCClusterResourcesParamsBodyParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetPXCClusterResourcesParamsBodyParamsBackup Backup configuration for a database cluster
swagger:model GetPXCClusterResourcesParamsBodyParamsBackup
*/
type GetPXCClusterResourcesParamsBodyParamsBackup struct {
	// Backup Location id of stored backup location in PMM.
	LocationID string `json:"location_id,omitempty"`

	// Keep copies represents how many copies should retain.
	KeepCopies int32 `json:"keep_copies,omitempty"`

	// Cron expression represents cron expression
	CronExpression string `json:"cron_expression,omitempty"`

	// Service acccount used for backups
	ServiceAccount string `json:"service_account,omitempty"`
}

// Validate validates this get PXC cluster resources params body params backup
func (o *GetPXCClusterResourcesParamsBodyParamsBackup) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get PXC cluster resources params body params backup based on context it is used
func (o *GetPXCClusterResourcesParamsBodyParamsBackup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetPXCClusterResourcesParamsBodyParamsBackup) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPXCClusterResourcesParamsBodyParamsBackup) UnmarshalBinary(b []byte) error {
	var res GetPXCClusterResourcesParamsBodyParamsBackup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetPXCClusterResourcesParamsBodyParamsHaproxy HAProxy container parameters.
// NOTE: HAProxy does not need disk size as ProxySQL does because the container does not require it.
swagger:model GetPXCClusterResourcesParamsBodyParamsHaproxy
*/
type GetPXCClusterResourcesParamsBodyParamsHaproxy struct {
	// Docker image used for HAProxy.
	Image string `json:"image,omitempty"`

	// compute resources
	ComputeResources *GetPXCClusterResourcesParamsBodyParamsHaproxyComputeResources `json:"compute_resources,omitempty"`
}

// Validate validates this get PXC cluster resources params body params haproxy
func (o *GetPXCClusterResourcesParamsBodyParamsHaproxy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateComputeResources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPXCClusterResourcesParamsBodyParamsHaproxy) validateComputeResources(formats strfmt.Registry) error {
	if swag.IsZero(o.ComputeResources) { // not required
		return nil
	}

	if o.ComputeResources != nil {
		if err := o.ComputeResources.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "haproxy" + "." + "compute_resources")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "params" + "." + "haproxy" + "." + "compute_resources")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get PXC cluster resources params body params haproxy based on the context it is used
func (o *GetPXCClusterResourcesParamsBodyParamsHaproxy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateComputeResources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPXCClusterResourcesParamsBodyParamsHaproxy) contextValidateComputeResources(ctx context.Context, formats strfmt.Registry) error {
	if o.ComputeResources != nil {
		if err := o.ComputeResources.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "haproxy" + "." + "compute_resources")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "params" + "." + "haproxy" + "." + "compute_resources")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPXCClusterResourcesParamsBodyParamsHaproxy) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPXCClusterResourcesParamsBodyParamsHaproxy) UnmarshalBinary(b []byte) error {
	var res GetPXCClusterResourcesParamsBodyParamsHaproxy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetPXCClusterResourcesParamsBodyParamsHaproxyComputeResources ComputeResources represents container computer resources requests or limits.
swagger:model GetPXCClusterResourcesParamsBodyParamsHaproxyComputeResources
*/
type GetPXCClusterResourcesParamsBodyParamsHaproxyComputeResources struct {
	// CPUs in milliCPUs; 1000m = 1 vCPU.
	CPUm int32 `json:"cpu_m,omitempty"`

	// Memory in bytes.
	MemoryBytes string `json:"memory_bytes,omitempty"`
}

// Validate validates this get PXC cluster resources params body params haproxy compute resources
func (o *GetPXCClusterResourcesParamsBodyParamsHaproxyComputeResources) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get PXC cluster resources params body params haproxy compute resources based on context it is used
func (o *GetPXCClusterResourcesParamsBodyParamsHaproxyComputeResources) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetPXCClusterResourcesParamsBodyParamsHaproxyComputeResources) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPXCClusterResourcesParamsBodyParamsHaproxyComputeResources) UnmarshalBinary(b []byte) error {
	var res GetPXCClusterResourcesParamsBodyParamsHaproxyComputeResources
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetPXCClusterResourcesParamsBodyParamsPXC PXC container parameters.
// TODO Do not use inner messages in all public APIs (for consistency).
swagger:model GetPXCClusterResourcesParamsBodyParamsPXC
*/
type GetPXCClusterResourcesParamsBodyParamsPXC struct {
	// Docker image used for PXC.
	Image string `json:"image,omitempty"`

	// Disk size in bytes.
	DiskSize string `json:"disk_size,omitempty"`

	// Configuration for PXC cluster
	Configuration string `json:"configuration,omitempty"`

	// Storage Class for PXC cluster.
	StorageClass string `json:"storage_class,omitempty"`

	// compute resources
	ComputeResources *GetPXCClusterResourcesParamsBodyParamsPXCComputeResources `json:"compute_resources,omitempty"`
}

// Validate validates this get PXC cluster resources params body params PXC
func (o *GetPXCClusterResourcesParamsBodyParamsPXC) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateComputeResources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPXCClusterResourcesParamsBodyParamsPXC) validateComputeResources(formats strfmt.Registry) error {
	if swag.IsZero(o.ComputeResources) { // not required
		return nil
	}

	if o.ComputeResources != nil {
		if err := o.ComputeResources.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "pxc" + "." + "compute_resources")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "params" + "." + "pxc" + "." + "compute_resources")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get PXC cluster resources params body params PXC based on the context it is used
func (o *GetPXCClusterResourcesParamsBodyParamsPXC) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateComputeResources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPXCClusterResourcesParamsBodyParamsPXC) contextValidateComputeResources(ctx context.Context, formats strfmt.Registry) error {
	if o.ComputeResources != nil {
		if err := o.ComputeResources.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "pxc" + "." + "compute_resources")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "params" + "." + "pxc" + "." + "compute_resources")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPXCClusterResourcesParamsBodyParamsPXC) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPXCClusterResourcesParamsBodyParamsPXC) UnmarshalBinary(b []byte) error {
	var res GetPXCClusterResourcesParamsBodyParamsPXC
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetPXCClusterResourcesParamsBodyParamsPXCComputeResources ComputeResources represents container computer resources requests or limits.
swagger:model GetPXCClusterResourcesParamsBodyParamsPXCComputeResources
*/
type GetPXCClusterResourcesParamsBodyParamsPXCComputeResources struct {
	// CPUs in milliCPUs; 1000m = 1 vCPU.
	CPUm int32 `json:"cpu_m,omitempty"`

	// Memory in bytes.
	MemoryBytes string `json:"memory_bytes,omitempty"`
}

// Validate validates this get PXC cluster resources params body params PXC compute resources
func (o *GetPXCClusterResourcesParamsBodyParamsPXCComputeResources) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get PXC cluster resources params body params PXC compute resources based on context it is used
func (o *GetPXCClusterResourcesParamsBodyParamsPXCComputeResources) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetPXCClusterResourcesParamsBodyParamsPXCComputeResources) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPXCClusterResourcesParamsBodyParamsPXCComputeResources) UnmarshalBinary(b []byte) error {
	var res GetPXCClusterResourcesParamsBodyParamsPXCComputeResources
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetPXCClusterResourcesParamsBodyParamsProxysql ProxySQL container parameters.
// TODO Do not use inner messages in all public APIs (for consistency).
swagger:model GetPXCClusterResourcesParamsBodyParamsProxysql
*/
type GetPXCClusterResourcesParamsBodyParamsProxysql struct {
	// Docker image used for ProxySQL.
	Image string `json:"image,omitempty"`

	// Disk size in bytes.
	DiskSize string `json:"disk_size,omitempty"`

	// compute resources
	ComputeResources *GetPXCClusterResourcesParamsBodyParamsProxysqlComputeResources `json:"compute_resources,omitempty"`
}

// Validate validates this get PXC cluster resources params body params proxysql
func (o *GetPXCClusterResourcesParamsBodyParamsProxysql) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateComputeResources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPXCClusterResourcesParamsBodyParamsProxysql) validateComputeResources(formats strfmt.Registry) error {
	if swag.IsZero(o.ComputeResources) { // not required
		return nil
	}

	if o.ComputeResources != nil {
		if err := o.ComputeResources.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "proxysql" + "." + "compute_resources")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "params" + "." + "proxysql" + "." + "compute_resources")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get PXC cluster resources params body params proxysql based on the context it is used
func (o *GetPXCClusterResourcesParamsBodyParamsProxysql) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateComputeResources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPXCClusterResourcesParamsBodyParamsProxysql) contextValidateComputeResources(ctx context.Context, formats strfmt.Registry) error {
	if o.ComputeResources != nil {
		if err := o.ComputeResources.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "proxysql" + "." + "compute_resources")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "params" + "." + "proxysql" + "." + "compute_resources")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPXCClusterResourcesParamsBodyParamsProxysql) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPXCClusterResourcesParamsBodyParamsProxysql) UnmarshalBinary(b []byte) error {
	var res GetPXCClusterResourcesParamsBodyParamsProxysql
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetPXCClusterResourcesParamsBodyParamsProxysqlComputeResources ComputeResources represents container computer resources requests or limits.
swagger:model GetPXCClusterResourcesParamsBodyParamsProxysqlComputeResources
*/
type GetPXCClusterResourcesParamsBodyParamsProxysqlComputeResources struct {
	// CPUs in milliCPUs; 1000m = 1 vCPU.
	CPUm int32 `json:"cpu_m,omitempty"`

	// Memory in bytes.
	MemoryBytes string `json:"memory_bytes,omitempty"`
}

// Validate validates this get PXC cluster resources params body params proxysql compute resources
func (o *GetPXCClusterResourcesParamsBodyParamsProxysqlComputeResources) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get PXC cluster resources params body params proxysql compute resources based on context it is used
func (o *GetPXCClusterResourcesParamsBodyParamsProxysqlComputeResources) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetPXCClusterResourcesParamsBodyParamsProxysqlComputeResources) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPXCClusterResourcesParamsBodyParamsProxysqlComputeResources) UnmarshalBinary(b []byte) error {
	var res GetPXCClusterResourcesParamsBodyParamsProxysqlComputeResources
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetPXCClusterResourcesParamsBodyParamsRestore Restore represents restoration payload to restore a database cluster from backup
swagger:model GetPXCClusterResourcesParamsBodyParamsRestore
*/
type GetPXCClusterResourcesParamsBodyParamsRestore struct {
	// Backup location in PMM.
	LocationID string `json:"location_id,omitempty"`

	// Destination filename.
	Destination string `json:"destination,omitempty"`

	// K8s Secrets name.
	SecretsName string `json:"secrets_name,omitempty"`
}

// Validate validates this get PXC cluster resources params body params restore
func (o *GetPXCClusterResourcesParamsBodyParamsRestore) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get PXC cluster resources params body params restore based on context it is used
func (o *GetPXCClusterResourcesParamsBodyParamsRestore) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetPXCClusterResourcesParamsBodyParamsRestore) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPXCClusterResourcesParamsBodyParamsRestore) UnmarshalBinary(b []byte) error {
	var res GetPXCClusterResourcesParamsBodyParamsRestore
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
