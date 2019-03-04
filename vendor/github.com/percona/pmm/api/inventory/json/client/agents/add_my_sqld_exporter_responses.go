// Code generated by go-swagger; DO NOT EDIT.

package agents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// AddMySqldExporterReader is a Reader for the AddMySqldExporter structure.
type AddMySqldExporterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddMySqldExporterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddMySqldExporterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddMySqldExporterOK creates a AddMySqldExporterOK with default headers values
func NewAddMySqldExporterOK() *AddMySqldExporterOK {
	return &AddMySqldExporterOK{}
}

/*AddMySqldExporterOK handles this case with default header values.

A successful response.
*/
type AddMySqldExporterOK struct {
	Payload *AddMySqldExporterOKBody
}

func (o *AddMySqldExporterOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/AddMySQLdExporter][%d] addMySqldExporterOK  %+v", 200, o.Payload)
}

func (o *AddMySqldExporterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddMySqldExporterOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AddMySqldExporterBody add my sqld exporter body
swagger:model AddMySqldExporterBody
*/
type AddMySqldExporterBody struct {

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// MySQL password for scraping metrics.
	Password string `json:"password,omitempty"`

	// Node identifier where this instance runs.
	RunsOnNodeID string `json:"runs_on_node_id,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// MySQL username for scraping metrics.
	Username string `json:"username,omitempty"`
}

// Validate validates this add my sqld exporter body
func (o *AddMySqldExporterBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddMySqldExporterBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddMySqldExporterBody) UnmarshalBinary(b []byte) error {
	var res AddMySqldExporterBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddMySqldExporterOKBody add my sqld exporter o k body
swagger:model AddMySqldExporterOKBody
*/
type AddMySqldExporterOKBody struct {

	// mysqld exporter
	MysqldExporter *AddMySqldExporterOKBodyMysqldExporter `json:"mysqld_exporter,omitempty"`
}

// Validate validates this add my sqld exporter o k body
func (o *AddMySqldExporterOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMysqldExporter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddMySqldExporterOKBody) validateMysqldExporter(formats strfmt.Registry) error {

	if swag.IsZero(o.MysqldExporter) { // not required
		return nil
	}

	if o.MysqldExporter != nil {
		if err := o.MysqldExporter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addMySqldExporterOK" + "." + "mysqld_exporter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddMySqldExporterOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddMySqldExporterOKBody) UnmarshalBinary(b []byte) error {
	var res AddMySqldExporterOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddMySqldExporterOKBodyMysqldExporter MySQLdExporter runs on Generic or Container Node and exposes MySQL and AmazonRDSMySQL Service metrics.
swagger:model AddMySqldExporterOKBodyMysqldExporter
*/
type AddMySqldExporterOKBodyMysqldExporter struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Listen port for scraping metrics.
	ListenPort int64 `json:"listen_port,omitempty"`

	// MySQL password for scraping metrics.
	Password string `json:"password,omitempty"`

	// Node identifier where this instance runs.
	RunsOnNodeID string `json:"runs_on_node_id,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// AgentStatus represents actual Agent process status.
	// Enum: [AGENT_STATUS_INVALID STARTING RUNNING WAITING STOPPING DONE]
	Status *string `json:"status,omitempty"`

	// MySQL username for scraping metrics.
	Username string `json:"username,omitempty"`
}

// Validate validates this add my sqld exporter o k body mysqld exporter
func (o *AddMySqldExporterOKBodyMysqldExporter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addMySqldExporterOKBodyMysqldExporterTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addMySqldExporterOKBodyMysqldExporterTypeStatusPropEnum = append(addMySqldExporterOKBodyMysqldExporterTypeStatusPropEnum, v)
	}
}

const (

	// AddMySqldExporterOKBodyMysqldExporterStatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	AddMySqldExporterOKBodyMysqldExporterStatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// AddMySqldExporterOKBodyMysqldExporterStatusSTARTING captures enum value "STARTING"
	AddMySqldExporterOKBodyMysqldExporterStatusSTARTING string = "STARTING"

	// AddMySqldExporterOKBodyMysqldExporterStatusRUNNING captures enum value "RUNNING"
	AddMySqldExporterOKBodyMysqldExporterStatusRUNNING string = "RUNNING"

	// AddMySqldExporterOKBodyMysqldExporterStatusWAITING captures enum value "WAITING"
	AddMySqldExporterOKBodyMysqldExporterStatusWAITING string = "WAITING"

	// AddMySqldExporterOKBodyMysqldExporterStatusSTOPPING captures enum value "STOPPING"
	AddMySqldExporterOKBodyMysqldExporterStatusSTOPPING string = "STOPPING"

	// AddMySqldExporterOKBodyMysqldExporterStatusDONE captures enum value "DONE"
	AddMySqldExporterOKBodyMysqldExporterStatusDONE string = "DONE"
)

// prop value enum
func (o *AddMySqldExporterOKBodyMysqldExporter) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, addMySqldExporterOKBodyMysqldExporterTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *AddMySqldExporterOKBodyMysqldExporter) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("addMySqldExporterOK"+"."+"mysqld_exporter"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddMySqldExporterOKBodyMysqldExporter) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddMySqldExporterOKBodyMysqldExporter) UnmarshalBinary(b []byte) error {
	var res AddMySqldExporterOKBodyMysqldExporter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
