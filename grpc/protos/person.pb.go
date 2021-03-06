// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: protos/person.proto

package grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Friend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Friend) Reset() {
	*x = Friend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_person_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Friend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Friend) ProtoMessage() {}

func (x *Friend) ProtoReflect() protoreflect.Message {
	mi := &file_protos_person_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Friend.ProtoReflect.Descriptor instead.
func (*Friend) Descriptor() ([]byte, []int) {
	return file_protos_person_proto_rawDescGZIP(), []int{0}
}

func (x *Friend) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Friend) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	XId            string    `protobuf:"bytes,1,opt,name=_id,json=Id,proto3" json:"_id,omitempty"`
	Index          int32     `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	Guid           string    `protobuf:"bytes,3,opt,name=guid,proto3" json:"guid,omitempty"`
	IsActive       bool      `protobuf:"varint,4,opt,name=isActive,proto3" json:"isActive,omitempty"`
	Balance        string    `protobuf:"bytes,5,opt,name=balance,proto3" json:"balance,omitempty"`
	Picture        string    `protobuf:"bytes,6,opt,name=picture,proto3" json:"picture,omitempty"`
	Age            int32     `protobuf:"varint,7,opt,name=age,proto3" json:"age,omitempty"`
	EyeColor       string    `protobuf:"bytes,8,opt,name=eyeColor,proto3" json:"eyeColor,omitempty"`
	Name           string    `protobuf:"bytes,9,opt,name=name,proto3" json:"name,omitempty"`
	Gender         string    `protobuf:"bytes,10,opt,name=gender,proto3" json:"gender,omitempty"`
	Company        string    `protobuf:"bytes,11,opt,name=company,proto3" json:"company,omitempty"`
	Email          string    `protobuf:"bytes,12,opt,name=email,proto3" json:"email,omitempty"`
	Phone          string    `protobuf:"bytes,13,opt,name=phone,proto3" json:"phone,omitempty"`
	Address        string    `protobuf:"bytes,14,opt,name=address,proto3" json:"address,omitempty"`
	About          string    `protobuf:"bytes,15,opt,name=about,proto3" json:"about,omitempty"`
	Registered     string    `protobuf:"bytes,16,opt,name=registered,proto3" json:"registered,omitempty"`
	Latitude       float32   `protobuf:"fixed32,17,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude      float32   `protobuf:"fixed32,18,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Tags           []string  `protobuf:"bytes,19,rep,name=tags,proto3" json:"tags,omitempty"`
	Friends        []*Friend `protobuf:"bytes,20,rep,name=friends,proto3" json:"friends,omitempty"`
	Greeting       string    `protobuf:"bytes,21,opt,name=greeting,proto3" json:"greeting,omitempty"`
	FavouriteFruit string    `protobuf:"bytes,22,opt,name=favouriteFruit,proto3" json:"favouriteFruit,omitempty"`
}

func (x *Person) Reset() {
	*x = Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_person_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_protos_person_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_protos_person_proto_rawDescGZIP(), []int{1}
}

func (x *Person) GetXId() string {
	if x != nil {
		return x.XId
	}
	return ""
}

func (x *Person) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Person) GetGuid() string {
	if x != nil {
		return x.Guid
	}
	return ""
}

func (x *Person) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *Person) GetBalance() string {
	if x != nil {
		return x.Balance
	}
	return ""
}

func (x *Person) GetPicture() string {
	if x != nil {
		return x.Picture
	}
	return ""
}

func (x *Person) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Person) GetEyeColor() string {
	if x != nil {
		return x.EyeColor
	}
	return ""
}

func (x *Person) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Person) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *Person) GetCompany() string {
	if x != nil {
		return x.Company
	}
	return ""
}

func (x *Person) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Person) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Person) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Person) GetAbout() string {
	if x != nil {
		return x.About
	}
	return ""
}

func (x *Person) GetRegistered() string {
	if x != nil {
		return x.Registered
	}
	return ""
}

func (x *Person) GetLatitude() float32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Person) GetLongitude() float32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *Person) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Person) GetFriends() []*Friend {
	if x != nil {
		return x.Friends
	}
	return nil
}

