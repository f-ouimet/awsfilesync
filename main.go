package main

import (
  "flag"
  "fmt"
  "os"
  //"io"
  //"path/filepath"
  "strings"
)

func isDirectory(path string) bool{
  info, err := os.Stat(path)
  if err != nil{
    return false
  }
  return info.IsDir()
}

func main(){
  operation := flag.String("operation", "", "Specify upload with 'up' or 'download with 'down'")
  filePath := flag.String("file", "", "Path to local file (upload) or dest (bucket download)")
  bucket := flag.String("bucket", "", "S3 bucket name")
  key := flag.String("key", "", "S3 object key (path and name)")

  flag.Parse()

  if *operation =="" || *filePath == "" || *bucket == "" || *key == ""{
    fmt.Println("ERR: One or more missing args!")
    fmt.Println("Usage: -operation=<up|down> -file=<file-path> -bucket=<bucket-name> -key=<object-key>")
    os.Exit(1)
  }

  switch *operation{
  case "up":
    if isDirectory(*filePath){
      fmt.Printf("Uploading folder %s to bucket %s with prefix %s...\n", *filePath, *bucket, *key)
      err := uploadFolder(*bucket, *filePath, *key)
      if err != nil{
        fmt.Println("Upload folder failed:", err)
        os.Exit(1)
      }
      fmt.Println("Upload folder successful!")
    } else {
        fmt.Printf("Uploading file %s to bucket %s with key %s", *filePath, *bucket, *key)
        err := uploadFile(*bucket, *key, *filePath)
        if err != nil{
          fmt.Println("Upload failed:", err)
          os.Exit(1)
        }
    
    fmt.Println("Upload file successful")
    }
  case "down":
    if strings.HasSuffix(*key, "/"){
      fmt.Printf("Downloading folder from bucket %s with prefix %s to %s...\n", *bucket, *key, *filePath)
      err := downloadFolder(*bucket, *key, *filePath)
      if err != nil{
        fmt.Println("Download folder failed:", err)
        os.Exit(1)
      }
      fmt.Println("Download folder successful!")
    } else {
        fmt.Printf("Downloading file from bucket %s with key %s to %s ... \n", *bucket, *key, *filePath)
        err := downloadFile(*bucket, *key, *filePath)
        if err != nil{
          fmt.Println("Download failed:", err)
          os.Exit(1)
        }
      fmt.Println("Download successful")
      }
  default:
    fmt.Println("Invalid operation, Use 'up' or 'down' .")
    os.Exit(1)
  }



}
