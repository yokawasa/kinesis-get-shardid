package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"math/big"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

var buildVersion string

func usage() {
	fmt.Println(usageText)
	os.Exit(0)
}

var usageText = `kinesis-get-shardid [options...]

Options:
-stream string       (Required) Kinesis stream name
-region string       Region for Kinesis stream
                     By default "ap-northeast-1"
-key string            (Required) Partition key
-version             Prints out build version information
-verbose             Verbose option
-h                   help message
`

func toBigInt(s string) *big.Int {
	n := big.NewInt(0)
	n.SetString(s, 10)
	return n
}

func includedInShardHashKeyRange(hashKey *big.Int, shard *kinesis.Shard) bool {
	return hashKey.Cmp(toBigInt(*shard.HashKeyRange.StartingHashKey)) >= 0 &&
		hashKey.Cmp(toBigInt(*shard.HashKeyRange.EndingHashKey)) <= 0
}

func getPartitionHashKey(key string) *big.Int {
	sum := md5.Sum([]byte(key))
	return big.NewInt(0).SetBytes(sum[:])
}

func getKinesisSession(region string) *kinesis.Kinesis {
	var sess *session.Session
	sess = session.New(&aws.Config{Region: aws.String(region)})
	return kinesis.New(sess)
}

func main() {

	var (
		streamName   string
		region       string
		partitionKey string
		version      bool
		verbose      bool
	)

	flag.StringVar(&streamName, "stream", "", "(Required) Kinesis stream name")
	flag.StringVar(&region, "region", "ap-northeast-1", "Region for Kinesis stream")
	flag.StringVar(&partitionKey, "key", "", "(Required) Partition Key")
	flag.BoolVar(&version, "version", false, "Build version")
	flag.BoolVar(&verbose, "verbose", false, "Verbose option")
	flag.Usage = usage
	flag.Parse()

	if version {
		fmt.Printf("version: %s\n", buildVersion)
		os.Exit(0)
	}

	if streamName == "" || partitionKey == "" {
		fmt.Println("[ERROR] Invalid Command Options! Minimum required options are \"-stream\", \"-k\" and \"-m\"")
		usage()
	}

	kc := getKinesisSession(region)
	var listShardsReq = kinesis.ListShardsInput{
		StreamName: &streamName,
	}
	listShardsResp, err := kc.ListShards(&listShardsReq)
	if err != nil {
		fmt.Printf("[ERROR] listing shards: %v", err)
	}

	hashKey := getPartitionHashKey(partitionKey)
	fmt.Printf("partion key hash %d\n", hashKey)
	//fmt.Printf("listShardsResp %v\n", listShardsResp)
	for _, shard := range listShardsResp.Shards {
		// Pick only the shard have a parent, that is child shard formed by either merge or split
		if shard.ParentShardId != nil {
			if includedInShardHashKeyRange(hashKey, shard) {
				fmt.Printf("PartionKey %s -> ShardId %s\n", partitionKey, *shard.ShardId)
				break
			}
		}
	}
}
