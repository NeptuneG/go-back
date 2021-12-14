// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: services/live/proto/live_event.proto

package proto

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

type LiveEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	LiveHouse       *LiveHouse             `protobuf:"bytes,2,opt,name=live_house,json=liveHouse,proto3" json:"live_house,omitempty"`
	Title           string                 `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Url             string                 `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	Description     string                 `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	PriceInfo       string                 `protobuf:"bytes,6,opt,name=price_info,json=priceInfo,proto3" json:"price_info,omitempty"`
	StageOneOpenAt  *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=stage_one_open_at,json=stageOneOpenAt,proto3" json:"stage_one_open_at,omitempty"`
	StageOneStartAt *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=stage_one_start_at,json=stageOneStartAt,proto3" json:"stage_one_start_at,omitempty"`
	StageTwoOpenAt  *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=stage_two_open_at,json=stageTwoOpenAt,proto3" json:"stage_two_open_at,omitempty"`
	StageTwoStartAt *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=stage_two_start_at,json=stageTwoStartAt,proto3" json:"stage_two_start_at,omitempty"`
	AvailableSeats  int32                  `protobuf:"varint,11,opt,name=available_seats,json=availableSeats,proto3" json:"available_seats,omitempty"`
}

func (x *LiveEvent) Reset() {
	*x = LiveEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_live_proto_live_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LiveEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LiveEvent) ProtoMessage() {}

func (x *LiveEvent) ProtoReflect() protoreflect.Message {
	mi := &file_services_live_proto_live_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LiveEvent.ProtoReflect.Descriptor instead.
func (*LiveEvent) Descriptor() ([]byte, []int) {
	return file_services_live_proto_live_event_proto_rawDescGZIP(), []int{0}
}

func (x *LiveEvent) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *LiveEvent) GetLiveHouse() *LiveHouse {
	if x != nil {
		return x.LiveHouse
	}
	return nil
}

func (x *LiveEvent) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *LiveEvent) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *LiveEvent) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *LiveEvent) GetPriceInfo() string {
	if x != nil {
		return x.PriceInfo
	}
	return ""
}

func (x *LiveEvent) GetStageOneOpenAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StageOneOpenAt
	}
	return nil
}

func (x *LiveEvent) GetStageOneStartAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StageOneStartAt
	}
	return nil
}

func (x *LiveEvent) GetStageTwoOpenAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StageTwoOpenAt
	}
	return nil
}

func (x *LiveEvent) GetStageTwoStartAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StageTwoStartAt
	}
	return nil
}

func (x *LiveEvent) GetAvailableSeats() int32 {
	if x != nil {
		return x.AvailableSeats
	}
	return 0
}

var File_services_live_proto_live_event_proto protoreflect.FileDescriptor

var file_services_live_proto_live_event_proto_rawDesc = []byte{
	0x0a, 0x24, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x6e, 0x65, 0x70, 0x74, 0x75, 0x6e, 0x65, 0x67,
	0x2e, 0x67, 0x6f, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x69, 0x76, 0x63, 0x65,
	0x73, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x5f, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x04, 0x0a, 0x09, 0x4c, 0x69, 0x76, 0x65, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x4e, 0x0a, 0x0a, 0x6c, 0x69, 0x76, 0x65, 0x5f, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x6e, 0x65, 0x70, 0x74, 0x75, 0x6e, 0x65, 0x67,
	0x2e, 0x67, 0x6f, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x69, 0x76, 0x63, 0x65,
	0x73, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x5f, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x4c, 0x69, 0x76,
	0x65, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x09, 0x6c, 0x69, 0x76, 0x65, 0x48, 0x6f, 0x75, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x72, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x45, 0x0a, 0x11, 0x73, 0x74,
	0x61, 0x67, 0x65, 0x5f, 0x6f, 0x6e, 0x65, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x61, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0e, 0x73, 0x74, 0x61, 0x67, 0x65, 0x4f, 0x6e, 0x65, 0x4f, 0x70, 0x65, 0x6e, 0x41,
	0x74, 0x12, 0x47, 0x0a, 0x12, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x6f, 0x6e, 0x65, 0x5f, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f, 0x73, 0x74, 0x61, 0x67, 0x65,
	0x4f, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x12, 0x45, 0x0a, 0x11, 0x73, 0x74,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x77, 0x6f, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x61, 0x74, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0e, 0x73, 0x74, 0x61, 0x67, 0x65, 0x54, 0x77, 0x6f, 0x4f, 0x70, 0x65, 0x6e, 0x41,
	0x74, 0x12, 0x47, 0x0a, 0x12, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x77, 0x6f, 0x5f, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f, 0x73, 0x74, 0x61, 0x67, 0x65,
	0x54, 0x77, 0x6f, 0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x61, 0x74, 0x73, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65,
	0x61, 0x74, 0x73, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4e, 0x65, 0x70, 0x74, 0x75, 0x6e, 0x65, 0x47, 0x2f, 0x67, 0x6f, 0x2d, 0x62, 0x61,
	0x63, 0x6b, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6c, 0x69, 0x76, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_live_proto_live_event_proto_rawDescOnce sync.Once
	file_services_live_proto_live_event_proto_rawDescData = file_services_live_proto_live_event_proto_rawDesc
)

func file_services_live_proto_live_event_proto_rawDescGZIP() []byte {
	file_services_live_proto_live_event_proto_rawDescOnce.Do(func() {
		file_services_live_proto_live_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_live_proto_live_event_proto_rawDescData)
	})
	return file_services_live_proto_live_event_proto_rawDescData
}

var file_services_live_proto_live_event_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_services_live_proto_live_event_proto_goTypes = []interface{}{
	(*LiveEvent)(nil),             // 0: neptuneg.go_back.serivces.live_event.LiveEvent
	(*LiveHouse)(nil),             // 1: neptuneg.go_back.serivces.live_house.LiveHouse
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_services_live_proto_live_event_proto_depIdxs = []int32{
	1, // 0: neptuneg.go_back.serivces.live_event.LiveEvent.live_house:type_name -> neptuneg.go_back.serivces.live_house.LiveHouse
	2, // 1: neptuneg.go_back.serivces.live_event.LiveEvent.stage_one_open_at:type_name -> google.protobuf.Timestamp
	2, // 2: neptuneg.go_back.serivces.live_event.LiveEvent.stage_one_start_at:type_name -> google.protobuf.Timestamp
	2, // 3: neptuneg.go_back.serivces.live_event.LiveEvent.stage_two_open_at:type_name -> google.protobuf.Timestamp
	2, // 4: neptuneg.go_back.serivces.live_event.LiveEvent.stage_two_start_at:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_services_live_proto_live_event_proto_init() }
func file_services_live_proto_live_event_proto_init() {
	if File_services_live_proto_live_event_proto != nil {
		return
	}
	file_services_live_proto_live_house_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_services_live_proto_live_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LiveEvent); i {
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
			RawDescriptor: file_services_live_proto_live_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_services_live_proto_live_event_proto_goTypes,
		DependencyIndexes: file_services_live_proto_live_event_proto_depIdxs,
		MessageInfos:      file_services_live_proto_live_event_proto_msgTypes,
	}.Build()
	File_services_live_proto_live_event_proto = out.File
	file_services_live_proto_live_event_proto_rawDesc = nil
	file_services_live_proto_live_event_proto_goTypes = nil
	file_services_live_proto_live_event_proto_depIdxs = nil
}
