package remoteconfig

import "errors"

type AWSRegion string

const (
	AWS_REGION_US_EAST_1      AWSRegion = "us-east-1"
	AWS_REGION_US_WEST_1      AWSRegion = "us-west-1"
	AWS_REGION_US_WEST_2      AWSRegion = "us-west-2"
	AWS_REGION_EU_WEST_1      AWSRegion = "eu-west-1"
	AWS_REGION_EU_CENTRAL_1   AWSRegion = "eu-central-1"
	AWS_REGION_AP_SOUTHEAST_1 AWSRegion = "ap-southeast-1"
	AWS_REGION_AP_SOUTHEAST_2 AWSRegion = "ap-southeast-2"
	AWS_REGION_AP_NORTHEAST_1 AWSRegion = "ap-northeast-1"
	AWS_REGION_SA_EAST_1      AWSRegion = "sa-east-1"
)

var AWSRegions = []AWSRegion{
	AWS_REGION_US_EAST_1,
	AWS_REGION_US_WEST_1,
	AWS_REGION_US_WEST_2,
	AWS_REGION_EU_WEST_1,
	AWS_REGION_EU_CENTRAL_1,
	AWS_REGION_AP_SOUTHEAST_1,
	AWS_REGION_AP_SOUTHEAST_2,
	AWS_REGION_AP_NORTHEAST_1,
	AWS_REGION_SA_EAST_1,
}

var (
	ErrAWSRegionEmptyString = errors.New("Region cannot be empty")
	ErrAWSRegionInvalid     = errors.New("Region is invalid")
)

func (r *AWSRegion) UnmarshalText(data []byte) error {
	rString := string(data[:])
	*r = (AWSRegion)(rString)
	return r.Validate()
}

func (r AWSRegion) Validate() error {
	if r == "" {
		return ErrAWSRegionEmptyString
	}

	if r != AWS_REGION_US_EAST_1 && r != AWS_REGION_US_WEST_1 && r != AWS_REGION_US_WEST_2 &&
		r != AWS_REGION_EU_WEST_1 && r != AWS_REGION_EU_CENTRAL_1 &&
		r != AWS_REGION_AP_SOUTHEAST_1 && r != AWS_REGION_AP_SOUTHEAST_2 && r != AWS_REGION_AP_NORTHEAST_1 &&
		r != AWS_REGION_SA_EAST_1 {
		return ErrAWSRegionInvalid
	}

	return nil
}
