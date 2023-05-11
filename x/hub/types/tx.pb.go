// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: hub/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgSendErasmusStudent struct {
	Creator          string         `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Port             string         `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
	ChannelID        string         `protobuf:"bytes,3,opt,name=channelID,proto3" json:"channelID,omitempty"`
	TimeoutTimestamp uint64         `protobuf:"varint,4,opt,name=timeoutTimestamp,proto3" json:"timeoutTimestamp,omitempty"`
	Student          *StoredStudent `protobuf:"bytes,5,opt,name=student,proto3" json:"student,omitempty"`
}

func (m *MsgSendErasmusStudent) Reset()         { *m = MsgSendErasmusStudent{} }
func (m *MsgSendErasmusStudent) String() string { return proto.CompactTextString(m) }
func (*MsgSendErasmusStudent) ProtoMessage()    {}
func (*MsgSendErasmusStudent) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b0dcb234192d0df, []int{0}
}
func (m *MsgSendErasmusStudent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSendErasmusStudent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSendErasmusStudent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSendErasmusStudent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSendErasmusStudent.Merge(m, src)
}
func (m *MsgSendErasmusStudent) XXX_Size() int {
	return m.Size()
}
func (m *MsgSendErasmusStudent) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSendErasmusStudent.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSendErasmusStudent proto.InternalMessageInfo

func (m *MsgSendErasmusStudent) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgSendErasmusStudent) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *MsgSendErasmusStudent) GetChannelID() string {
	if m != nil {
		return m.ChannelID
	}
	return ""
}

func (m *MsgSendErasmusStudent) GetTimeoutTimestamp() uint64 {
	if m != nil {
		return m.TimeoutTimestamp
	}
	return 0
}

func (m *MsgSendErasmusStudent) GetStudent() *StoredStudent {
	if m != nil {
		return m.Student
	}
	return nil
}

type MsgSendErasmusStudentResponse struct {
}

func (m *MsgSendErasmusStudentResponse) Reset()         { *m = MsgSendErasmusStudentResponse{} }
func (m *MsgSendErasmusStudentResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSendErasmusStudentResponse) ProtoMessage()    {}
func (*MsgSendErasmusStudentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b0dcb234192d0df, []int{1}
}
func (m *MsgSendErasmusStudentResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSendErasmusStudentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSendErasmusStudentResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSendErasmusStudentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSendErasmusStudentResponse.Merge(m, src)
}
func (m *MsgSendErasmusStudentResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSendErasmusStudentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSendErasmusStudentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSendErasmusStudentResponse proto.InternalMessageInfo

type MsgConfigureChain struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *MsgConfigureChain) Reset()         { *m = MsgConfigureChain{} }
func (m *MsgConfigureChain) String() string { return proto.CompactTextString(m) }
func (*MsgConfigureChain) ProtoMessage()    {}
func (*MsgConfigureChain) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b0dcb234192d0df, []int{2}
}
func (m *MsgConfigureChain) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgConfigureChain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgConfigureChain.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgConfigureChain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgConfigureChain.Merge(m, src)
}
func (m *MsgConfigureChain) XXX_Size() int {
	return m.Size()
}
func (m *MsgConfigureChain) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgConfigureChain.DiscardUnknown(m)
}

var xxx_messageInfo_MsgConfigureChain proto.InternalMessageInfo

func (m *MsgConfigureChain) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

