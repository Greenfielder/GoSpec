// Code generated by protoc-gen-go. DO NOT EDIT.
// source: consignment.proto

package consignment

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

//Коэффициенты квадратного уравнения
type Coefficients struct {
	A                    int32    `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B                    int32    `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
	C                    int32    `protobuf:"varint,3,opt,name=c,proto3" json:"c,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Coefficients) Reset()         { *m = Coefficients{} }
func (m *Coefficients) String() string { return proto.CompactTextString(m) }
func (*Coefficients) ProtoMessage()    {}
func (*Coefficients) Descriptor() ([]byte, []int) {
	return fileDescriptor_3804bf87090b51a9, []int{0}
}

func (m *Coefficients) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Coefficients.Unmarshal(m, b)
}
func (m *Coefficients) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Coefficients.Marshal(b, m, deterministic)
}
func (m *Coefficients) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Coefficients.Merge(m, src)
}
func (m *Coefficients) XXX_Size() int {
	return xxx_messageInfo_Coefficients.Size(m)
}
func (m *Coefficients) XXX_DiscardUnknown() {
	xxx_messageInfo_Coefficients.DiscardUnknown(m)
}

var xxx_messageInfo_Coefficients proto.InternalMessageInfo

func (m *Coefficients) GetA() int32 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *Coefficients) GetB() int32 {
	if m != nil {
		return m.B
	}
	return 0
}

func (m *Coefficients) GetC() int32 {
	if m != nil {
		return m.C
	}
	return 0
}

//Решение квадратного уравнения
type Solution struct {
	Coefs                *Coefficients `protobuf:"bytes,1,opt,name=coefs,proto3" json:"coefs,omitempty"`
	NRoots               int32         `protobuf:"varint,2,opt,name=n_roots,json=nRoots,proto3" json:"n_roots,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Solution) Reset()         { *m = Solution{} }
func (m *Solution) String() string { return proto.CompactTextString(m) }
func (*Solution) ProtoMessage()    {}
func (*Solution) Descriptor() ([]byte, []int) {
	return fileDescriptor_3804bf87090b51a9, []int{1}
}

func (m *Solution) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Solution.Unmarshal(m, b)
}
func (m *Solution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Solution.Marshal(b, m, deterministic)
}
func (m *Solution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Solution.Merge(m, src)
}
func (m *Solution) XXX_Size() int {
	return xxx_messageInfo_Solution.Size(m)
}
func (m *Solution) XXX_DiscardUnknown() {
	xxx_messageInfo_Solution.DiscardUnknown(m)
}

var xxx_messageInfo_Solution proto.InternalMessageInfo

func (m *Solution) GetCoefs() *Coefficients {
	if m != nil {
		return m.Coefs
	}
	return nil
}

func (m *Solution) GetNRoots() int32 {
	if m != nil {
		return m.NRoots
	}
	return 0
}

