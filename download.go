package main

import(
  
  "fmt"
  "io"
  "os"
  "path/filepath"
  "strings"

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

func downloadFolder(bucket, prefix, destination string) error{
  sess, err := session.NewSession(&aws.Config{
    Region: aws.String("us-east-2"),
  })
  if err != nil{
    return fmt.Errorf("failed to create session: %v", err)
  }

  svc := s3.New(sess)
  listOutput, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{
    Bucket: aws.String(bucket),
    Prefix: aws.String(prefix),
  })
  if err != nil{
    return fmt.Errorf("failed to list objects: %v", err)
  }
  for _, object := range listOutput.Contents{
    relativePath := strings.TrimPrefix(*object.Key, prefix)
    localPath := filepath.Join(destination, relativePath)

    err := os.MkdirAll(filepath.Dir(localPath),0755)
    if err != nil{
      return fmt.Errorf("Failed to create local directory for %s: %v", localPath, err)
    }

    fmt.Printf("Downloading %s to %s\n", *object.Key, localPath)
    err = downloadFile(bucket, *object.Key, localPath)
    if err != nil{
      return fmt.Errorf("failed to download %s : %v\n", *object.Key, err)
    }
  }
  return nil
}
