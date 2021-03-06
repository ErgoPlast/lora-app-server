// Code generated by protoc-gen-go. DO NOT EDIT.
// source: deviceProfile.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateDeviceProfileRequest struct {
	// Device-profile object to create.
	DeviceProfile        *DeviceProfile `protobuf:"bytes,1,opt,name=device_profile,json=deviceProfile,proto3" json:"device_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CreateDeviceProfileRequest) Reset()         { *m = CreateDeviceProfileRequest{} }
func (m *CreateDeviceProfileRequest) String() string { return proto.CompactTextString(m) }
func (*CreateDeviceProfileRequest) ProtoMessage()    {}
func (*CreateDeviceProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_deviceProfile_b4dd410e0cee7d90, []int{0}
}
func (m *CreateDeviceProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDeviceProfileRequest.Unmarshal(m, b)
}
func (m *CreateDeviceProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDeviceProfileRequest.Marshal(b, m, deterministic)
}
func (dst *CreateDeviceProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDeviceProfileRequest.Merge(dst, src)
}
func (m *CreateDeviceProfileRequest) XXX_Size() int {
	return xxx_messageInfo_CreateDeviceProfileRequest.Size(m)
}
func (m *CreateDeviceProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDeviceProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDeviceProfileRequest proto.InternalMessageInfo

func (m *CreateDeviceProfileRequest) GetDeviceProfile() *DeviceProfile {
	if m != nil {
		return m.DeviceProfile
	}
	return nil
}

type CreateDeviceProfileResponse struct {
	// Device-profile ID (UUID string).
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateDeviceProfileResponse) Reset()         { *m = CreateDeviceProfileResponse{} }
func (m *CreateDeviceProfileResponse) String() string { return proto.CompactTextString(m) }
func (*CreateDeviceProfileResponse) ProtoMessage()    {}
func (*CreateDeviceProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_deviceProfile_b4dd410e0cee7d90, []int{1}
}
func (m *CreateDeviceProfileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDeviceProfileResponse.Unmarshal(m, b)
}
func (m *CreateDeviceProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDeviceProfileResponse.Marshal(b, m, deterministic)
}
func (dst *CreateDeviceProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDeviceProfileResponse.Merge(dst, src)
}
func (m *CreateDeviceProfileResponse) XXX_Size() int {
	return xxx_messageInfo_CreateDeviceProfileResponse.Size(m)
}
func (m *CreateDeviceProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDeviceProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDeviceProfileResponse proto.InternalMessageInfo

func (m *CreateDeviceProfileResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetDeviceProfileRequest struct {
	// Device-profile ID (UUID string).
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDeviceProfileRequest) Reset()         { *m = GetDeviceProfileRequest{} }
func (m *GetDeviceProfileRequest) String() string { return proto.CompactTextString(m) }
func (*GetDeviceProfileRequest) ProtoMessage()    {}
func (*GetDeviceProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_deviceProfile_b4dd410e0cee7d90, []int{2}
}
func (m *GetDeviceProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDeviceProfileRequest.Unmarshal(m, b)
}
func (m *GetDeviceProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDeviceProfileRequest.Marshal(b, m, deterministic)
}
func (dst *GetDeviceProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDeviceProfileRequest.Merge(dst, src)
}
func (m *GetDeviceProfileRequest) XXX_Size() int {
	return xxx_messageInfo_GetDeviceProfileRequest.Size(m)
}
func (m *GetDeviceProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDeviceProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDeviceProfileRequest proto.InternalMessageInfo

func (m *GetDeviceProfileRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetDeviceProfileResponse struct {
	// Device-profile object.
	DeviceProfile *DeviceProfile `protobuf:"bytes,1,opt,name=device_profile,json=deviceProfile,proto3" json:"device_profile,omitempty"`
	// Created at timestamp.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Last update timestamp.
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetDeviceProfileResponse) Reset()         { *m = GetDeviceProfileResponse{} }
func (m *GetDeviceProfileResponse) String() string { return proto.CompactTextString(m) }
func (*GetDeviceProfileResponse) ProtoMessage()    {}
func (*GetDeviceProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_deviceProfile_b4dd410e0cee7d90, []int{3}
}
func (m *GetDeviceProfileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDeviceProfileResponse.Unmarshal(m, b)
}
func (m *GetDeviceProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDeviceProfileResponse.Marshal(b, m, deterministic)
}
func (dst *GetDeviceProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDeviceProfileResponse.Merge(dst, src)
}
func (m *GetDeviceProfileResponse) XXX_Size() int {
	return xxx_messageInfo_GetDeviceProfileResponse.Size(m)
}
func (m *GetDeviceProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDeviceProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDeviceProfileResponse proto.InternalMessageInfo

func (m *GetDeviceProfileResponse) GetDeviceProfile() *DeviceProfile {
	if m != nil {
		return m.DeviceProfile
	}
	return nil
}

func (m *GetDeviceProfileResponse) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *GetDeviceProfileResponse) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type UpdateDeviceProfileRequest struct {
	// Device-profile object to update.
	DeviceProfile        *DeviceProfile `protobuf:"bytes,1,opt,name=device_profile,json=deviceProfile,proto3" json:"device_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UpdateDeviceProfileRequest) Reset()         { *m = UpdateDeviceProfileRequest{} }