//Набор уже решенных уравнений
type Solutions struct {
	Solutions            []*Solution `protobuf:"bytes,1,rep,name=solutions,proto3" json:"solutions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Solutions) Reset()         { *m = Solutions{} }
func (m *Solutions) String() string { return proto.CompactTextString(m) }
func (*Solutions) ProtoMessage()    {}
func (*Solutions) Descriptor() ([]byte, []int) {
	return fileDescriptor_3804bf87090b51a9, []int{2}
}

func (m *Solutions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Solutions.Unmarshal(m, b)
}
func (m *Solutions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Solutions.Marshal(b, m, deterministic)
}
func (m *Solutions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Solutions.Merge(m, src)
}
func (m *Solutions) XXX_Size() int {
	return xxx_messageInfo_Solutions.Size(m)
}
func (m *Solutions) XXX_DiscardUnknown() {
	xxx_messageInfo_Solutions.DiscardUnknown(m)
}

var xxx_messageInfo_Solutions proto.InternalMessageInfo

func (m *Solutions) GetSolutions() []*Solution {
	if m != nil {
		return m.Solutions
	}
	return nil
}

//Пустышка для GetAll()
type GetRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3804bf87090b51a9, []int{3}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Coefficients)(nil), "consignment.Coefficients")
	proto.RegisterType((*Solution)(nil), "consignment.Solution")
	proto.RegisterType((*Solutions)(nil), "consignment.Solutions")
	proto.RegisterType((*GetRequest)(nil), "consignment.GetRequest")
}

func init() {
	proto.RegisterFile("consignment.proto", fileDescriptor_3804bf87090b51a9)
}

var fileDescriptor_3804bf87090b51a9 = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x40, 0x6b, 0xaa, 0x18, 0x7a, 0xcd, 0xc2, 0x49, 0xd0, 0xd0, 0x09, 0x79, 0x62, 0x2a, 0x52,
	0xbb, 0x20, 0xb1, 0x80, 0x18, 0xd8, 0x0d, 0x3b, 0x6a, 0x2c, 0x07, 0x59, 0x0a, 0x3e, 0xc8, 0x39,
	0x7c, 0x00, 0x5f, 0x8e, 0x93, 0x10, 0x25, 0x91, 0xc2, 0xe6, 0x67, 0xbd, 0x7b, 0x67, 0x19, 0xce,
	0x0d, 0x79, 0x76, 0xef, 0xfe, 0xc3, 0xfa, 0xb0, 0xfb, 0xac, 0x28, 0x10, 0xae, 0x47, 0x57, 0xea,
	0x0e, 0xd2, 0x27, 0xb2, 0x45, 0xe1, 0x8c, 0x8b, 0xc8, 0x98, 0x82, 0x38, 0x66, 0xe2, 0x5a, 0xdc,
	0x24, 0x5a, 0x1c, 0x1b, 0xca, 0xb3, 0x93, 0x8e, 0xf2, 0x86, 0x4c, 0xb6, 0xec, 0xc8, 0xa8, 0x57,
	0x38, 0x7b, 0xa1, 0xb2, 0x0e, 0x8e, 0x3c, 0xde, 0x42, 0x62, 0x62, 0x85, 0xdb, 0xc9, 0xf5, 0xfe,
	0x6a, 0x37, 0xde, 0x3a, 0xee, 0xeb, 0xce, 0xc3, 0x0d, 0x9c, 0xfa, 0xb7, 0x8a, 0x28, 0xf0, 0x5f,
	0x5e, 0x7a, 0xdd, 0x90, 0x7a, 0x80, 0x55, 0x5f, 0x65, 0x3c, 0xc0, 0x8a, 0x7b, 0x88, 0xe9, 0x65,
	0x4c, 0x5f, 0x4c, 0xd2, 0xbd, 0xaa, 0x07, 0x4f, 0xa5, 0x00, 0xcf, 0x36, 0x68, 0xfb, 0x55, 0x5b,
	0x0e, 0xfb, 0x1f, 0x01, 0x32, 0x5a, 0xdf, 0xb6, 0xc2, 0x7b, 0x48, 0xda, 0x13, 0xfe, 0xff, 0xbc,
	0xed, 0x7c, 0x5e, 0x2d, 0xe2, 0xb0, 0x8c, 0xd5, 0xc7, 0xb2, 0xc4, 0xcd, 0x44, 0x19, 0x56, 0x6d,
	0x2f, 0x67, 0x67, 0x59, 0x2d, 0x72, 0xd9, 0x7e, 0xfc, 0xe1, 0x37, 0x00, 0x00, 0xff, 0xff, 0x76,
	0xf4, 0xf3, 0x10, 0x8d, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SolverClient is the client API for Solver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SolverClient interface {
	Solve(ctx context.Context, in *Coefficients, opts ...grpc.CallOption) (*Solution, error)
	GetAll(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Solutions, error)
}

type solverClient struct {
	cc grpc.ClientConnInterface
}

func NewSolverClient(cc grpc.ClientConnInterface) SolverClient {
	return &solverClient{cc}
}

func (c *solverClient) Solve(ctx context.Context, in *Coefficients, opts ...grpc.CallOption) (*Solution, error) {
	out := new(Solution)
	err := c.cc.Invoke(ctx, "/consignment.Solver/Solve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *solverClient) GetAll(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Solutions, error) {
	out := new(Solutions)
	err := c.cc.Invoke(ctx, "/consignment.Solver/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SolverServer is the server API for Solver service.
type SolverServer interface {
	Solve(context.Context, *Coefficients) (*Solution, error)
	GetAll(context.Context, *GetRequest) (*Solutions, error)
}

// UnimplementedSolverServer can be embedded to have forward compatible implementations.
type UnimplementedSolverServer struct {
}

func (*UnimplementedSolverServer) Solve(ctx context.Context, req *Coefficients) (*Solution, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Solve not implemented")
}
func (*UnimplementedSolverServer) GetAll(ctx context.Context, req *GetRequest) (*Solutions, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}

func RegisterSolverServer(s *grpc.Server, srv SolverServer) {
	s.RegisterService(&_Solver_serviceDesc, srv)
}

func _Solver_Solve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Coefficients)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SolverServer).Solve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consignment.Solver/Solve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SolverServer).Solve(ctx, req.(*Coefficients))
	}
	return interceptor(ctx, in, info, handler)
}

func _Solver_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SolverServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consignment.Solver/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SolverServer).GetAll(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Solver_serviceDesc = grpc.ServiceDesc{
	ServiceName: "consignment.Solver",
	HandlerType: (*SolverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Solve",
			Handler:    _Solver_Solve_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _Solver_GetAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "consignment.proto",
}
