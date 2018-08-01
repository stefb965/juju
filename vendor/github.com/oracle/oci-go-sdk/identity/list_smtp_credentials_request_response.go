// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package identity

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListSmtpCredentialsRequest wrapper for the ListSmtpCredentials operation
type ListSmtpCredentialsRequest struct {

	// The OCID of the user.
	UserId *string `mandatory:"true" contributesTo:"path" name:"userId"`
}

func (request ListSmtpCredentialsRequest) String() string {
	return common.PointerString(request)
}

// ListSmtpCredentialsResponse wrapper for the ListSmtpCredentials operation
type ListSmtpCredentialsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The []SmtpCredentialSummary instance
	Items []SmtpCredentialSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListSmtpCredentialsResponse) String() string {
	return common.PointerString(response)
}
