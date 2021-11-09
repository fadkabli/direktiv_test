// Code generated by go-swagger; DO NOT EDIT.

package workflow_services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ListWorkflowServiceRevisionPodsReader is a Reader for the ListWorkflowServiceRevisionPods structure.
type ListWorkflowServiceRevisionPodsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListWorkflowServiceRevisionPodsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListWorkflowServiceRevisionPodsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListWorkflowServiceRevisionPodsOK creates a ListWorkflowServiceRevisionPodsOK with default headers values
func NewListWorkflowServiceRevisionPodsOK() *ListWorkflowServiceRevisionPodsOK {
	return &ListWorkflowServiceRevisionPodsOK{}
}

/* ListWorkflowServiceRevisionPodsOK describes a response with status code 200, with default header values.

successfully got list of a service revision pods
*/
type ListWorkflowServiceRevisionPodsOK struct {
}

func (o *ListWorkflowServiceRevisionPodsOK) Error() string {
	return fmt.Sprintf("[GET /api/functions/namespaces/{namespace}/tree/{workflow}?op=pods][%d] listWorkflowServiceRevisionPodsOK ", 200)
}

func (o *ListWorkflowServiceRevisionPodsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}