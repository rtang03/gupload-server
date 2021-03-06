// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: service.proto

package core

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type StatusCode int32

const (
	StatusCode_Unknown StatusCode = 0
	StatusCode_Ok      StatusCode = 1
	StatusCode_Failed  StatusCode = 2
)

// Enum value maps for StatusCode.
var (
	StatusCode_name = map[int32]string{
		0: "Unknown",
		1: "Ok",
		2: "Failed",
	}
	StatusCode_value = map[string]int32{
		"Unknown": 0,
		"Ok":      1,
		"Failed":  2,
	}
)

func (x StatusCode) Enum() *StatusCode {
	p := new(StatusCode)
	*p = x
	return p
}

func (x StatusCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatusCode) Descriptor() protoreflect.EnumDescriptor {
	return file_service_proto_enumTypes[0].Descriptor()
}

func (StatusCode) Type() protoreflect.EnumType {
	return &file_service_proto_enumTypes[0]
}

func (x StatusCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatusCode.Descriptor instead.
func (StatusCode) EnumDescriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

type Chunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//	*Chunk_Content
	//	*Chunk_Info
	Data isChunk_Data `protobuf_oneof:"data"`
}

func (x *Chunk) Reset() {
	*x = Chunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunk) ProtoMessage() {}

func (x *Chunk) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chunk.ProtoReflect.Descriptor instead.
func (*Chunk) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

func (m *Chunk) GetData() isChunk_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *Chunk) GetContent() []byte {
	if x, ok := x.GetData().(*Chunk_Content); ok {
		return x.Content
	}
	return nil
}

func (x *Chunk) GetInfo() *UploadFileInfo {
	if x, ok := x.GetData().(*Chunk_Info); ok {
		return x.Info
	}
	return nil
}

type isChunk_Data interface {
	isChunk_Data()
}

type Chunk_Content struct {
	Content []byte `protobuf:"bytes,1,opt,name=Content,proto3,oneof"`
}

type Chunk_Info struct {
	Info *UploadFileInfo `protobuf:"bytes,2,opt,name=info,proto3,oneof"`
}

func (*Chunk_Content) isChunk_Data() {}

func (*Chunk_Info) isChunk_Data() {}

// Download
type FileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
}

func (x *FileRequest) Reset() {
	*x = FileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileRequest) ProtoMessage() {}

func (x *FileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileRequest.ProtoReflect.Descriptor instead.
func (*FileRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{1}
}

func (x *FileRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

type FileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shard []byte `protobuf:"bytes,1,opt,name=shard,proto3" json:"shard,omitempty"`
}

func (x *FileResponse) Reset() {
	*x = FileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileResponse) ProtoMessage() {}

func (x *FileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileResponse.ProtoReflect.Descriptor instead.
func (*FileResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{2}
}

func (x *FileResponse) GetShard() []byte {
	if x != nil {
		return x.Shard
	}
	return nil
}

// Upload
type UploadFileInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	FileType string `protobuf:"bytes,2,opt,name=fileType,proto3" json:"fileType,omitempty"`
}

func (x *UploadFileInfo) Reset() {
	*x = UploadFileInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadFileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFileInfo) ProtoMessage() {}

func (x *UploadFileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFileInfo.ProtoReflect.Descriptor instead.
func (*UploadFileInfo) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{3}
}

func (x *UploadFileInfo) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *UploadFileInfo) GetFileType() string {
	if x != nil {
		return x.FileType
	}
	return ""
}

type UploadStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string     `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	Code    StatusCode `protobuf:"varint,2,opt,name=Code,proto3,enum=StatusCode" json:"Code,omitempty"`
}

func (x *UploadStatus) Reset() {
	*x = UploadStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadStatus) ProtoMessage() {}

func (x *UploadStatus) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadStatus.ProtoReflect.Descriptor instead.
func (*UploadStatus) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{4}
}

func (x *UploadStatus) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UploadStatus) GetCode() StatusCode {
	if x != nil {
		return x.Code
	}
	return StatusCode_Unknown
}

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x52, 0x0a, 0x05, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x1a, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x07, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x42, 0x06, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x29, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x24,
	0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x68, 0x61, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x73,
	0x68, 0x61, 0x72, 0x64, 0x22, 0x48, 0x0a, 0x0e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69,
	0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x49,
	0x0a, 0x0c, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x2a, 0x2d, 0x0a, 0x0a, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f,
	0x77, 0x6e, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x6b, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06,
	0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x02, 0x32, 0x62, 0x0a, 0x0e, 0x47, 0x75, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x06, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x12, 0x06, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x1a, 0x0d, 0x2e, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x28, 0x01, 0x12,
	0x2b, 0x0a, 0x08, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x0c, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x25, 0x5a, 0x23,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x74, 0x61, 0x6e, 0x67,
	0x30, 0x33, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_proto_rawDescOnce sync.Once
	file_service_proto_rawDescData = file_service_proto_rawDesc
)

func file_service_proto_rawDescGZIP() []byte {
	file_service_proto_rawDescOnce.Do(func() {
		file_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_proto_rawDescData)
	})
	return file_service_proto_rawDescData
}

var file_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_service_proto_goTypes = []interface{}{
	(StatusCode)(0),        // 0: StatusCode
	(*Chunk)(nil),          // 1: Chunk
	(*FileRequest)(nil),    // 2: FileRequest
	(*FileResponse)(nil),   // 3: FileResponse
	(*UploadFileInfo)(nil), // 4: UploadFileInfo
	(*UploadStatus)(nil),   // 5: UploadStatus
}
var file_service_proto_depIdxs = []int32{
	4, // 0: Chunk.info:type_name -> UploadFileInfo
	0, // 1: UploadStatus.Code:type_name -> StatusCode
	1, // 2: GuploadService.Upload:input_type -> Chunk
	2, // 3: GuploadService.Download:input_type -> FileRequest
	5, // 4: GuploadService.Upload:output_type -> UploadStatus
	3, // 5: GuploadService.Download:output_type -> FileResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chunk); i {
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
		file_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileRequest); i {
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
		file_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileResponse); i {
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
		file_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadFileInfo); i {
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
		file_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadStatus); i {
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
	file_service_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Chunk_Content)(nil),
		(*Chunk_Info)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
		EnumInfos:         file_service_proto_enumTypes,
		MessageInfos:      file_service_proto_msgTypes,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GuploadServiceClient is the client API for GuploadService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GuploadServiceClient interface {
	Upload(ctx context.Context, opts ...grpc.CallOption) (GuploadService_UploadClient, error)
	Download(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (GuploadService_DownloadClient, error)
}

type guploadServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGuploadServiceClient(cc grpc.ClientConnInterface) GuploadServiceClient {
	return &guploadServiceClient{cc}
}

func (c *guploadServiceClient) Upload(ctx context.Context, opts ...grpc.CallOption) (GuploadService_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GuploadService_serviceDesc.Streams[0], "/GuploadService/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &guploadServiceUploadClient{stream}
	return x, nil
}

type GuploadService_UploadClient interface {
	Send(*Chunk) error
	CloseAndRecv() (*UploadStatus, error)
	grpc.ClientStream
}

type guploadServiceUploadClient struct {
	grpc.ClientStream
}

func (x *guploadServiceUploadClient) Send(m *Chunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *guploadServiceUploadClient) CloseAndRecv() (*UploadStatus, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *guploadServiceClient) Download(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (GuploadService_DownloadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GuploadService_serviceDesc.Streams[1], "/GuploadService/Download", opts...)
	if err != nil {
		return nil, err
	}
	x := &guploadServiceDownloadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GuploadService_DownloadClient interface {
	Recv() (*FileResponse, error)
	grpc.ClientStream
}

type guploadServiceDownloadClient struct {
	grpc.ClientStream
}

func (x *guploadServiceDownloadClient) Recv() (*FileResponse, error) {
	m := new(FileResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GuploadServiceServer is the server API for GuploadService service.
type GuploadServiceServer interface {
	Upload(GuploadService_UploadServer) error
	Download(*FileRequest, GuploadService_DownloadServer) error
}

// UnimplementedGuploadServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGuploadServiceServer struct {
}

func (*UnimplementedGuploadServiceServer) Upload(GuploadService_UploadServer) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (*UnimplementedGuploadServiceServer) Download(*FileRequest, GuploadService_DownloadServer) error {
	return status.Errorf(codes.Unimplemented, "method Download not implemented")
}

func RegisterGuploadServiceServer(s *grpc.Server, srv GuploadServiceServer) {
	s.RegisterService(&_GuploadService_serviceDesc, srv)
}

func _GuploadService_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GuploadServiceServer).Upload(&guploadServiceUploadServer{stream})
}

type GuploadService_UploadServer interface {
	SendAndClose(*UploadStatus) error
	Recv() (*Chunk, error)
	grpc.ServerStream
}

type guploadServiceUploadServer struct {
	grpc.ServerStream
}

func (x *guploadServiceUploadServer) SendAndClose(m *UploadStatus) error {
	return x.ServerStream.SendMsg(m)
}

func (x *guploadServiceUploadServer) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GuploadService_Download_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FileRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GuploadServiceServer).Download(m, &guploadServiceDownloadServer{stream})
}

type GuploadService_DownloadServer interface {
	Send(*FileResponse) error
	grpc.ServerStream
}

type guploadServiceDownloadServer struct {
	grpc.ServerStream
}

func (x *guploadServiceDownloadServer) Send(m *FileResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _GuploadService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "GuploadService",
	HandlerType: (*GuploadServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _GuploadService_Upload_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Download",
			Handler:       _GuploadService_Download_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "service.proto",
}
