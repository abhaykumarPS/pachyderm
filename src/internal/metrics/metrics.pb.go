// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.2
// source: internal/metrics/metrics.proto

package metrics

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Metrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterId           string  `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	PodId               string  `protobuf:"bytes,2,opt,name=pod_id,json=podId,proto3" json:"pod_id,omitempty"`
	Nodes               int64   `protobuf:"varint,3,opt,name=nodes,proto3" json:"nodes,omitempty"`
	Version             string  `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	Repos               int64   `protobuf:"varint,5,opt,name=repos,proto3" json:"repos,omitempty"`                                                              // Number of repos
	Commits             int64   `protobuf:"varint,6,opt,name=commits,proto3" json:"commits,omitempty"`                                                          // Number of commits -- not used
	Files               int64   `protobuf:"varint,7,opt,name=files,proto3" json:"files,omitempty"`                                                              // Number of files -- not used
	Bytes               uint64  `protobuf:"varint,8,opt,name=bytes,proto3" json:"bytes,omitempty"`                                                              // Number of bytes in all repos
	Jobs                int64   `protobuf:"varint,9,opt,name=jobs,proto3" json:"jobs,omitempty"`                                                                // Number of jobs
	Pipelines           int64   `protobuf:"varint,10,opt,name=pipelines,proto3" json:"pipelines,omitempty"`                                                     // Number of pipelines in the cluster -- not the same as DAG
	ArchivedCommits     int64   `protobuf:"varint,11,opt,name=archived_commits,json=archivedCommits,proto3" json:"archived_commits,omitempty"`                  // Number of archived commit -- not used
	CancelledCommits    int64   `protobuf:"varint,12,opt,name=cancelled_commits,json=cancelledCommits,proto3" json:"cancelled_commits,omitempty"`               // Number of cancelled commits -- not used
	ActivationCode      string  `protobuf:"bytes,13,opt,name=activation_code,json=activationCode,proto3" json:"activation_code,omitempty"`                      // Activation code
	MaxBranches         uint64  `protobuf:"varint,14,opt,name=max_branches,json=maxBranches,proto3" json:"max_branches,omitempty"`                              // Max branches in across all the repos
	PpsSpout            int64   `protobuf:"varint,15,opt,name=pps_spout,json=ppsSpout,proto3" json:"pps_spout,omitempty"`                                       // Number of spout pipelines
	PpsSpoutService     int64   `protobuf:"varint,16,opt,name=pps_spout_service,json=ppsSpoutService,proto3" json:"pps_spout_service,omitempty"`                // Number of spout services
	PpsBuild            int64   `protobuf:"varint,17,opt,name=pps_build,json=ppsBuild,proto3" json:"pps_build,omitempty"`                                       // Number of build pipelines
	CfgEgress           int64   `protobuf:"varint,18,opt,name=cfg_egress,json=cfgEgress,proto3" json:"cfg_egress,omitempty"`                                    // Number of pipelines with Egress configured
	CfgStandby          int64   `protobuf:"varint,19,opt,name=cfg_standby,json=cfgStandby,proto3" json:"cfg_standby,omitempty"`                                 // Number of pipelines with Standby congigured
	CfgS3Gateway        int64   `protobuf:"varint,20,opt,name=cfg_s3gateway,json=cfgS3gateway,proto3" json:"cfg_s3gateway,omitempty"`                           // Number of pipelines with S3 Gateway configured
	CfgServices         int64   `protobuf:"varint,21,opt,name=cfg_services,json=cfgServices,proto3" json:"cfg_services,omitempty"`                              // Number of pipelines with services configured
	CfgErrcmd           int64   `protobuf:"varint,22,opt,name=cfg_errcmd,json=cfgErrcmd,proto3" json:"cfg_errcmd,omitempty"`                                    // Number of pipelines with error cmd set
	CfgStats            int64   `protobuf:"varint,23,opt,name=cfg_stats,json=cfgStats,proto3" json:"cfg_stats,omitempty"`                                       // Number of pipelines with stats configured
	CfgTfjob            int64   `protobuf:"varint,24,opt,name=cfg_tfjob,json=cfgTfjob,proto3" json:"cfg_tfjob,omitempty"`                                       // Number of pipelines with TFJobs configured
	InputGroup          int64   `protobuf:"varint,25,opt,name=input_group,json=inputGroup,proto3" json:"input_group,omitempty"`                                 // Number of pipelines with group inputs
	InputJoin           int64   `protobuf:"varint,26,opt,name=input_join,json=inputJoin,proto3" json:"input_join,omitempty"`                                    // Number of pipelines with join inputs
	InputCross          int64   `protobuf:"varint,27,opt,name=input_cross,json=inputCross,proto3" json:"input_cross,omitempty"`                                 // Number of pipelines with cross inputs
	InputUnion          int64   `protobuf:"varint,28,opt,name=input_union,json=inputUnion,proto3" json:"input_union,omitempty"`                                 // Number of pipelines with union inputs
	InputCron           int64   `protobuf:"varint,29,opt,name=input_cron,json=inputCron,proto3" json:"input_cron,omitempty"`                                    // Number of pipelines with cron inputs
	InputGit            int64   `protobuf:"varint,30,opt,name=input_git,json=inputGit,proto3" json:"input_git,omitempty"`                                       // Number of pipelines with git inputs
	InputPfs            int64   `protobuf:"varint,31,opt,name=input_pfs,json=inputPfs,proto3" json:"input_pfs,omitempty"`                                       // Number of pfs inputs
	InputCommit         int64   `protobuf:"varint,32,opt,name=input_commit,json=inputCommit,proto3" json:"input_commit,omitempty"`                              // Number of pfs inputs with commits
	InputJoinOn         int64   `protobuf:"varint,33,opt,name=input_join_on,json=inputJoinOn,proto3" json:"input_join_on,omitempty"`                            // Number of pfs inputs with join_on
	InputOuterJoin      int64   `protobuf:"varint,34,opt,name=input_outer_join,json=inputOuterJoin,proto3" json:"input_outer_join,omitempty"`                   // Number of pipelines with outer joins
	InputLazy           int64   `protobuf:"varint,35,opt,name=input_lazy,json=inputLazy,proto3" json:"input_lazy,omitempty"`                                    // Number of pipelines with lazy set
	InputEmptyFiles     int64   `protobuf:"varint,36,opt,name=input_empty_files,json=inputEmptyFiles,proto3" json:"input_empty_files,omitempty"`                // Number of pipelines with empty files set
	InputS3             int64   `protobuf:"varint,37,opt,name=input_s3,json=inputS3,proto3" json:"input_s3,omitempty"`                                          // Number of pipelines with S3 input
	InputTrigger        int64   `protobuf:"varint,38,opt,name=input_trigger,json=inputTrigger,proto3" json:"input_trigger,omitempty"`                           // Number of pipelines with triggers set
	ResourceCpuReq      float32 `protobuf:"fixed32,39,opt,name=resource_cpu_req,json=resourceCpuReq,proto3" json:"resource_cpu_req,omitempty"`                  // Total CPU request across all pipelines
	ResourceCpuReqMax   float32 `protobuf:"fixed32,40,opt,name=resource_cpu_req_max,json=resourceCpuReqMax,proto3" json:"resource_cpu_req_max,omitempty"`       // Max CPU resource requests set
	ResourceMemReq      string  `protobuf:"bytes,41,opt,name=resource_mem_req,json=resourceMemReq,proto3" json:"resource_mem_req,omitempty"`                    // Sting of memory requests set across all pipelines
	ResourceGpuReq      int64   `protobuf:"varint,42,opt,name=resource_gpu_req,json=resourceGpuReq,proto3" json:"resource_gpu_req,omitempty"`                   // Total GPU requests across all pipelines
	ResourceGpuReqMax   int64   `protobuf:"varint,43,opt,name=resource_gpu_req_max,json=resourceGpuReqMax,proto3" json:"resource_gpu_req_max,omitempty"`        // Max GPU request across all pipelines
	ResourceDiskReq     string  `protobuf:"bytes,44,opt,name=resource_disk_req,json=resourceDiskReq,proto3" json:"resource_disk_req,omitempty"`                 // String of disk requests set across all pipelines
	ResourceCpuLimit    float32 `protobuf:"fixed32,45,opt,name=resource_cpu_limit,json=resourceCpuLimit,proto3" json:"resource_cpu_limit,omitempty"`            // Total CPU limits set across all pipelines
	ResourceCpuLimitMax float32 `protobuf:"fixed32,46,opt,name=resource_cpu_limit_max,json=resourceCpuLimitMax,proto3" json:"resource_cpu_limit_max,omitempty"` // Max CPU limit set
	ResourceMemLimit    string  `protobuf:"bytes,47,opt,name=resource_mem_limit,json=resourceMemLimit,proto3" json:"resource_mem_limit,omitempty"`              // String of memory limits set
	ResourceGpuLimit    int64   `protobuf:"varint,48,opt,name=resource_gpu_limit,json=resourceGpuLimit,proto3" json:"resource_gpu_limit,omitempty"`             // Number of pipelines with
	ResourceGpuLimitMax int64   `protobuf:"varint,49,opt,name=resource_gpu_limit_max,json=resourceGpuLimitMax,proto3" json:"resource_gpu_limit_max,omitempty"`  // Max GPU limit set
	ResourceDiskLimit   string  `protobuf:"bytes,50,opt,name=resource_disk_limit,json=resourceDiskLimit,proto3" json:"resource_disk_limit,omitempty"`           // String of disk limits set across all pipelines
	MaxParallelism      uint64  `protobuf:"varint,51,opt,name=max_parallelism,json=maxParallelism,proto3" json:"max_parallelism,omitempty"`                     // Max parallelism set
	MinParallelism      uint64  `protobuf:"varint,52,opt,name=min_parallelism,json=minParallelism,proto3" json:"min_parallelism,omitempty"`                     // Min parallelism set
	NumParallelism      uint64  `protobuf:"varint,53,opt,name=num_parallelism,json=numParallelism,proto3" json:"num_parallelism,omitempty"`                     // Number of pipelines with parallelism set
}

