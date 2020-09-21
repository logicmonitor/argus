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

// GetImmediateWebsiteListByWebsiteGroupIDReader is a Reader for the GetImmediateWebsiteListByWebsiteGroupID structure.
type GetImmediateWebsiteListByWebsiteGroupIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetImmediateWebsiteListByWebsiteGroupIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetImmediateWebsiteListByWebsiteGroupIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetImmediateWebsiteListByWebsiteGroupIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetImmediateWebsiteListByWebsiteGroupIDOK creates a GetImmediateWebsiteListByWebsiteGroupIDOK with default headers values
func NewGetImmediateWebsiteListByWebsiteGroupIDOK() *GetImmediateWebsiteListByWebsiteGroupIDOK {
	return &GetImmediateWebsiteListByWebsiteGroupIDOK{}
}

/*GetImmediateWebsiteListByWebsiteGroupIDOK handles this case with default header values.

successful operation
*/
type GetImmediateWebsiteListByWebsiteGroupIDOK struct {
	Payload *models.WebsitePaginationResponse
}

func (o *GetImmediateWebsiteListByWebsiteGroupIDOK) Error() string {
	return fmt.Sprintf("[GET /website/groups/{id}/websites][%d] getImmediateWebsiteListByWebsiteGroupIdOK  %+v", 200, o.Payload)
}

func (o *GetImmediateWebsiteListByWebsiteGroupIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebsitePaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetImmediateWebsiteListByWebsiteGroupIDDefault creates a GetImmediateWebsiteListByWebsiteGroupIDDefault with default headers values
func NewGetImmediateWebsiteListByWebsiteGroupIDDefault(code int) *GetImmediateWebsiteListByWebsiteGroupIDDefault {
	return &GetImmediateWebsiteListByWebsiteGroupIDDefault{
		_statusCode: code,
	}
}

/*GetImmediateWebsiteListByWebsiteGroupIDDefault handles this case with default header values.

Error
*/
type GetImmediateWebsiteListByWebsiteGroupIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get immediate website list by website group Id default response
func (o *GetImmediateWebsiteListByWebsiteGroupIDDefault) Code() int {
	return o._statusCode
}

func (o *GetImmediateWebsiteListByWebsiteGroupIDDefault) Error() string {
	return fmt.Sprintf("[GET /website/groups/{id}/websites][%d] getImmediateWebsiteListByWebsiteGroupId default  %+v", o._statusCode, o.Payload)
}

func (o *GetImmediateWebsiteListByWebsiteGroupIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}