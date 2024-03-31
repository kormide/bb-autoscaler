// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.23.1
// source: pkg/proto/configuration/bb_autoscaler/bb_autoscaler.proto

package bb_autoscaler

import (
	v2 "github.com/bazelbuild/remote-apis/build/bazel/remote/execution/v2"
	aws "github.com/buildbarn/bb-storage/pkg/proto/configuration/cloud/aws"
	http "github.com/buildbarn/bb-storage/pkg/proto/configuration/http"
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

type ApplicationConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrometheusHttpClient *http.ClientConfiguration `protobuf:"bytes,5,opt,name=prometheus_http_client,json=prometheusHttpClient,proto3" json:"prometheus_http_client,omitempty"`
	PrometheusEndpoint   string                    `protobuf:"bytes,1,opt,name=prometheus_endpoint,json=prometheusEndpoint,proto3" json:"prometheus_endpoint,omitempty"`
	PrometheusQuery      string                    `protobuf:"bytes,2,opt,name=prometheus_query,json=prometheusQuery,proto3" json:"prometheus_query,omitempty"`
	NodeGroups           []*NodeGroupConfiguration `protobuf:"bytes,3,rep,name=node_groups,json=nodeGroups,proto3" json:"node_groups,omitempty"`
	AwsSession           *aws.SessionConfiguration `protobuf:"bytes,4,opt,name=aws_session,json=awsSession,proto3" json:"aws_session,omitempty"`
}

func (x *ApplicationConfiguration) Reset() {
	*x = ApplicationConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_configuration_bb_autoscaler_bb_autoscaler_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplicationConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationConfiguration) ProtoMessage() {}