func (m *UpdateDeviceProfileRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateDeviceProfileRequest) ProtoMessage()    {}
func (*UpdateDeviceProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_deviceProfile_b4dd410e0cee7d90, []int{4}
}
func (m *UpdateDeviceProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateDeviceProfileRequest.Unmarshal(m, b)
}
func (m *UpdateDeviceProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateDeviceProfileRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateDeviceProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateDeviceProfileRequest.Merge(dst, src)
}
func (m *UpdateDeviceProfileRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateDeviceProfileRequest.Size(m)
}
func (m *UpdateDeviceProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateDeviceProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateDeviceProfileRequest proto.InternalMessageInfo

func (m *UpdateDeviceProfileRequest) GetDeviceProfile() *DeviceProfile {
	if m != nil {
		return m.DeviceProfile
	}
	return nil
}

type DeleteDeviceProfileRequest struct {
	// Device-profile ID (UUID string).
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteDeviceProfileRequest) Reset()         { *m = DeleteDeviceProfileRequest{} }
func (m *DeleteDeviceProfileRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteDeviceProfileRequest) ProtoMessage()    {}
func (*DeleteDeviceProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_deviceProfile_b4dd410e0cee7d90, []int{5}
}
func (m *DeleteDeviceProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteDeviceProfileRequest.Unmarshal(m, b)
}
func (m *DeleteDeviceProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteDeviceProfileRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteDeviceProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteDeviceProfileRequest.Merge(dst, src)
}
func (m *DeleteDeviceProfileRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteDeviceProfileRequest.Size(m)
}
func (m *DeleteDeviceProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteDeviceProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteDeviceProfileRequest proto.InternalMessageInfo

func (m *DeleteDeviceProfileRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeviceProfileListItem struct {
	// Device-profile ID (UUID string).
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Device-profile name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Organization ID.
	OrganizationId int64 `protobuf:"varint,3,opt,name=organization_id,json=organizationID,proto3" json:"organization_id,omitempty"`
	// Network-server ID.
	NetworkServerId int64 `protobuf:"varint,4,opt,name=network_server_id,json=networkServerID,proto3" json:"network_server_id,omitempty"`
	// Created at timestamp.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Last update timestamp.
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *DeviceProfileListItem) Reset()         { *m = DeviceProfileListItem{} }
func (m *DeviceProfileListItem) String() string { return proto.CompactTextString(m) }
func (*DeviceProfileListItem) ProtoMessage()    {}
func (*DeviceProfileListItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_deviceProfile_b4dd410e0cee7d90, []int{6}
}
func (m *DeviceProfileListItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceProfileListItem.Unmarshal(m, b)
}
func (m *DeviceProfileListItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceProfileListItem.Marshal(b, m, deterministic)
}
func (dst *DeviceProfileListItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceProfileListItem.Merge(dst, src)
}
func (m *DeviceProfileListItem) XXX_Size() int {
	return xxx_messageInfo_DeviceProfileListItem.Size(m)
}
func (m *DeviceProfileListItem) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceProfileListItem.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceProfileListItem proto.InternalMessageInfo

func (m *DeviceProfileListItem) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DeviceProfileListItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeviceProfileListItem) GetOrganizationId() int64 {
	if m != nil {
		return m.OrganizationId
	}
	return 0
}

func (m *DeviceProfileListItem) GetNetworkServerId() int64 {
	if m != nil {
		return m.NetworkServerId
	}
	return 0
}

func (m *DeviceProfileListItem) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *DeviceProfileListItem) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type ListDeviceProfileRequest struct {
	// Max number of items to return.
	Limit int64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	// Offset in the result-set (for pagination).
	Offset int64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	// Organization id to filter on.
	OrganizationId int64 `protobuf:"varint,3,opt,name=organization_id,json=organizationID,proto3" json:"organization_id,omitempty"`
	// Application id to filter on.
	ApplicationId        int64    `protobuf:"varint,4,opt,name=application_id,json=applicationID,proto3" json:"application_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListDeviceProfileRequest) Reset()         { *m = ListDeviceProfileRequest{} }
func (m *ListDeviceProfileRequest) String() string { return proto.CompactTextString(m) }
func (*ListDeviceProfileRequest) ProtoMessage()    {}
func (*ListDeviceProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_deviceProfile_b4dd410e0cee7d90, []int{7}
}
func (m *ListDeviceProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListDeviceProfileRequest.Unmarshal(m, b)
}
func (m *ListDeviceProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListDeviceProfileRequest.Marshal(b, m, deterministic)
}
func (dst *ListDeviceProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListDeviceProfileRequest.Merge(dst, src)
}
func (m *ListDeviceProfileRequest) XXX_Size() int {
	return xxx_messageInfo_ListDeviceProfileRequest.Size(m)
}
func (m *ListDeviceProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListDeviceProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListDeviceProfileRequest proto.InternalMessageInfo

func (m *ListDeviceProfileRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListDeviceProfileRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ListDeviceProfileRequest) GetOrganizationId() int64 {
	if m != nil {
		return m.OrganizationId
	}
	return 0
}

func (m *ListDeviceProfileRequest) GetApplicationId() int64 {
	if m != nil {
		return m.ApplicationId
	}
	return 0
}

type ListDeviceProfileResponse struct {
	// Total number of device-profiles.
	TotalCount           int64                    `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	Result               []*DeviceProfileListItem `protobuf:"bytes,2,rep,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ListDeviceProfileResponse) Reset()         { *m = ListDeviceProfileResponse{} }
func (m *ListDeviceProfileResponse) String() string { return proto.CompactTextString(m) }
func (*ListDeviceProfileResponse) ProtoMessage()    {}
func (*ListDeviceProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_deviceProfile_b4dd410e0cee7d90, []int{8}
}
func (m *ListDeviceProfileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListDeviceProfileResponse.Unmarshal(m, b)
}
func (m *ListDeviceProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListDeviceProfileResponse.Marshal(b, m, deterministic)
}
func (dst *ListDeviceProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListDeviceProfileResponse.Merge(dst, src)
}
func (m *ListDeviceProfileResponse) XXX_Size() int {
	return xxx_messageInfo_ListDeviceProfileResponse.Size(m)
}
func (m *ListDeviceProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListDeviceProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListDeviceProfileResponse proto.InternalMessageInfo

func (m *ListDeviceProfileResponse) GetTotalCount() int64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *ListDeviceProfileResponse) GetResult() []*DeviceProfileListItem {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateDeviceProfileRequest)(nil), "api.CreateDeviceProfileRequest")
	proto.RegisterType((*CreateDeviceProfileResponse)(nil), "api.CreateDeviceProfileResponse")
	proto.RegisterType((*GetDeviceProfileRequest)(nil), "api.GetDeviceProfileRequest")
	proto.RegisterType((*GetDeviceProfileResponse)(nil), "api.GetDeviceProfileResponse")
	proto.RegisterType((*UpdateDeviceProfileRequest)(nil), "api.UpdateDeviceProfileRequest")
	proto.RegisterType((*DeleteDeviceProfileRequest)(nil), "api.DeleteDeviceProfileRequest")
	proto.RegisterType((*DeviceProfileListItem)(nil), "api.DeviceProfileListItem")
	proto.RegisterType((*ListDeviceProfileRequest)(nil), "api.ListDeviceProfileRequest")
	proto.RegisterType((*ListDeviceProfileResponse)(nil), "api.ListDeviceProfileResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DeviceProfileServiceClient is the client API for DeviceProfileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DeviceProfileServiceClient interface {
	// Create creates the given device-profile.
	Create(ctx context.Context, in *CreateDeviceProfileRequest, opts ...grpc.CallOption) (*CreateDeviceProfileResponse, error)
	// Get returns the device-profile matching the given id.
	Get(ctx context.Context, in *GetDeviceProfileRequest, opts ...grpc.CallOption) (*GetDeviceProfileResponse, error)
	// Update updates the given device-profile.
	Update(ctx context.Context, in *UpdateDeviceProfileRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Delete deletes the device-profile matching the given id.
	Delete(ctx context.Context, in *DeleteDeviceProfileRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// List lists the available device-profiles.
	List(ctx context.Context, in *ListDeviceProfileRequest, opts ...grpc.CallOption) (*ListDeviceProfileResponse, error)
}

type deviceProfileServiceClient struct {
	cc *grpc.ClientConn
}

func NewDeviceProfileServiceClient(cc *grpc.ClientConn) DeviceProfileServiceClient {
	return &deviceProfileServiceClient{cc}
}

func (c *deviceProfileServiceClient) Create(ctx context.Context, in *CreateDeviceProfileRequest, opts ...grpc.CallOption) (*CreateDeviceProfileResponse, error) {
	out := new(CreateDeviceProfileResponse)
	err := c.cc.Invoke(ctx, "/api.DeviceProfileService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceProfileServiceClient) Get(ctx context.Context, in *GetDeviceProfileRequest, opts ...grpc.CallOption) (*GetDeviceProfileResponse, error) {
	out := new(GetDeviceProfileResponse)
	err := c.cc.Invoke(ctx, "/api.DeviceProfileService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceProfileServiceClient) Update(ctx context.Context, in *UpdateDeviceProfileRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.DeviceProfileService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceProfileServiceClient) Delete(ctx context.Context, in *DeleteDeviceProfileRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.DeviceProfileService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceProfileServiceClient) List(ctx context.Context, in *ListDeviceProfileRequest, opts ...grpc.CallOption) (*ListDeviceProfileResponse, error) {
	out := new(ListDeviceProfileResponse)
	err := c.cc.Invoke(ctx, "/api.DeviceProfileService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeviceProfileServiceServer is the server API for DeviceProfileService service.
type DeviceProfileServiceServer interface {
	// Create creates the given device-profile.
	Create(context.Context, *CreateDeviceProfileRequest) (*CreateDeviceProfileResponse, error)
	// Get returns the device-profile matching the given id.
	Get(context.Context, *GetDeviceProfileRequest) (*GetDeviceProfileResponse, error)
	// Update updates the given device-profile.
	Update(context.Context, *UpdateDeviceProfileRequest) (*empty.Empty, error)
	// Delete deletes the device-profile matching the given id.
	Delete(context.Context, *DeleteDeviceProfileRequest) (*empty.Empty, error)
	// List lists the available device-profiles.
	List(context.Context, *ListDeviceProfileRequest) (*ListDeviceProfileResponse, error)
}

func RegisterDeviceProfileServiceServer(s *grpc.Server, srv DeviceProfileServiceServer) {
	s.RegisterService(&_DeviceProfileService_serviceDesc, srv)
}

func _DeviceProfileService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDeviceProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceProfileServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceProfileService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceProfileServiceServer).Create(ctx, req.(*CreateDeviceProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceProfileService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceProfileServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceProfileService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceProfileServiceServer).Get(ctx, req.(*GetDeviceProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceProfileService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDeviceProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceProfileServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceProfileService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceProfileServiceServer).Update(ctx, req.(*UpdateDeviceProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceProfileService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDeviceProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceProfileServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceProfileService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceProfileServiceServer).Delete(ctx, req.(*DeleteDeviceProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceProfileService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDeviceProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceProfileServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceProfileService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceProfileServiceServer).List(ctx, req.(*ListDeviceProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DeviceProfileService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.DeviceProfileService",
	HandlerType: (*DeviceProfileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _DeviceProfileService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _DeviceProfileService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _DeviceProfileService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _DeviceProfileService_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _DeviceProfileService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "deviceProfile.proto",
}

func init() { proto.RegisterFile("deviceProfile.proto", fileDescriptor_deviceProfile_b4dd410e0cee7d90) }

var fileDescriptor_deviceProfile_b4dd410e0cee7d90 = []byte{
	// 616 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x56, 0xe2, 0xd4, 0x52, 0xa7, 0x6a, 0x2a, 0x96, 0x52, 0x52, 0x27, 0x90, 0x62, 0x09, 0x51,
	0x22, 0xe2, 0x48, 0xe9, 0xa9, 0xdc, 0xaa, 0x06, 0x55, 0x91, 0x38, 0x20, 0x03, 0xe2, 0x18, 0x6d,
	0xed, 0x49, 0xb4, 0xc2, 0xf6, 0x2e, 0xf6, 0x26, 0x08, 0x50, 0x2f, 0x1c, 0x78, 0x01, 0x2e, 0x3c,
	0x05, 0x2f, 0xc2, 0x91, 0x57, 0xe0, 0x41, 0x90, 0x77, 0xd7, 0x28, 0x3f, 0x36, 0x3f, 0x15, 0x37,
	0x7b, 0xe7, 0x9b, 0x99, 0x6f, 0xbe, 0xf9, 0x81, 0x9b, 0x21, 0x2e, 0x58, 0x80, 0xcf, 0x52, 0x3e,
	0x65, 0x11, 0x7a, 0x22, 0xe5, 0x92, 0x13, 0x8b, 0x0a, 0xe6, 0x74, 0x66, 0x9c, 0xcf, 0x22, 0x1c,
	0x50, 0xc1, 0x06, 0x34, 0x49, 0xb8, 0xa4, 0x92, 0xf1, 0x24, 0xd3, 0x10, 0xa7, 0x6b, 0xac, 0xea,
	0xef, 0x72, 0x3e, 0x1d, 0x48, 0x16, 0x63, 0x26, 0x69, 0x2c, 0x0c, 0xa0, 0xbd, 0x0e, 0xc0, 0x58,
	0xc8, 0x77, 0xc6, 0xd8, 0x14, 0x3a, 0x9f, 0x89, 0xe6, 0xbe, 0x02, 0xe7, 0x3c, 0x45, 0x2a, 0x71,
	0xb4, 0xcc, 0xc6, 0xc7, 0x37, 0x73, 0xcc, 0x24, 0x39, 0x85, 0xa6, 0x66, 0x39, 0x31, 0x6e, 0xad,
	0xda, 0x51, 0xed, 0x78, 0x67, 0x48, 0x3c, 0x2a, 0x98, 0xb7, 0xea, 0xb2, 0xbb, 0x52, 0x8f, 0xdb,
	0x87, 0x76, 0x69, 0xe0, 0x4c, 0xf0, 0x24, 0x43, 0xd2, 0x84, 0x3a, 0x0b, 0x55, 0xb4, 0x6d, 0xbf,
	0xce, 0x42, 0xf7, 0x21, 0xdc, 0xbe, 0x40, 0x59, 0x4a, 0x62, 0x1d, 0xfa, 0xad, 0x06, 0xad, 0x4d,
	0xac, 0x89, 0x7b, 0x7d, 0xc6, 0xe4, 0x14, 0x20, 0x50, 0x8c, 0xc3, 0x09, 0x95, 0xad, 0xba, 0x72,
	0x73, 0x3c, 0x2d, 0xa6, 0x57, 0x88, 0xe9, 0xbd, 0x28, 0xd4, 0xf6, 0xb7, 0x0d, 0xfa, 0x2c, 0xd7,
	0x09, 0xe6, 0x22, 0x2c, 0x5c, 0xad, 0x3f, 0xbb, 0x1a, 0xf4, 0x99, 0xcc, 0x1b, 0xf0, 0x52, 0xfd,
	0xfc, 0xef, 0x06, 0x3c, 0x02, 0x67, 0x84, 0x11, 0x56, 0x04, 0x5e, 0x17, 0xf5, 0x53, 0x1d, 0x6e,
	0xad, 0x00, 0x9f, 0xb2, 0x4c, 0x8e, 0x25, 0xc6, 0xeb, 0x48, 0x42, 0xa0, 0x91, 0xd0, 0x18, 0x95,
	0x40, 0xdb, 0xbe, 0xfa, 0x26, 0x0f, 0x60, 0x8f, 0xa7, 0x33, 0x9a, 0xb0, 0xf7, 0x6a, 0x54, 0x27,
	0x2c, 0x54, 0x22, 0x58, 0x7e, 0x73, 0xf9, 0x79, 0x3c, 0x22, 0x3d, 0xb8, 0x91, 0xa0, 0x7c, 0xcb,
	0xd3, 0xd7, 0x93, 0x0c, 0xd3, 0x05, 0xa6, 0x39, 0xb4, 0xa1, 0xa0, 0x7b, 0xc6, 0xf0, 0x5c, 0xbd,
	0x8f, 0x47, 0x6b, 0xfd, 0xd8, 0xba, 0x7e, 0x3f, 0xec, 0x7f, 0xe9, 0xc7, 0x97, 0x1a, 0xb4, 0xf2,
	0xda, 0x4b, 0x55, 0xdb, 0x87, 0xad, 0x88, 0xc5, 0x4c, 0x2a, 0x39, 0x2c, 0x5f, 0xff, 0x90, 0x03,
	0xb0, 0xf9, 0x74, 0x9a, 0xa1, 0x1e, 0x1a, 0xcb, 0x37, 0x7f, 0x7f, 0xaf, 0xca, 0x7d, 0x68, 0x52,
	0x21, 0x22, 0x16, 0xfc, 0xc2, 0x69, 0x49, 0x76, 0x97, 0x5e, 0xc7, 0x23, 0x57, 0xc0, 0x61, 0x09,
	0x33, 0x33, 0xf8, 0x5d, 0xd8, 0x91, 0x5c, 0xd2, 0x68, 0x12, 0xf0, 0x79, 0x52, 0x10, 0x04, 0xf5,
	0x74, 0x9e, 0xbf, 0x90, 0x21, 0xd8, 0x29, 0x66, 0xf3, 0x28, 0x67, 0x69, 0x29, 0x3d, 0x36, 0x46,
	0xa8, 0xe8, 0xb9, 0x6f, 0x90, 0xc3, 0xaf, 0x0d, 0xd8, 0x5f, 0x41, 0xe4, 0xcd, 0x61, 0x01, 0x92,
	0x08, 0x6c, 0xbd, 0xdd, 0xa4, 0xab, 0xc2, 0x54, 0xdf, 0x10, 0xe7, 0xa8, 0x1a, 0xa0, 0xa9, 0xbb,
	0xdd, 0x8f, 0xdf, 0x7f, 0x7c, 0xae, 0x1f, 0xba, 0xfb, 0xea, 0xe2, 0xe9, 0x29, 0xee, 0x17, 0x77,
	0xea, 0x71, 0xad, 0x47, 0x10, 0xac, 0x0b, 0x94, 0xa4, 0xa3, 0x22, 0x55, 0x9c, 0x09, 0xe7, 0x4e,
	0x85, 0xd5, 0x24, 0xb9, 0xa7, 0x92, 0xb4, 0xc9, 0x61, 0x59, 0x92, 0xc1, 0x07, 0x16, 0x5e, 0x91,
	0x05, 0xd8, 0x7a, 0x15, 0x4d, 0x51, 0xd5, 0x7b, 0xe9, 0x1c, 0x6c, 0x0c, 0xd3, 0x93, 0xfc, 0xc8,
	0xba, 0x27, 0x2a, 0x4b, 0xdf, 0x39, 0x2e, 0xcf, 0xb2, 0xba, 0xcb, 0x1e, 0x0b, 0xaf, 0xf2, 0xf2,
	0x42, 0xb0, 0xf5, 0xa6, 0x9a, 0xbc, 0xd5, 0x6b, 0x5b, 0x99, 0xd7, 0x54, 0xd7, 0xfb, 0x4d, 0x75,
	0x01, 0x34, 0xf2, 0xfe, 0x12, 0xad, 0x53, 0xd5, 0x88, 0x3b, 0x77, 0xab, 0xcc, 0x46, 0xc7, 0x8e,
	0xca, 0x74, 0x40, 0x4a, 0x9b, 0x75, 0x69, 0x2b, 0x5e, 0x27, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff,
	0xce, 0xf1, 0xe7, 0xa5, 0xdd, 0x06, 0x00, 0x00,
}
