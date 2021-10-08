# kinesis-get-shardid

A Golang tool that get a shard ID to be assigned in AWS Kinesis Data Stream from a partition key.

## Usage

```
kinesis-get-shardid [options...]

Options:
-stream string       (Required) Kinesis stream name
-region string       Region for Kinesis stream
                     By default "ap-northeast-1"
-key string            (Required) Partition key
-version             Prints out build version information
-verbose             Verbose option
-h                   help message
```
