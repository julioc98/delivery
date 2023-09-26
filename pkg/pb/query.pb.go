// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.3
// source: query.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DriverPosition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DriverId  uint64                 `protobuf:"varint,1,opt,name=driver_id,json=driverId,proto3" json:"driver_id,omitempty"`
	Latitude  float64                `protobuf:"fixed64,2,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float64                `protobuf:"fixed64,3,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Id        int64                  `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *DriverPosition) Reset() {
	*x = DriverPosition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_query_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DriverPosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DriverPosition) ProtoMessage() {}

func (x *DriverPosition) ProtoReflect() protoreflect.Message {
	mi := &file_query_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DriverPosition.ProtoReflect.Descriptor instead.
func (*DriverPosition) Descriptor() ([]byte, []int) {
	return file_query_proto_rawDescGZIP(), []int{0}
}

func (x *DriverPosition) GetDriverId() uint64 {
	if x != nil {
		return x.DriverId
	}
	return 0
}

func (x *DriverPosition) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *DriverPosition) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *DriverPosition) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DriverPosition) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

// Request message to find a driver's position.
type FindDriverPositionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DriverId uint64 `protobuf:"varint,1,opt,name=driver_id,json=driverId,proto3" json:"driver_id,omitempty"`
}

func (x *FindDriverPositionRequest) Reset() {
	*x = FindDriverPositionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_query_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindDriverPositionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindDriverPositionRequest) ProtoMessage() {}

func (x *FindDriverPositionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_query_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindDriverPositionRequest.ProtoReflect.Descriptor instead.
func (*FindDriverPositionRequest) Descriptor() ([]byte, []int) {
	return file_query_proto_rawDescGZIP(), []int{1}
}

func (x *FindDriverPositionRequest) GetDriverId() uint64 {
	if x != nil {
		return x.DriverId
	}
	return 0
}

// Response message to find a driver's position.
type FindDriverPositionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Position *DriverPosition `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
}

func (x *FindDriverPositionResponse) Reset() {
	*x = FindDriverPositionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_query_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindDriverPositionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindDriverPositionResponse) ProtoMessage() {}

func (x *FindDriverPositionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_query_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindDriverPositionResponse.ProtoReflect.Descriptor instead.
func (*FindDriverPositionResponse) Descriptor() ([]byte, []int) {
	return file_query_proto_rawDescGZIP(), []int{2}
}

func (x *FindDriverPositionResponse) GetPosition() *DriverPosition {
	if x != nil {
		return x.Position
	}
	return nil
}

// Request message to find a driver's position history.
type HistoryDriverPositionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DriverId uint64 `protobuf:"varint,1,opt,name=driver_id,json=driverId,proto3" json:"driver_id,omitempty"`
}

func (x *HistoryDriverPositionRequest) Reset() {
	*x = HistoryDriverPositionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_query_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoryDriverPositionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoryDriverPositionRequest) ProtoMessage() {}

func (x *HistoryDriverPositionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_query_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoryDriverPositionRequest.ProtoReflect.Descriptor instead.
func (*HistoryDriverPositionRequest) Descriptor() ([]byte, []int) {
	return file_query_proto_rawDescGZIP(), []int{3}
}

func (x *HistoryDriverPositionRequest) GetDriverId() uint64 {
	if x != nil {
		return x.DriverId
	}
	return 0
}

// Response message to find a driver's position history.
type HistoryDriverPositionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Positions []*DriverPosition `protobuf:"bytes,1,rep,name=positions,proto3" json:"positions,omitempty"`
}

func (x *HistoryDriverPositionResponse) Reset() {
	*x = HistoryDriverPositionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_query_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoryDriverPositionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoryDriverPositionResponse) ProtoMessage() {}

func (x *HistoryDriverPositionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_query_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoryDriverPositionResponse.ProtoReflect.Descriptor instead.
func (*HistoryDriverPositionResponse) Descriptor() ([]byte, []int) {
	return file_query_proto_rawDescGZIP(), []int{4}
}

func (x *HistoryDriverPositionResponse) GetPositions() []*DriverPosition {
	if x != nil {
		return x.Positions
	}
	return nil
}

// Request message to find nearby drivers.
type FindDriversNearbyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  float64 `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Radius    int32   `protobuf:"varint,3,opt,name=radius,proto3" json:"radius,omitempty"`
}

func (x *FindDriversNearbyRequest) Reset() {
	*x = FindDriversNearbyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_query_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindDriversNearbyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindDriversNearbyRequest) ProtoMessage() {}

func (x *FindDriversNearbyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_query_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindDriversNearbyRequest.ProtoReflect.Descriptor instead.
func (*FindDriversNearbyRequest) Descriptor() ([]byte, []int) {
	return file_query_proto_rawDescGZIP(), []int{5}
}

func (x *FindDriversNearbyRequest) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *FindDriversNearbyRequest) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *FindDriversNearbyRequest) GetRadius() int32 {
	if x != nil {
		return x.Radius
	}
	return 0
}

// Response message to find nearby drivers.
type FindDriversNearbyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Positions []*DriverPosition `protobuf:"bytes,1,rep,name=positions,proto3" json:"positions,omitempty"`
}

