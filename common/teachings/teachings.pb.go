// Code generated by protoc-gen-go. DO NOT EDIT.
// source: teachings/teachings.proto

package teachings

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type DefaultResponse_Status int32

const (
	DefaultResponse_UNKOWN                DefaultResponse_Status = 0
	DefaultResponse_OK                    DefaultResponse_Status = 200
	DefaultResponse_CREATED               DefaultResponse_Status = 201
	DefaultResponse_NO_CONTENT            DefaultResponse_Status = 204
	DefaultResponse_BAD_REQUEST           DefaultResponse_Status = 400
	DefaultResponse_UNAUTHORIZED          DefaultResponse_Status = 401
	DefaultResponse_FORBIDDEN             DefaultResponse_Status = 403
	DefaultResponse_NOT_FOUND             DefaultResponse_Status = 404
	DefaultResponse_METHOD_NOT_ALLOWED    DefaultResponse_Status = 405
	DefaultResponse_GONE                  DefaultResponse_Status = 410
	DefaultResponse_UNPROCESSABLE_ENTITY  DefaultResponse_Status = 422
	DefaultResponse_INTERNAL_SERVER_ERROR DefaultResponse_Status = 500
	DefaultResponse_BAD_GATEWAY           DefaultResponse_Status = 502
)

var DefaultResponse_Status_name = map[int32]string{
	0:   "UNKOWN",
	200: "OK",
	201: "CREATED",
	204: "NO_CONTENT",
	400: "BAD_REQUEST",
	401: "UNAUTHORIZED",
	403: "FORBIDDEN",
	404: "NOT_FOUND",
	405: "METHOD_NOT_ALLOWED",
	410: "GONE",
	422: "UNPROCESSABLE_ENTITY",
	500: "INTERNAL_SERVER_ERROR",
	502: "BAD_GATEWAY",
}

var DefaultResponse_Status_value = map[string]int32{
	"UNKOWN":                0,
	"OK":                    200,
	"CREATED":               201,
	"NO_CONTENT":            204,
	"BAD_REQUEST":           400,
	"UNAUTHORIZED":          401,
	"FORBIDDEN":             403,
	"NOT_FOUND":             404,
	"METHOD_NOT_ALLOWED":    405,
	"GONE":                  410,
	"UNPROCESSABLE_ENTITY":  422,
	"INTERNAL_SERVER_ERROR": 500,
	"BAD_GATEWAY":           502,
}

func (x DefaultResponse_Status) String() string {
	return proto.EnumName(DefaultResponse_Status_name, int32(x))
}

func (DefaultResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0c9aefcc6dd6cc46, []int{3, 0}
}

type TeachingRequest struct {
	CategorySlug         string   `protobuf:"bytes,1,opt,name=categorySlug,proto3" json:"categorySlug,omitempty"`
	Text                 []string `protobuf:"bytes,2,rep,name=text,proto3" json:"text,omitempty"`
	MediaBase64          []string `protobuf:"bytes,3,rep,name=mediaBase64,proto3" json:"mediaBase64,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeachingRequest) Reset()         { *m = TeachingRequest{} }
func (m *TeachingRequest) String() string { return proto.CompactTextString(m) }
func (*TeachingRequest) ProtoMessage()    {}
func (*TeachingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c9aefcc6dd6cc46, []int{0}
}

func (m *TeachingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeachingRequest.Unmarshal(m, b)
}
func (m *TeachingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeachingRequest.Marshal(b, m, deterministic)
}
func (m *TeachingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeachingRequest.Merge(m, src)
}
func (m *TeachingRequest) XXX_Size() int {
	return xxx_messageInfo_TeachingRequest.Size(m)
}
func (m *TeachingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TeachingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TeachingRequest proto.InternalMessageInfo

func (m *TeachingRequest) GetCategorySlug() string {
	if m != nil {
		return m.CategorySlug
	}
	return ""
}

func (m *TeachingRequest) GetText() []string {
	if m != nil {
		return m.Text
	}
	return nil
}

func (m *TeachingRequest) GetMediaBase64() []string {
	if m != nil {
		return m.MediaBase64
	}
	return nil
}

type TeachingSlug struct {
	Slug                 string   `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeachingSlug) Reset()         { *m = TeachingSlug{} }