func (x *Person) GetGreeting() string {
	if x != nil {
		return x.Greeting
	}
	return ""
}

func (x *Person) GetFavouriteFruit() string {
	if x != nil {
		return x.FavouriteFruit
	}
	return ""
}

type Persons struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Persons []*Person `protobuf:"bytes,1,rep,name=persons,proto3" json:"persons,omitempty"`
}

func (x *Persons) Reset() {
	*x = Persons{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_person_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Persons) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Persons) ProtoMessage() {}

func (x *Persons) ProtoReflect() protoreflect.Message {
	mi := &file_protos_person_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Persons.ProtoReflect.Descriptor instead.
func (*Persons) Descriptor() ([]byte, []int) {
	return file_protos_person_proto_rawDescGZIP(), []int{2}
}

func (x *Persons) GetPersons() []*Person {
	if x != nil {
		return x.Persons
	}
	return nil
}

var File_protos_person_proto protoreflect.FileDescriptor

var file_protos_person_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x06, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0xb8, 0x04, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x0f, 0x0a, 0x03, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x67, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70,
	0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x79, 0x65, 0x43,
	0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x79, 0x65, 0x43,
	0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x65, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18,
	0x12, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x13, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x12, 0x21, 0x0a, 0x07, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x18,
	0x14, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x07,
	0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74,
	0x69, 0x6e, 0x67, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74,
	0x69, 0x6e, 0x67, 0x12, 0x26, 0x0a, 0x0e, 0x66, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x69, 0x74, 0x65,
	0x46, 0x72, 0x75, 0x69, 0x74, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x66, 0x61, 0x76,
	0x6f, 0x75, 0x72, 0x69, 0x74, 0x65, 0x46, 0x72, 0x75, 0x69, 0x74, 0x22, 0x2c, 0x0a, 0x07, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x12, 0x21, 0x0a, 0x07, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x52, 0x07, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x32, 0xba, 0x01, 0x0a, 0x0d, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x53, 0x6d, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x08, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x73, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x64, 0x69, 0x75,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x08, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x22, 0x00, 0x12, 0x36,
	0x0a, 0x10, 0x47, 0x65, 0x4c, 0x61, 0x65, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x08, 0x2e, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x73, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2e, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_person_proto_rawDescOnce sync.Once
	file_protos_person_proto_rawDescData = file_protos_person_proto_rawDesc
)

func file_protos_person_proto_rawDescGZIP() []byte {
	file_protos_person_proto_rawDescOnce.Do(func() {
		file_protos_person_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_person_proto_rawDescData)
	})
	return file_protos_person_proto_rawDescData
}

var file_protos_person_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protos_person_proto_goTypes = []interface{}{
	(*Friend)(nil),        // 0: Friend
	(*Person)(nil),        // 1: Person
	(*Persons)(nil),       // 2: Persons
	(*emptypb.Empty)(nil), // 3: google.protobuf.Empty
}
var file_protos_person_proto_depIdxs = []int32{
	0, // 0: Person.friends:type_name -> Friend
	1, // 1: Persons.persons:type_name -> Person
	3, // 2: PersonService.GetSmallRespponse:input_type -> google.protobuf.Empty
	3, // 3: PersonService.GetMediumRespponse:input_type -> google.protobuf.Empty
	3, // 4: PersonService.GeLaegeRespponse:input_type -> google.protobuf.Empty
	2, // 5: PersonService.GetSmallRespponse:output_type -> Persons
	2, // 6: PersonService.GetMediumRespponse:output_type -> Persons
	2, // 7: PersonService.GeLaegeRespponse:output_type -> Persons
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_person_proto_init() }
func file_protos_person_proto_init() {
	if File_protos_person_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_person_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Friend); i {
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
		file_protos_person_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person); i {
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
		file_protos_person_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Persons); i {
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
			RawDescriptor: file_protos_person_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_person_proto_goTypes,
		DependencyIndexes: file_protos_person_proto_depIdxs,
		MessageInfos:      file_protos_person_proto_msgTypes,
	}.Build()
	File_protos_person_proto = out.File
	file_protos_person_proto_rawDesc = nil
	file_protos_person_proto_goTypes = nil
	file_protos_person_proto_depIdxs = nil
}
