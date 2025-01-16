package main

import(
  "bytes"
  "fmt"
  "io"
  "os"

  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/s3"
)
//type AWSService struct{
  //S3Client *s3.S3Client
//} 


func uploadFile(bucket, key, filePath string) error{
  region := "us-east-2"

  sess, err := session.NewSession(&aws.Config{
    Region: aws.String(region),
  })
  if err != nil {
    return fmt.Errorf("Error creating session:", err)
    
  }
  client := s3.New(sess)

  _, err = client.HeadBucket(&s3.HeadBucketInput{
    Bucket: aws.String(bucket),
  })
  if err != nil{
    return fmt.Errorf("Failed to access bucket:", err)
  }

  fmt.Println("Bucket Reached!")

  file, err := os.Open(filePath)
  if err != nil{
    return fmt.Errorf("Error opening file:", err)
    
  }
  defer file.Close()

  
  var buf bytes.Buffer
  if _, err := io.Copy(&buf, file); err != nil{
    return fmt.Errorf("Error reading file:", err)
    
  }

  _, err = client.PutObject(&s3.PutObjectInput{
    Bucket: aws.String(bucket),
    Key:    aws.String(key),
    Body:   bytes.NewReader(buf.Bytes()),
  })
  if err != nil{
    return fmt.Errorf("Error uploading file:", err)
    
  }

  fmt.Println("Upload successful!")
  return nil

}
