// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.2
// source: internal/config/config.proto

package config

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

type ContextSource int32

const (
	ContextSource_NONE      ContextSource = 0
	ContextSource_CONFIG_V1 ContextSource = 1
	ContextSource_HUB       ContextSource = 2
	ContextSource_IMPORTED  ContextSource = 3
)

// Enum value maps for ContextSource.
var (
	ContextSource_name = map[int32]string{
		0: "NONE",
		1: "CONFIG_V1",
		2: "HUB",
		3: "IMPORTED",
	}
	ContextSource_value = map[string]int32{
		"NONE":      0,
		"CONFIG_V1": 1,
		"HUB":       2,
		"IMPORTED":  3,
	}
)

func (x ContextSource) Enum() *ContextSource {
	p := new(ContextSource)
	*p = x
	return p
}

func (x ContextSource) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ContextSource) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_config_config_proto_enumTypes[0].Descriptor()
}

func (ContextSource) Type() protoreflect.EnumType {
	return &file_internal_config_config_proto_enumTypes[0]
}

func (x ContextSource) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ContextSource.Descriptor instead.
func (ContextSource) EnumDescriptor() ([]byte, []int) {
	return file_internal_config_config_proto_rawDescGZIP(), []int{0}
}

// Config specifies the pachyderm config that is read and interpreted by the
// pachctl command-line tool. Right now, this is stored at
// $HOME/.pachyderm/config.
//
// Different versions of the pachyderm config are specified as subfields of this
// message (this allows us to make significant changes to the config structure
// without breaking existing users by defining a new config version).
//
// DO NOT change or remove field numbers from this proto, otherwise ALL
// pachyderm user configs will fail to parse.
type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// Configuration options. Exactly one of these fields should be set
	// (depending on which version of the config is being used)
	V1 *ConfigV1 `protobuf:"bytes,2,opt,name=v1,proto3" json:"v1,omitempty"`
	V2 *ConfigV2 `protobuf:"bytes,3,opt,name=v2,proto3" json:"v2,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_config_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_internal_config_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_internal_config_config_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Config) GetV1() *ConfigV1 {
	if x != nil {
		return x.V1
	}
	return nil
}

func (x *Config) GetV2() *ConfigV2 {
	if x != nil {
		return x.V2
	}
	return nil
}

// ConfigV1 specifies v1 of the pachyderm config (June 30 2017 - June 2019)
// DO NOT change or remove field numbers from this proto, as if you do, v1 user
// configs will become unparseable.
type ConfigV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A host:port pointing pachd at a pachyderm cluster.
	PachdAddress string `protobuf:"bytes,2,opt,name=pachd_address,json=pachdAddress,proto3" json:"pachd_address,omitempty"`
	// Trusted root certificates (overrides installed certificates), formatted
	// as base64-encoded PEM
	ServerCas string `protobuf:"bytes,3,opt,name=server_cas,json=serverCas,proto3" json:"server_cas,omitempty"`
	// A secret token identifying the current pachctl user within their
	// pachyderm cluster. This is included in all RPCs sent by pachctl, and used
	// to determine if pachctl actions are authorized.
	SessionToken string `protobuf:"bytes,1,opt,name=session_token,json=sessionToken,proto3" json:"session_token,omitempty"`
	// The currently active transaction for batching together pachctl commands.
	// This can be set or cleared via many of the `pachctl * transaction` commands.
	// This is the ID of the transaction object stored in the pachyderm etcd.
	ActiveTransaction string `protobuf:"bytes,4,opt,name=active_transaction,json=activeTransaction,proto3" json:"active_transaction,omitempty"`
}

func (x *ConfigV1) Reset() {
	*x = ConfigV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_config_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigV1) ProtoMessage() {}

func (x *ConfigV1) ProtoReflect() protoreflect.Message {
	mi := &file_internal_config_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigV1.ProtoReflect.Descriptor instead.
func (*ConfigV1) Descriptor() ([]byte, []int) {
	return file_internal_config_config_proto_rawDescGZIP(), []int{1}
}

func (x *ConfigV1) GetPachdAddress() string {
	if x != nil {
		return x.PachdAddress
	}
	return ""
}

func (x *ConfigV1) GetServerCas() string {
	if x != nil {
		return x.ServerCas
	}
	return ""
}

func (x *ConfigV1) GetSessionToken() string {
	if x != nil {
		return x.SessionToken
	}
	return ""
}

func (x *ConfigV1) GetActiveTransaction() string {
	if x != nil {
		return x.ActiveTransaction
	}
	return ""
}

// ConfigV2 specifies v2 of the pachyderm config (June 2019 - present)
type ConfigV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActiveContext           string              `protobuf:"bytes,1,opt,name=active_context,json=activeContext,proto3" json:"active_context,omitempty"`
	ActiveEnterpriseContext string              `protobuf:"bytes,5,opt,name=active_enterprise_context,json=activeEnterpriseContext,proto3" json:"active_enterprise_context,omitempty"`
	Contexts                map[string]*Context `protobuf:"bytes,2,rep,name=contexts,proto3" json:"contexts,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Metrics                 bool                `protobuf:"varint,3,opt,name=metrics,proto3" json:"metrics,omitempty"`
	MaxShellCompletions     int64               `protobuf:"varint,4,opt,name=max_shell_completions,json=maxShellCompletions,proto3" json:"max_shell_completions,omitempty"`
}

