package config

import (
	"time"

	"github.com/sionreview/sion/common/logger"

	protocol "github.com/sionreview/sion/common/types"
)

// FLAG_FIXED_INTERVAL_WARMUP warms up lambda with fixed interval regardless workload.
const FLAG_FIXED_INTERVAL_WARMUP = 0x0001

// FLAG_CLIENTSIDE_FIRSTD_OPTIMIZATION enables D+P optimization on client side.
const FLAG_CLIENTSIDE_FIRSTD_OPTIMIZATION = 0x0002

// FLAG_ENABLE_LOCAL_CACHE enables local cache.
const FLAG_ENABLE_LOCAL_CACHE = 0x0004

// const LambdaFeatures = protocol.FLAG_ENABLE_WARMUP | protocol.FLAG_DISABLE_WAIT_FOR_COS | protocol.FLAG_ENABLE_PERSISTENT

const LambdaFeatures = protocol.FLAG_ENABLE_WARMUP | protocol.FLAG_DISABLE_WAIT_FOR_COS

//const LambdaFeatures = protocol.FLAG_DISABLE_WAIT_FOR_COS

const ProxyFeatures = FLAG_CLIENTSIDE_FIRSTD_OPTIMIZATION | FLAG_ENABLE_LOCAL_CACHE
const LambdaLogLevel = logger.LOG_LEVEL_ALL // Set to logger.LOG_LEVEL_ALL to keep Lambda log level aligned with proxy log level.

// LambdaPrefix Prefix of Lambda function, overridable with command line parameter -lambda-prefix.
const LambdaPrefix = "CacheNodeA" //

// AWSRegion Region of AWS services.
const AWSRegion = "us-west-2"

// LambdaMaxDeployments Number of Lambda function deployments available.
const LambdaMaxDeployments = 400

// Mode of cluster.
const StaticCluster = "static"
const WindowCluster = "window"
const Cluster = WindowCluster

// Size of a slice if the cluster implementation support. Client library use this value to initialize chunk placements.
const SliceSize = 100

// NumLambdaClusters Number of Lambda function deployments initiated on launching.
// For window cluster, this must be at least D+P
const NumLambdaClusters = 12

// LambdaStoreName Obsoleted. Name of Lambda function for replica version.
const LambdaStoreName = "LambdaStore"

// InstanceWarmTimout Interval to warmup Lambda functions.
const InstanceWarmTimeout = 1 * time.Minute //ping instance every 1 minute

// Instance degrade warmup interval
const InstanceDegradeWarmTimeout = 5 * time.Minute //ping instance every 5 minutes

// InstanceCapacity Capacity of deployed Lambda functions.
// TODO: Detectable on invocation. Can be specified by option -funcap for now.
const DefaultInstanceCapacity = 2048 * 1000000 //1536 * 1000000 // 1GB

// InstanceOverhead Memory reserved for running program on Lambda functions.
const InstanceOverhead = 100 * 1000000 // MB

// Threshold Scaling out avg instance size threshold
const Threshold = 0.9 // Don't set beyond 0.9

// Maximum chunk per instance
const ChunkThreshold = 125000 // Fraction, ChunkThreshold = InstanceCapacity / 100K * Threshold

// ServerPublicIp Public IP of proxy, leave empty if running Lambda functions in VPC.
const ServerPublicIp = "" // Leave it empty if Lambda VPC is enabled.

// RecoverRate Empirical S3 download rate for specified InstanceCapacity.
// 40MB for 512, 1024, 1536MB instance, 70MB for 3008MB instance.
const RecoverRate = 40 * 1000000 // Not actually used.

// BackupsPerInstance  Number of backup instances used for parallel recovery.
const BackupsPerInstance = 10 // (InstanceCapacity - InstanceOverhead) / RecoverRate

// Each bucket's active duration
const BucketDuration = 10 // min

// Number of buckets that warmup every InstanceWarmTimeout
const NumActiveBuckets = 6

// Number of buckets before expiring
// Buckets beyond NumActiveBuckets but within ExpireBucketsNum will get degraded warmup: InstanceDegradeWarmTimeout
const NumAvailableBuckets = 18

// Async migrate control
const ActiveReplica = 2 //min

// ProxyList Ip addresses and ports in the format "ip:port" of proxies.
// If running on one proxy, then can be left empty. For multi-proxies deployment, build static proxy list here.
// Private ip should be used if Lambda VPC is enabled.
// var ProxyList = []string{"10.0.25.0:6378", "10.0.24.50:6378", "10.0.25.43:6378", "10.0.21.60:6378"}

var ProxyList []string