func (m *TeachingSlug) String() string { return proto.CompactTextString(m) }
func (*TeachingSlug) ProtoMessage()    {}
func (*TeachingSlug) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c9aefcc6dd6cc46, []int{1}
}

func (m *TeachingSlug) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeachingSlug.Unmarshal(m, b)
}
func (m *TeachingSlug) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeachingSlug.Marshal(b, m, deterministic)
}
func (m *TeachingSlug) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeachingSlug.Merge(m, src)
}
func (m *TeachingSlug) XXX_Size() int {
	return xxx_messageInfo_TeachingSlug.Size(m)
}
func (m *TeachingSlug) XXX_DiscardUnknown() {
	xxx_messageInfo_TeachingSlug.DiscardUnknown(m)
}

var xxx_messageInfo_TeachingSlug proto.InternalMessageInfo

func (m *TeachingSlug) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

type EmptyPayload struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyPayload) Reset()         { *m = EmptyPayload{} }
func (m *EmptyPayload) String() string { return proto.CompactTextString(m) }
func (*EmptyPayload) ProtoMessage()    {}
func (*EmptyPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c9aefcc6dd6cc46, []int{2}
}

func (m *EmptyPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyPayload.Unmarshal(m, b)
}
func (m *EmptyPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyPayload.Marshal(b, m, deterministic)
}
func (m *EmptyPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyPayload.Merge(m, src)
}
func (m *EmptyPayload) XXX_Size() int {
	return xxx_messageInfo_EmptyPayload.Size(m)
}
func (m *EmptyPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyPayload.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyPayload proto.InternalMessageInfo

type DefaultResponse struct {
	Status               DefaultResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=teachings.DefaultResponse_Status" json:"status,omitempty"`
	Message              string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *DefaultResponse) Reset()         { *m = DefaultResponse{} }
func (m *DefaultResponse) String() string { return proto.CompactTextString(m) }
func (*DefaultResponse) ProtoMessage()    {}
func (*DefaultResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c9aefcc6dd6cc46, []int{3}
}

func (m *DefaultResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DefaultResponse.Unmarshal(m, b)
}
func (m *DefaultResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DefaultResponse.Marshal(b, m, deterministic)
}
func (m *DefaultResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DefaultResponse.Merge(m, src)
}
func (m *DefaultResponse) XXX_Size() int {
	return xxx_messageInfo_DefaultResponse.Size(m)
}
func (m *DefaultResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DefaultResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DefaultResponse proto.InternalMessageInfo

func (m *DefaultResponse) GetStatus() DefaultResponse_Status {
	if m != nil {
		return m.Status
	}
	return DefaultResponse_UNKOWN
}

func (m *DefaultResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Teaching struct {
	Slug                 string   `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug,omitempty"`
	CategorySlug         string   `protobuf:"bytes,2,opt,name=categorySlug,proto3" json:"categorySlug,omitempty"`
	Text                 []string `protobuf:"bytes,3,rep,name=text,proto3" json:"text,omitempty"`
	MediaBase64          []string `protobuf:"bytes,4,rep,name=mediaBase64,proto3" json:"mediaBase64,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Teaching) Reset()         { *m = Teaching{} }
func (m *Teaching) String() string { return proto.CompactTextString(m) }
func (*Teaching) ProtoMessage()    {}
func (*Teaching) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c9aefcc6dd6cc46, []int{4}
}

func (m *Teaching) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Teaching.Unmarshal(m, b)
}
func (m *Teaching) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Teaching.Marshal(b, m, deterministic)
}
func (m *Teaching) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Teaching.Merge(m, src)
}
func (m *Teaching) XXX_Size() int {
	return xxx_messageInfo_Teaching.Size(m)
}
func (m *Teaching) XXX_DiscardUnknown() {
	xxx_messageInfo_Teaching.DiscardUnknown(m)
}

var xxx_messageInfo_Teaching proto.InternalMessageInfo

func (m *Teaching) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *Teaching) GetCategorySlug() string {
	if m != nil {
		return m.CategorySlug
	}
	return ""
}

func (m *Teaching) GetText() []string {
	if m != nil {
		return m.Text
	}
	return nil
}

func (m *Teaching) GetMediaBase64() []string {
	if m != nil {
		return m.MediaBase64
	}
	return nil
}

func init() {
	proto.RegisterEnum("teachings.DefaultResponse_Status", DefaultResponse_Status_name, DefaultResponse_Status_value)
	proto.RegisterType((*TeachingRequest)(nil), "teachings.TeachingRequest")
	proto.RegisterType((*TeachingSlug)(nil), "teachings.TeachingSlug")
	proto.RegisterType((*EmptyPayload)(nil), "teachings.EmptyPayload")
	proto.RegisterType((*DefaultResponse)(nil), "teachings.DefaultResponse")
	proto.RegisterType((*Teaching)(nil), "teachings.Teaching")
}

func init() { proto.RegisterFile("teachings/teachings.proto", fileDescriptor_0c9aefcc6dd6cc46) }

var fileDescriptor_0c9aefcc6dd6cc46 = []byte{
	// 571 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xc1, 0x6e, 0xd3, 0x4c,
	0x10, 0xfe, 0x1d, 0x47, 0xe9, 0x9f, 0x89, 0x95, 0x2c, 0x03, 0xa8, 0xa9, 0x4f, 0xc5, 0xa7, 0x9c,
	0x4a, 0x55, 0x10, 0x12, 0x20, 0x81, 0x9c, 0x7a, 0x9b, 0x46, 0x0d, 0xbb, 0x65, 0xbd, 0xa6, 0x2a,
	0x17, 0xcb, 0x34, 0x4b, 0xa8, 0x48, 0x9b, 0x10, 0x6f, 0x50, 0xf3, 0x16, 0x20, 0xe0, 0xc2, 0x03,
	0x20, 0xf1, 0x26, 0x20, 0x71, 0xe7, 0x2d, 0x38, 0x21, 0xce, 0x28, 0xa6, 0x6e, 0x9c, 0xd6, 0x41,
	0xa8, 0xb7, 0x99, 0xf1, 0xcc, 0x37, 0xdf, 0x7e, 0xf3, 0xc9, 0xb0, 0xa2, 0x55, 0x74, 0xf0, 0xe2,
	0xf0, 0xb8, 0x17, 0xdf, 0x3c, 0x8b, 0xd6, 0x86, 0xa3, 0x81, 0x1e, 0x60, 0xf9, 0xac, 0xe0, 0xbc,
	0x84, 0x9a, 0x3c, 0x4d, 0x84, 0x7a, 0x35, 0x56, 0xb1, 0x46, 0x07, 0xac, 0x83, 0x48, 0xab, 0xde,
	0x60, 0x34, 0xf1, 0xfb, 0xe3, 0x5e, 0xdd, 0x58, 0x35, 0x1a, 0x65, 0x31, 0x57, 0x43, 0x84, 0xa2,
	0x56, 0x27, 0xba, 0x5e, 0x58, 0x35, 0x1b, 0x65, 0x91, 0xc4, 0xb8, 0x0a, 0x95, 0x23, 0xd5, 0x3d,
	0x8c, 0x9a, 0x51, 0xac, 0xee, 0xdc, 0xae, 0x9b, 0xc9, 0xa7, 0x6c, 0xc9, 0x71, 0xc0, 0x4a, 0x97,
	0xa5, 0x28, 0xf1, 0x6c, 0x43, 0x12, 0x3b, 0x55, 0xb0, 0xe8, 0xd1, 0x50, 0x4f, 0x76, 0xa3, 0x49,
	0x7f, 0x10, 0x75, 0x9d, 0xef, 0x05, 0xa8, 0x79, 0xea, 0x79, 0x34, 0xee, 0x6b, 0xa1, 0xe2, 0xe1,
	0xe0, 0x38, 0x56, 0x78, 0x17, 0x4a, 0xb1, 0x8e, 0xf4, 0x38, 0x4e, 0x26, 0xab, 0x1b, 0x37, 0xd6,
	0x66, 0x2f, 0x3c, 0xd7, 0xbb, 0xe6, 0x27, 0x8d, 0xe2, 0x74, 0x00, 0xeb, 0xb0, 0x74, 0xa4, 0xe2,
	0x38, 0xea, 0xa9, 0x7a, 0x21, 0xd9, 0x9a, 0xa6, 0xce, 0x0f, 0x03, 0x4a, 0x7f, 0x9a, 0x11, 0xa0,
	0x14, 0xb0, 0x1d, 0xbe, 0xc7, 0xc8, 0x7f, 0xb8, 0x04, 0x05, 0xbe, 0x43, 0xbe, 0x18, 0x68, 0xc1,
	0xd2, 0xa6, 0xa0, 0xae, 0xa4, 0x1e, 0xf9, 0x6a, 0x60, 0x0d, 0x80, 0xf1, 0x70, 0x93, 0x33, 0x49,
	0x99, 0x24, 0xdf, 0x0c, 0x24, 0x50, 0x69, 0xba, 0x5e, 0x28, 0xe8, 0xe3, 0x80, 0xfa, 0x92, 0xbc,
	0x31, 0xf1, 0x0a, 0x58, 0x01, 0x73, 0x03, 0xb9, 0xcd, 0x45, 0xfb, 0x29, 0xf5, 0xc8, 0x5b, 0x13,
	0xab, 0x50, 0xde, 0xe2, 0xa2, 0xd9, 0xf6, 0x3c, 0xca, 0xc8, 0xbb, 0x24, 0x67, 0x5c, 0x86, 0x5b,
	0x3c, 0x60, 0x1e, 0x79, 0x6f, 0xe2, 0x32, 0xe0, 0x23, 0x2a, 0xb7, 0xb9, 0x17, 0x4e, 0xcb, 0x6e,
	0xa7, 0xc3, 0xf7, 0xa8, 0x47, 0x3e, 0x98, 0x58, 0x86, 0x62, 0x8b, 0x33, 0x4a, 0x3e, 0x9a, 0xb8,
	0x02, 0xd7, 0x02, 0xb6, 0x2b, 0xf8, 0x26, 0xf5, 0x7d, 0xb7, 0xd9, 0xa1, 0x21, 0x65, 0xb2, 0x2d,
	0xf7, 0xc9, 0x27, 0x13, 0x6d, 0xb8, 0xde, 0x66, 0x92, 0x0a, 0xe6, 0x76, 0x42, 0x9f, 0x8a, 0x27,
	0x54, 0x84, 0x54, 0x08, 0x2e, 0xc8, 0x4f, 0x33, 0xe5, 0xd7, 0x72, 0x25, 0xdd, 0x73, 0xf7, 0xc9,
	0x2f, 0xd3, 0x39, 0x81, 0xff, 0xd3, 0x6b, 0xe4, 0x5d, 0xe2, 0x82, 0x0f, 0x0a, 0x7f, 0xf1, 0x81,
	0xb9, 0xd8, 0x07, 0xc5, 0x0b, 0x3e, 0xd8, 0xf8, 0x5c, 0x9c, 0xb9, 0xce, 0x57, 0xa3, 0xd7, 0x87,
	0x07, 0x0a, 0x5b, 0x60, 0xb9, 0xdd, 0x6e, 0x5a, 0x8d, 0xd1, 0xce, 0xdc, 0xf4, 0x9c, 0x43, 0xed,
	0xe5, 0x9c, 0x6f, 0x53, 0x3a, 0x0d, 0x63, 0xdd, 0xc0, 0x07, 0x60, 0xb5, 0x94, 0x9e, 0x01, 0x65,
	0x9b, 0xb3, 0xce, 0xb2, 0xaf, 0xe6, 0xa0, 0xac, 0x1b, 0xf8, 0x10, 0x6a, 0xc1, 0xb0, 0x1b, 0x69,
	0x35, 0x83, 0xc8, 0xeb, 0xcc, 0x1d, 0x4f, 0x08, 0x6c, 0x4f, 0x0d, 0xdb, 0x57, 0x59, 0x80, 0x45,
	0x84, 0x6d, 0x7b, 0xb1, 0x73, 0x1b, 0x06, 0x36, 0xa1, 0x92, 0xd1, 0xe4, 0x52, 0x92, 0xe0, 0x7d,
	0xa8, 0x64, 0xe4, 0x58, 0xcc, 0x24, 0xef, 0x39, 0x78, 0x0f, 0xaa, 0xf3, 0x5a, 0xfc, 0xbb, 0x14,
	0x48, 0xa1, 0x3a, 0x2f, 0xc3, 0xa5, 0x54, 0x78, 0x56, 0x4a, 0x7e, 0x59, 0xb7, 0x7e, 0x07, 0x00,
	0x00, 0xff, 0xff, 0x08, 0xa0, 0xe2, 0x59, 0xcf, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TeachingServiceClient is the client API for TeachingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TeachingServiceClient interface {
	AddTeachings(ctx context.Context, opts ...grpc.CallOption) (TeachingService_AddTeachingsClient, error)
	GetTeachings(ctx context.Context, in *EmptyPayload, opts ...grpc.CallOption) (TeachingService_GetTeachingsClient, error)
	UpdateTeachings(ctx context.Context, opts ...grpc.CallOption) (TeachingService_UpdateTeachingsClient, error)
	DeleteTeachings(ctx context.Context, opts ...grpc.CallOption) (TeachingService_DeleteTeachingsClient, error)
	AddTeaching(ctx context.Context, in *TeachingRequest, opts ...grpc.CallOption) (*TeachingSlug, error)
	GetTeaching(ctx context.Context, in *TeachingSlug, opts ...grpc.CallOption) (*Teaching, error)
	UpdateTeaching(ctx context.Context, in *Teaching, opts ...grpc.CallOption) (*Teaching, error)
	DeleteTeaching(ctx context.Context, in *TeachingSlug, opts ...grpc.CallOption) (*DefaultResponse, error)
}

type teachingServiceClient struct {
	cc *grpc.ClientConn
}

func NewTeachingServiceClient(cc *grpc.ClientConn) TeachingServiceClient {
	return &teachingServiceClient{cc}
}

func (c *teachingServiceClient) AddTeachings(ctx context.Context, opts ...grpc.CallOption) (TeachingService_AddTeachingsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TeachingService_serviceDesc.Streams[0], "/teachings.TeachingService/AddTeachings", opts...)
	if err != nil {
		return nil, err
	}
	x := &teachingServiceAddTeachingsClient{stream}
	return x, nil
}

type TeachingService_AddTeachingsClient interface {
	Send(*TeachingRequest) error
	Recv() (*TeachingSlug, error)
	grpc.ClientStream
}

type teachingServiceAddTeachingsClient struct {
	grpc.ClientStream
}

func (x *teachingServiceAddTeachingsClient) Send(m *TeachingRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *teachingServiceAddTeachingsClient) Recv() (*TeachingSlug, error) {
	m := new(TeachingSlug)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *teachingServiceClient) GetTeachings(ctx context.Context, in *EmptyPayload, opts ...grpc.CallOption) (TeachingService_GetTeachingsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TeachingService_serviceDesc.Streams[1], "/teachings.TeachingService/GetTeachings", opts...)
	if err != nil {
		return nil, err
	}
	x := &teachingServiceGetTeachingsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TeachingService_GetTeachingsClient interface {
	Recv() (*Teaching, error)
	grpc.ClientStream
}

type teachingServiceGetTeachingsClient struct {
	grpc.ClientStream
}

func (x *teachingServiceGetTeachingsClient) Recv() (*Teaching, error) {
	m := new(Teaching)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *teachingServiceClient) UpdateTeachings(ctx context.Context, opts ...grpc.CallOption) (TeachingService_UpdateTeachingsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TeachingService_serviceDesc.Streams[2], "/teachings.TeachingService/UpdateTeachings", opts...)
	if err != nil {
		return nil, err
	}
	x := &teachingServiceUpdateTeachingsClient{stream}
	return x, nil
}

type TeachingService_UpdateTeachingsClient interface {
	Send(*Teaching) error
	Recv() (*Teaching, error)
	grpc.ClientStream
}

type teachingServiceUpdateTeachingsClient struct {
	grpc.ClientStream
}

func (x *teachingServiceUpdateTeachingsClient) Send(m *Teaching) error {
	return x.ClientStream.SendMsg(m)
}

func (x *teachingServiceUpdateTeachingsClient) Recv() (*Teaching, error) {
	m := new(Teaching)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *teachingServiceClient) DeleteTeachings(ctx context.Context, opts ...grpc.CallOption) (TeachingService_DeleteTeachingsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TeachingService_serviceDesc.Streams[3], "/teachings.TeachingService/DeleteTeachings", opts...)
	if err != nil {
		return nil, err
	}
	x := &teachingServiceDeleteTeachingsClient{stream}
	return x, nil
}

type TeachingService_DeleteTeachingsClient interface {
	Send(*TeachingSlug) error
	CloseAndRecv() (*DefaultResponse, error)
	grpc.ClientStream
}

type teachingServiceDeleteTeachingsClient struct {
	grpc.ClientStream
}

func (x *teachingServiceDeleteTeachingsClient) Send(m *TeachingSlug) error {
	return x.ClientStream.SendMsg(m)
}

func (x *teachingServiceDeleteTeachingsClient) CloseAndRecv() (*DefaultResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(DefaultResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *teachingServiceClient) AddTeaching(ctx context.Context, in *TeachingRequest, opts ...grpc.CallOption) (*TeachingSlug, error) {
	out := new(TeachingSlug)
	err := c.cc.Invoke(ctx, "/teachings.TeachingService/AddTeaching", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teachingServiceClient) GetTeaching(ctx context.Context, in *TeachingSlug, opts ...grpc.CallOption) (*Teaching, error) {
	out := new(Teaching)
	err := c.cc.Invoke(ctx, "/teachings.TeachingService/GetTeaching", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teachingServiceClient) UpdateTeaching(ctx context.Context, in *Teaching, opts ...grpc.CallOption) (*Teaching, error) {
	out := new(Teaching)
	err := c.cc.Invoke(ctx, "/teachings.TeachingService/UpdateTeaching", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teachingServiceClient) DeleteTeaching(ctx context.Context, in *TeachingSlug, opts ...grpc.CallOption) (*DefaultResponse, error) {
	out := new(DefaultResponse)
	err := c.cc.Invoke(ctx, "/teachings.TeachingService/DeleteTeaching", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TeachingServiceServer is the server API for TeachingService service.
type TeachingServiceServer interface {
	AddTeachings(TeachingService_AddTeachingsServer) error
	GetTeachings(*EmptyPayload, TeachingService_GetTeachingsServer) error
	UpdateTeachings(TeachingService_UpdateTeachingsServer) error
	DeleteTeachings(TeachingService_DeleteTeachingsServer) error
	AddTeaching(context.Context, *TeachingRequest) (*TeachingSlug, error)
	GetTeaching(context.Context, *TeachingSlug) (*Teaching, error)
	UpdateTeaching(context.Context, *Teaching) (*Teaching, error)
	DeleteTeaching(context.Context, *TeachingSlug) (*DefaultResponse, error)
}

func RegisterTeachingServiceServer(s *grpc.Server, srv TeachingServiceServer) {
	s.RegisterService(&_TeachingService_serviceDesc, srv)
}

func _TeachingService_AddTeachings_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TeachingServiceServer).AddTeachings(&teachingServiceAddTeachingsServer{stream})
}

type TeachingService_AddTeachingsServer interface {
	Send(*TeachingSlug) error
	Recv() (*TeachingRequest, error)
	grpc.ServerStream
}

type teachingServiceAddTeachingsServer struct {
	grpc.ServerStream
}

func (x *teachingServiceAddTeachingsServer) Send(m *TeachingSlug) error {
	return x.ServerStream.SendMsg(m)
}

func (x *teachingServiceAddTeachingsServer) Recv() (*TeachingRequest, error) {
	m := new(TeachingRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TeachingService_GetTeachings_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EmptyPayload)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TeachingServiceServer).GetTeachings(m, &teachingServiceGetTeachingsServer{stream})
}

type TeachingService_GetTeachingsServer interface {
	Send(*Teaching) error
	grpc.ServerStream
}

type teachingServiceGetTeachingsServer struct {
	grpc.ServerStream
}

func (x *teachingServiceGetTeachingsServer) Send(m *Teaching) error {
	return x.ServerStream.SendMsg(m)
}

func _TeachingService_UpdateTeachings_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TeachingServiceServer).UpdateTeachings(&teachingServiceUpdateTeachingsServer{stream})
}

type TeachingService_UpdateTeachingsServer interface {
	Send(*Teaching) error
	Recv() (*Teaching, error)
	grpc.ServerStream
}

type teachingServiceUpdateTeachingsServer struct {
	grpc.ServerStream
}

func (x *teachingServiceUpdateTeachingsServer) Send(m *Teaching) error {
	return x.ServerStream.SendMsg(m)
}

func (x *teachingServiceUpdateTeachingsServer) Recv() (*Teaching, error) {
	m := new(Teaching)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TeachingService_DeleteTeachings_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TeachingServiceServer).DeleteTeachings(&teachingServiceDeleteTeachingsServer{stream})
}

type TeachingService_DeleteTeachingsServer interface {
	SendAndClose(*DefaultResponse) error
	Recv() (*TeachingSlug, error)
	grpc.ServerStream
}

type teachingServiceDeleteTeachingsServer struct {
	grpc.ServerStream
}

func (x *teachingServiceDeleteTeachingsServer) SendAndClose(m *DefaultResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *teachingServiceDeleteTeachingsServer) Recv() (*TeachingSlug, error) {
	m := new(TeachingSlug)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TeachingService_AddTeaching_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeachingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeachingServiceServer).AddTeaching(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/teachings.TeachingService/AddTeaching",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeachingServiceServer).AddTeaching(ctx, req.(*TeachingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeachingService_GetTeaching_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeachingSlug)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeachingServiceServer).GetTeaching(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/teachings.TeachingService/GetTeaching",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeachingServiceServer).GetTeaching(ctx, req.(*TeachingSlug))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeachingService_UpdateTeaching_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Teaching)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeachingServiceServer).UpdateTeaching(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/teachings.TeachingService/UpdateTeaching",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeachingServiceServer).UpdateTeaching(ctx, req.(*Teaching))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeachingService_DeleteTeaching_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeachingSlug)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeachingServiceServer).DeleteTeaching(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/teachings.TeachingService/DeleteTeaching",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeachingServiceServer).DeleteTeaching(ctx, req.(*TeachingSlug))
	}
	return interceptor(ctx, in, info, handler)
}

var _TeachingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "teachings.TeachingService",
	HandlerType: (*TeachingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddTeaching",
			Handler:    _TeachingService_AddTeaching_Handler,
		},
		{
			MethodName: "GetTeaching",
			Handler:    _TeachingService_GetTeaching_Handler,
		},
		{
			MethodName: "UpdateTeaching",
			Handler:    _TeachingService_UpdateTeaching_Handler,
		},
		{
			MethodName: "DeleteTeaching",
			Handler:    _TeachingService_DeleteTeaching_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AddTeachings",
			Handler:       _TeachingService_AddTeachings_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "GetTeachings",
			Handler:       _TeachingService_GetTeachings_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "UpdateTeachings",
			Handler:       _TeachingService_UpdateTeachings_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DeleteTeachings",
			Handler:       _TeachingService_DeleteTeachings_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "teachings/teachings.proto",
}
