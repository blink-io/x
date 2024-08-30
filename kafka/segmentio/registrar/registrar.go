package registrar

import (
	"context"

	kafkago "github.com/segmentio/kafka-go"
)

type RegisterFunc func(context.Context, ServiceRegistrar) error

type WithRegistrar interface {
	GoKafkaRegistrar(context.Context) RegisterFunc
}

type ServiceRegistrar interface {
	RawProduce(ctx context.Context, req *kafkago.RawProduceRequest) (*kafkago.ProduceResponse, error)
	CreateACLs(ctx context.Context, req *kafkago.CreateACLsRequest) (*kafkago.CreateACLsResponse, error)
	DescribeGroups(
		ctx context.Context,
		req *kafkago.DescribeGroupsRequest,
	) (*kafkago.DescribeGroupsResponse, error)
	Produce(ctx context.Context, req *kafkago.ProduceRequest) (*kafkago.ProduceResponse, error)
	AlterConfigs(
		ctx context.Context,
		req *kafkago.AlterConfigsRequest,
	) (*kafkago.AlterConfigsResponse, error)
	AlterUserScramCredentials(
		ctx context.Context,
		req *kafkago.AlterUserScramCredentialsRequest,
	) (*kafkago.AlterUserScramCredentialsResponse, error)
	AlterClientQuotas(
		ctx context.Context,
		req *kafkago.AlterClientQuotasRequest,
	) (*kafkago.AlterClientQuotasResponse, error)
	DeleteACLs(ctx context.Context,
		req *kafkago.DeleteACLsRequest,
	) (*kafkago.DeleteACLsResponse, error)
	ListGroups(
		ctx context.Context,
		req *kafkago.ListGroupsRequest,
	) (*kafkago.ListGroupsResponse, error)
	DescribeACLs(ctx context.Context,
		req *kafkago.DescribeACLsRequest,
	) (*kafkago.DescribeACLsResponse, error)
	FindCoordinator(ctx context.Context,
		req *kafkago.FindCoordinatorRequest,
	) (*kafkago.FindCoordinatorResponse, error)
	JoinGroup(ctx context.Context,
		req *kafkago.JoinGroupRequest,
	) (*kafkago.JoinGroupResponse, error)
	ElectLeaders(
		ctx context.Context,
		req *kafkago.ElectLeadersRequest,
	) (*kafkago.ElectLeadersResponse, error)
	DeleteGroups(
		ctx context.Context,
		req *kafkago.DeleteGroupsRequest,
	) (*kafkago.DeleteGroupsResponse, error)
	AlterPartitionReassignments(
		ctx context.Context,
		req *kafkago.AlterPartitionReassignmentsRequest,
	) (*kafkago.AlterPartitionReassignmentsResponse, error)
	OffsetCommit(ctx context.Context, req *kafkago.OffsetCommitRequest) (*kafkago.OffsetCommitResponse, error)
	Metadata(ctx context.Context, req *kafkago.MetadataRequest) (*kafkago.MetadataResponse, error)
	Fetch(ctx context.Context, req *kafkago.FetchRequest) (*kafkago.FetchResponse, error)
	IncrementalAlterConfigs(
		ctx context.Context,
		req *kafkago.IncrementalAlterConfigsRequest,
	) (*kafkago.IncrementalAlterConfigsResponse, error)
	LeaveGroup(
		ctx context.Context,
		req *kafkago.LeaveGroupRequest,
	) (*kafkago.LeaveGroupResponse, error)
	OffsetDelete(
		ctx context.Context,
		req *kafkago.OffsetDeleteRequest,
	) (*kafkago.OffsetDeleteResponse, error)
	DescribeConfigs(
		ctx context.Context,
		req *kafkago.DescribeConfigsRequest,
	) (*kafkago.DescribeConfigsResponse, error)
	DescribeUserScramCredentials(
		ctx context.Context,
		req *kafkago.DescribeUserScramCredentialsRequest,
	) (*kafkago.DescribeUserScramCredentialsResponse, error)
	ApiVersions(
		ctx context.Context,
		req *kafkago.ApiVersionsRequest,
	) (*kafkago.ApiVersionsResponse, error)
	CreateTopics(
		ctx context.Context,
		req *kafkago.CreateTopicsRequest,
	) (*kafkago.CreateTopicsResponse, error)
	EndTxn(ctx context.Context, req *kafkago.EndTxnRequest) (*kafkago.EndTxnResponse, error)
	ConsumerOffsets(ctx context.Context, tg kafkago.TopicAndGroup) (map[int]int64, error)
	AddPartitionsToTxn(
		ctx context.Context,
		req *kafkago.AddPartitionsToTxnRequest,
	) (*kafkago.AddPartitionsToTxnResponse, error)
	ListOffsets(
		ctx context.Context,
		req *kafkago.ListOffsetsRequest,
	) (*kafkago.ListOffsetsResponse, error)
	DescribeClientQuotas(
		ctx context.Context,
		req *kafkago.DescribeClientQuotasRequest,
	) (*kafkago.DescribeClientQuotasResponse, error)
	SyncGroup(
		ctx context.Context,
		req *kafkago.SyncGroupRequest,
	) (*kafkago.SyncGroupResponse, error)
	TxnOffsetCommit(
		ctx context.Context,
		req *kafkago.TxnOffsetCommitRequest,
	) (*kafkago.TxnOffsetCommitResponse, error)
	AddOffsetsToTxn(
		ctx context.Context,
		req *kafkago.AddOffsetsToTxnRequest,
	) (*kafkago.AddOffsetsToTxnResponse, error)
	DeleteTopics(
		ctx context.Context,
		req *kafkago.DeleteTopicsRequest,
	) (*kafkago.DeleteTopicsResponse, error)
	InitProducerID(
		ctx context.Context,
		req *kafkago.InitProducerIDRequest,
	) (*kafkago.InitProducerIDResponse, error)
	OffsetFetch(ctx context.Context,
		req *kafkago.OffsetFetchRequest,
	) (*kafkago.OffsetFetchResponse, error)
	ListPartitionReassignments(
		ctx context.Context,
		req *kafkago.ListPartitionReassignmentsRequest,
	) (*kafkago.ListPartitionReassignmentsResponse, error)
	Heartbeat(
		ctx context.Context,
		req *kafkago.HeartbeatRequest,
	) (*kafkago.HeartbeatResponse, error)
	CreatePartitions(
		ctx context.Context,
		req *kafkago.CreatePartitionsRequest,
	) (*kafkago.CreatePartitionsResponse, error)
}

