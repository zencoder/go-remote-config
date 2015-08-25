package remoteconfig

import (
	"errors"
	"net/url"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	ErrS3SigningInvalidS3URLScheme error = errors.New("S3 URL does have s3:// scheme.")
)

func BuildSignedS3URL(s3URL string, s3Region AWSRegion, expiry uint, endpoint string) (string, error) {
	pURL, err := url.Parse(s3URL)
	if err != nil {
		return "", err
	}
	if pURL.Scheme != "s3" {
		return "", errors.New("S3 URL does not start with the s3:// scheme")
	}

	if err = s3Region.Validate(); err != nil {
		return "", err
	}

	bucket := pURL.Host
	key := strings.TrimPrefix(pURL.Path, "/")

	return generateSignedS3URL(s3Region, bucket, key, expiry, endpoint)
}

func generateSignedS3URL(region AWSRegion, bucket string, key string, expiry uint, endpoint string) (string, error) {
	s3ForcePathStyle := false
	if endpoint != "" {
		s3ForcePathStyle = true
	}

	// We want to use the default credentials chain so that it will attempt Env & Instance role creds
	svc := s3.New(&aws.Config{
		Region:           aws.String(string(region)),
		Endpoint:         aws.String(endpoint),
		S3ForcePathStyle: aws.Bool(s3ForcePathStyle),
	})

	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})

	url, err := req.Presign(time.Duration(expiry) * time.Second)
	if err != nil {
		return "", err
	}
	return url, nil
}
