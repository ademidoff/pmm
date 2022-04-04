// Code generated by go-swagger; DO NOT EDIT.

package security_checks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListFailedServicesReader is a Reader for the ListFailedServices structure.
type ListFailedServicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListFailedServicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListFailedServicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListFailedServicesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListFailedServicesOK creates a ListFailedServicesOK with default headers values
func NewListFailedServicesOK() *ListFailedServicesOK {
	return &ListFailedServicesOK{}
}

/*ListFailedServicesOK handles this case with default header values.

A successful response.
*/
type ListFailedServicesOK struct {
	Payload *ListFailedServicesOKBody
}

func (o *ListFailedServicesOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/SecurityChecks/ListFailedServices][%d] listFailedServicesOk  %+v", 200, o.Payload)
}

func (o *ListFailedServicesOK) GetPayload() *ListFailedServicesOKBody {
	return o.Payload
}

func (o *ListFailedServicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListFailedServicesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListFailedServicesDefault creates a ListFailedServicesDefault with default headers values
func NewListFailedServicesDefault(code int) *ListFailedServicesDefault {
	return &ListFailedServicesDefault{
		_statusCode: code,
	}
}

/*ListFailedServicesDefault handles this case with default header values.

An unexpected error response.
*/
type ListFailedServicesDefault struct {
	_statusCode int

	Payload *ListFailedServicesDefaultBody
}

// Code gets the status code for the list failed services default response
func (o *ListFailedServicesDefault) Code() int {
	return o._statusCode
}

func (o *ListFailedServicesDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/SecurityChecks/ListFailedServices][%d] ListFailedServices default  %+v", o._statusCode, o.Payload)
}

func (o *ListFailedServicesDefault) GetPayload() *ListFailedServicesDefaultBody {
	return o.Payload
}

func (o *ListFailedServicesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListFailedServicesDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ListFailedServicesDefaultBody list failed services default body
swagger:model ListFailedServicesDefaultBody
*/
type ListFailedServicesDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this list failed services default body
func (o *ListFailedServicesDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListFailedServicesDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("ListFailedServices default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListFailedServicesDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListFailedServicesDefaultBody) UnmarshalBinary(b []byte) error {
	var res ListFailedServicesDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListFailedServicesOKBody list failed services OK body
swagger:model ListFailedServicesOKBody
*/
type ListFailedServicesOKBody struct {

	// result
	Result []*ResultItems0 `json:"result"`
}

// Validate validates this list failed services OK body
func (o *ListFailedServicesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListFailedServicesOKBody) validateResult(formats strfmt.Registry) error {

	if swag.IsZero(o.Result) { // not required
		return nil
	}

	for i := 0; i < len(o.Result); i++ {
		if swag.IsZero(o.Result[i]) { // not required
			continue
		}

		if o.Result[i] != nil {
			if err := o.Result[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listFailedServicesOk" + "." + "result" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListFailedServicesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListFailedServicesOKBody) UnmarshalBinary(b []byte) error {
	var res ListFailedServicesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ResultItems0 CheckResultSummary is a summary of check results.
swagger:model ResultItems0
*/
type ResultItems0 struct {

	// service name
	ServiceName string `json:"service_name,omitempty"`

	// service id
	ServiceID string `json:"service_id,omitempty"`

	// Number of failed checks for this service with severity level "EMERGENCY".
	EmergencyCount int64 `json:"emergency_count,omitempty"`

	// Number of failed checks for this service with severity level "ALERT".
	AlertCount int64 `json:"alert_count,omitempty"`

	// Number of failed checks for this service with severity level "CRITICAL".
	CriticalCount int64 `json:"critical_count,omitempty"`

	// Number of failed checks for this service with severity level "ERROR".
	ErrorCount int64 `json:"error_count,omitempty"`

	// Number of failed checks for this service with severity level "WARNING".
	WarningCount int64 `json:"warning_count,omitempty"`

	// Number of failed checks for this service with severity level "NOTICE".
	NoticeCount int64 `json:"notice_count,omitempty"`

	// Number of failed checks for this service with severity level "INFO".
	InfoCount int64 `json:"info_count,omitempty"`

	// Number of failed checks for this service with severity level "DEBUG".
	DebugCount int64 `json:"debug_count,omitempty"`
}

// Validate validates this result items0
func (o *ResultItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ResultItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ResultItems0) UnmarshalBinary(b []byte) error {
	var res ResultItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}