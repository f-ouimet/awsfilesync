# awsfilesync
AWS File Sync

Allows user to upload files to a configured AWS S3 bucket\
Currently for region us-2\
Credentials need to be in the default directory ~/.aws with the default name credentials . See: https://docs.aws.amazon.com/cli/v1/userguide/cli-configure-files.html\
Mainly for my own personal use atm.

# Building code
go build -o <pathtofolder/execname>

# Usage
./pathtoexec/execname -operation=<up|down> -file=<localPath> -bucket=<yourawsbucketname> -key=<bucketpath>
