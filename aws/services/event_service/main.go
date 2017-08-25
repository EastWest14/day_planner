package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	awsSession, err := session.NewSession()
	if err != nil {
		panic(err.Error())
	}
	svc := s3.New(awsSession)

	inputparams := &s3.ListObjectsInput{
		Bucket:  aws.String("kops-state-12345"),
		MaxKeys: aws.Int64(10),
	}
	pageNum := 0
	svc.ListObjectsPages(inputparams, func(page *s3.ListObjectsOutput, lastPage bool) bool {
		pageNum++
		for _, value := range page.Contents {
			fmt.Println(*value.Key)
		}
		return pageNum < 3
	})
}
