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


func main(){
  region := "us-east-2"

  sess, err := session.NewSession(&aws.Config{
    Region: aws.String(region),
  })
  if err != nil {
    fmt.Println("Error creating session:", err)
    return
  }
  client := s3.New(sess)

  bucket := "ouimetfelixfilestorage"
  filePath := "test.txt"

  _, err = client.HeadBucket(&s3.HeadBucketInput{
    Bucket: aws.String(bucket),
  })
  if err != nil{
    fmt.Println("Failed to access bucket:", err)
  }

  fmt.Println("Bucket Reached!")

  file, err := os.Open(filePath)
  if err != nil{
    fmt.Fprintln(os.Stderr, "Error opening file:", err)
    return
  }
  defer file.Close()

  key := "filetest/test.txt"
  var buf bytes.Buffer
  if _, err := io.Copy(&buf, file); err != nil{
    fmt.Fprintln(os.Stderr, "Error reading file:", err)
    return
  }

  _, err = client.PutObject(&s3.PutObjectInput{
    Bucket: aws.String(bucket),
    Key:    aws.String(key),
    Body:   bytes.NewReader(buf.Bytes()),
  })
  if err != nil{
    fmt.Println("Error uploading file:", err)
    return
  }

  fmt.Println("Upload successful!")

}
