# kinesis-get-shardid

[![Upload Release Asset](https://github.com/yokawasa/kinesis-get-shardid/actions/workflows/release.yml/badge.svg)](https://github.com/yokawasa/kinesis-get-shardid/actions/workflows/release.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/yokawasa/kinesis-get-shardid)](https://goreportcard.com/report/github.com/yokawasa/kinesis-get-shardid) [![GoDoc](https://godoc.org/github.com/yokawasa/kinesis-get-shardid?status.svg)](https://godoc.org/github.com/yokawasa/kinesis-get-shardid)

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

## Download

You can download the compiled command with [downloader](https://github.com/yokawasa/kinesis-get-shardid/blob/main/downloader) like this:

```bash
# Download latest command
./downloader

# Download the command with a specified version
./downloader v0.0.2
```
Or you can download it on the fly with the following commmand:

```bash
curl -sS https://raw.githubusercontent.com/yokawasa/kinesis-get-shardid/main/downloader | bash --
```

## Execute the command

```bash
stream_name=test-kds01
region=ap-northeast-1
partitionkey=mykey001

kinesis-get-shardid -stream ${stream_name} -region ${region} -key ${partitionkey}
```

Sample output would be like this:

```
PartionKey mykey001 -> ShardId shardId-000000000004
```


## Build and Run

To build, simply run `make` like below
```
make

golint /Users/yoichika/dev/github/kinesis-get-shardid
GOOS=linux GOARCH=amd64 GO111MODULE=on go build -ldflags "-X main.buildVersion=v0.0.1" -o kinesis-get-shardid/dist/kinesis-get-shardid_linux_amd64 kinesis-get-shardid/src
GOOS=darwin GOARCH=amd64 GO111MODULE=on go build -ldflags "-X main.buildVersion=v0.0.1" -o kinesis-get-shardid/dist/kinesis-get-shardid_darwin_amd64 kinesis-get-shardid/src
GOOS=windows GOARCH=amd64 GO111MODULE=on go build -ldflags "-X main.buildVersion=v0.0.1" -o kinesis-get-shardid/dist/kinesis-get-shardid_windows_amd64 kinesis-get-shardid/src
```

Suppose you are using macOS, run the `kinesis-get-shardid_darwin` (while `kinesis-get-shardid_linux` if you are using Linux, or `kinesis-get-shardid_windows` if using Windows) like below

```bash
./dist/kinesis-get-shardid_darwin_amd64 -stream test-kds01 -key hoge
```

Finally clean built commands

```
make clean
```

## Relevant projects

- [kinsis-bulk-loader](https://github.com/yokawasa/kinesis-bulk-loader): A Golang tool that sends bulk messages in parallel to Amazon SQS
- [kinesis-consumer](https://github.com/yokawasa/kinesis-consumer): The Kinesis Consumer side can be tested with [kinesis-consumer](https://github.com/yokawasa/kinesis-consumer)

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/yokawasa/kinesis-get-shardid.
