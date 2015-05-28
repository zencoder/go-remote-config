package s3_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/aws/awsutil"
	"github.com/awslabs/aws-sdk-go/service/s3"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleS3_AbortMultipartUpload() {
	svc := s3.New(nil)

	params := &s3.AbortMultipartUploadInput{
		Bucket:       aws.String("BucketName"),        // Required
		Key:          aws.String("ObjectKey"),         // Required
		UploadID:     aws.String("MultipartUploadId"), // Required
		RequestPayer: aws.String("RequestPayer"),
	}
	resp, err := svc.AbortMultipartUpload(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleS3_CompleteMultipartUpload() {
	svc := s3.New(nil)

	params := &s3.CompleteMultipartUploadInput{
		Bucket:   aws.String("BucketName"),        // Required
		Key:      aws.String("ObjectKey"),         // Required
		UploadID: aws.String("MultipartUploadId"), // Required
		MultipartUpload: &s3.CompletedMultipartUpload{
			Parts: []*s3.CompletedPart{
				&s3.CompletedPart{ // Required
					ETag:       aws.String("ETag"),
					PartNumber: aws.Long(1),
				},
				// More values...
			},
		},
		RequestPayer: aws.String("RequestPayer"),
	}
	resp, err := svc.CompleteMultipartUpload(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleS3_CopyObject() {
	svc := s3.New(nil)

	params := &s3.CopyObjectInput{
		Bucket:                         aws.String("BucketName"), // Required
		CopySource:                     aws.String("CopySource"), // Required
		Key:                            aws.String("ObjectKey"),  // Required
		ACL:                            aws.String("ObjectCannedACL"),
		CacheControl:                   aws.String("CacheControl"),
		ContentDisposition:             aws.String("ContentDisposition"),
		ContentEncoding:                aws.String("ContentEncoding"),
		ContentLanguage:                aws.String("ContentLanguage"),
		ContentType:                    aws.String("ContentType"),
		CopySourceIfMatch:              aws.String("CopySourceIfMatch"),
		CopySourceIfModifiedSince:      aws.Time(time.Now()),
		CopySourceIfNoneMatch:          aws.String("CopySourceIfNoneMatch"),
		CopySourceIfUnmodifiedSince:    aws.Time(time.Now()),
		CopySourceSSECustomerAlgorithm: aws.String("CopySourceSSECustomerAlgorithm"),
		CopySourceSSECustomerKey:       aws.String("CopySourceSSECustomerKey"),
		CopySourceSSECustomerKeyMD5:    aws.String("CopySourceSSECustomerKeyMD5"),
		Expires:                        aws.Time(time.Now()),
		GrantFullControl:               aws.String("GrantFullControl"),
		GrantRead:                      aws.String("GrantRead"),
		GrantReadACP:                   aws.String("GrantReadACP"),
		GrantWriteACP:                  aws.String("GrantWriteACP"),
		Metadata: &map[string]*string{
			"Key": aws.String("MetadataValue"), // Required
			// More values...
		},
		MetadataDirective:       aws.String("MetadataDirective"),
		RequestPayer:            aws.String("RequestPayer"),
		SSECustomerAlgorithm:    aws.String("SSECustomerAlgorithm"),
		SSECustomerKey:          aws.String("SSECustomerKey"),
		SSECustomerKeyMD5:       aws.String("SSECustomerKeyMD5"),
		SSEKMSKeyID:             aws.String("SSEKMSKeyId"),
		ServerSideEncryption:    aws.String("ServerSideEncryption"),
		StorageClass:            aws.String("StorageClass"),
		WebsiteRedirectLocation: aws.String("WebsiteRedirectLocation"),
	}
	resp, err := svc.CopyObject(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleS3_CreateBucket() {
	svc := s3.New(nil)

	params := &s3.CreateBucketInput{
		Bucket: aws.String("BucketName"), // Required
		ACL:    aws.String("BucketCannedACL"),
		CreateBucketConfiguration: &s3.CreateBucketConfiguration{
			LocationConstraint: aws.String("BucketLocationConstraint"),
		},
		GrantFullControl: aws.String("GrantFullControl"),
		GrantRead:        aws.String("GrantRead"),
		GrantReadACP:     aws.String("GrantReadACP"),
		GrantWrite:       aws.String("GrantWrite"),
		GrantWriteACP:    aws.String("GrantWriteACP"),
	}
	resp, err := svc.CreateBucket(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleS3_CreateMultipartUpload() {
	svc := s3.New(nil)

	params := &s3.CreateMultipartUploadInput{
		Bucket:             aws.String("BucketName"), // Required
		Key:                aws.String("ObjectKey"),  // Required
		ACL:                aws.String("ObjectCannedACL"),
		CacheControl:       aws.String("CacheControl"),
		ContentDisposition: aws.String("ContentDisposition"),
		ContentEncoding:    aws.String("ContentEncoding"),
		ContentLanguage:    aws.String("ContentLanguage"),
		ContentType:        aws.String("ContentType"),
		Expires:            aws.Time(time.Now()),
		GrantFullControl:   aws.String("GrantFullControl"),
		GrantRead:          aws.String("GrantRead"),
		GrantReadACP:       aws.String("GrantReadACP"),
		GrantWriteACP:      aws.String("GrantWriteACP"),
		Metadata: &map[string]*string{
			"Key": aws.String("MetadataValue"), // Required
			// More values...
		},
		RequestPayer:            aws.String("RequestPayer"),
		SSECustomerAlgorithm:    aws.String("SSECustomerAlgorithm"),
		SSECustomerKey:          aws.String("SSECustomerKey"),
		SSECustomerKeyMD5:       aws.String("SSECustomerKeyMD5"),
		SSEKMSKeyID:             aws.String("SSEKMSKeyId"),
		ServerSideEncryption:    aws.String("ServerSideEncryption"),
		StorageClass:            aws.String("StorageClass"),
		WebsiteRedirectLocation: aws.String("WebsiteRedirectLocation"),
	}
	resp, err := svc.CreateMultipartUpload(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleS3_DeleteBucket() {
	svc := s3.New(nil)

	params := &s3.DeleteBucketInput{
		Bucket: aws.String("BucketName"), // Required
	}
	resp, err := svc.DeleteBucket(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleS3_DeleteBucketCORS() {
	svc := s3.New(nil)

	params := &s3.DeleteBucketCORSInput{
		Bucket: aws.String("BucketName"), // Required
	}
	resp, err := svc.DeleteBucketCORS(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleS3_DeleteBucketLifecycle() {
	svc := s3.New(nil)

	params := &s3.DeleteBucketLifecycleInput{
		Bucket: aws.String("BucketName"), // Required
	}
	resp, err := svc.DeleteBucketLifecycle(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleS3_DeleteBucketPolicy() {
	svc := s3.New(nil)

	params := &s3.DeleteBucketPolicyInput{
		Bucket: aws.String("BucketName"), // Required
	}
	resp, err := svc.DeleteBucketPolicy(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleS3_DeleteBucketReplication() {
	svc := s3.New(nil)

	params := &s3.DeleteBucketReplicationInput{
		Bucket: aws.String("BucketName"), // Required
	}
	resp, err := svc.DeleteBucketReplication(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleS3_DeleteBucketTagging() {
	svc := s3.New(nil)

	params := &s3.DeleteBucketTaggingInput{
		Bucket: aws.String("BucketName"), // Required
	}
	resp, err := svc.DeleteBucketTagging(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleS3_DeleteBucketWebsite() {
	svc := s3.New(nil)

	params := &s3.DeleteBucketWebsiteInput{
		Bucket: aws.String("BucketName"), // Required
	}
	resp, err := svc.DeleteBucketWebsite(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleS3_DeleteObject() {
	svc := s3.New(nil)

	params := &s3.DeleteObjectInput{
		Bucket:       aws.String("BucketName"), // Required
		Key:          aws.String("ObjectKey"),  // Required
		MFA:          aws.String("MFA"),
		RequestPayer: aws.String("RequestPayer"),
		VersionID:    aws.String("ObjectVersionId"),
	}
	resp, err := svc.DeleteObject(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleS3_DeleteObjects() {
	svc := s3.New(nil)

	params := &s3.DeleteObjectsInput{
		Bucket: aws.String("BucketName"), // Required
		Delete: &s3.Delete{ // Required
			Objects: []*s3.ObjectIdentifier{ // Required
				&s3.ObjectIdentifier{ // Required
					Key:       aws.String("ObjectKey"), // Required
					VersionID: aws.String("ObjectVersionId"),
				},
				// More values...
			},
			Quiet: aws.Boolean(true),
		},
		MFA:          aws.String("MFA"),
		RequestPayer: aws.String("RequestPayer"),
	}
	resp, err := svc.DeleteObjects(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleS3_GetBucketACL() {
	svc := s3.New(nil)

	params := &s3.GetBucketACLInput{
		Bucket: aws.String("BucketName"), // Required
	}
	resp, err := svc.GetBucketACL(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleS3_GetBucketCORS() {
	svc := s3.New(nil)

	params := &s3.GetBucketCORSInput{
		Bucket: aws.String("BucketName"), // Required
	}
	resp, err := svc.GetBucketCORS(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleS3_GetBucketLifecycle() {
	svc := s3.New(nil)

	params := &s3.GetBucketLifecycleInput{
		Bucket: aws.String("BucketName"), // Required
	}
	resp, err := svc.GetBucketLifecycle(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleS3_GetBucketLocation() {
	svc := s3.New(nil)

	params := &s3.GetBucketLocationInput{
		Bucket: aws.String("BucketName"), // Required
	}
	resp, err := svc.GetBucketLocation(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleS3_GetBucketLogging() {
	svc := s3.New(nil)

	params := &s3.GetBucketLoggingInput{
		Bucket: aws.String("BucketName"), // Required
	}
	resp, err := svc.GetBucketLogging(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleS3_GetBucketNotification() {
	svc := s3.New(nil)

	params := &s3.GetBucketNotificationInput{
		Bucket: aws.String("BucketName"), // Required
	}
	resp, err := svc.GetBucketNotification(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleS3_GetBucketPolicy() {
	svc := s3.New(nil)

	params := &s3.GetBucketPolicyInput{
		Bucket: aws.String("BucketName"), // Required
	}
	resp, err := svc.GetBucketPolicy(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleS3_GetBucketReplication() {
	svc := s3.New(nil)

	params := &s3.GetBucketReplicationInput{
		Bucket: aws.String("BucketName"), // Required
	}
	resp, err := svc.GetBucketReplication(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleS3_GetBucketRequestPayment() {
	svc := s3.New(nil)

	params := &s3.GetBucketRequestPaymentInput{
		Bucket: aws.String("BucketName"), // Required
	}
	resp, err := svc.GetBucketRequestPayment(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleS3_GetBucketTagging() {
	svc := s3.New(nil)

	params := &s3.GetBucketTaggingInput{
		Bucket: aws.String("BucketName"), // Required
	}
	resp, err := svc.GetBucketTagging(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleS3_GetBucketVersioning() {
	svc := s3.New(nil)

	params := &s3.GetBucketVersioningInput{
		Bucket: aws.String("BucketName"), // Required
	}
	resp, err := svc.GetBucketVersioning(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleS3_GetBucketWebsite() {
	svc := s3.New(nil)

	params := &s3.GetBucketWebsiteInput{
		Bucket: aws.String("BucketName"), // Required
	}
	resp, err := svc.GetBucketWebsite(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleS3_GetObject() {
	svc := s3.New(nil)

	params := &s3.GetObjectInput{
		Bucket:                     aws.String("BucketName"), // Required
		Key:                        aws.String("ObjectKey"),  // Required
		IfMatch:                    aws.String("IfMatch"),
		IfModifiedSince:            aws.Time(time.Now()),
		IfNoneMatch:                aws.String("IfNoneMatch"),
		IfUnmodifiedSince:          aws.Time(time.Now()),
		Range:                      aws.String("Range"),
		RequestPayer:               aws.String("RequestPayer"),
		ResponseCacheControl:       aws.String("ResponseCacheControl"),
		ResponseContentDisposition: aws.String("ResponseContentDisposition"),
		ResponseContentEncoding:    aws.String("ResponseContentEncoding"),
		ResponseContentLanguage:    aws.String("ResponseContentLanguage"),
		ResponseContentType:        aws.String("ResponseContentType"),
		ResponseExpires:            aws.Time(time.Now()),
		SSECustomerAlgorithm:       aws.String("SSECustomerAlgorithm"),
		SSECustomerKey:             aws.String("SSECustomerKey"),
		SSECustomerKeyMD5:          aws.String("SSECustomerKeyMD5"),
		VersionID:                  aws.String("ObjectVersionId"),
	}
	resp, err := svc.GetObject(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleS3_GetObjectACL() {
	svc := s3.New(nil)

	params := &s3.GetObjectACLInput{
		Bucket:       aws.String("BucketName"), // Required
		Key:          aws.String("ObjectKey"),  // Required
		RequestPayer: aws.String("RequestPayer"),
		VersionID:    aws.String("ObjectVersionId"),
	}
	resp, err := svc.GetObjectACL(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleS3_GetObjectTorrent() {
	svc := s3.New(nil)

	params := &s3.GetObjectTorrentInput{
		Bucket:       aws.String("BucketName"), // Required
		Key:          aws.String("ObjectKey"),  // Required
		RequestPayer: aws.String("RequestPayer"),
	}
	resp, err := svc.GetObjectTorrent(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleS3_HeadBucket() {
	svc := s3.New(nil)

	params := &s3.HeadBucketInput{
		Bucket: aws.String("BucketName"), // Required
	}
	resp, err := svc.HeadBucket(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleS3_HeadObject() {
	svc := s3.New(nil)

	params := &s3.HeadObjectInput{
		Bucket:               aws.String("BucketName"), // Required
		Key:                  aws.String("ObjectKey"),  // Required
		IfMatch:              aws.String("IfMatch"),
		IfModifiedSince:      aws.Time(time.Now()),
		IfNoneMatch:          aws.String("IfNoneMatch"),
		IfUnmodifiedSince:    aws.Time(time.Now()),
		Range:                aws.String("Range"),
		RequestPayer:         aws.String("RequestPayer"),
		SSECustomerAlgorithm: aws.String("SSECustomerAlgorithm"),
		SSECustomerKey:       aws.String("SSECustomerKey"),
		SSECustomerKeyMD5:    aws.String("SSECustomerKeyMD5"),
		VersionID:            aws.String("ObjectVersionId"),
	}
	resp, err := svc.HeadObject(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleS3_ListBuckets() {
	svc := s3.New(nil)

	var params *s3.ListBucketsInput
	resp, err := svc.ListBuckets(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleS3_ListMultipartUploads() {
	svc := s3.New(nil)

	params := &s3.ListMultipartUploadsInput{
		Bucket:         aws.String("BucketName"), // Required
		Delimiter:      aws.String("Delimiter"),
		EncodingType:   aws.String("EncodingType"),
		KeyMarker:      aws.String("KeyMarker"),
		MaxUploads:     aws.Long(1),
		Prefix:         aws.String("Prefix"),
		UploadIDMarker: aws.String("UploadIdMarker"),
	}
	resp, err := svc.ListMultipartUploads(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleS3_ListObjectVersions() {
	svc := s3.New(nil)

	params := &s3.ListObjectVersionsInput{
		Bucket:          aws.String("BucketName"), // Required
		Delimiter:       aws.String("Delimiter"),
		EncodingType:    aws.String("EncodingType"),
		KeyMarker:       aws.String("KeyMarker"),
		MaxKeys:         aws.Long(1),
		Prefix:          aws.String("Prefix"),
		VersionIDMarker: aws.String("VersionIdMarker"),
	}
	resp, err := svc.ListObjectVersions(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleS3_ListObjects() {
	svc := s3.New(nil)

	params := &s3.ListObjectsInput{
		Bucket:       aws.String("BucketName"), // Required
		Delimiter:    aws.String("Delimiter"),
		EncodingType: aws.String("EncodingType"),
		Marker:       aws.String("Marker"),
		MaxKeys:      aws.Long(1),
		Prefix:       aws.String("Prefix"),
	}
	resp, err := svc.ListObjects(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleS3_ListParts() {
	svc := s3.New(nil)

	params := &s3.ListPartsInput{
		Bucket:           aws.String("BucketName"),        // Required
		Key:              aws.String("ObjectKey"),         // Required
		UploadID:         aws.String("MultipartUploadId"), // Required
		MaxParts:         aws.Long(1),
		PartNumberMarker: aws.Long(1),
		RequestPayer:     aws.String("RequestPayer"),
	}
	resp, err := svc.ListParts(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleS3_PutBucketACL() {
	svc := s3.New(nil)

	params := &s3.PutBucketACLInput{
		Bucket: aws.String("BucketName"), // Required
		ACL:    aws.String("BucketCannedACL"),
		AccessControlPolicy: &s3.AccessControlPolicy{
			Grants: []*s3.Grant{
				&s3.Grant{ // Required
					Grantee: &s3.Grantee{
						Type:         aws.String("Type"), // Required
						DisplayName:  aws.String("DisplayName"),
						EmailAddress: aws.String("EmailAddress"),
						ID:           aws.String("ID"),
						URI:          aws.String("URI"),
					},
					Permission: aws.String("Permission"),
				},
				// More values...
			},
			Owner: &s3.Owner{
				DisplayName: aws.String("DisplayName"),
				ID:          aws.String("ID"),
			},
		},
		ContentMD5:       aws.String("ContentMD5"),
		GrantFullControl: aws.String("GrantFullControl"),
		GrantRead:        aws.String("GrantRead"),
		GrantReadACP:     aws.String("GrantReadACP"),
		GrantWrite:       aws.String("GrantWrite"),
		GrantWriteACP:    aws.String("GrantWriteACP"),
	}
	resp, err := svc.PutBucketACL(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleS3_PutBucketCORS() {
	svc := s3.New(nil)

	params := &s3.PutBucketCORSInput{
		Bucket: aws.String("BucketName"), // Required
		CORSConfiguration: &s3.CORSConfiguration{
			CORSRules: []*s3.CORSRule{
				&s3.CORSRule{ // Required
					AllowedHeaders: []*string{
						aws.String("AllowedHeader"), // Required
						// More values...
					},
					AllowedMethods: []*string{
						aws.String("AllowedMethod"), // Required
						// More values...
					},
					AllowedOrigins: []*string{
						aws.String("AllowedOrigin"), // Required
						// More values...
					},
					ExposeHeaders: []*string{
						aws.String("ExposeHeader"), // Required
						// More values...
					},
					MaxAgeSeconds: aws.Long(1),
				},
				// More values...
			},
		},
		ContentMD5: aws.String("ContentMD5"),
	}
	resp, err := svc.PutBucketCORS(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleS3_PutBucketLifecycle() {
	svc := s3.New(nil)

	params := &s3.PutBucketLifecycleInput{
		Bucket:     aws.String("BucketName"), // Required
		ContentMD5: aws.String("ContentMD5"),
		LifecycleConfiguration: &s3.LifecycleConfiguration{
			Rules: []*s3.Rule{ // Required
				&s3.Rule{ // Required
					Prefix: aws.String("Prefix"),           // Required
					Status: aws.String("ExpirationStatus"), // Required
					Expiration: &s3.LifecycleExpiration{
						Date: aws.Time(time.Now()),
						Days: aws.Long(1),
					},
					ID: aws.String("ID"),
					NoncurrentVersionExpiration: &s3.NoncurrentVersionExpiration{
						NoncurrentDays: aws.Long(1),
					},
					NoncurrentVersionTransition: &s3.NoncurrentVersionTransition{
						NoncurrentDays: aws.Long(1),
						StorageClass:   aws.String("TransitionStorageClass"),
					},
					Transition: &s3.Transition{
						Date:         aws.Time(time.Now()),
						Days:         aws.Long(1),
						StorageClass: aws.String("TransitionStorageClass"),
					},
				},
				// More values...
			},
		},
	}
	resp, err := svc.PutBucketLifecycle(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleS3_PutBucketLogging() {
	svc := s3.New(nil)

	params := &s3.PutBucketLoggingInput{
		Bucket: aws.String("BucketName"), // Required
		BucketLoggingStatus: &s3.BucketLoggingStatus{ // Required
			LoggingEnabled: &s3.LoggingEnabled{
				TargetBucket: aws.String("TargetBucket"),
				TargetGrants: []*s3.TargetGrant{
					&s3.TargetGrant{ // Required
						Grantee: &s3.Grantee{
							Type:         aws.String("Type"), // Required
							DisplayName:  aws.String("DisplayName"),
							EmailAddress: aws.String("EmailAddress"),
							ID:           aws.String("ID"),
							URI:          aws.String("URI"),
						},
						Permission: aws.String("BucketLogsPermission"),
					},
					// More values...
				},
				TargetPrefix: aws.String("TargetPrefix"),
			},
		},
		ContentMD5: aws.String("ContentMD5"),
	}
	resp, err := svc.PutBucketLogging(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleS3_PutBucketNotification() {
	svc := s3.New(nil)

	params := &s3.PutBucketNotificationInput{
		Bucket: aws.String("BucketName"), // Required
		NotificationConfiguration: &s3.NotificationConfiguration{ // Required
			CloudFunctionConfiguration: &s3.CloudFunctionConfiguration{
				CloudFunction: aws.String("CloudFunction"),
				Event:         aws.String("Event"),
				Events: []*string{
					aws.String("Event"), // Required
					// More values...
				},
				ID:             aws.String("NotificationId"),
				InvocationRole: aws.String("CloudFunctionInvocationRole"),
			},
			QueueConfiguration: &s3.QueueConfiguration{
				Event: aws.String("Event"),
				Events: []*string{
					aws.String("Event"), // Required
					// More values...
				},
				ID:    aws.String("NotificationId"),
				Queue: aws.String("Queue"),
			},
			TopicConfiguration: &s3.TopicConfiguration{
				Event: aws.String("Event"),
				Events: []*string{
					aws.String("Event"), // Required
					// More values...
				},
				ID:    aws.String("NotificationId"),
				Topic: aws.String("Topic"),
			},
		},
		ContentMD5: aws.String("ContentMD5"),
	}
	resp, err := svc.PutBucketNotification(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleS3_PutBucketPolicy() {
	svc := s3.New(nil)

	params := &s3.PutBucketPolicyInput{
		Bucket:     aws.String("BucketName"), // Required
		Policy:     aws.String("Policy"),     // Required
		ContentMD5: aws.String("ContentMD5"),
	}
	resp, err := svc.PutBucketPolicy(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleS3_PutBucketReplication() {
	svc := s3.New(nil)

	params := &s3.PutBucketReplicationInput{
		Bucket: aws.String("BucketName"), // Required
		ReplicationConfiguration: &s3.ReplicationConfiguration{ // Required
			Role: aws.String("Role"), // Required
			Rules: []*s3.ReplicationRule{ // Required
				&s3.ReplicationRule{ // Required
					Destination: &s3.Destination{ // Required
						Bucket: aws.String("BucketName"), // Required
					},
					Prefix: aws.String("Prefix"),                // Required
					Status: aws.String("ReplicationRuleStatus"), // Required
					ID:     aws.String("ID"),
				},
				// More values...
			},
		},
		ContentMD5: aws.String("ContentMD5"),
	}
	resp, err := svc.PutBucketReplication(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleS3_PutBucketRequestPayment() {
	svc := s3.New(nil)

	params := &s3.PutBucketRequestPaymentInput{
		Bucket: aws.String("BucketName"), // Required
		RequestPaymentConfiguration: &s3.RequestPaymentConfiguration{ // Required
			Payer: aws.String("Payer"), // Required
		},
		ContentMD5: aws.String("ContentMD5"),
	}
	resp, err := svc.PutBucketRequestPayment(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleS3_PutBucketTagging() {
	svc := s3.New(nil)

	params := &s3.PutBucketTaggingInput{
		Bucket: aws.String("BucketName"), // Required
		Tagging: &s3.Tagging{ // Required
			TagSet: []*s3.Tag{ // Required
				&s3.Tag{ // Required
					Key:   aws.String("ObjectKey"), // Required
					Value: aws.String("Value"),     // Required
				},
				// More values...
			},
		},
		ContentMD5: aws.String("ContentMD5"),
	}
	resp, err := svc.PutBucketTagging(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleS3_PutBucketVersioning() {
	svc := s3.New(nil)

	params := &s3.PutBucketVersioningInput{
		Bucket: aws.String("BucketName"), // Required
		VersioningConfiguration: &s3.VersioningConfiguration{ // Required
			MFADelete: aws.String("MFADelete"),
			Status:    aws.String("BucketVersioningStatus"),
		},
		ContentMD5: aws.String("ContentMD5"),
		MFA:        aws.String("MFA"),
	}
	resp, err := svc.PutBucketVersioning(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleS3_PutBucketWebsite() {
	svc := s3.New(nil)

	params := &s3.PutBucketWebsiteInput{
		Bucket: aws.String("BucketName"), // Required
		WebsiteConfiguration: &s3.WebsiteConfiguration{ // Required
			ErrorDocument: &s3.ErrorDocument{
				Key: aws.String("ObjectKey"), // Required
			},
			IndexDocument: &s3.IndexDocument{
				Suffix: aws.String("Suffix"), // Required
			},
			RedirectAllRequestsTo: &s3.RedirectAllRequestsTo{
				HostName: aws.String("HostName"), // Required
				Protocol: aws.String("Protocol"),
			},
			RoutingRules: []*s3.RoutingRule{
				&s3.RoutingRule{ // Required
					Redirect: &s3.Redirect{ // Required
						HTTPRedirectCode:     aws.String("HttpRedirectCode"),
						HostName:             aws.String("HostName"),
						Protocol:             aws.String("Protocol"),
						ReplaceKeyPrefixWith: aws.String("ReplaceKeyPrefixWith"),
						ReplaceKeyWith:       aws.String("ReplaceKeyWith"),
					},
					Condition: &s3.Condition{
						HTTPErrorCodeReturnedEquals: aws.String("HttpErrorCodeReturnedEquals"),
						KeyPrefixEquals:             aws.String("KeyPrefixEquals"),
					},
				},
				// More values...
			},
		},
		ContentMD5: aws.String("ContentMD5"),
	}
	resp, err := svc.PutBucketWebsite(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleS3_PutObject() {
	svc := s3.New(nil)

	params := &s3.PutObjectInput{
		Bucket:             aws.String("BucketName"), // Required
		Key:                aws.String("ObjectKey"),  // Required
		ACL:                aws.String("ObjectCannedACL"),
		Body:               bytes.NewReader([]byte("PAYLOAD")),
		CacheControl:       aws.String("CacheControl"),
		ContentDisposition: aws.String("ContentDisposition"),
		ContentEncoding:    aws.String("ContentEncoding"),
		ContentLanguage:    aws.String("ContentLanguage"),
		ContentLength:      aws.Long(1),
		ContentMD5:         aws.String("ContentMD5"),
		ContentType:        aws.String("ContentType"),
		Expires:            aws.Time(time.Now()),
		GrantFullControl:   aws.String("GrantFullControl"),
		GrantRead:          aws.String("GrantRead"),
		GrantReadACP:       aws.String("GrantReadACP"),
		GrantWriteACP:      aws.String("GrantWriteACP"),
		Metadata: &map[string]*string{
			"Key": aws.String("MetadataValue"), // Required
			// More values...
		},
		RequestPayer:            aws.String("RequestPayer"),
		SSECustomerAlgorithm:    aws.String("SSECustomerAlgorithm"),
		SSECustomerKey:          aws.String("SSECustomerKey"),
		SSECustomerKeyMD5:       aws.String("SSECustomerKeyMD5"),
		SSEKMSKeyID:             aws.String("SSEKMSKeyId"),
		ServerSideEncryption:    aws.String("ServerSideEncryption"),
		StorageClass:            aws.String("StorageClass"),
		WebsiteRedirectLocation: aws.String("WebsiteRedirectLocation"),
	}
	resp, err := svc.PutObject(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleS3_PutObjectACL() {
	svc := s3.New(nil)

	params := &s3.PutObjectACLInput{
		Bucket: aws.String("BucketName"), // Required
		Key:    aws.String("ObjectKey"),  // Required
		ACL:    aws.String("ObjectCannedACL"),
		AccessControlPolicy: &s3.AccessControlPolicy{
			Grants: []*s3.Grant{
				&s3.Grant{ // Required
					Grantee: &s3.Grantee{
						Type:         aws.String("Type"), // Required
						DisplayName:  aws.String("DisplayName"),
						EmailAddress: aws.String("EmailAddress"),
						ID:           aws.String("ID"),
						URI:          aws.String("URI"),
					},
					Permission: aws.String("Permission"),
				},
				// More values...
			},
			Owner: &s3.Owner{
				DisplayName: aws.String("DisplayName"),
				ID:          aws.String("ID"),
			},
		},
		ContentMD5:       aws.String("ContentMD5"),
		GrantFullControl: aws.String("GrantFullControl"),
		GrantRead:        aws.String("GrantRead"),
		GrantReadACP:     aws.String("GrantReadACP"),
		GrantWrite:       aws.String("GrantWrite"),
		GrantWriteACP:    aws.String("GrantWriteACP"),
		RequestPayer:     aws.String("RequestPayer"),
	}
	resp, err := svc.PutObjectACL(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleS3_RestoreObject() {
	svc := s3.New(nil)

	params := &s3.RestoreObjectInput{
		Bucket:       aws.String("BucketName"), // Required
		Key:          aws.String("ObjectKey"),  // Required
		RequestPayer: aws.String("RequestPayer"),
		RestoreRequest: &s3.RestoreRequest{
			Days: aws.Long(1), // Required
		},
		VersionID: aws.String("ObjectVersionId"),
	}
	resp, err := svc.RestoreObject(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleS3_UploadPart() {
	svc := s3.New(nil)

	params := &s3.UploadPartInput{
		Bucket:               aws.String("BucketName"),        // Required
		Key:                  aws.String("ObjectKey"),         // Required
		PartNumber:           aws.Long(1),                     // Required
		UploadID:             aws.String("MultipartUploadId"), // Required
		Body:                 bytes.NewReader([]byte("PAYLOAD")),
		ContentLength:        aws.Long(1),
		ContentMD5:           aws.String("ContentMD5"),
		RequestPayer:         aws.String("RequestPayer"),
		SSECustomerAlgorithm: aws.String("SSECustomerAlgorithm"),
		SSECustomerKey:       aws.String("SSECustomerKey"),
		SSECustomerKeyMD5:    aws.String("SSECustomerKeyMD5"),
	}
	resp, err := svc.UploadPart(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleS3_UploadPartCopy() {
	svc := s3.New(nil)

	params := &s3.UploadPartCopyInput{
		Bucket:                         aws.String("BucketName"),        // Required
		CopySource:                     aws.String("CopySource"),        // Required
		Key:                            aws.String("ObjectKey"),         // Required
		PartNumber:                     aws.Long(1),                     // Required
		UploadID:                       aws.String("MultipartUploadId"), // Required
		CopySourceIfMatch:              aws.String("CopySourceIfMatch"),
		CopySourceIfModifiedSince:      aws.Time(time.Now()),
		CopySourceIfNoneMatch:          aws.String("CopySourceIfNoneMatch"),
		CopySourceIfUnmodifiedSince:    aws.Time(time.Now()),
		CopySourceRange:                aws.String("CopySourceRange"),
		CopySourceSSECustomerAlgorithm: aws.String("CopySourceSSECustomerAlgorithm"),
		CopySourceSSECustomerKey:       aws.String("CopySourceSSECustomerKey"),
		CopySourceSSECustomerKeyMD5:    aws.String("CopySourceSSECustomerKeyMD5"),
		RequestPayer:                   aws.String("RequestPayer"),
		SSECustomerAlgorithm:           aws.String("SSECustomerAlgorithm"),
		SSECustomerKey:                 aws.String("SSECustomerKey"),
		SSECustomerKeyMD5:              aws.String("SSECustomerKeyMD5"),
	}
	resp, err := svc.UploadPartCopy(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}