type MsgConfigureChainResponse struct {
	Status int32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (m *MsgConfigureChainResponse) Reset()         { *m = MsgConfigureChainResponse{} }
func (m *MsgConfigureChainResponse) String() string { return proto.CompactTextString(m) }
func (*MsgConfigureChainResponse) ProtoMessage()    {}
func (*MsgConfigureChainResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b0dcb234192d0df, []int{3}
}
func (m *MsgConfigureChainResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgConfigureChainResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgConfigureChainResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgConfigureChainResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgConfigureChainResponse.Merge(m, src)
}
func (m *MsgConfigureChainResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgConfigureChainResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgConfigureChainResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgConfigureChainResponse proto.InternalMessageInfo

func (m *MsgConfigureChainResponse) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func init() {
	proto.RegisterType((*MsgSendErasmusStudent)(nil), "hub.hub.MsgSendErasmusStudent")
	proto.RegisterType((*MsgSendErasmusStudentResponse)(nil), "hub.hub.MsgSendErasmusStudentResponse")
	proto.RegisterType((*MsgConfigureChain)(nil), "hub.hub.MsgConfigureChain")
	proto.RegisterType((*MsgConfigureChainResponse)(nil), "hub.hub.MsgConfigureChainResponse")
}

func init() { proto.RegisterFile("hub/tx.proto", fileDescriptor_7b0dcb234192d0df) }

var fileDescriptor_7b0dcb234192d0df = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcb, 0x4a, 0xfb, 0x40,
	0x14, 0xc6, 0x3b, 0xff, 0xde, 0xe8, 0xf9, 0x8b, 0x97, 0x01, 0x4b, 0x0c, 0x3a, 0x96, 0x2c, 0xa4,
	0x0a, 0xa6, 0xd2, 0xbe, 0x81, 0xd5, 0x85, 0x8b, 0x82, 0xa4, 0xae, 0x44, 0x90, 0xa4, 0x1d, 0x93,
	0x80, 0x99, 0x09, 0x73, 0x81, 0xfa, 0x16, 0xbe, 0x8c, 0x0f, 0xe0, 0xce, 0x65, 0x97, 0x2e, 0xa5,
	0x7d, 0x11, 0xe9, 0x34, 0xa9, 0x97, 0xb6, 0xee, 0x72, 0xbe, 0xef, 0x9c, 0x93, 0xf3, 0x1b, 0x3e,
	0xd8, 0x88, 0x74, 0xd0, 0x52, 0x23, 0x37, 0x15, 0x5c, 0x71, 0x5c, 0x8d, 0x74, 0xe0, 0x46, 0x3a,
	0xb0, 0xad, 0x99, 0x2c, 0x15, 0x17, 0x74, 0x78, 0x2f, 0x95, 0x1e, 0x52, 0xa6, 0xe6, 0x2d, 0xce,
	0x2b, 0x82, 0xdd, 0x9e, 0x0c, 0xfb, 0x94, 0x0d, 0x2f, 0x85, 0x2f, 0x13, 0x2d, 0xfb, 0x73, 0x1f,
	0x5b, 0x50, 0x1d, 0x08, 0xea, 0x2b, 0x2e, 0x2c, 0xd4, 0x40, 0xcd, 0x9a, 0x97, 0x97, 0x18, 0x43,
	0x29, 0xe5, 0x42, 0x59, 0xff, 0x8c, 0x6c, 0xbe, 0xf1, 0x3e, 0xd4, 0x06, 0x91, 0xcf, 0x18, 0x7d,
	0xbc, 0xba, 0xb0, 0x8a, 0xc6, 0xf8, 0x12, 0xf0, 0x09, 0x6c, 0xab, 0x38, 0xa1, 0x5c, 0xab, 0x9b,
	0x38, 0xa1, 0x52, 0xf9, 0x49, 0x6a, 0x95, 0x1a, 0xa8, 0x59, 0xf2, 0x96, 0x74, 0x7c, 0x06, 0xd5,
	0xec, 0x44, 0xab, 0xdc, 0x40, 0xcd, 0xff, 0xed, 0xba, 0x9b, 0x61, 0xb8, 0x7d, 0x43, 0x90, 0x1d,
	0xe8, 0xe5, 0x6d, 0xce, 0x21, 0x1c, 0xac, 0x44, 0xf0, 0xa8, 0x4c, 0x39, 0x93, 0xd4, 0x39, 0x85,
	0x9d, 0x9e, 0x0c, 0xbb, 0x9c, 0x3d, 0xc4, 0xa1, 0x16, 0xb4, 0x1b, 0xf9, 0x31, 0x5b, 0xcf, 0xe7,
	0x74, 0x60, 0x6f, 0xa9, 0x3d, 0xdf, 0x85, 0xeb, 0x50, 0x91, 0xca, 0x57, 0x5a, 0x9a, 0xa9, 0xb2,
	0x97, 0x55, 0xed, 0x17, 0x04, 0xc5, 0x9e, 0x0c, 0xf1, 0x1d, 0xe0, 0x15, 0x8f, 0x49, 0x16, 0x0c,
	0x2b, 0x2f, 0xb5, 0x8f, 0xfe, 0xf6, 0x17, 0x7f, 0xbf, 0x86, 0xcd, 0x5f, 0x18, 0xf6, 0xf7, 0xc9,
	0x9f, 0x9e, 0xed, 0xac, 0xf7, 0xf2, 0x8d, 0xe7, 0xc7, 0x6f, 0x13, 0x82, 0xc6, 0x13, 0x82, 0x3e,
	0x26, 0x04, 0x3d, 0x4f, 0x49, 0x61, 0x3c, 0x25, 0x85, 0xf7, 0x29, 0x29, 0xdc, 0x6e, 0xcd, 0x42,
	0x33, 0x6a, 0x99, 0x44, 0x3d, 0xa5, 0x54, 0x06, 0x15, 0x13, 0x99, 0xce, 0x67, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x79, 0x65, 0xca, 0x58, 0x65, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	SendErasmusStudent(ctx context.Context, in *MsgSendErasmusStudent, opts ...grpc.CallOption) (*MsgSendErasmusStudentResponse, error)
	ConfigureChain(ctx context.Context, in *MsgConfigureChain, opts ...grpc.CallOption) (*MsgConfigureChainResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SendErasmusStudent(ctx context.Context, in *MsgSendErasmusStudent, opts ...grpc.CallOption) (*MsgSendErasmusStudentResponse, error) {
	out := new(MsgSendErasmusStudentResponse)
	err := c.cc.Invoke(ctx, "/hub.hub.Msg/SendErasmusStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ConfigureChain(ctx context.Context, in *MsgConfigureChain, opts ...grpc.CallOption) (*MsgConfigureChainResponse, error) {
	out := new(MsgConfigureChainResponse)
	err := c.cc.Invoke(ctx, "/hub.hub.Msg/ConfigureChain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	SendErasmusStudent(context.Context, *MsgSendErasmusStudent) (*MsgSendErasmusStudentResponse, error)
	ConfigureChain(context.Context, *MsgConfigureChain) (*MsgConfigureChainResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SendErasmusStudent(ctx context.Context, req *MsgSendErasmusStudent) (*MsgSendErasmusStudentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendErasmusStudent not implemented")
}
func (*UnimplementedMsgServer) ConfigureChain(ctx context.Context, req *MsgConfigureChain) (*MsgConfigureChainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigureChain not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SendErasmusStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSendErasmusStudent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SendErasmusStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hub.hub.Msg/SendErasmusStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SendErasmusStudent(ctx, req.(*MsgSendErasmusStudent))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ConfigureChain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgConfigureChain)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ConfigureChain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hub.hub.Msg/ConfigureChain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ConfigureChain(ctx, req.(*MsgConfigureChain))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hub.hub.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendErasmusStudent",
			Handler:    _Msg_SendErasmusStudent_Handler,
		},
		{
			MethodName: "ConfigureChain",
			Handler:    _Msg_ConfigureChain_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hub/tx.proto",
}

func (m *MsgSendErasmusStudent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSendErasmusStudent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSendErasmusStudent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Student != nil {
		{
			size, err := m.Student.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.TimeoutTimestamp != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.TimeoutTimestamp))
		i--
		dAtA[i] = 0x20
	}
	if len(m.ChannelID) > 0 {
		i -= len(m.ChannelID)
		copy(dAtA[i:], m.ChannelID)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ChannelID)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Port) > 0 {
		i -= len(m.Port)
		copy(dAtA[i:], m.Port)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Port)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSendErasmusStudentResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSendErasmusStudentResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSendErasmusStudentResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgConfigureChain) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgConfigureChain) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgConfigureChain) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgConfigureChainResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgConfigureChainResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgConfigureChainResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSendErasmusStudent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Port)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ChannelID)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.TimeoutTimestamp != 0 {
		n += 1 + sovTx(uint64(m.TimeoutTimestamp))
	}
	if m.Student != nil {
		l = m.Student.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSendErasmusStudentResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgConfigureChain) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgConfigureChainResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Status != 0 {
		n += 1 + sovTx(uint64(m.Status))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSendErasmusStudent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSendErasmusStudent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSendErasmusStudent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Port = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChannelID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutTimestamp", wireType)
			}
			m.TimeoutTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutTimestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Student", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Student == nil {
				m.Student = &StoredStudent{}
			}
			if err := m.Student.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSendErasmusStudentResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSendErasmusStudentResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSendErasmusStudentResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgConfigureChain) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgConfigureChain: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgConfigureChain: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgConfigureChainResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgConfigureChainResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgConfigureChainResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