func (x *ConfigV2) Reset() {
	*x = ConfigV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_config_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigV2) ProtoMessage() {}

func (x *ConfigV2) ProtoReflect() protoreflect.Message {
	mi := &file_internal_config_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigV2.ProtoReflect.Descriptor instead.
func (*ConfigV2) Descriptor() ([]byte, []int) {
	return file_internal_config_config_proto_rawDescGZIP(), []int{2}
}

func (x *ConfigV2) GetActiveContext() string {
	if x != nil {
		return x.ActiveContext
	}
	return ""
}

func (x *ConfigV2) GetActiveEnterpriseContext() string {
	if x != nil {
		return x.ActiveEnterpriseContext
	}
	return ""
}

func (x *ConfigV2) GetContexts() map[string]*Context {
	if x != nil {
		return x.Contexts
	}
	return nil
}

func (x *ConfigV2) GetMetrics() bool {
	if x != nil {
		return x.Metrics
	}
	return false
}

func (x *ConfigV2) GetMaxShellCompletions() int64 {
	if x != nil {
		return x.MaxShellCompletions
	}
	return 0
}

type Context struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Where this context came from
	Source ContextSource `protobuf:"varint,1,opt,name=source,proto3,enum=config.ContextSource" json:"source,omitempty"`
	// The hostname or IP address pointing pachd at a pachyderm cluster.
	PachdAddress string `protobuf:"bytes,2,opt,name=pachd_address,json=pachdAddress,proto3" json:"pachd_address,omitempty"`
	// Trusted root certificates (overrides installed certificates), formatted
	// as base64-encoded PEM.
	ServerCas string `protobuf:"bytes,3,opt,name=server_cas,json=serverCas,proto3" json:"server_cas,omitempty"`
	// A secret token identifying the current pachctl user within their
	// pachyderm cluster. This is included in all RPCs sent by pachctl, and used
	// to determine if pachctl actions are authorized.
	SessionToken string `protobuf:"bytes,4,opt,name=session_token,json=sessionToken,proto3" json:"session_token,omitempty"`
	// The currently active transaction for batching together pachctl commands.
	// This can be set or cleared via many of the `pachctl * transaction` commands.
	// This is the ID of the transaction object stored in the pachyderm etcd.
	ActiveTransaction string `protobuf:"bytes,5,opt,name=active_transaction,json=activeTransaction,proto3" json:"active_transaction,omitempty"`
	// The k8s cluster name - used to construct a k8s context.
	ClusterName string `protobuf:"bytes,6,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	// The k8s auth info - used to construct a k8s context.
	AuthInfo string `protobuf:"bytes,7,opt,name=auth_info,json=authInfo,proto3" json:"auth_info,omitempty"`
	// The k8s namespace - used to construct a k8s context.
	Namespace string `protobuf:"bytes,8,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// A mapping of service -> port number, when port forwarding is
	// running for this context.
	PortForwarders map[string]uint32 `protobuf:"bytes,10,rep,name=port_forwarders,json=portForwarders,proto3" json:"port_forwarders,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// A unique ID for the cluster deployment. At client initialization time,
	// we ensure this is the same as what the cluster reports back, to prevent
	// us from connecting to the wrong cluster.
	ClusterDeploymentId string `protobuf:"bytes,11,opt,name=cluster_deployment_id,json=clusterDeploymentId,proto3" json:"cluster_deployment_id,omitempty"`
}

func (x *Context) Reset() {
	*x = Context{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_config_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Context) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Context) ProtoMessage() {}

func (x *Context) ProtoReflect() protoreflect.Message {
	mi := &file_internal_config_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Context.ProtoReflect.Descriptor instead.
func (*Context) Descriptor() ([]byte, []int) {
	return file_internal_config_config_proto_rawDescGZIP(), []int{3}
}

func (x *Context) GetSource() ContextSource {
	if x != nil {
		return x.Source
	}
	return ContextSource_NONE
}

func (x *Context) GetPachdAddress() string {
	if x != nil {
		return x.PachdAddress
	}
	return ""
}

func (x *Context) GetServerCas() string {
	if x != nil {
		return x.ServerCas
	}
	return ""
}

func (x *Context) GetSessionToken() string {
	if x != nil {
		return x.SessionToken
	}
	return ""
}

func (x *Context) GetActiveTransaction() string {
	if x != nil {
		return x.ActiveTransaction
	}
	return ""
}

func (x *Context) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *Context) GetAuthInfo() string {
	if x != nil {
		return x.AuthInfo
	}
	return ""
}

func (x *Context) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Context) GetPortForwarders() map[string]uint32 {
	if x != nil {
		return x.PortForwarders
	}
	return nil
}

func (x *Context) GetClusterDeploymentId() string {
	if x != nil {
		return x.ClusterDeploymentId
	}
	return ""
}

var File_internal_config_config_proto protoreflect.FileDescriptor

var file_internal_config_config_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x65, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x02, 0x76, 0x31, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x56, 0x31, 0x52, 0x02, 0x76, 0x31, 0x12, 0x20, 0x0a, 0x02, 0x76,
	0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x56, 0x32, 0x52, 0x02, 0x76, 0x32, 0x22, 0xa2, 0x01,
	0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x56, 0x31, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x61,
	0x63, 0x68, 0x64, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x70, 0x61, 0x63, 0x68, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x63, 0x61, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x61, 0x73, 0x12, 0x23,
	0x0a, 0x0d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x2d, 0x0a, 0x12, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x11, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0xc5, 0x02, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x56, 0x32, 0x12,
	0x25, 0x0a, 0x0e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x3a, 0x0a, 0x19, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x5f, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x12, 0x3a, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x56, 0x32, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x32, 0x0a, 0x15, 0x6d, 0x61, 0x78, 0x5f,
	0x73, 0x68, 0x65, 0x6c, 0x6c, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x13, 0x6d, 0x61, 0x78, 0x53, 0x68, 0x65, 0x6c,
	0x6c, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x4c, 0x0a, 0x0d,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x25, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xf9, 0x03, 0x0a, 0x07, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x2d, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x61, 0x63, 0x68, 0x64, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x61,
	0x63, 0x68, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x5f, 0x63, 0x61, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x61, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2d,
	0x0a, 0x12, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x0a,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x0f, 0x70,
	0x6f, 0x72, 0x74, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x0a,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72,
	0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e, 0x70, 0x6f, 0x72, 0x74, 0x46,
	0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x32, 0x0a, 0x15, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x1a, 0x41, 0x0a,
	0x13, 0x50, 0x6f, 0x72, 0x74, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x4a, 0x04, 0x08, 0x09, 0x10, 0x0a, 0x2a, 0x3f, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10,
	0x00, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x56, 0x31, 0x10, 0x01,
	0x12, 0x07, 0x0a, 0x03, 0x48, 0x55, 0x42, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x4d, 0x50,
	0x4f, 0x52, 0x54, 0x45, 0x44, 0x10, 0x03, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x61, 0x63, 0x68, 0x79, 0x64, 0x65, 0x72, 0x6d, 0x2f,
	0x70, 0x61, 0x63, 0x68, 0x79, 0x64, 0x65, 0x72, 0x6d, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x72, 0x63,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_config_config_proto_rawDescOnce sync.Once
	file_internal_config_config_proto_rawDescData = file_internal_config_config_proto_rawDesc
)

func file_internal_config_config_proto_rawDescGZIP() []byte {
	file_internal_config_config_proto_rawDescOnce.Do(func() {
		file_internal_config_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_config_config_proto_rawDescData)
	})
	return file_internal_config_config_proto_rawDescData
}

var file_internal_config_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_internal_config_config_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_internal_config_config_proto_goTypes = []interface{}{
	(ContextSource)(0), // 0: config.ContextSource
	(*Config)(nil),     // 1: config.Config
	(*ConfigV1)(nil),   // 2: config.ConfigV1
	(*ConfigV2)(nil),   // 3: config.ConfigV2
	(*Context)(nil),    // 4: config.Context
	nil,                // 5: config.ConfigV2.ContextsEntry
	nil,                // 6: config.Context.PortForwardersEntry
}
var file_internal_config_config_proto_depIdxs = []int32{
	2, // 0: config.Config.v1:type_name -> config.ConfigV1
	3, // 1: config.Config.v2:type_name -> config.ConfigV2
	5, // 2: config.ConfigV2.contexts:type_name -> config.ConfigV2.ContextsEntry
	0, // 3: config.Context.source:type_name -> config.ContextSource
	6, // 4: config.Context.port_forwarders:type_name -> config.Context.PortForwardersEntry
	4, // 5: config.ConfigV2.ContextsEntry.value:type_name -> config.Context
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_internal_config_config_proto_init() }
func file_internal_config_config_proto_init() {
	if File_internal_config_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_config_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_internal_config_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigV1); i {
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
		file_internal_config_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigV2); i {
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
		file_internal_config_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Context); i {
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
			RawDescriptor: file_internal_config_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_config_config_proto_goTypes,
		DependencyIndexes: file_internal_config_config_proto_depIdxs,
		EnumInfos:         file_internal_config_config_proto_enumTypes,
		MessageInfos:      file_internal_config_config_proto_msgTypes,
	}.Build()
	File_internal_config_config_proto = out.File
	file_internal_config_config_proto_rawDesc = nil
	file_internal_config_config_proto_goTypes = nil
	file_internal_config_config_proto_depIdxs = nil
}
