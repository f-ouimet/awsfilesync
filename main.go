package main

import (
  "flag"
  "fmt"
  "os"
)

func main(){
  operation := flag.String("operation", "", "Specify 'upload' or 'download'")
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
  case "upload":
  case "download":
  }



}
