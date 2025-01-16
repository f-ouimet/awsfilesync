package main

import(
  
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
  key := "filetest/test.txt"

  _, err = client.HeadBucket(&s3.HeadBucketInput{
    Bucket: aws.String(bucket),
  })
  if err != nil{
    fmt.Println("Failed to access bucket:", err)
  }

  fmt.Println("Bucket Reached!")

  output, err := client.GetObject(&s3.GetObjectInput{
    Bucket: aws.String(bucket),
    Key:    aws.String(key),
  })
  if err != nil{
    fmt.Println("Error fetching object:", err)
    return
  }

  defer output.Body.Close()

  fileName := "testRx.txt"
  localFile, err := os.Create(fileName)
  if err != nil{
    fmt.Println("Error creating local file:", err)
    return
  }

  _, err = io.Copy(localFile, output.Body)
  if err != nil{
    fmt.Println("Error copying data to file:", err)
    return
  }
  fmt.Printf("Successfully downloaded file from bucket")
}
