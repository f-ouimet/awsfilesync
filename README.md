# awsfilesync
AWS File Sync

Allows user to upload files to a configured AWS S3 bucket (with default options)

#Building code
go build -o <pathtofolder/execname>

#Usage
./pathtoexec/execname -operation=<up|down> -file=<localPath> -bucket=<yourawsbucketname> -key=<bucketpath>