func (x *FindDriversNearbyResponse) Reset() {
	*x = FindDriversNearbyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_query_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindDriversNearbyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindDriversNearbyResponse) ProtoMessage() {}

func (x *FindDriversNearbyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_query_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindDriversNearbyResponse.ProtoReflect.Descriptor instead.
func (*FindDriversNearbyResponse) Descriptor() ([]byte, []int) {
	return file_query_proto_rawDescGZIP(), []int{6}
}

func (x *FindDriversNearbyResponse) GetPositions() []*DriverPosition {
	if x != nil {
		return x.Positions
	}
	return nil
}

var File_query_proto protoreflect.FileDescriptor

var file_query_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x01, 0x0a, 0x0e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x64, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x38, 0x0a, 0x19, 0x46, 0x69, 0x6e,
	0x64, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x4f, 0x0a, 0x1a, 0x46, 0x69, 0x6e, 0x64, 0x44, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x31, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x44, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3b, 0x0a, 0x1c, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x44,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x54, 0x0a, 0x1d, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x44, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x33, 0x0a, 0x09, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x44, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x6c, 0x0a, 0x18, 0x46, 0x69, 0x6e, 0x64, 0x44,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x4e, 0x65, 0x61, 0x72, 0x62, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72,
	0x61, 0x64, 0x69, 0x75, 0x73, 0x22, 0x50, 0x0a, 0x19, 0x46, 0x69, 0x6e, 0x64, 0x44, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x73, 0x4e, 0x65, 0x61, 0x72, 0x62, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x33, 0x0a, 0x09, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x44, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0xa5, 0x02, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x59, 0x0a, 0x12, 0x46, 0x69, 0x6e, 0x64,
	0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20,
	0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x44, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x44, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a, 0x15, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x44, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x2e, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x44, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x24, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x11, 0x46, 0x69, 0x6e, 0x64, 0x44,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x4e, 0x65, 0x61, 0x72, 0x62, 0x79, 0x12, 0x1f, 0x2e, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73,
	0x4e, 0x65, 0x61, 0x72, 0x62, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x73, 0x4e, 0x65, 0x61, 0x72, 0x62, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x75,
	0x6c, 0x69, 0x6f, 0x63, 0x39, 0x38, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_query_proto_rawDescOnce sync.Once
	file_query_proto_rawDescData = file_query_proto_rawDesc
)

func file_query_proto_rawDescGZIP() []byte {
	file_query_proto_rawDescOnce.Do(func() {
		file_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_query_proto_rawDescData)
	})
	return file_query_proto_rawDescData
}

var file_query_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_query_proto_goTypes = []interface{}{
	(*DriverPosition)(nil),                // 0: query.DriverPosition
	(*FindDriverPositionRequest)(nil),     // 1: query.FindDriverPositionRequest
	(*FindDriverPositionResponse)(nil),    // 2: query.FindDriverPositionResponse
	(*HistoryDriverPositionRequest)(nil),  // 3: query.HistoryDriverPositionRequest
	(*HistoryDriverPositionResponse)(nil), // 4: query.HistoryDriverPositionResponse
	(*FindDriversNearbyRequest)(nil),      // 5: query.FindDriversNearbyRequest
	(*FindDriversNearbyResponse)(nil),     // 6: query.FindDriversNearbyResponse
	(*timestamppb.Timestamp)(nil),         // 7: google.protobuf.Timestamp
}
var file_query_proto_depIdxs = []int32{
	7, // 0: query.DriverPosition.timestamp:type_name -> google.protobuf.Timestamp
	0, // 1: query.FindDriverPositionResponse.position:type_name -> query.DriverPosition
	0, // 2: query.HistoryDriverPositionResponse.positions:type_name -> query.DriverPosition
	0, // 3: query.FindDriversNearbyResponse.positions:type_name -> query.DriverPosition
	1, // 4: query.QueryService.FindDriverPosition:input_type -> query.FindDriverPositionRequest
	3, // 5: query.QueryService.HistoryDriverPosition:input_type -> query.HistoryDriverPositionRequest
	5, // 6: query.QueryService.FindDriversNearby:input_type -> query.FindDriversNearbyRequest
	2, // 7: query.QueryService.FindDriverPosition:output_type -> query.FindDriverPositionResponse
	4, // 8: query.QueryService.HistoryDriverPosition:output_type -> query.HistoryDriverPositionResponse
	6, // 9: query.QueryService.FindDriversNearby:output_type -> query.FindDriversNearbyResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_query_proto_init() }
func file_query_proto_init() {
	if File_query_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_query_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DriverPosition); i {
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
		file_query_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindDriverPositionRequest); i {
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
		file_query_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindDriverPositionResponse); i {
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
		file_query_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistoryDriverPositionRequest); i {
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
		file_query_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistoryDriverPositionResponse); i {
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
		file_query_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindDriversNearbyRequest); i {
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
		file_query_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindDriversNearbyResponse); i {
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
			RawDescriptor: file_query_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_query_proto_goTypes,
		DependencyIndexes: file_query_proto_depIdxs,
		MessageInfos:      file_query_proto_msgTypes,
	}.Build()
	File_query_proto = out.File
	file_query_proto_rawDesc = nil
	file_query_proto_goTypes = nil
	file_query_proto_depIdxs = nil
}