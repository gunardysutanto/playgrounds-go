// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/teams.proto

package teampb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Team struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Department           string   `protobuf:"bytes,4,opt,name=department,proto3" json:"department,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Team) Reset()         { *m = Team{} }
func (m *Team) String() string { return proto.CompactTextString(m) }
func (*Team) ProtoMessage()    {}
func (*Team) Descriptor() ([]byte, []int) {
	return fileDescriptor_772c174fe3166e3d, []int{0}
}

func (m *Team) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Team.Unmarshal(m, b)
}
func (m *Team) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Team.Marshal(b, m, deterministic)
}
func (m *Team) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Team.Merge(m, src)
}
func (m *Team) XXX_Size() int {
	return xxx_messageInfo_Team.Size(m)
}
func (m *Team) XXX_DiscardUnknown() {
	xxx_messageInfo_Team.DiscardUnknown(m)
}

var xxx_messageInfo_Team proto.InternalMessageInfo

func (m *Team) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Team) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Team) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Team) GetDepartment() string {
	if m != nil {
		return m.Department
	}
	return ""
}

type QueryRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryRequest) Reset()         { *m = QueryRequest{} }
func (m *QueryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()    {}
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_772c174fe3166e3d, []int{1}
}

func (m *QueryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryRequest.Unmarshal(m, b)
}
func (m *QueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryRequest.Marshal(b, m, deterministic)
}
func (m *QueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRequest.Merge(m, src)
}
func (m *QueryRequest) XXX_Size() int {
	return xxx_messageInfo_QueryRequest.Size(m)
}
func (m *QueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRequest proto.InternalMessageInfo

type QueryResponse struct {
	Team                 *Team    `protobuf:"bytes,1,opt,name=team,proto3" json:"team,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryResponse) Reset()         { *m = QueryResponse{} }