var _ ServiceRegistrar = (*kafkago.Client)(nil)

type Func[S any] func(ServiceRegistrar, S)

type FuncWithErr[S any] func(ServiceRegistrar, S) error

type CtxFunc[S any] func(context.Context, ServiceRegistrar, S)

type CtxFuncWithErr[S any] func(context.Context, ServiceRegistrar, S) error

type Registrar interface {
	RegisterToGoKafka(context.Context, ServiceRegistrar) error
}

var _ Registrar = (*registrar[any])(nil)

func New[S any](s S, f Func[S]) Registrar {
	cf := func(ctx context.Context, r ServiceRegistrar, s S) error {
		f(r, s)
		return nil
	}
	return NewCtxWithErr(s, cf)
}

func NewWithErr[S any](s S, f FuncWithErr[S]) Registrar {
	cf := func(ctx context.Context, r ServiceRegistrar, s S) error {
		return f(r, s)
	}
	return NewCtxWithErr(s, cf)
}

func NewCtx[S any](s S, f CtxFunc[S]) Registrar {
	cf := func(ctx context.Context, r ServiceRegistrar, s S) error {
		f(ctx, r, s)
		return nil
	}
	return NewCtxWithErr(s, cf)
}

func NewCtxWithErr[S any](s S, f CtxFuncWithErr[S]) Registrar {
	h := &registrar[S]{
		s: s,
		f: f,
	}
	return h
}

type registrar[S any] struct {
	s S
	f CtxFuncWithErr[S]
}

func (h *registrar[S]) RegisterToGoKafka(ctx context.Context, r ServiceRegistrar) error {
	return h.f(ctx, r, h.s)
}