func (x *Metrics) Reset() {
	*x = Metrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_metrics_metrics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metrics) ProtoMessage() {}

func (x *Metrics) ProtoReflect() protoreflect.Message {
	mi := &file_internal_metrics_metrics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metrics.ProtoReflect.Descriptor instead.
func (*Metrics) Descriptor() ([]byte, []int) {
	return file_internal_metrics_metrics_proto_rawDescGZIP(), []int{0}
}

func (x *Metrics) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *Metrics) GetPodId() string {
	if x != nil {
		return x.PodId
	}
	return ""
}

func (x *Metrics) GetNodes() int64 {
	if x != nil {
		return x.Nodes
	}
	return 0
}

func (x *Metrics) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Metrics) GetRepos() int64 {
	if x != nil {
		return x.Repos
	}
	return 0
}

func (x *Metrics) GetCommits() int64 {
	if x != nil {
		return x.Commits
	}
	return 0
}

func (x *Metrics) GetFiles() int64 {
	if x != nil {
		return x.Files
	}
	return 0
}

func (x *Metrics) GetBytes() uint64 {
	if x != nil {
		return x.Bytes
	}
	return 0
}

func (x *Metrics) GetJobs() int64 {
	if x != nil {
		return x.Jobs
	}
	return 0
}

func (x *Metrics) GetPipelines() int64 {
	if x != nil {
		return x.Pipelines
	}
	return 0
}

