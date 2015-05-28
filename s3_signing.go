package remoteconfig

import (
	"errors"
	"net/url"
	"strings"
	"time"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/service/s3"
)

var (
	ErrS3SigningInvalidS3URLScheme error = errors.New("S3 URL does have s3:// scheme.")
)

func BuildSignedS3URL(s3URL string, s3Region AWSRegion, expiry uint) (string, error) {
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

	return generateSignedS3URL(s3Region, bucket, key, expiry)
}

func generateSignedS3URL(region AWSRegion, bucket string, key string, expiry uint) (string, error) {
	svc := s3.New(&aws.Config{
		Credentials: aws.DefaultCreds(),
		Region:      string(region),
	})

	var req *aws.Request
	req, _ = svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})

	url, err := req.Presign(time.Duration(expiry) * time.Second)
	if err != nil {
		return "", err
	}
	return url, nil
}
