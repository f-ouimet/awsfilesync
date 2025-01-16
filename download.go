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


func downloadFile(bucket, key, fileName string) error{
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

  output, err := client.GetObject(&s3.GetObjectInput{
    Bucket: aws.String(bucket),
    Key:    aws.String(key),
  })
  if err != nil{
    return fmt.Errorf("Error fetching object:", err)
    
  }

  defer output.Body.Close()

  localFile, err := os.Create(fileName)
  if err != nil{
   return fmt.Errorf("Error creating local file:", err)
  
  }

  _, err = io.Copy(localFile, output.Body)
  if err != nil{
   return fmt.Errorf("Error copying data to file:", err)
  
  }
  fmt.Printf("Successfully downloaded file from bucket")
  return nil
}