func (x *Metrics) GetArchivedCommits() int64 {
	if x != nil {
		return x.ArchivedCommits
	}
	return 0
}

func (x *Metrics) GetCancelledCommits() int64 {
	if x != nil {
		return x.CancelledCommits
	}
	return 0
}

func (x *Metrics) GetActivationCode() string {
	if x != nil {
		return x.ActivationCode
	}
	return ""
}

func (x *Metrics) GetMaxBranches() uint64 {
	if x != nil {
		return x.MaxBranches
	}
	return 0
}

func (x *Metrics) GetPpsSpout() int64 {
	if x != nil {
		return x.PpsSpout
	}
	return 0
}

func (x *Metrics) GetPpsSpoutService() int64 {
	if x != nil {
		return x.PpsSpoutService
	}
	return 0
}

func (x *Metrics) GetPpsBuild() int64 {
	if x != nil {
		return x.PpsBuild
	}
	return 0
}

func (x *Metrics) GetCfgEgress() int64 {
	if x != nil {
		return x.CfgEgress
	}
	return 0
}

func (x *Metrics) GetCfgStandby() int64 {
	if x != nil {
		return x.CfgStandby
	}
	return 0
}

func (x *Metrics) GetCfgS3Gateway() int64 {
	if x != nil {
		return x.CfgS3Gateway
	}
	return 0
}

func (x *Metrics) GetCfgServices() int64 {
	if x != nil {
		return x.CfgServices
	}
	return 0
}

func (x *Metrics) GetCfgErrcmd() int64 {
	if x != nil {
		return x.CfgErrcmd
	}
	return 0
}

