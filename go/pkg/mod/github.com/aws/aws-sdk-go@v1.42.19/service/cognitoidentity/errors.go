// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentity

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeConcurrentModificationException for service response error code
	// "ConcurrentModificationException".
	//
	// Thrown if there are parallel requests to modify a resource.
	ErrCodeConcurrentModificationException = "ConcurrentModificationException"

	// ErrCodeDeveloperUserAlreadyRegisteredException for service response error code
	// "DeveloperUserAlreadyRegisteredException".
	//
	// The provided developer user identifier is already registered with Cognito
	// under a different identity ID.
	ErrCodeDeveloperUserAlreadyRegisteredException = "DeveloperUserAlreadyRegisteredException"

	// ErrCodeExternalServiceException for service response error code
	// "ExternalServiceException".
	//
	// An exception thrown when a dependent service such as Facebook or Twitter
	// is not responding
	ErrCodeExternalServiceException = "ExternalServiceException"

	// ErrCodeInternalErrorException for service response error code
	// "InternalErrorException".
	//
	// Thrown when the service encounters an error during processing the request.
	ErrCodeInternalErrorException = "InternalErrorException"

	// ErrCodeInvalidIdentityPoolConfigurationException for service response error code
	// "InvalidIdentityPoolConfigurationException".
	//
	// Thrown if the identity pool has no role associated for the given auth type
	// (auth/unauth) or if the AssumeRole fails.
	ErrCodeInvalidIdentityPoolConfigurationException = "InvalidIdentityPoolConfigurationException"

	// ErrCodeInvalidParameterException for service response error code
	// "InvalidParameterException".
	//
	// Thrown for missing or bad input parameter(s).
	ErrCodeInvalidParameterException = "InvalidParameterException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// Thrown when the total number of user pools has exceeded a preset limit.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeNotAuthorizedException for service response error code
	// "NotAuthorizedException".
	//
	// Thrown when a user is not authorized to access the requested resource.
	ErrCodeNotAuthorizedException = "NotAuthorizedException"

	// ErrCodeResourceConflictException for service response error code
	// "ResourceConflictException".
	//
	// Thrown when a user tries to use a login which is already linked to another
	// account.
	ErrCodeResourceConflictException = "ResourceConflictException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// Thrown when the requested resource (for example, a dataset or record) does
	// not exist.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeTooManyRequestsException for service response error code
	// "TooManyRequestsException".
	//
	// Thrown when a request is throttled.
	ErrCodeTooManyRequestsException = "TooManyRequestsException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"ConcurrentModificationException":           newErrorConcurrentModificationException,
	"DeveloperUserAlreadyRegisteredException":   newErrorDeveloperUserAlreadyRegisteredException,
	"ExternalServiceException":                  newErrorExternalServiceException,
	"InternalErrorException":                    newErrorInternalErrorException,
	"InvalidIdentityPoolConfigurationException": newErrorInvalidIdentityPoolConfigurationException,
	"InvalidParameterException":                 newErrorInvalidParameterException,
	"LimitExceededException":                    newErrorLimitExceededException,
	"NotAuthorizedException":                    newErrorNotAuthorizedException,
	"ResourceConflictException":                 newErrorResourceConflictException,
	"ResourceNotFoundException":                 newErrorResourceNotFoundException,
	"TooManyRequestsException":                  newErrorTooManyRequestsException,
}
