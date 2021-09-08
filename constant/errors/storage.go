package errors

import "errors"

var (
	ErrBucketDeleted                     = errors.New("bucket is already deleted")
	ErrAccessDenied                      = errors.New("access Denied")
	ErrBadDigest                         = errors.New("the Content-Md5 you specified did not match what we received")
	ErrEntityTooSmall                    = errors.New("your proposed upload is smaller than the minimum allowed object size")
	ErrEntityTooLarge                    = errors.New("your proposed upload exceeds the maximum allowed object size")
	ErrIncompleteBody                    = errors.New("you did not provide the number of bytes specified by the Content-Length HTTP header")
	ErrInternalError                     = errors.New("we encountered an internal error, please try again")
	ErrInvalidAccessKeyId                = errors.New("the access key ID you provided does not exist in our records")
	ErrInvalidBucketName                 = errors.New("the specified bucket is not valid")
	ErrInvalidDigest                     = errors.New("the Content-Md5 you specified is not valid")
	ErrInvalidRange                      = errors.New("the requested range is not satisfiable")
	ErrMalformedXML                      = errors.New("the XML you provided was not well-formed or did not validate against our published schema")
	ErrMissingContentLength              = errors.New("you must provide the Content-Length HTTP header")
	ErrMissingContentMD5                 = errors.New("missing required header for this request: Content-Md5")
	ErrMissingRequestBodyError           = errors.New("request body is empty")
	ErrNoSuchBucket                      = errors.New("the specified bucket does not exist")
	ErrNoSuchBucketPolicy                = errors.New("the bucket policy does not exist")
	ErrNoSuchKey                         = errors.New("the specified key does not exist")
	ErrNoSuchUpload                      = errors.New("the specified multipart upload does not exist. The upload ID may be invalid, or the upload may have been aborted or completed")
	ErrStorageNotImplemented             = errors.New("a header you provided implies functionality that is not implemented")
	ErrPreconditionFailed                = errors.New("at least one of the pre-conditions you specified did not hold")
	ErrRequestTimeTooSkewed              = errors.New("the difference between the request time and the server's time is too large")
	ErrSignatureDoesNotMatch             = errors.New("the request signature we calculated does not match the signature you provided. Check your key and signing method")
	ErrMethodNotAllowed                  = errors.New("the specified method is not allowed against this resource")
	ErrInvalidPart                       = errors.New("one or more of the specified parts could not be found")
	ErrInvalidPartOrder                  = errors.New("the list of parts was not in ascending order. The parts list must be specified in order by part number")
	ErrInvalidObjectState                = errors.New("the operation is not valid for the current state of the object")
	ErrAuthorizationHeaderMalformed      = errors.New("the authorization header is malformed; the region is wrong")
	ErrMalformedPOSTRequest              = errors.New("the body of your POST request is not well-formed multipart/form-data")
	ErrBucketNotEmpty                    = errors.New("the bucket you tried to delete is not empty")
	ErrAllAccessDisabled                 = errors.New("all access to this bucket has been disabled")
	ErrMalformedPolicy                   = errors.New("policy has invalid resource")
	ErrMissingFields                     = errors.New("missing fields in request")
	ErrAuthorizationQueryParametersError = errors.New("error parsing the X-Amz-Credential parameter; the Credential is mal-formed; expecting \"<YOUR-AKID>/YYYYMMDD/REGION/SERVICE/aws4_request\"")
	ErrMalformedDate                     = errors.New("invalid date format header, expected to be in ISO8601, RFC1123 or RFC1123Z time format")
	ErrBucketAlreadyOwnedByYou           = errors.New("your previous request to create the named bucket succeeded and you already own it")
	ErrInvalidDuration                   = errors.New("duration provided in the request is invalid")
	ErrXAmzContentSHA256Mismatch         = errors.New("the provided 'x-amz-content-sha256' header does not match what was computed")
)

var StorageCodeErrorMapping = map[string]error{
	"AccessDenied":                      ErrAccessDenied,
	"BadDigest":                         ErrBadDigest,
	"EntityTooSmall":                    ErrEntityTooSmall,
	"EntityTooLarge":                    ErrEntityTooLarge,
	"IncompleteBody":                    ErrIncompleteBody,
	"InternalError":                     ErrInternalError,
	"InvalidAccessKeyId":                ErrInvalidAccessKeyId,
	"InvalidBucketName":                 ErrInvalidBucketName,
	"InvalidDigest":                     ErrInvalidDigest,
	"InvalidRange":                      ErrInvalidRange,
	"MalformedXML":                      ErrMalformedXML,
	"MissingContentLength":              ErrMissingContentLength,
	"MissingContentMD5":                 ErrMissingContentMD5,
	"MissingRequestBodyError":           ErrMissingRequestBodyError,
	"NoSuchBucket":                      ErrNoSuchBucket,
	"NoSuchBucketPolicy":                ErrNoSuchBucketPolicy,
	"NoSuchKey":                         ErrNoSuchKey,
	"NoSuchUpload":                      ErrNoSuchUpload,
	"NotImplemented":                    ErrStorageNotImplemented,
	"PreconditionFailed":                ErrPreconditionFailed,
	"RequestTimeTooSkewed":              ErrRequestTimeTooSkewed,
	"SignatureDoesNotMatch":             ErrSignatureDoesNotMatch,
	"MethodNotAllowed":                  ErrMethodNotAllowed,
	"InvalidPart":                       ErrInvalidPart,
	"InvalidPartOrder":                  ErrInvalidPartOrder,
	"InvalidObjectState":                ErrInvalidObjectState,
	"AuthorizationHeaderMalformed":      ErrAuthorizationHeaderMalformed,
	"MalformedPOSTRequest":              ErrMalformedPOSTRequest,
	"BucketNotEmpty":                    ErrBucketNotEmpty,
	"AllAccessDisabled":                 ErrAllAccessDisabled,
	"MalformedPolicy":                   ErrMalformedPolicy,
	"MissingFields":                     ErrMissingFields,
	"AuthorizationQueryParametersError": ErrAuthorizationQueryParametersError,
	"MalformedDate":                     ErrMalformedDate,
	"BucketAlreadyOwnedByYou":           ErrBucketAlreadyOwnedByYou,
	"InvalidDuration":                   ErrInvalidDuration,
	"XAmzContentSHA256Mismatch":         ErrXAmzContentSHA256Mismatch,
}