func (x *ApplicationConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_configuration_bb_autoscaler_bb_autoscaler_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationConfiguration.ProtoReflect.Descriptor instead.
func (*ApplicationConfiguration) Descriptor() ([]byte, []int) {
	return file_pkg_proto_configuration_bb_autoscaler_bb_autoscaler_proto_rawDescGZIP(), []int{0}
}

func (x *ApplicationConfiguration) GetPrometheusHttpClient() *http.ClientConfiguration {
	if x != nil {
		return x.PrometheusHttpClient
	}
	return nil
}

func (x *ApplicationConfiguration) GetPrometheusEndpoint() string {
	if x != nil {
		return x.PrometheusEndpoint
	}
	return ""
}

func (x *ApplicationConfiguration) GetPrometheusQuery() string {
	if x != nil {
		return x.PrometheusQuery
	}
	return ""
}

func (x *ApplicationConfiguration) GetNodeGroups() []*NodeGroupConfiguration {
	if x != nil {
		return x.NodeGroups
	}
	return nil
}

func (x *ApplicationConfiguration) GetAwsSession() *aws.SessionConfiguration {
	if x != nil {
		return x.AwsSession
	}
	return nil
}

type EKSManagedNodeGroupConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterName   string `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	NodeGroupName string `protobuf:"bytes,2,opt,name=node_group_name,json=nodeGroupName,proto3" json:"node_group_name,omitempty"`
}

func (x *EKSManagedNodeGroupConfiguration) Reset() {
	*x = EKSManagedNodeGroupConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_configuration_bb_autoscaler_bb_autoscaler_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EKSManagedNodeGroupConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EKSManagedNodeGroupConfiguration) ProtoMessage() {}

func (x *EKSManagedNodeGroupConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_configuration_bb_autoscaler_bb_autoscaler_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EKSManagedNodeGroupConfiguration.ProtoReflect.Descriptor instead.
func (*EKSManagedNodeGroupConfiguration) Descriptor() ([]byte, []int) {
	return file_pkg_proto_configuration_bb_autoscaler_bb_autoscaler_proto_rawDescGZIP(), []int{1}
}

func (x *EKSManagedNodeGroupConfiguration) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *EKSManagedNodeGroupConfiguration) GetNodeGroupName() string {
	if x != nil {
		return x.NodeGroupName
	}
	return ""
}

type KubernetesDeploymentConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace       string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name            string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	MinimumReplicas int32  `protobuf:"varint,3,opt,name=minimum_replicas,json=minimumReplicas,proto3" json:"minimum_replicas,omitempty"`
	MaximumReplicas int32  `protobuf:"varint,4,opt,name=maximum_replicas,json=maximumReplicas,proto3" json:"maximum_replicas,omitempty"`
}

func (x *KubernetesDeploymentConfiguration) Reset() {
	*x = KubernetesDeploymentConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_configuration_bb_autoscaler_bb_autoscaler_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubernetesDeploymentConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubernetesDeploymentConfiguration) ProtoMessage() {}

func (x *KubernetesDeploymentConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_configuration_bb_autoscaler_bb_autoscaler_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubernetesDeploymentConfiguration.ProtoReflect.Descriptor instead.
func (*KubernetesDeploymentConfiguration) Descriptor() ([]byte, []int) {
	return file_pkg_proto_configuration_bb_autoscaler_bb_autoscaler_proto_rawDescGZIP(), []int{2}
}

func (x *KubernetesDeploymentConfiguration) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *KubernetesDeploymentConfiguration) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *KubernetesDeploymentConfiguration) GetMinimumReplicas() int32 {
	if x != nil {
		return x.MinimumReplicas
	}
	return 0
}

func (x *KubernetesDeploymentConfiguration) GetMaximumReplicas() int32 {
	if x != nil {
		return x.MaximumReplicas
	}
	return 0
}

type NodeGroupConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstanceNamePrefix string       `protobuf:"bytes,1,opt,name=instance_name_prefix,json=instanceNamePrefix,proto3" json:"instance_name_prefix,omitempty"`
	Platform           *v2.Platform `protobuf:"bytes,2,opt,name=platform,proto3" json:"platform,omitempty"`
	SizeClass          uint32       `protobuf:"varint,6,opt,name=size_class,json=sizeClass,proto3" json:"size_class,omitempty"`
	// Types that are assignable to Kind:
	//
	//	*NodeGroupConfiguration_AutoScalingGroupName
	//	*NodeGroupConfiguration_EksManagedNodeGroup
	//	*NodeGroupConfiguration_KubernetesDeployment
	Kind                   isNodeGroupConfiguration_Kind `protobuf_oneof:"kind"`
	WorkersPerCapacityUnit int32                         `protobuf:"varint,4,opt,name=workers_per_capacity_unit,json=workersPerCapacityUnit,proto3" json:"workers_per_capacity_unit,omitempty"`
}

func (x *NodeGroupConfiguration) Reset() {
	*x = NodeGroupConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_configuration_bb_autoscaler_bb_autoscaler_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeGroupConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeGroupConfiguration) ProtoMessage() {}

func (x *NodeGroupConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_configuration_bb_autoscaler_bb_autoscaler_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeGroupConfiguration.ProtoReflect.Descriptor instead.
func (*NodeGroupConfiguration) Descriptor() ([]byte, []int) {
	return file_pkg_proto_configuration_bb_autoscaler_bb_autoscaler_proto_rawDescGZIP(), []int{3}
}

func (x *NodeGroupConfiguration) GetInstanceNamePrefix() string {
	if x != nil {
		return x.InstanceNamePrefix
	}
	return ""
}

func (x *NodeGroupConfiguration) GetPlatform() *v2.Platform {
	if x != nil {
		return x.Platform
	}
	return nil
}

func (x *NodeGroupConfiguration) GetSizeClass() uint32 {
	if x != nil {
		return x.SizeClass
	}
	return 0
}

func (m *NodeGroupConfiguration) GetKind() isNodeGroupConfiguration_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (x *NodeGroupConfiguration) GetAutoScalingGroupName() string {
	if x, ok := x.GetKind().(*NodeGroupConfiguration_AutoScalingGroupName); ok {
		return x.AutoScalingGroupName
	}
	return ""
}

func (x *NodeGroupConfiguration) GetEksManagedNodeGroup() *EKSManagedNodeGroupConfiguration {
	if x, ok := x.GetKind().(*NodeGroupConfiguration_EksManagedNodeGroup); ok {
		return x.EksManagedNodeGroup
	}
	return nil
}

func (x *NodeGroupConfiguration) GetKubernetesDeployment() *KubernetesDeploymentConfiguration {
	if x, ok := x.GetKind().(*NodeGroupConfiguration_KubernetesDeployment); ok {
		return x.KubernetesDeployment
	}
	return nil
}

func (x *NodeGroupConfiguration) GetWorkersPerCapacityUnit() int32 {
	if x != nil {
		return x.WorkersPerCapacityUnit
	}
	return 0
}

type isNodeGroupConfiguration_Kind interface {
	isNodeGroupConfiguration_Kind()
}

type NodeGroupConfiguration_AutoScalingGroupName struct {
	AutoScalingGroupName string `protobuf:"bytes,3,opt,name=auto_scaling_group_name,json=autoScalingGroupName,proto3,oneof"`
}

type NodeGroupConfiguration_EksManagedNodeGroup struct {
	EksManagedNodeGroup *EKSManagedNodeGroupConfiguration `protobuf:"bytes,5,opt,name=eks_managed_node_group,json=eksManagedNodeGroup,proto3,oneof"`
}

type NodeGroupConfiguration_KubernetesDeployment struct {
	KubernetesDeployment *KubernetesDeploymentConfiguration `protobuf:"bytes,7,opt,name=kubernetes_deployment,json=kubernetesDeployment,proto3,oneof"`
}

func (*NodeGroupConfiguration_AutoScalingGroupName) isNodeGroupConfiguration_Kind() {}

func (*NodeGroupConfiguration_EksManagedNodeGroup) isNodeGroupConfiguration_Kind() {}

func (*NodeGroupConfiguration_KubernetesDeployment) isNodeGroupConfiguration_Kind() {}

var File_pkg_proto_configuration_bb_autoscaler_bb_autoscaler_proto protoreflect.FileDescriptor

var file_pkg_proto_configuration_bb_autoscaler_bb_autoscaler_proto_rawDesc = []byte{
	0x0a, 0x39, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x62, 0x62, 0x5f, 0x61, 0x75, 0x74,
	0x6f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x2f, 0x62, 0x62, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x73,
	0x63, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x25, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x62, 0x62, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c,
	0x65, 0x72, 0x1a, 0x36, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x62, 0x61, 0x7a, 0x65, 0x6c, 0x2f,
	0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x70, 0x6b, 0x67, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x77, 0x73, 0x2f, 0x61, 0x77,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x99, 0x03, 0x0a, 0x18, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x67, 0x0a,
	0x16, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x5f, 0x68, 0x74, 0x74, 0x70,
	0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x14, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x48, 0x74, 0x74, 0x70,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x13, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74,
	0x68, 0x65, 0x75, 0x73, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x12, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x6d, 0x65,
	0x74, 0x68, 0x65, 0x75, 0x73, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x12, 0x5e, 0x0a, 0x0b, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62,
	0x61, 0x72, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x62, 0x62, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x2e,
	0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x12, 0x58, 0x0a, 0x0b, 0x61, 0x77, 0x73, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62,
	0x61, 0x72, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0a, 0x61, 0x77, 0x73, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x6d, 0x0a, 0x20,
	0x45, 0x4b, 0x53, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x6f,
	0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xab, 0x01, 0x0a, 0x21,
	0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x5f, 0x72,
	0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x6d,
	0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x12, 0x29,
	0x0a, 0x10, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75,
	0x6d, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x22, 0xad, 0x04, 0x0a, 0x16, 0x4e, 0x6f,
	0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x14, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x12, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x45, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x2e, 0x62, 0x61, 0x7a, 0x65, 0x6c, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x65, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x73, 0x69, 0x7a, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x37, 0x0a, 0x17,
	0x61, 0x75, 0x74, 0x6f, 0x5f, 0x73, 0x63, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x14, 0x61, 0x75, 0x74, 0x6f, 0x53, 0x63, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x7e, 0x0a, 0x16, 0x65, 0x6b, 0x73, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x64, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x47, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72,
	0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x62, 0x62, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x45, 0x4b,
	0x53, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00,
	0x52, 0x13, 0x65, 0x6b, 0x73, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x4e, 0x6f, 0x64, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x7f, 0x0a, 0x15, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x5f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x48, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x62,
	0x62, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x4b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00,
	0x52, 0x14, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x44, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x39, 0x0a, 0x19, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x75,
	0x6e, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x16, 0x77, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x73, 0x50, 0x65, 0x72, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x55, 0x6e, 0x69,
	0x74, 0x42, 0x06, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x42, 0x4a, 0x5a, 0x48, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72,
	0x6e, 0x2f, 0x62, 0x62, 0x2d, 0x61, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x62, 0x62, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x73,
	0x63, 0x61, 0x6c, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_configuration_bb_autoscaler_bb_autoscaler_proto_rawDescOnce sync.Once
	file_pkg_proto_configuration_bb_autoscaler_bb_autoscaler_proto_rawDescData = file_pkg_proto_configuration_bb_autoscaler_bb_autoscaler_proto_rawDesc
)

func file_pkg_proto_configuration_bb_autoscaler_bb_autoscaler_proto_rawDescGZIP() []byte {
	file_pkg_proto_configuration_bb_autoscaler_bb_autoscaler_proto_rawDescOnce.Do(func() {
		file_pkg_proto_configuration_bb_autoscaler_bb_autoscaler_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_configuration_bb_autoscaler_bb_autoscaler_proto_rawDescData)
	})
	return file_pkg_proto_configuration_bb_autoscaler_bb_autoscaler_proto_rawDescData
}

var file_pkg_proto_configuration_bb_autoscaler_bb_autoscaler_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_proto_configuration_bb_autoscaler_bb_autoscaler_proto_goTypes = []interface{}{
	(*ApplicationConfiguration)(nil),          // 0: buildbarn.configuration.bb_autoscaler.ApplicationConfiguration
	(*EKSManagedNodeGroupConfiguration)(nil),  // 1: buildbarn.configuration.bb_autoscaler.EKSManagedNodeGroupConfiguration
	(*KubernetesDeploymentConfiguration)(nil), // 2: buildbarn.configuration.bb_autoscaler.KubernetesDeploymentConfiguration
	(*NodeGroupConfiguration)(nil),            // 3: buildbarn.configuration.bb_autoscaler.NodeGroupConfiguration
	(*http.ClientConfiguration)(nil),          // 4: buildbarn.configuration.http.ClientConfiguration
	(*aws.SessionConfiguration)(nil),          // 5: buildbarn.configuration.cloud.aws.SessionConfiguration
	(*v2.Platform)(nil),                       // 6: build.bazel.remote.execution.v2.Platform
}
var file_pkg_proto_configuration_bb_autoscaler_bb_autoscaler_proto_depIdxs = []int32{
	4, // 0: buildbarn.configuration.bb_autoscaler.ApplicationConfiguration.prometheus_http_client:type_name -> buildbarn.configuration.http.ClientConfiguration
	3, // 1: buildbarn.configuration.bb_autoscaler.ApplicationConfiguration.node_groups:type_name -> buildbarn.configuration.bb_autoscaler.NodeGroupConfiguration
	5, // 2: buildbarn.configuration.bb_autoscaler.ApplicationConfiguration.aws_session:type_name -> buildbarn.configuration.cloud.aws.SessionConfiguration
	6, // 3: buildbarn.configuration.bb_autoscaler.NodeGroupConfiguration.platform:type_name -> build.bazel.remote.execution.v2.Platform
	1, // 4: buildbarn.configuration.bb_autoscaler.NodeGroupConfiguration.eks_managed_node_group:type_name -> buildbarn.configuration.bb_autoscaler.EKSManagedNodeGroupConfiguration
	2, // 5: buildbarn.configuration.bb_autoscaler.NodeGroupConfiguration.kubernetes_deployment:type_name -> buildbarn.configuration.bb_autoscaler.KubernetesDeploymentConfiguration
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_pkg_proto_configuration_bb_autoscaler_bb_autoscaler_proto_init() }
func file_pkg_proto_configuration_bb_autoscaler_bb_autoscaler_proto_init() {
	if File_pkg_proto_configuration_bb_autoscaler_bb_autoscaler_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_configuration_bb_autoscaler_bb_autoscaler_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplicationConfiguration); i {
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
		file_pkg_proto_configuration_bb_autoscaler_bb_autoscaler_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EKSManagedNodeGroupConfiguration); i {
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
		file_pkg_proto_configuration_bb_autoscaler_bb_autoscaler_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubernetesDeploymentConfiguration); i {
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
		file_pkg_proto_configuration_bb_autoscaler_bb_autoscaler_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeGroupConfiguration); i {
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
	file_pkg_proto_configuration_bb_autoscaler_bb_autoscaler_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*NodeGroupConfiguration_AutoScalingGroupName)(nil),
		(*NodeGroupConfiguration_EksManagedNodeGroup)(nil),
		(*NodeGroupConfiguration_KubernetesDeployment)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_proto_configuration_bb_autoscaler_bb_autoscaler_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_proto_configuration_bb_autoscaler_bb_autoscaler_proto_goTypes,
		DependencyIndexes: file_pkg_proto_configuration_bb_autoscaler_bb_autoscaler_proto_depIdxs,
		MessageInfos:      file_pkg_proto_configuration_bb_autoscaler_bb_autoscaler_proto_msgTypes,
	}.Build()
	File_pkg_proto_configuration_bb_autoscaler_bb_autoscaler_proto = out.File
	file_pkg_proto_configuration_bb_autoscaler_bb_autoscaler_proto_rawDesc = nil
	file_pkg_proto_configuration_bb_autoscaler_bb_autoscaler_proto_goTypes = nil
	file_pkg_proto_configuration_bb_autoscaler_bb_autoscaler_proto_depIdxs = nil
}
