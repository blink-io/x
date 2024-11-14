package registrar

import (
	"context"

	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/kmsg"
)

type Client interface {
	LeaveGroup()
	LeaveGroupContext(ctx context.Context) error
	GroupMetadata() (string, int32)
	ForceRebalance()
	UncommittedOffsets() map[string]map[int32]kgo.EpochOffset
	MarkedOffsets() map[string]map[int32]kgo.EpochOffset
	CommittedOffsets() map[string]map[int32]kgo.EpochOffset
	CommitRecords(ctx context.Context, rs ...*kgo.Record) error
	MarkCommitRecords(rs ...*kgo.Record)
	MarkCommitOffsets(unmarked map[string]map[int32]kgo.EpochOffset)
	CommitUncommittedOffsets(ctx context.Context) error
	CommitMarkedOffsets(ctx context.Context) error
	CommitOffsetsSync(
		ctx context.Context,
		uncommitted map[string]map[int32]kgo.EpochOffset,
		onDone func(*kgo.Client, *kmsg.OffsetCommitRequest, *kmsg.OffsetCommitResponse, error),
	)
	CommitOffsets(
		ctx context.Context,
		uncommitted map[string]map[int32]kgo.EpochOffset,
		onDone func(*kgo.Client, *kmsg.OffsetCommitRequest, *kmsg.OffsetCommitResponse, error),
	)
	BeginTransaction() error
	EndAndBeginTransaction(
		ctx context.Context,
		how kgo.EndBeginTxnHow,
		commit kgo.TransactionEndTry,
		onEnd func(context.Context, error) error,
	) (rerr error)
	AbortBufferedRecords(ctx context.Context) error
	UnsafeAbortBufferedRecords()
	EndTransaction(ctx context.Context, commit kgo.TransactionEndTry) error
	BufferedProduceRecords() int64
	BufferedProduceBytes() int64
	ProduceSync(ctx context.Context, rs ...*kgo.Record) kgo.ProduceResults
	TryProduce(
		ctx context.Context,
		r *kgo.Record,
		promise func(*kgo.Record, error),
	)
	Produce(
		ctx context.Context,
		r *kgo.Record,
		promise func(*kgo.Record, error),
	)
	ProducerID(ctx context.Context) (int64, int16, error)
	Flush(ctx context.Context) error
	OptValue(opt any) any
	OptValues(opt any) []any
	Opts() []kgo.Opt
	Ping(ctx context.Context) error
	PurgeTopicsFromClient(topics ...string)
	PurgeTopicsFromProducing(topics ...string)
	PurgeTopicsFromConsuming(topics ...string)
	CloseAllowingRebalance()
	Close()
	Request(ctx context.Context, req kmsg.Request) (kmsg.Response, error)
	RequestSharded(ctx context.Context, req kmsg.Request) []kgo.ResponseShard
	Broker(id int) *kgo.Broker
	DiscoveredBrokers() []*kgo.Broker
	SeedBrokers() []*kgo.Broker
	UpdateSeedBrokers(addrs ...string) error
	BufferedFetchRecords() int64
	BufferedFetchBytes() int64
	PollFetches(ctx context.Context) kgo.Fetches
	PollRecords(ctx context.Context, maxPollRecords int) kgo.Fetches
	AllowRebalance()
	UpdateFetchMaxBytes(maxBytes, maxPartBytes int32)
	PauseFetchTopics(topics ...string) []string
	PauseFetchPartitions(topicPartitions map[string][]int32) map[string][]int32
	ResumeFetchTopics(topics ...string)
	ResumeFetchPartitions(topicPartitions map[string][]int32)
	SetOffsets(setOffsets map[string]map[int32]kgo.EpochOffset)
	AddConsumeTopics(topics ...string)
	GetConsumeTopics() []string
	AddConsumePartitions(partitions map[string]map[int32]kgo.Offset)
	RemoveConsumePartitions(partitions map[string][]int32)
	ForceMetadataRefresh()
	PartitionLeader(topic string, partition int32) (leader, leaderEpoch int32, err error)
}
