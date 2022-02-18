// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: hello.proto

package hello

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// The request message containing the user's name.
type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Hey  *HelloRequest_Hey `protobuf:"bytes,2,opt,name=hey,proto3" json:"hey,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hello_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_hello_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HelloRequest) GetHey() *HelloRequest_Hey {
	if x != nil {
		return x.Hey
	}
	return nil
}

// The response message containing the greetings
type HelloReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	First  string `protobuf:"bytes,1,opt,name=first,proto3" json:"first,omitempty"`
	Second string `protobuf:"bytes,2,opt,name=second,proto3" json:"second,omitempty"`
}

func (x *HelloReply) Reset() {
	*x = HelloReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReply) ProtoMessage() {}

func (x *HelloReply) ProtoReflect() protoreflect.Message {
	mi := &file_hello_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReply.ProtoReflect.Descriptor instead.
func (*HelloReply) Descriptor() ([]byte, []int) {
	return file_hello_proto_rawDescGZIP(), []int{1}
}

func (x *HelloReply) GetFirst() string {
	if x != nil {
		return x.First
	}
	return ""
}

func (x *HelloReply) GetSecond() string {
	if x != nil {
		return x.Second
	}
	return ""
}

type HelloRequest_Hey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//	*HelloRequest_Hey_You_
	//	*HelloRequest_Hey_Me_
	Value isHelloRequest_Hey_Value `protobuf_oneof:"value"`
}

func (x *HelloRequest_Hey) Reset() {
	*x = HelloRequest_Hey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest_Hey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest_Hey) ProtoMessage() {}

func (x *HelloRequest_Hey) ProtoReflect() protoreflect.Message {
	mi := &file_hello_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest_Hey.ProtoReflect.Descriptor instead.
func (*HelloRequest_Hey) Descriptor() ([]byte, []int) {
	return file_hello_proto_rawDescGZIP(), []int{0, 0}
}

func (m *HelloRequest_Hey) GetValue() isHelloRequest_Hey_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *HelloRequest_Hey) GetYou() *HelloRequest_Hey_You {
	if x, ok := x.GetValue().(*HelloRequest_Hey_You_); ok {
		return x.You
	}
	return nil
}

func (x *HelloRequest_Hey) GetMe() *HelloRequest_Hey_Me {
	if x, ok := x.GetValue().(*HelloRequest_Hey_Me_); ok {
		return x.Me
	}
	return nil
}

type isHelloRequest_Hey_Value interface {
	isHelloRequest_Hey_Value()
}

type HelloRequest_Hey_You_ struct {
	You *HelloRequest_Hey_You `protobuf:"bytes,1,opt,name=you,proto3,oneof"`
}

type HelloRequest_Hey_Me_ struct {
	Me *HelloRequest_Hey_Me `protobuf:"bytes,2,opt,name=me,proto3,oneof"`
}

func (*HelloRequest_Hey_You_) isHelloRequest_Hey_Value() {}

func (*HelloRequest_Hey_Me_) isHelloRequest_Hey_Value() {}

type HelloRequest_Hey_You struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Abc string `protobuf:"bytes,1,opt,name=abc,proto3" json:"abc,omitempty"`
}

func (x *HelloRequest_Hey_You) Reset() {
	*x = HelloRequest_Hey_You{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest_Hey_You) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest_Hey_You) ProtoMessage() {}

func (x *HelloRequest_Hey_You) ProtoReflect() protoreflect.Message {
	mi := &file_hello_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest_Hey_You.ProtoReflect.Descriptor instead.
func (*HelloRequest_Hey_You) Descriptor() ([]byte, []int) {
	return file_hello_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *HelloRequest_Hey_You) GetAbc() string {
	if x != nil {
		return x.Abc
	}
	return ""
}

type HelloRequest_Hey_Me struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Def int32 `protobuf:"varint,1,opt,name=def,proto3" json:"def,omitempty"`
}

func (x *HelloRequest_Hey_Me) Reset() {
	*x = HelloRequest_Hey_Me{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest_Hey_Me) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest_Hey_Me) ProtoMessage() {}

func (x *HelloRequest_Hey_Me) ProtoReflect() protoreflect.Message {
	mi := &file_hello_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest_Hey_Me.ProtoReflect.Descriptor instead.
func (*HelloRequest_Hey_Me) Descriptor() ([]byte, []int) {
	return file_hello_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (x *HelloRequest_Hey_Me) GetDef() int32 {
	if x != nil {
		return x.Def
	}
	return 0
}

var File_hello_proto protoreflect.FileDescriptor

var file_hello_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xee, 0x01, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x03, 0x68, 0x65, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x48, 0x65, 0x79, 0x52, 0x03, 0x68,
	0x65, 0x79, 0x1a, 0x9e, 0x01, 0x0a, 0x03, 0x48, 0x65, 0x79, 0x12, 0x2f, 0x0a, 0x03, 0x79, 0x6f,
	0x75, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x48, 0x65, 0x79,
	0x2e, 0x59, 0x6f, 0x75, 0x48, 0x00, 0x52, 0x03, 0x79, 0x6f, 0x75, 0x12, 0x2c, 0x0a, 0x02, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x48, 0x65, 0x79,
	0x2e, 0x4d, 0x65, 0x48, 0x00, 0x52, 0x02, 0x6d, 0x65, 0x1a, 0x17, 0x0a, 0x03, 0x59, 0x6f, 0x75,
	0x12, 0x10, 0x0a, 0x03, 0x61, 0x62, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61,
	0x62, 0x63, 0x1a, 0x16, 0x0a, 0x02, 0x4d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x65, 0x66, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x64, 0x65, 0x66, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x3a, 0x0a, 0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x72, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x66, 0x69, 0x72, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x6f, 0x6e,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x32,
	0x4b, 0x0a, 0x07, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x12, 0x40, 0x0a, 0x08, 0x53, 0x61,
	0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x13, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x0c,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x06, 0x12, 0x04, 0x2f, 0x73, 0x61, 0x79, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hello_proto_rawDescOnce sync.Once
	file_hello_proto_rawDescData = file_hello_proto_rawDesc
)

func file_hello_proto_rawDescGZIP() []byte {
	file_hello_proto_rawDescOnce.Do(func() {
		file_hello_proto_rawDescData = protoimpl.X.CompressGZIP(file_hello_proto_rawDescData)
	})
	return file_hello_proto_rawDescData
}

var file_hello_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_hello_proto_goTypes = []interface{}{
	(*HelloRequest)(nil),         // 0: hello.HelloRequest
	(*HelloReply)(nil),           // 1: hello.HelloReply
	(*HelloRequest_Hey)(nil),     // 2: hello.HelloRequest.Hey
	(*HelloRequest_Hey_You)(nil), // 3: hello.HelloRequest.Hey.You
	(*HelloRequest_Hey_Me)(nil),  // 4: hello.HelloRequest.Hey.Me
}
var file_hello_proto_depIdxs = []int32{
	2, // 0: hello.HelloRequest.hey:type_name -> hello.HelloRequest.Hey
	3, // 1: hello.HelloRequest.Hey.you:type_name -> hello.HelloRequest.Hey.You
	4, // 2: hello.HelloRequest.Hey.me:type_name -> hello.HelloRequest.Hey.Me
	0, // 3: hello.Greeter.SayHello:input_type -> hello.HelloRequest
	1, // 4: hello.Greeter.SayHello:output_type -> hello.HelloReply
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_hello_proto_init() }
func file_hello_proto_init() {
	if File_hello_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hello_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_hello_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReply); i {
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
		file_hello_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest_Hey); i {
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
		file_hello_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest_Hey_You); i {
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
		file_hello_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest_Hey_Me); i {
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
	file_hello_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*HelloRequest_Hey_You_)(nil),
		(*HelloRequest_Hey_Me_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_hello_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hello_proto_goTypes,
		DependencyIndexes: file_hello_proto_depIdxs,
		MessageInfos:      file_hello_proto_msgTypes,
	}.Build()
	File_hello_proto = out.File
	file_hello_proto_rawDesc = nil
	file_hello_proto_goTypes = nil
	file_hello_proto_depIdxs = nil
}