func (m *QueryResponse) String() string { return proto.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()    {}
func (*QueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_772c174fe3166e3d, []int{2}
}

func (m *QueryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryResponse.Unmarshal(m, b)
}
func (m *QueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryResponse.Marshal(b, m, deterministic)
}
func (m *QueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryResponse.Merge(m, src)
}
func (m *QueryResponse) XXX_Size() int {
	return xxx_messageInfo_QueryResponse.Size(m)
}
func (m *QueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryResponse proto.InternalMessageInfo

func (m *QueryResponse) GetTeam() *Team {
	if m != nil {
		return m.Team
	}
	return nil
}

type ViewRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ViewRequest) Reset()         { *m = ViewRequest{} }
func (m *ViewRequest) String() string { return proto.CompactTextString(m) }
func (*ViewRequest) ProtoMessage()    {}
func (*ViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_772c174fe3166e3d, []int{3}
}

func (m *ViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ViewRequest.Unmarshal(m, b)
}
func (m *ViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ViewRequest.Marshal(b, m, deterministic)
}
func (m *ViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ViewRequest.Merge(m, src)
}
func (m *ViewRequest) XXX_Size() int {
	return xxx_messageInfo_ViewRequest.Size(m)
}
func (m *ViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ViewRequest proto.InternalMessageInfo

func (m *ViewRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ViewResponse struct {
	Team                 *Team    `protobuf:"bytes,1,opt,name=team,proto3" json:"team,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ViewResponse) Reset()         { *m = ViewResponse{} }
func (m *ViewResponse) String() string { return proto.CompactTextString(m) }
func (*ViewResponse) ProtoMessage()    {}
func (*ViewResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_772c174fe3166e3d, []int{4}
}

func (m *ViewResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ViewResponse.Unmarshal(m, b)
}
func (m *ViewResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ViewResponse.Marshal(b, m, deterministic)
}
func (m *ViewResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ViewResponse.Merge(m, src)
}
func (m *ViewResponse) XXX_Size() int {
	return xxx_messageInfo_ViewResponse.Size(m)
}
func (m *ViewResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ViewResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ViewResponse proto.InternalMessageInfo

func (m *ViewResponse) GetTeam() *Team {
	if m != nil {
		return m.Team
	}
	return nil
}

type NewTeamRequest struct {
	Team                 *Team    `protobuf:"bytes,1,opt,name=team,proto3" json:"team,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewTeamRequest) Reset()         { *m = NewTeamRequest{} }
func (m *NewTeamRequest) String() string { return proto.CompactTextString(m) }
func (*NewTeamRequest) ProtoMessage()    {}
func (*NewTeamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_772c174fe3166e3d, []int{5}
}

func (m *NewTeamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewTeamRequest.Unmarshal(m, b)
}
func (m *NewTeamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewTeamRequest.Marshal(b, m, deterministic)
}
func (m *NewTeamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewTeamRequest.Merge(m, src)
}
func (m *NewTeamRequest) XXX_Size() int {
	return xxx_messageInfo_NewTeamRequest.Size(m)
}
func (m *NewTeamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NewTeamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NewTeamRequest proto.InternalMessageInfo

func (m *NewTeamRequest) GetTeam() *Team {
	if m != nil {
		return m.Team
	}
	return nil
}

type NewTeamResponse struct {
	Team                 *Team    `protobuf:"bytes,1,opt,name=team,proto3" json:"team,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewTeamResponse) Reset()         { *m = NewTeamResponse{} }
func (m *NewTeamResponse) String() string { return proto.CompactTextString(m) }
func (*NewTeamResponse) ProtoMessage()    {}
func (*NewTeamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_772c174fe3166e3d, []int{6}
}

func (m *NewTeamResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewTeamResponse.Unmarshal(m, b)
}
func (m *NewTeamResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewTeamResponse.Marshal(b, m, deterministic)
}
func (m *NewTeamResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewTeamResponse.Merge(m, src)
}
func (m *NewTeamResponse) XXX_Size() int {
	return xxx_messageInfo_NewTeamResponse.Size(m)
}
func (m *NewTeamResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NewTeamResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NewTeamResponse proto.InternalMessageInfo

func (m *NewTeamResponse) GetTeam() *Team {
	if m != nil {
		return m.Team
	}
	return nil
}

type UpdateTeamRequest struct {
	Team                 *Team    `protobuf:"bytes,1,opt,name=team,proto3" json:"team,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTeamRequest) Reset()         { *m = UpdateTeamRequest{} }
func (m *UpdateTeamRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateTeamRequest) ProtoMessage()    {}
func (*UpdateTeamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_772c174fe3166e3d, []int{7}
}

func (m *UpdateTeamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTeamRequest.Unmarshal(m, b)
}
func (m *UpdateTeamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTeamRequest.Marshal(b, m, deterministic)
}
func (m *UpdateTeamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTeamRequest.Merge(m, src)
}
func (m *UpdateTeamRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateTeamRequest.Size(m)
}
func (m *UpdateTeamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTeamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTeamRequest proto.InternalMessageInfo

func (m *UpdateTeamRequest) GetTeam() *Team {
	if m != nil {
		return m.Team
	}
	return nil
}

type UpdateTeamResponse struct {
	Team                 *Team    `protobuf:"bytes,1,opt,name=team,proto3" json:"team,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTeamResponse) Reset()         { *m = UpdateTeamResponse{} }
func (m *UpdateTeamResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateTeamResponse) ProtoMessage()    {}
func (*UpdateTeamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_772c174fe3166e3d, []int{8}
}

func (m *UpdateTeamResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTeamResponse.Unmarshal(m, b)
}
func (m *UpdateTeamResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTeamResponse.Marshal(b, m, deterministic)
}
func (m *UpdateTeamResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTeamResponse.Merge(m, src)
}
func (m *UpdateTeamResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateTeamResponse.Size(m)
}
func (m *UpdateTeamResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTeamResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTeamResponse proto.InternalMessageInfo

func (m *UpdateTeamResponse) GetTeam() *Team {
	if m != nil {
		return m.Team
	}
	return nil
}

type RemoveTeamRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveTeamRequest) Reset()         { *m = RemoveTeamRequest{} }
func (m *RemoveTeamRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveTeamRequest) ProtoMessage()    {}
func (*RemoveTeamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_772c174fe3166e3d, []int{9}
}

func (m *RemoveTeamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveTeamRequest.Unmarshal(m, b)
}
func (m *RemoveTeamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveTeamRequest.Marshal(b, m, deterministic)
}
func (m *RemoveTeamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveTeamRequest.Merge(m, src)
}
func (m *RemoveTeamRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveTeamRequest.Size(m)
}
func (m *RemoveTeamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveTeamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveTeamRequest proto.InternalMessageInfo

func (m *RemoveTeamRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type RemoveTeamResponse struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveTeamResponse) Reset()         { *m = RemoveTeamResponse{} }
func (m *RemoveTeamResponse) String() string { return proto.CompactTextString(m) }
func (*RemoveTeamResponse) ProtoMessage()    {}
func (*RemoveTeamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_772c174fe3166e3d, []int{10}
}

func (m *RemoveTeamResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveTeamResponse.Unmarshal(m, b)
}
func (m *RemoveTeamResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveTeamResponse.Marshal(b, m, deterministic)
}
func (m *RemoveTeamResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveTeamResponse.Merge(m, src)
}
func (m *RemoveTeamResponse) XXX_Size() int {
	return xxx_messageInfo_RemoveTeamResponse.Size(m)
}
func (m *RemoveTeamResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveTeamResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveTeamResponse proto.InternalMessageInfo

func (m *RemoveTeamResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*Team)(nil), "teams.Team")
	proto.RegisterType((*QueryRequest)(nil), "teams.QueryRequest")
	proto.RegisterType((*QueryResponse)(nil), "teams.QueryResponse")
	proto.RegisterType((*ViewRequest)(nil), "teams.ViewRequest")
	proto.RegisterType((*ViewResponse)(nil), "teams.ViewResponse")
	proto.RegisterType((*NewTeamRequest)(nil), "teams.NewTeamRequest")
	proto.RegisterType((*NewTeamResponse)(nil), "teams.NewTeamResponse")
	proto.RegisterType((*UpdateTeamRequest)(nil), "teams.UpdateTeamRequest")
	proto.RegisterType((*UpdateTeamResponse)(nil), "teams.UpdateTeamResponse")
	proto.RegisterType((*RemoveTeamRequest)(nil), "teams.RemoveTeamRequest")
	proto.RegisterType((*RemoveTeamResponse)(nil), "teams.RemoveTeamResponse")
}

func init() { proto.RegisterFile("proto/teams.proto", fileDescriptor_772c174fe3166e3d) }

var fileDescriptor_772c174fe3166e3d = []byte{
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4f, 0x4f, 0xfa, 0x40,
	0x10, 0x4d, 0xfb, 0x2b, 0xcd, 0x8f, 0x29, 0x62, 0x18, 0x94, 0xd4, 0x26, 0x2a, 0xa9, 0x17, 0x0e,
	0x06, 0x10, 0x31, 0x9e, 0xbc, 0xf8, 0x01, 0x4c, 0xac, 0x7f, 0x0e, 0xde, 0x16, 0x3a, 0x87, 0x4d,
	0xec, 0x1f, 0x76, 0xb7, 0x10, 0x3f, 0x93, 0x5f, 0xd2, 0x74, 0xd9, 0x60, 0x0b, 0x26, 0xe2, 0x6d,
	0xe7, 0xbd, 0x79, 0x33, 0x2f, 0x6f, 0xb2, 0xd0, 0xc9, 0x45, 0xa6, 0xb2, 0x91, 0x22, 0x96, 0xc8,
	0xa1, 0x7e, 0x63, 0x43, 0x17, 0xe1, 0x3b, 0x38, 0xcf, 0xc4, 0x12, 0x6c, 0x83, 0xcd, 0x63, 0xdf,
	0xea, 0x5b, 0x83, 0x46, 0x64, 0xf3, 0x18, 0x11, 0x9c, 0x94, 0x25, 0xe4, 0xdb, 0x7d, 0x6b, 0xd0,
	0x8c, 0xf4, 0x1b, 0xfb, 0xe0, 0xc5, 0x24, 0xe7, 0x82, 0xe7, 0x8a, 0x67, 0xa9, 0xff, 0x4f, 0x53,
	0x55, 0x08, 0xcf, 0x00, 0x62, 0xca, 0x99, 0x50, 0x09, 0xa5, 0xca, 0x77, 0x74, 0x43, 0x05, 0x09,
	0xdb, 0xd0, 0x7a, 0x2c, 0x48, 0x7c, 0x44, 0xb4, 0x28, 0x48, 0xaa, 0x70, 0x0c, 0x07, 0xa6, 0x96,
	0x79, 0x96, 0x4a, 0xc2, 0x73, 0x70, 0x4a, 0x5f, 0xda, 0x88, 0x37, 0xf1, 0x86, 0x6b, 0xc7, 0xa5,
	0xc3, 0x48, 0x13, 0xe1, 0x29, 0x78, 0xaf, 0x9c, 0x56, 0x66, 0xc0, 0xb6, 0xed, 0x70, 0x04, 0xad,
	0x35, 0xbd, 0xef, 0xbc, 0x2b, 0x68, 0x3f, 0xd0, 0x4a, 0x03, 0x66, 0xe4, 0xaf, 0x92, 0x09, 0x1c,
	0x6e, 0x24, 0xfb, 0xae, 0x99, 0x42, 0xe7, 0x25, 0x8f, 0x99, 0xa2, 0x3f, 0x6d, 0xba, 0x01, 0xac,
	0xaa, 0xf6, 0x5d, 0x76, 0x01, 0x9d, 0x88, 0x92, 0x6c, 0x59, 0x5b, 0xb6, 0x9d, 0xd4, 0x25, 0x60,
	0xb5, 0xc9, 0xcc, 0xee, 0x81, 0x2b, 0x15, 0x53, 0x85, 0xd4, 0x9d, 0xcd, 0xc8, 0x54, 0x93, 0x4f,
	0x1b, 0xbc, 0xb2, 0xf1, 0x89, 0xc4, 0x92, 0xcf, 0x09, 0xa7, 0xd0, 0x58, 0x94, 0x87, 0xc3, 0xae,
	0x59, 0x5f, 0x3d, 0x6b, 0x70, 0x54, 0x07, 0xd7, 0xb3, 0xc7, 0x16, 0x8e, 0xc0, 0x59, 0x72, 0x5a,
	0x21, 0x1a, 0xbe, 0x72, 0xc9, 0xa0, 0x5b, 0xc3, 0x8c, 0x9d, 0x5b, 0x70, 0xe7, 0x82, 0x98, 0x22,
	0x3c, 0x36, 0x74, 0xfd, 0x58, 0x41, 0x6f, 0x1b, 0x36, 0xc2, 0x3b, 0x70, 0x0b, 0x9d, 0x1c, 0xfa,
	0xa6, 0x63, 0x27, 0xfe, 0xe0, 0xe4, 0x07, 0xe6, 0x5b, 0x2e, 0x74, 0x38, 0x1b, 0xf9, 0x4e, 0xa0,
	0x1b, 0xf9, 0x6e, 0x8a, 0xf7, 0xff, 0xdf, 0xdc, 0x92, 0xcb, 0x67, 0x33, 0x57, 0x7f, 0xb6, 0xeb,
	0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x97, 0xbb, 0x90, 0x81, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TeamServiceClient is the client API for TeamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TeamServiceClient interface {
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (TeamService_QueryClient, error)
	View(ctx context.Context, in *ViewRequest, opts ...grpc.CallOption) (*ViewResponse, error)
	Create(ctx context.Context, in *NewTeamRequest, opts ...grpc.CallOption) (*NewTeamResponse, error)
	Update(ctx context.Context, in *UpdateTeamRequest, opts ...grpc.CallOption) (*UpdateTeamResponse, error)
	Remove(ctx context.Context, in *RemoveTeamRequest, opts ...grpc.CallOption) (*RemoveTeamResponse, error)
}

type teamServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTeamServiceClient(cc grpc.ClientConnInterface) TeamServiceClient {
	return &teamServiceClient{cc}
}

func (c *teamServiceClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (TeamService_QueryClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TeamService_serviceDesc.Streams[0], "/teams.TeamService/query", opts...)
	if err != nil {
		return nil, err
	}
	x := &teamServiceQueryClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TeamService_QueryClient interface {
	Recv() (*QueryResponse, error)
	grpc.ClientStream
}

type teamServiceQueryClient struct {
	grpc.ClientStream
}

func (x *teamServiceQueryClient) Recv() (*QueryResponse, error) {
	m := new(QueryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *teamServiceClient) View(ctx context.Context, in *ViewRequest, opts ...grpc.CallOption) (*ViewResponse, error) {
	out := new(ViewResponse)
	err := c.cc.Invoke(ctx, "/teams.TeamService/view", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) Create(ctx context.Context, in *NewTeamRequest, opts ...grpc.CallOption) (*NewTeamResponse, error) {
	out := new(NewTeamResponse)
	err := c.cc.Invoke(ctx, "/teams.TeamService/create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) Update(ctx context.Context, in *UpdateTeamRequest, opts ...grpc.CallOption) (*UpdateTeamResponse, error) {
	out := new(UpdateTeamResponse)
	err := c.cc.Invoke(ctx, "/teams.TeamService/update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) Remove(ctx context.Context, in *RemoveTeamRequest, opts ...grpc.CallOption) (*RemoveTeamResponse, error) {
	out := new(RemoveTeamResponse)
	err := c.cc.Invoke(ctx, "/teams.TeamService/remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TeamServiceServer is the server API for TeamService service.
type TeamServiceServer interface {
	Query(*QueryRequest, TeamService_QueryServer) error
	View(context.Context, *ViewRequest) (*ViewResponse, error)
	Create(context.Context, *NewTeamRequest) (*NewTeamResponse, error)
	Update(context.Context, *UpdateTeamRequest) (*UpdateTeamResponse, error)
	Remove(context.Context, *RemoveTeamRequest) (*RemoveTeamResponse, error)
}

// UnimplementedTeamServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTeamServiceServer struct {
}

func (*UnimplementedTeamServiceServer) Query(req *QueryRequest, srv TeamService_QueryServer) error {
	return status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (*UnimplementedTeamServiceServer) View(ctx context.Context, req *ViewRequest) (*ViewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method View not implemented")
}
func (*UnimplementedTeamServiceServer) Create(ctx context.Context, req *NewTeamRequest) (*NewTeamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedTeamServiceServer) Update(ctx context.Context, req *UpdateTeamRequest) (*UpdateTeamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedTeamServiceServer) Remove(ctx context.Context, req *RemoveTeamRequest) (*RemoveTeamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}

func RegisterTeamServiceServer(s *grpc.Server, srv TeamServiceServer) {
	s.RegisterService(&_TeamService_serviceDesc, srv)
}

func _TeamService_Query_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(QueryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TeamServiceServer).Query(m, &teamServiceQueryServer{stream})
}

type TeamService_QueryServer interface {
	Send(*QueryResponse) error
	grpc.ServerStream
}

type teamServiceQueryServer struct {
	grpc.ServerStream
}

func (x *teamServiceQueryServer) Send(m *QueryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _TeamService_View_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).View(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/teams.TeamService/View",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).View(ctx, req.(*ViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/teams.TeamService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).Create(ctx, req.(*NewTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/teams.TeamService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).Update(ctx, req.(*UpdateTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/teams.TeamService/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).Remove(ctx, req.(*RemoveTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TeamService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "teams.TeamService",
	HandlerType: (*TeamServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "view",
			Handler:    _TeamService_View_Handler,
		},
		{
			MethodName: "create",
			Handler:    _TeamService_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _TeamService_Update_Handler,
		},
		{
			MethodName: "remove",
			Handler:    _TeamService_Remove_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "query",
			Handler:       _TeamService_Query_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/teams.proto",
}
