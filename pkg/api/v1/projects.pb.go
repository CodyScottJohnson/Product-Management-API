// Code generated by protoc-gen-go. DO NOT EDIT.
// source: projects.proto

package v1

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

type Project struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Company              string   `protobuf:"bytes,2,opt,name=company,proto3" json:"company,omitempty"`
	Background           string   `protobuf:"bytes,3,opt,name=background,proto3" json:"background,omitempty"`
	Image                string   `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	Title                string   `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Project) Reset()         { *m = Project{} }
func (m *Project) String() string { return proto.CompactTextString(m) }
func (*Project) ProtoMessage()    {}
func (*Project) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2826043bb82544a, []int{0}
}

func (m *Project) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Project.Unmarshal(m, b)
}
func (m *Project) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Project.Marshal(b, m, deterministic)
}
func (m *Project) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Project.Merge(m, src)
}
func (m *Project) XXX_Size() int {
	return xxx_messageInfo_Project.Size(m)
}
func (m *Project) XXX_DiscardUnknown() {
	xxx_messageInfo_Project.DiscardUnknown(m)
}

var xxx_messageInfo_Project proto.InternalMessageInfo

func (m *Project) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Project) GetCompany() string {
	if m != nil {
		return m.Company
	}
	return ""
}

func (m *Project) GetBackground() string {
	if m != nil {
		return m.Background
	}
	return ""
}

func (m *Project) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Project) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Project) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type CreateRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Project              *Project `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2826043bb82544a, []int{1}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *CreateRequest) GetProject() *Project {
	if m != nil {
		return m.Project
	}
	return nil
}

type CreateResponse struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Id                   int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2826043bb82544a, []int{2}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *CreateResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ReadRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Id                   int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2826043bb82544a, []int{3}
}

func (m *ReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRequest.Unmarshal(m, b)
}
func (m *ReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRequest.Marshal(b, m, deterministic)
}
func (m *ReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRequest.Merge(m, src)
}
func (m *ReadRequest) XXX_Size() int {
	return xxx_messageInfo_ReadRequest.Size(m)
}
func (m *ReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRequest proto.InternalMessageInfo

func (m *ReadRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ReadRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ReadResponse struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Project              *Project `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadResponse) Reset()         { *m = ReadResponse{} }
func (m *ReadResponse) String() string { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()    {}
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2826043bb82544a, []int{4}
}

func (m *ReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadResponse.Unmarshal(m, b)
}
func (m *ReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadResponse.Marshal(b, m, deterministic)
}
func (m *ReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadResponse.Merge(m, src)
}
func (m *ReadResponse) XXX_Size() int {
	return xxx_messageInfo_ReadResponse.Size(m)
}
func (m *ReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadResponse proto.InternalMessageInfo

func (m *ReadResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ReadResponse) GetProject() *Project {
	if m != nil {
		return m.Project
	}
	return nil
}

type ReadAllRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadAllRequest) Reset()         { *m = ReadAllRequest{} }
func (m *ReadAllRequest) String() string { return proto.CompactTextString(m) }
func (*ReadAllRequest) ProtoMessage()    {}
func (*ReadAllRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2826043bb82544a, []int{5}
}

