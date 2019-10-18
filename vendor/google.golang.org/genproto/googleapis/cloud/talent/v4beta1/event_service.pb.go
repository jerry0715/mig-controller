// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/talent/v4beta1/event_service.proto

package talent

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

// The report event request.
type CreateClientEventRequest struct {
	// Required. Resource name of the tenant under which the event is created.
	//
	// The format is "projects/{project_id}/tenants/{tenant_id}", for example,
	// "projects/foo/tenant/bar". If tenant id is unspecified, a default tenant
	// is created, for example, "projects/foo".
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. Events issued when end user interacts with customer's application
	// that uses Cloud Talent Solution.
	ClientEvent          *ClientEvent `protobuf:"bytes,2,opt,name=client_event,json=clientEvent,proto3" json:"client_event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateClientEventRequest) Reset()         { *m = CreateClientEventRequest{} }
func (m *CreateClientEventRequest) String() string { return proto.CompactTextString(m) }
func (*CreateClientEventRequest) ProtoMessage()    {}
func (*CreateClientEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe732bea1e1e1bea, []int{0}
}

func (m *CreateClientEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateClientEventRequest.Unmarshal(m, b)
}
func (m *CreateClientEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateClientEventRequest.Marshal(b, m, deterministic)
}
func (m *CreateClientEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateClientEventRequest.Merge(m, src)
}
func (m *CreateClientEventRequest) XXX_Size() int {
	return xxx_messageInfo_CreateClientEventRequest.Size(m)
}
func (m *CreateClientEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateClientEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateClientEventRequest proto.InternalMessageInfo

func (m *CreateClientEventRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *CreateClientEventRequest) GetClientEvent() *ClientEvent {
	if m != nil {
		return m.ClientEvent
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateClientEventRequest)(nil), "google.cloud.talent.v4beta1.CreateClientEventRequest")
}

func init() {
	proto.RegisterFile("google/cloud/talent/v4beta1/event_service.proto", fileDescriptor_fe732bea1e1e1bea)
}

var fileDescriptor_fe732bea1e1e1bea = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0xce, 0xd2, 0x40,
	0x10, 0xc7, 0xd3, 0x7e, 0xc9, 0x97, 0xd8, 0x8f, 0x0b, 0xf5, 0x20, 0x29, 0x26, 0x90, 0x5e, 0x44,
	0xa2, 0xbb, 0x11, 0xf4, 0x82, 0xf1, 0x50, 0x88, 0x77, 0x02, 0x9e, 0xb8, 0x90, 0x6d, 0x19, 0xda,
	0x92, 0x65, 0xb7, 0xee, 0x0e, 0xe5, 0x60, 0x8c, 0x89, 0x37, 0xcf, 0xbe, 0x81, 0x8f, 0xe3, 0x51,
	0x5f, 0xc0, 0x44, 0x1e, 0xc4, 0xb4, 0xdb, 0x48, 0x23, 0x5a, 0xbc, 0xed, 0xee, 0xcc, 0x6f, 0xfe,
	0xff, 0x99, 0x1d, 0x87, 0xc6, 0x52, 0xc6, 0x1c, 0x68, 0xc4, 0xe5, 0x61, 0x43, 0x91, 0x71, 0x10,
	0x48, 0xf3, 0xe7, 0x21, 0x20, 0x7b, 0x46, 0x21, 0x07, 0x81, 0x6b, 0x0d, 0x2a, 0x4f, 0x23, 0x20,
	0x99, 0x92, 0x28, 0xdd, 0xae, 0x01, 0x48, 0x09, 0x10, 0x03, 0x90, 0x0a, 0xf0, 0x1e, 0x56, 0xd5,
	0x58, 0x96, 0x52, 0x26, 0x84, 0x44, 0x86, 0xa9, 0x14, 0xda, 0xa0, 0xde, 0x83, 0x5a, 0x34, 0xe2,
	0x69, 0x01, 0x9a, 0x40, 0xaf, 0x16, 0xd8, 0xa6, 0xc0, 0x37, 0xeb, 0x10, 0x12, 0x96, 0xa7, 0x52,
	0x55, 0x09, 0x8f, 0xae, 0xba, 0x34, 0x89, 0xfe, 0x27, 0xcb, 0xe9, 0xcc, 0x14, 0x30, 0x84, 0x59,
	0x29, 0xf0, 0xba, 0x88, 0x2d, 0xe0, 0xed, 0x01, 0x34, 0xba, 0x5d, 0xe7, 0x36, 0x63, 0x0a, 0x04,
	0x76, 0xac, 0xbe, 0x35, 0xb8, 0x37, 0xbd, 0xf9, 0x11, 0xd8, 0x8b, 0xea, 0xc9, 0x9d, 0x3b, 0x2d,
	0xe3, 0x69, 0x5d, 0xd6, 0xeb, 0xd8, 0x7d, 0x6b, 0x70, 0x37, 0x1a, 0x90, 0x86, 0x76, 0x49, 0x4d,
	0xc3, 0x14, 0xbb, 0x8b, 0xce, 0x2f, 0xa3, 0x9f, 0xb6, 0xd3, 0x2a, 0x4f, 0x4b, 0x33, 0x40, 0xf7,
	0x64, 0x39, 0xed, 0x0b, 0x73, 0xee, 0x8b, 0x66, 0x89, 0x7f, 0x34, 0xe3, 0xfd, 0xb7, 0x33, 0x5f,
	0x7c, 0xfc, 0x7e, 0xfa, 0x6c, 0x27, 0xfe, 0xf8, 0xf7, 0xc4, 0xde, 0x99, 0x9e, 0x5f, 0x65, 0x4a,
	0xee, 0x20, 0x42, 0x4d, 0x87, 0x14, 0x41, 0x30, 0x51, 0x9c, 0xde, 0xd3, 0x5a, 0x0f, 0x7a, 0x62,
	0x0d, 0x57, 0xc4, 0x7f, 0xdc, 0x40, 0x5e, 0xe4, 0x7b, 0xfc, 0x6b, 0x70, 0x7f, 0x27, 0x43, 0x5d,
	0x19, 0x64, 0x59, 0xaa, 0x49, 0x24, 0xf7, 0xdf, 0x82, 0x65, 0x82, 0x98, 0xe9, 0x09, 0xa5, 0xc7,
	0xe3, 0xf1, 0x8f, 0x20, 0x65, 0x07, 0x4c, 0xcc, 0xf7, 0x3e, 0xcd, 0x38, 0xc3, 0xad, 0x54, 0xfb,
	0x27, 0xd7, 0xd2, 0x0b, 0x91, 0xe9, 0x07, 0xa7, 0x17, 0xc9, 0x7d, 0xd3, 0x30, 0xa6, 0xed, 0xfa,
	0x2f, 0xcc, 0x8b, 0x3d, 0x99, 0x5b, 0xab, 0xa0, 0x22, 0x62, 0xc9, 0x99, 0x88, 0x89, 0x54, 0x31,
	0x8d, 0x41, 0x94, 0x5b, 0x44, 0xcf, 0x62, 0x7f, 0xdd, 0xb8, 0x97, 0xe6, 0xfa, 0xc5, 0xbe, 0x99,
	0xbd, 0x59, 0x86, 0xb7, 0x25, 0x33, 0xfe, 0x15, 0x00, 0x00, 0xff, 0xff, 0xfe, 0x38, 0x01, 0xc1,
	0x4a, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EventServiceClient is the client API for EventService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventServiceClient interface {
	// Report events issued when end user interacts with customer's application
	// that uses Cloud Talent Solution. You may inspect the created events in
	// [self service
	// tools](https://console.cloud.google.com/talent-solution/overview).
	// [Learn
	// more](https://cloud.google.com/talent-solution/docs/management-tools)
	// about self service tools.
	CreateClientEvent(ctx context.Context, in *CreateClientEventRequest, opts ...grpc.CallOption) (*ClientEvent, error)
}

type eventServiceClient struct {
	cc *grpc.ClientConn
}

func NewEventServiceClient(cc *grpc.ClientConn) EventServiceClient {
	return &eventServiceClient{cc}
}

func (c *eventServiceClient) CreateClientEvent(ctx context.Context, in *CreateClientEventRequest, opts ...grpc.CallOption) (*ClientEvent, error) {
	out := new(ClientEvent)
	err := c.cc.Invoke(ctx, "/google.cloud.talent.v4beta1.EventService/CreateClientEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventServiceServer is the server API for EventService service.
type EventServiceServer interface {
	// Report events issued when end user interacts with customer's application
	// that uses Cloud Talent Solution. You may inspect the created events in
	// [self service
	// tools](https://console.cloud.google.com/talent-solution/overview).
	// [Learn
	// more](https://cloud.google.com/talent-solution/docs/management-tools)
	// about self service tools.
	CreateClientEvent(context.Context, *CreateClientEventRequest) (*ClientEvent, error)
}

func RegisterEventServiceServer(s *grpc.Server, srv EventServiceServer) {
	s.RegisterService(&_EventService_serviceDesc, srv)
}

func _EventService_CreateClientEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClientEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).CreateClientEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.talent.v4beta1.EventService/CreateClientEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).CreateClientEvent(ctx, req.(*CreateClientEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.talent.v4beta1.EventService",
	HandlerType: (*EventServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateClientEvent",
			Handler:    _EventService_CreateClientEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/talent/v4beta1/event_service.proto",
}