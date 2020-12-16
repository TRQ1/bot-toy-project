package tools

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func DownloadS3(fileName string) string {
	ak := os.Getenv("AwsAccessKey")
	sk := os.Getenv("AwsSecretKey")
	sb := os.Getenv("AwsS3BucketName")
	e := os.Getenv("Environment")

	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String("ap-northeast-2"),
		Credentials: credentials.NewStaticCredentials(ak, sk, ""),
	})

	svc := s3.New(sess, &aws.Config{
		Region:                         aws.String(endpoints.ApNortheast2RegionID),
		DisableRestProtocolURICleaning: aws.Bool(true),
	})

	input := &s3.GetObjectInput{
		Bucket: aws.String(sb),
		Key:    aws.String(fileName),
	}

	_, err := svc.GetObject(input)

	if err != nil {
		panic(fmt.Sprintf("failed to get object for bucket, %s, %v", sb, err))
	}

	// Download from s3
	downloader := s3manager.NewDownloader(sess)

	filePath := ""
	if e == "local" {
		filePath = fileName
	} else {
		filePath = "/tmp/" + fileName
	}

	f, err := os.Create(filePath)

	if err != nil {
		panic(fmt.Errorf("failed to create file %q, %v", fileName, err))
	}

	n, err := downloader.Download(f, input)

	if err != nil {
		panic(fmt.Errorf("failed to download file %q, %v", fileName, err))
	} else {
		fmt.Printf("file downloaded, %d bytes\n", n)
	}

	return fileName
}