func (m *ReadAllRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadAllRequest.Unmarshal(m, b)
}
func (m *ReadAllRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadAllRequest.Marshal(b, m, deterministic)
}
func (m *ReadAllRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadAllRequest.Merge(m, src)
}
func (m *ReadAllRequest) XXX_Size() int {
	return xxx_messageInfo_ReadAllRequest.Size(m)
}
func (m *ReadAllRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadAllRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadAllRequest proto.InternalMessageInfo

func (m *ReadAllRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

type ReadAllResponse struct {
	Api                  string     `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Projects             []*Project `protobuf:"bytes,2,rep,name=projects,proto3" json:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ReadAllResponse) Reset()         { *m = ReadAllResponse{} }
func (m *ReadAllResponse) String() string { return proto.CompactTextString(m) }
func (*ReadAllResponse) ProtoMessage()    {}
func (*ReadAllResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2826043bb82544a, []int{6}
}

func (m *ReadAllResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadAllResponse.Unmarshal(m, b)
}
func (m *ReadAllResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadAllResponse.Marshal(b, m, deterministic)
}
func (m *ReadAllResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadAllResponse.Merge(m, src)
}
func (m *ReadAllResponse) XXX_Size() int {
	return xxx_messageInfo_ReadAllResponse.Size(m)
}
func (m *ReadAllResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadAllResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadAllResponse proto.InternalMessageInfo

func (m *ReadAllResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ReadAllResponse) GetProjects() []*Project {
	if m != nil {
		return m.Projects
	}
	return nil
}

func init() {
	proto.RegisterType((*Project)(nil), "v1.Project")
	proto.RegisterType((*CreateRequest)(nil), "v1.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "v1.CreateResponse")
	proto.RegisterType((*ReadRequest)(nil), "v1.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "v1.ReadResponse")
	proto.RegisterType((*ReadAllRequest)(nil), "v1.ReadAllRequest")
	proto.RegisterType((*ReadAllResponse)(nil), "v1.ReadAllResponse")
}

func init() { proto.RegisterFile("projects.proto", fileDescriptor_a2826043bb82544a) }

var fileDescriptor_a2826043bb82544a = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x4e, 0xbb, 0x40,
	0x10, 0xc6, 0x03, 0xb4, 0xe5, 0xdf, 0xe1, 0x2f, 0xad, 0xa3, 0x87, 0x4d, 0x0f, 0x86, 0x90, 0x18,
	0xeb, 0xa5, 0x4d, 0xf1, 0x09, 0x8c, 0x07, 0x3d, 0x78, 0x30, 0xf8, 0x04, 0x14, 0x26, 0xcd, 0x6a,
	0xcb, 0x22, 0x6c, 0x9b, 0xf8, 0x36, 0xfa, 0xa6, 0x86, 0x5d, 0xb6, 0xa5, 0x8d, 0x24, 0xde, 0xd8,
	0xdf, 0xb7, 0x33, 0xdf, 0x37, 0xb3, 0x80, 0x5f, 0x94, 0xe2, 0x8d, 0x52, 0x59, 0xcd, 0x8a, 0x52,
	0x48, 0x81, 0xf6, 0x6e, 0x11, 0x7e, 0x5b, 0xe0, 0xbe, 0x68, 0x8c, 0x3e, 0xd8, 0x3c, 0x63, 0x56,
	0x60, 0x4d, 0x9d, 0xd8, 0xe6, 0x19, 0x32, 0x70, 0x53, 0xb1, 0x29, 0x92, 0xfc, 0x93, 0xd9, 0x81,
	0x35, 0x1d, 0xc6, 0xe6, 0x88, 0x57, 0x00, 0xcb, 0x24, 0x7d, 0x5f, 0x95, 0x62, 0x9b, 0x67, 0xcc,
	0x51, 0x62, 0x8b, 0xe0, 0x25, 0xf4, 0xf9, 0x26, 0x59, 0x11, 0xeb, 0x29, 0x49, 0x1f, 0x6a, 0x2a,
	0xb9, 0x5c, 0x13, 0xeb, 0x6b, 0xaa, 0x0e, 0x18, 0x80, 0x97, 0x51, 0x95, 0x96, 0xbc, 0x90, 0x5c,
	0xe4, 0x6c, 0xa0, 0xb4, 0x36, 0x0a, 0x9f, 0xe0, 0xec, 0xa1, 0xa4, 0x44, 0x52, 0x4c, 0x1f, 0x5b,
	0xaa, 0x24, 0x8e, 0xc1, 0x49, 0x0a, 0xae, 0x92, 0x0e, 0xe3, 0xfa, 0x13, 0xaf, 0xc1, 0x6d, 0x86,
	0x53, 0x51, 0xbd, 0xc8, 0x9b, 0xed, 0x16, 0xb3, 0x66, 0xb0, 0xd8, 0x68, 0x61, 0x04, 0xbe, 0xe9,
	0x54, 0x15, 0x22, 0xaf, 0xe8, 0x97, 0x56, 0x7a, 0x0b, 0xb6, 0xd9, 0x42, 0x38, 0x07, 0x2f, 0xa6,
	0x24, 0xeb, 0xf6, 0x3e, 0x2d, 0x78, 0x84, 0xff, 0xba, 0xa0, 0xd3, 0xe2, 0x8f, 0x69, 0x43, 0xf0,
	0xeb, 0x46, 0xf7, 0xeb, 0x75, 0xa7, 0x79, 0xf8, 0x0c, 0xa3, 0xfd, 0x9d, 0x4e, 0xbf, 0x1b, 0xf8,
	0x67, 0x9e, 0x9e, 0xd9, 0x81, 0x73, 0x6a, 0xb8, 0x17, 0xa3, 0x2f, 0x0b, 0xfc, 0x86, 0xbe, 0x52,
	0xb9, 0xe3, 0x29, 0xe1, 0x1c, 0x06, 0x7a, 0x65, 0x78, 0x5e, 0xd7, 0x1c, 0x3d, 0xc4, 0x04, 0xdb,
	0xa8, 0xb1, 0xbf, 0x85, 0x5e, 0x9d, 0x08, 0x47, 0xb5, 0xd6, 0xda, 0xdc, 0x64, 0x7c, 0x00, 0xcd,
	0xd5, 0x08, 0xdc, 0x26, 0x3c, 0xa2, 0x11, 0x0f, 0xd3, 0x4e, 0x2e, 0x8e, 0x98, 0xae, 0x59, 0x0e,
	0xd4, 0xbf, 0x7b, 0xf7, 0x13, 0x00, 0x00, 0xff, 0xff, 0x96, 0x02, 0xb1, 0x35, 0xcd, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProjectServiceClient is the client API for ProjectService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProjectServiceClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	// Read todo task
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
	// Read all todo tasks
	ReadAll(ctx context.Context, in *ReadAllRequest, opts ...grpc.CallOption) (*ReadAllResponse, error)
}

type projectServiceClient struct {
	cc *grpc.ClientConn
}

func NewProjectServiceClient(cc *grpc.ClientConn) ProjectServiceClient {
	return &projectServiceClient{cc}
}

func (c *projectServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/v1.ProjectService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := c.cc.Invoke(ctx, "/v1.ProjectService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) ReadAll(ctx context.Context, in *ReadAllRequest, opts ...grpc.CallOption) (*ReadAllResponse, error) {
	out := new(ReadAllResponse)
	err := c.cc.Invoke(ctx, "/v1.ProjectService/ReadAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectServiceServer is the server API for ProjectService service.
type ProjectServiceServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	// Read todo task
	Read(context.Context, *ReadRequest) (*ReadResponse, error)
	// Read all todo tasks
	ReadAll(context.Context, *ReadAllRequest) (*ReadAllResponse, error)
}

func RegisterProjectServiceServer(s *grpc.Server, srv ProjectServiceServer) {
	s.RegisterService(&_ProjectService_serviceDesc, srv)
}

func _ProjectService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ProjectService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ProjectService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_ReadAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).ReadAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ProjectService/ReadAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).ReadAll(ctx, req.(*ReadAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProjectService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.ProjectService",
	HandlerType: (*ProjectServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ProjectService_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _ProjectService_Read_Handler,
		},
		{
			MethodName: "ReadAll",
			Handler:    _ProjectService_ReadAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "projects.proto",
}