func (x *Metrics) GetCfgStats() int64 {
	if x != nil {
		return x.CfgStats
	}
	return 0
}

func (x *Metrics) GetCfgTfjob() int64 {
	if x != nil {
		return x.CfgTfjob
	}
	return 0
}

func (x *Metrics) GetInputGroup() int64 {
	if x != nil {
		return x.InputGroup
	}
	return 0
}

func (x *Metrics) GetInputJoin() int64 {
	if x != nil {
		return x.InputJoin
	}
	return 0
}

func (x *Metrics) GetInputCross() int64 {
	if x != nil {
		return x.InputCross
	}
	return 0
}

func (x *Metrics) GetInputUnion() int64 {
	if x != nil {
		return x.InputUnion
	}
	return 0
}

func (x *Metrics) GetInputCron() int64 {
	if x != nil {
		return x.InputCron
	}
	return 0
}

func (x *Metrics) GetInputGit() int64 {
	if x != nil {
		return x.InputGit
	}
	return 0
}

func (x *Metrics) GetInputPfs() int64 {
	if x != nil {
		return x.InputPfs
	}
	return 0
}

func (x *Metrics) GetInputCommit() int64 {
	if x != nil {
		return x.InputCommit
	}
	return 0
}

func (x *Metrics) GetInputJoinOn() int64 {
	if x != nil {
		return x.InputJoinOn
	}
	return 0
}

func (x *Metrics) GetInputOuterJoin() int64 {
	if x != nil {
		return x.InputOuterJoin
	}
	return 0
}

func (x *Metrics) GetInputLazy() int64 {
	if x != nil {
		return x.InputLazy
	}
	return 0
}

func (x *Metrics) GetInputEmptyFiles() int64 {
	if x != nil {
		return x.InputEmptyFiles
	}
	return 0
}

func (x *Metrics) GetInputS3() int64 {
	if x != nil {
		return x.InputS3
	}
	return 0
}

func (x *Metrics) GetInputTrigger() int64 {
	if x != nil {
		return x.InputTrigger
	}
	return 0
}

func (x *Metrics) GetResourceCpuReq() float32 {
	if x != nil {
		return x.ResourceCpuReq
	}
	return 0
}

func (x *Metrics) GetResourceCpuReqMax() float32 {
	if x != nil {
		return x.ResourceCpuReqMax
	}
	return 0
}

func (x *Metrics) GetResourceMemReq() string {
	if x != nil {
		return x.ResourceMemReq
	}
	return ""
}

func (x *Metrics) GetResourceGpuReq() int64 {
	if x != nil {
		return x.ResourceGpuReq
	}
	return 0
}

func (x *Metrics) GetResourceGpuReqMax() int64 {
	if x != nil {
		return x.ResourceGpuReqMax
	}
	return 0
}

func (x *Metrics) GetResourceDiskReq() string {
	if x != nil {
		return x.ResourceDiskReq
	}
	return ""
}

func (x *Metrics) GetResourceCpuLimit() float32 {
	if x != nil {
		return x.ResourceCpuLimit
	}
	return 0
}

func (x *Metrics) GetResourceCpuLimitMax() float32 {
	if x != nil {
		return x.ResourceCpuLimitMax
	}
	return 0
}

func (x *Metrics) GetResourceMemLimit() string {
	if x != nil {
		return x.ResourceMemLimit
	}
	return ""
}

func (x *Metrics) GetResourceGpuLimit() int64 {
	if x != nil {
		return x.ResourceGpuLimit
	}
	return 0
}

func (x *Metrics) GetResourceGpuLimitMax() int64 {
	if x != nil {
		return x.ResourceGpuLimitMax
	}
	return 0
}

func (x *Metrics) GetResourceDiskLimit() string {
	if x != nil {
		return x.ResourceDiskLimit
	}
	return ""
}

func (x *Metrics) GetMaxParallelism() uint64 {
	if x != nil {
		return x.MaxParallelism
	}
	return 0
}

func (x *Metrics) GetMinParallelism() uint64 {
	if x != nil {
		return x.MinParallelism
	}
	return 0
}

func (x *Metrics) GetNumParallelism() uint64 {
	if x != nil {
		return x.NumParallelism
	}
	return 0
}

var File_internal_metrics_metrics_proto protoreflect.FileDescriptor

