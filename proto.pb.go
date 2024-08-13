// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.27.2
// source: proto.proto

package rpcproto

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

// 系统中所有的消息交互都以此为包装
type RpcCmd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// rpc command code
	Cmd int32 `protobuf:"varint,1,opt,name=Cmd,proto3" json:"Cmd,omitempty"`
	// sequence no
	SeqNo int32 `protobuf:"varint,2,opt,name=SeqNo,proto3" json:"SeqNo,omitempty"`
	// session id
	Sid string `protobuf:"bytes,3,opt,name=Sid,proto3" json:"Sid,omitempty"`
	// call id
	CallId string `protobuf:"bytes,4,opt,name=CallId,proto3" json:"CallId,omitempty"`
	// response code, ref http status code,200=ok,400-599=error
	Res int32 `protobuf:"varint,5,opt,name=Res,proto3" json:"Res,omitempty"`
	// command body
	Body []byte `protobuf:"bytes,6,opt,name=Body,proto3" json:"Body,omitempty"`
	// optional str parameter,额外的信息,一般不会有,有些特殊命令里面可能用到
	ParaStr string `protobuf:"bytes,7,opt,name=ParaStr,proto3" json:"ParaStr,omitempty"`
	// optional binary parameter,额外的信息,一般不会有,有些特殊命令里面可能用到
	ParaBin []byte `protobuf:"bytes,8,opt,name=ParaBin,proto3" json:"ParaBin,omitempty"`
	// optional int64 parameter
	ParaInt int64 `protobuf:"varint,9,opt,name=ParaInt,proto3" json:"ParaInt,omitempty"`
	// optional headers,one key followed by one value,Headers[0]=key,Headers[1]=value,Headers[2]=key,Headers[3]=value ...
	Headers []string `protobuf:"bytes,10,rep,name=Headers,proto3" json:"Headers,omitempty"`
}

func (x *RpcCmd) Reset() {
	*x = RpcCmd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcCmd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcCmd) ProtoMessage() {}

func (x *RpcCmd) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpcCmd.ProtoReflect.Descriptor instead.
func (*RpcCmd) Descriptor() ([]byte, []int) {
	return file_proto_proto_rawDescGZIP(), []int{0}
}

func (x *RpcCmd) GetCmd() int32 {
	if x != nil {
		return x.Cmd
	}
	return 0
}

func (x *RpcCmd) GetSeqNo() int32 {
	if x != nil {
		return x.SeqNo
	}
	return 0
}

func (x *RpcCmd) GetSid() string {
	if x != nil {
		return x.Sid
	}
	return ""
}

func (x *RpcCmd) GetCallId() string {
	if x != nil {
		return x.CallId
	}
	return ""
}

func (x *RpcCmd) GetRes() int32 {
	if x != nil {
		return x.Res
	}
	return 0
}

func (x *RpcCmd) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *RpcCmd) GetParaStr() string {
	if x != nil {
		return x.ParaStr
	}
	return ""
}

func (x *RpcCmd) GetParaBin() []byte {
	if x != nil {
		return x.ParaBin
	}
	return nil
}

func (x *RpcCmd) GetParaInt() int64 {
	if x != nil {
		return x.ParaInt
	}
	return 0
}

func (x *RpcCmd) GetHeaders() []string {
	if x != nil {
		return x.Headers
	}
	return nil
}

var File_proto_proto protoreflect.FileDescriptor

var file_proto_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x72,
	0x70, 0x63, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe8, 0x01, 0x0a, 0x06, 0x52, 0x70, 0x63, 0x43,
	0x6d, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x43, 0x6d, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x43, 0x6d, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x65, 0x71, 0x4e, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x53, 0x65, 0x71, 0x4e, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x53, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x43, 0x61, 0x6c, 0x6c, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x43, 0x61,
	0x6c, 0x6c, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x52, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x52, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x61,
	0x72, 0x61, 0x53, 0x74, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x50, 0x61, 0x72,
	0x61, 0x53, 0x74, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x61, 0x72, 0x61, 0x42, 0x69, 0x6e, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x50, 0x61, 0x72, 0x61, 0x42, 0x69, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x50, 0x61, 0x72, 0x61, 0x49, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x50, 0x61, 0x72, 0x61, 0x49, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x42, 0x1d, 0x48, 0x03, 0x5a, 0x19, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x79, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x70, 0x63, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_proto_rawDescOnce sync.Once
	file_proto_proto_rawDescData = file_proto_proto_rawDesc
)

func file_proto_proto_rawDescGZIP() []byte {
	file_proto_proto_rawDescOnce.Do(func() {
		file_proto_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_proto_rawDescData)
	})
	return file_proto_proto_rawDescData
}

var file_proto_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_proto_goTypes = []interface{}{
	(*RpcCmd)(nil), // 0: rpcproto.RpcCmd
}
var file_proto_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_proto_init() }
func file_proto_proto_init() {
	if File_proto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RpcCmd); i {
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
			RawDescriptor: file_proto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_proto_goTypes,
		DependencyIndexes: file_proto_proto_depIdxs,
		MessageInfos:      file_proto_proto_msgTypes,
	}.Build()
	File_proto_proto = out.File
	file_proto_proto_rawDesc = nil
	file_proto_proto_goTypes = nil
	file_proto_proto_depIdxs = nil
}
