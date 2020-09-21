// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vkumbhar94/lm-sdk-go/models"
)

// GetAllSDTListByWebsiteGroupIDReader is a Reader for the GetAllSDTListByWebsiteGroupID structure.
type GetAllSDTListByWebsiteGroupIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllSDTListByWebsiteGroupIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAllSDTListByWebsiteGroupIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetAllSDTListByWebsiteGroupIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAllSDTListByWebsiteGroupIDOK creates a GetAllSDTListByWebsiteGroupIDOK with default headers values
func NewGetAllSDTListByWebsiteGroupIDOK() *GetAllSDTListByWebsiteGroupIDOK {
	return &GetAllSDTListByWebsiteGroupIDOK{}
}

/*GetAllSDTListByWebsiteGroupIDOK handles this case with default header values.

successful operation
*/
type GetAllSDTListByWebsiteGroupIDOK struct {
	Payload *models.SDTPaginationResponse
}

func (o *GetAllSDTListByWebsiteGroupIDOK) Error() string {
	return fmt.Sprintf("[GET /website/groups/{id}/sdts][%d] getAllSdtListByWebsiteGroupIdOK  %+v", 200, o.Payload)
}

func (o *GetAllSDTListByWebsiteGroupIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SDTPaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllSDTListByWebsiteGroupIDDefault creates a GetAllSDTListByWebsiteGroupIDDefault with default headers values
func NewGetAllSDTListByWebsiteGroupIDDefault(code int) *GetAllSDTListByWebsiteGroupIDDefault {
	return &GetAllSDTListByWebsiteGroupIDDefault{
		_statusCode: code,
	}
}

/*GetAllSDTListByWebsiteGroupIDDefault handles this case with default header values.

Error
*/
type GetAllSDTListByWebsiteGroupIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get all SDT list by website group Id default response
func (o *GetAllSDTListByWebsiteGroupIDDefault) Code() int {
	return o._statusCode
}

func (o *GetAllSDTListByWebsiteGroupIDDefault) Error() string {
	return fmt.Sprintf("[GET /website/groups/{id}/sdts][%d] getAllSDTListByWebsiteGroupId default  %+v", o._statusCode, o.Payload)
}

func (o *GetAllSDTListByWebsiteGroupIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}