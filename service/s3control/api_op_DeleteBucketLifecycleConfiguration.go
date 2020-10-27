// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"strings"
)

// This API action deletes an Amazon S3 on Outposts bucket's lifecycle
// configuration. To delete an S3 bucket's lifecycle configuration, see
// DeleteBucketLifecycle
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketLifecycle.html)
// in the Amazon Simple Storage Service API. Deletes the lifecycle configuration
// from the specified Outposts bucket. Amazon S3 on Outposts removes all the
// lifecycle configuration rules in the lifecycle subresource associated with the
// bucket. Your objects never expire, and Amazon S3 on Outposts no longer
// automatically deletes any objects on the basis of rules contained in the deleted
// lifecycle configuration. For more information, see Using Amazon S3 on Outposts
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/S3onOutposts.html) in Amazon
// Simple Storage Service Developer Guide. To use this operation, you must have
// permission to perform the s3outposts:DeleteLifecycleConfiguration action. By
// default, the bucket owner has this permission and the Outposts bucket owner can
// grant this permission to others. All Amazon S3 on Outposts REST API requests for
// this action require an additional parameter of outpost-id to be passed with the
// request and an S3 on Outposts endpoint hostname prefix instead of s3-control.
// For an example of the request syntax for Amazon S3 on Outposts that uses the S3
// on Outposts endpoint hostname prefix and the outpost-id derived using the access
// point ARN, see the  Example
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API__control_DeleteBucketLifecycleConfiguration.html#API_control_DeleteBucketLifecycleConfiguration_Examples)
// section below. For more information about object expiration, see  Elements to
// Describe Lifecycle Actions
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/intro-lifecycle-rules.html#intro-lifecycle-rules-actions).
// Related actions include:
//
//     * PutBucketLifecycleConfiguration
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketLifecycleConfiguration.html)
//
//
// * GetBucketLifecycleConfiguration
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketLifecycleConfiguration.html)
func (c *Client) DeleteBucketLifecycleConfiguration(ctx context.Context, params *DeleteBucketLifecycleConfigurationInput, optFns ...func(*Options)) (*DeleteBucketLifecycleConfigurationOutput, error) {
	if params == nil {
		params = &DeleteBucketLifecycleConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteBucketLifecycleConfiguration", params, optFns, addOperationDeleteBucketLifecycleConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteBucketLifecycleConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteBucketLifecycleConfigurationInput struct {

	// The account ID of the lifecycle configuration to delete.
	//
	// This member is required.
	AccountId *string

	// The bucket ARN of the bucket. For Amazon S3 on Outposts specify the ARN of the
	// bucket accessed in the format arn:aws:s3-outposts:::outpost//bucket/. For
	// example, to access the bucket reports through outpost my-outpost owned by
	// account 123456789012 in Region us-west-2, use the URL encoding of
	// arn:aws:s3-outposts:us-west-2:123456789012:outpost/my-outpost/bucket/reports.
	// The value must be URL encoded.
	//
	// This member is required.
	Bucket *string
}

type DeleteBucketLifecycleConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteBucketLifecycleConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpDeleteBucketLifecycleConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteBucketLifecycleConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addEndpointPrefix_opDeleteBucketLifecycleConfigurationMiddleware(stack)
	addOpDeleteBucketLifecycleConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteBucketLifecycleConfiguration(options.Region), middleware.Before)
	addMetadataRetrieverMiddleware(stack)
	addUpdateEndpointMiddleware(stack, options)
	addResponseErrorMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)
	return nil
}

type endpointPrefix_opDeleteBucketLifecycleConfigurationMiddleware struct {
}

func (*endpointPrefix_opDeleteBucketLifecycleConfigurationMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opDeleteBucketLifecycleConfigurationMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*DeleteBucketLifecycleConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input type %T", in.Parameters)
	}

	var prefix strings.Builder
	if input.AccountId == nil {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so may not be nil")}
	} else if !smithyhttp.ValidHostLabel(*input.AccountId) {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so must match \"[a-zA-Z0-9-]{1,63}\", but was \"%s\"", *input.AccountId)}
	} else {
		prefix.WriteString(*input.AccountId)
	}
	prefix.WriteString(".")
	req.HostPrefix = prefix.String()

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opDeleteBucketLifecycleConfigurationMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opDeleteBucketLifecycleConfigurationMiddleware{}, `OperationSerializer`, middleware.Before)
}

func newServiceMetadataMiddleware_opDeleteBucketLifecycleConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "DeleteBucketLifecycleConfiguration",
	}
}
