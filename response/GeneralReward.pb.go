// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: GeneralReward.proto

package response

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

// 通用奖励消息
type GeneralReward struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32             `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg     string            `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Changes map[uint32]uint64 `protobuf:"bytes,3,rep,name=changes,proto3" json:"changes,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"` // 客户端展示奖励的部分 : 道具ID -> 道具数量
	Balance map[uint32]uint64 `protobuf:"bytes,4,rep,name=balance,proto3" json:"balance,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"` // 道具有变化部分的当前余额 : 道具ID -> 道具数量
	Counter map[uint32]uint64 `protobuf:"bytes,5,rep,name=counter,proto3" json:"counter,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"` // 计数器当前值 : counterType -> 计数
	Ext     string            `protobuf:"bytes,6,opt,name=ext,proto3" json:"ext,omitempty"`                                                                                                   // 扩展字段，IAP使用
}

func (x *GeneralReward) Reset() {
	*x = GeneralReward{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GeneralReward_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeneralReward) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneralReward) ProtoMessage() {}

func (x *GeneralReward) ProtoReflect() protoreflect.Message {
	mi := &file_GeneralReward_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneralReward.ProtoReflect.Descriptor instead.
func (*GeneralReward) Descriptor() ([]byte, []int) {
	return file_GeneralReward_proto_rawDescGZIP(), []int{0}
}

func (x *GeneralReward) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GeneralReward) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GeneralReward) GetChanges() map[uint32]uint64 {
	if x != nil {
		return x.Changes
	}
	return nil
}

func (x *GeneralReward) GetBalance() map[uint32]uint64 {
	if x != nil {
		return x.Balance
	}
	return nil
}

func (x *GeneralReward) GetCounter() map[uint32]uint64 {
	if x != nil {
		return x.Counter
	}
	return nil
}

func (x *GeneralReward) GetExt() string {
	if x != nil {
		return x.Ext
	}
	return ""
}

var File_GeneralReward_proto protoreflect.FileDescriptor

var file_GeneralReward_proto_rawDesc = []byte{
	0x0a, 0x13, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0xbb, 0x03, 0x0a, 0x0d, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x3e, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x12, 0x3e, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07,
	0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x65, 0x72, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x78, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x78, 0x74, 0x1a, 0x3a, 0x0a, 0x0c, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3a, 0x0a, 0x0c, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x3a, 0x0a, 0x0c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0b, 0x5a,
	0x09, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_GeneralReward_proto_rawDescOnce sync.Once
	file_GeneralReward_proto_rawDescData = file_GeneralReward_proto_rawDesc
)

func file_GeneralReward_proto_rawDescGZIP() []byte {
	file_GeneralReward_proto_rawDescOnce.Do(func() {
		file_GeneralReward_proto_rawDescData = protoimpl.X.CompressGZIP(file_GeneralReward_proto_rawDescData)
	})
	return file_GeneralReward_proto_rawDescData
}

var file_GeneralReward_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_GeneralReward_proto_goTypes = []interface{}{
	(*GeneralReward)(nil), // 0: response.GeneralReward
	nil,                   // 1: response.GeneralReward.ChangesEntry
	nil,                   // 2: response.GeneralReward.BalanceEntry
	nil,                   // 3: response.GeneralReward.CounterEntry
}
var file_GeneralReward_proto_depIdxs = []int32{
	1, // 0: response.GeneralReward.changes:type_name -> response.GeneralReward.ChangesEntry
	2, // 1: response.GeneralReward.balance:type_name -> response.GeneralReward.BalanceEntry
	3, // 2: response.GeneralReward.counter:type_name -> response.GeneralReward.CounterEntry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_GeneralReward_proto_init() }
func file_GeneralReward_proto_init() {
	if File_GeneralReward_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GeneralReward_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeneralReward); i {
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
			RawDescriptor: file_GeneralReward_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GeneralReward_proto_goTypes,
		DependencyIndexes: file_GeneralReward_proto_depIdxs,
		MessageInfos:      file_GeneralReward_proto_msgTypes,
	}.Build()
	File_GeneralReward_proto = out.File
	file_GeneralReward_proto_rawDesc = nil
	file_GeneralReward_proto_goTypes = nil
	file_GeneralReward_proto_depIdxs = nil
}