var file_internal_metrics_metrics_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x22, 0xea, 0x0e, 0x0a, 0x07, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x70, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x6f, 0x64, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6e,
	0x6f, 0x64, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x72,
	0x65, 0x70, 0x6f, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6f,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x70,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x72, 0x63,
	0x68, 0x69, 0x76, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x73, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x64, 0x43, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65,
	0x64, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x10, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x73, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x61,
	0x78, 0x5f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0b, 0x6d, 0x61, 0x78, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x12, 0x1b, 0x0a,
	0x09, 0x70, 0x70, 0x73, 0x5f, 0x73, 0x70, 0x6f, 0x75, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x70, 0x70, 0x73, 0x53, 0x70, 0x6f, 0x75, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x70, 0x70,
	0x73, 0x5f, 0x73, 0x70, 0x6f, 0x75, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x70, 0x70, 0x73, 0x53, 0x70, 0x6f, 0x75, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x70, 0x73, 0x5f, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x70, 0x73, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x66, 0x67, 0x5f, 0x65, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x12, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x66, 0x67, 0x45, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x66, 0x67, 0x5f, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x62,
	0x79, 0x18, 0x13, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x66, 0x67, 0x53, 0x74, 0x61, 0x6e,
	0x64, 0x62, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x66, 0x67, 0x5f, 0x73, 0x33, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x18, 0x14, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x66, 0x67, 0x53,
	0x33, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x66, 0x67, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x15, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x63, 0x66, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x66, 0x67, 0x5f, 0x65, 0x72, 0x72, 0x63, 0x6d, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x63, 0x66, 0x67, 0x45, 0x72, 0x72, 0x63, 0x6d, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x66,
	0x67, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x17, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63,
	0x66, 0x67, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x66, 0x67, 0x5f, 0x74,
	0x66, 0x6a, 0x6f, 0x62, 0x18, 0x18, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x66, 0x67, 0x54,
	0x66, 0x6a, 0x6f, 0x62, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x18, 0x19, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x69, 0x6e, 0x70, 0x75, 0x74,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x6a,
	0x6f, 0x69, 0x6e, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x69, 0x6e, 0x70, 0x75, 0x74,
	0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x63, 0x72,
	0x6f, 0x73, 0x73, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x69, 0x6e, 0x70, 0x75, 0x74,
	0x43, 0x72, 0x6f, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x75,
	0x6e, 0x69, 0x6f, 0x6e, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x55, 0x6e, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f,
	0x63, 0x72, 0x6f, 0x6e, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x43, 0x72, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x67,
	0x69, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x47,
	0x69, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x70, 0x66, 0x73, 0x18,
	0x1f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x50, 0x66, 0x73, 0x12,
	0x21, 0x0a, 0x0c, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18,
	0x20, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x43, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x6a, 0x6f, 0x69, 0x6e,
	0x5f, 0x6f, 0x6e, 0x18, 0x21, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x69, 0x6e, 0x70, 0x75, 0x74,
	0x4a, 0x6f, 0x69, 0x6e, 0x4f, 0x6e, 0x12, 0x28, 0x0a, 0x10, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f,
	0x6f, 0x75, 0x74, 0x65, 0x72, 0x5f, 0x6a, 0x6f, 0x69, 0x6e, 0x18, 0x22, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0e, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x4f, 0x75, 0x74, 0x65, 0x72, 0x4a, 0x6f, 0x69, 0x6e,
	0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x6c, 0x61, 0x7a, 0x79, 0x18, 0x23,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x4c, 0x61, 0x7a, 0x79, 0x12,
	0x2a, 0x0a, 0x11, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x5f, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x18, 0x24, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x5f, 0x73, 0x33, 0x18, 0x25, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x53, 0x33, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f,
	0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x18, 0x26, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x10, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x70, 0x75, 0x5f, 0x72, 0x65, 0x71, 0x18,
	0x27, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43,
	0x70, 0x75, 0x52, 0x65, 0x71, 0x12, 0x2f, 0x0a, 0x14, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x63, 0x70, 0x75, 0x5f, 0x72, 0x65, 0x71, 0x5f, 0x6d, 0x61, 0x78, 0x18, 0x28, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x11, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x70, 0x75,
	0x52, 0x65, 0x71, 0x4d, 0x61, 0x78, 0x12, 0x28, 0x0a, 0x10, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x6d, 0x65, 0x6d, 0x5f, 0x72, 0x65, 0x71, 0x18, 0x29, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x65, 0x6d, 0x52, 0x65, 0x71,
	0x12, 0x28, 0x0a, 0x10, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x67, 0x70, 0x75,
	0x5f, 0x72, 0x65, 0x71, 0x18, 0x2a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x47, 0x70, 0x75, 0x52, 0x65, 0x71, 0x12, 0x2f, 0x0a, 0x14, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x67, 0x70, 0x75, 0x5f, 0x72, 0x65, 0x71, 0x5f, 0x6d,
	0x61, 0x78, 0x18, 0x2b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x47, 0x70, 0x75, 0x52, 0x65, 0x71, 0x4d, 0x61, 0x78, 0x12, 0x2a, 0x0a, 0x11, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x72, 0x65, 0x71,
	0x18, 0x2c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x44, 0x69, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x2c, 0x0a, 0x12, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x63, 0x70, 0x75, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x2d, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x10, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x70, 0x75,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x33, 0x0a, 0x16, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x63, 0x70, 0x75, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x6d, 0x61, 0x78, 0x18,
	0x2e, 0x20, 0x01, 0x28, 0x02, 0x52, 0x13, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43,
	0x70, 0x75, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4d, 0x61, 0x78, 0x12, 0x2c, 0x0a, 0x12, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x6d, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x2f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x4d, 0x65, 0x6d, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x67, 0x70, 0x75, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x30,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x47, 0x70,
	0x75, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x33, 0x0a, 0x16, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x67, 0x70, 0x75, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x6d, 0x61, 0x78,
	0x18, 0x31, 0x20, 0x01, 0x28, 0x03, 0x52, 0x13, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x47, 0x70, 0x75, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4d, 0x61, 0x78, 0x12, 0x2e, 0x0a, 0x13, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x44, 0x69, 0x73, 0x6b, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x6d,
	0x61, 0x78, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6c, 0x6c, 0x65, 0x6c, 0x69, 0x73, 0x6d, 0x18, 0x33,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x6d, 0x61, 0x78, 0x50, 0x61, 0x72, 0x61, 0x6c, 0x6c, 0x65,
	0x6c, 0x69, 0x73, 0x6d, 0x12, 0x27, 0x0a, 0x0f, 0x6d, 0x69, 0x6e, 0x5f, 0x70, 0x61, 0x72, 0x61,
	0x6c, 0x6c, 0x65, 0x6c, 0x69, 0x73, 0x6d, 0x18, 0x34, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x6d,
	0x69, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6c, 0x6c, 0x65, 0x6c, 0x69, 0x73, 0x6d, 0x12, 0x27, 0x0a,
	0x0f, 0x6e, 0x75, 0x6d, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6c, 0x6c, 0x65, 0x6c, 0x69, 0x73, 0x6d,
	0x18, 0x35, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x6e, 0x75, 0x6d, 0x50, 0x61, 0x72, 0x61, 0x6c,
	0x6c, 0x65, 0x6c, 0x69, 0x73, 0x6d, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x61, 0x63, 0x68, 0x79, 0x64, 0x65, 0x72, 0x6d, 0x2f, 0x70,
	0x61, 0x63, 0x68, 0x79, 0x64, 0x65, 0x72, 0x6d, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x72, 0x63, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_metrics_metrics_proto_rawDescOnce sync.Once
	file_internal_metrics_metrics_proto_rawDescData = file_internal_metrics_metrics_proto_rawDesc
)

func file_internal_metrics_metrics_proto_rawDescGZIP() []byte {
	file_internal_metrics_metrics_proto_rawDescOnce.Do(func() {
		file_internal_metrics_metrics_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_metrics_metrics_proto_rawDescData)
	})
	return file_internal_metrics_metrics_proto_rawDescData
}

var file_internal_metrics_metrics_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_internal_metrics_metrics_proto_goTypes = []interface{}{
	(*Metrics)(nil), // 0: metrics.Metrics
}
var file_internal_metrics_metrics_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_metrics_metrics_proto_init() }
func file_internal_metrics_metrics_proto_init() {
	if File_internal_metrics_metrics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_metrics_metrics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metrics); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_metrics_metrics_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_metrics_metrics_proto_goTypes,
		DependencyIndexes: file_internal_metrics_metrics_proto_depIdxs,
		MessageInfos:      file_internal_metrics_metrics_proto_msgTypes,
	}.Build()
	File_internal_metrics_metrics_proto = out.File
	file_internal_metrics_metrics_proto_rawDesc = nil
	file_internal_metrics_metrics_proto_goTypes = nil
	file_internal_metrics_metrics_proto_depIdxs = nil
}
