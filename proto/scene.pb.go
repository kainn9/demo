// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.2
// source: scene.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TerrainBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type     float32 `protobuf:"fixed32,1,opt,name=type,proto3" json:"type,omitempty"`
	X        float32 `protobuf:"fixed32,2,opt,name=x,proto3" json:"x,omitempty"`
	Y        float32 `protobuf:"fixed32,3,opt,name=y,proto3" json:"y,omitempty"`
	Width    float32 `protobuf:"fixed32,4,opt,name=width,proto3" json:"width,omitempty"`
	Height   float32 `protobuf:"fixed32,5,opt,name=height,proto3" json:"height,omitempty"`
	Rotation float32 `protobuf:"fixed32,6,opt,name=rotation,proto3" json:"rotation,omitempty"`
}

func (x *TerrainBlock) Reset() {
	*x = TerrainBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scene_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TerrainBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerrainBlock) ProtoMessage() {}

func (x *TerrainBlock) ProtoReflect() protoreflect.Message {
	mi := &file_scene_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerrainBlock.ProtoReflect.Descriptor instead.
func (*TerrainBlock) Descriptor() ([]byte, []int) {
	return file_scene_proto_rawDescGZIP(), []int{0}
}

func (x *TerrainBlock) GetType() float32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *TerrainBlock) GetX() float32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *TerrainBlock) GetY() float32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *TerrainBlock) GetWidth() float32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *TerrainBlock) GetHeight() float32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *TerrainBlock) GetRotation() float32 {
	if x != nil {
		return x.Rotation
	}
	return 0
}

type TerrainBlocks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TerrainBlocks []*TerrainBlock `protobuf:"bytes,1,rep,name=terrainBlocks,proto3" json:"terrainBlocks,omitempty"`
}

func (x *TerrainBlocks) Reset() {
	*x = TerrainBlocks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scene_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TerrainBlocks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerrainBlocks) ProtoMessage() {}

func (x *TerrainBlocks) ProtoReflect() protoreflect.Message {
	mi := &file_scene_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerrainBlocks.ProtoReflect.Descriptor instead.
func (*TerrainBlocks) Descriptor() ([]byte, []int) {
	return file_scene_proto_rawDescGZIP(), []int{1}
}

func (x *TerrainBlocks) GetTerrainBlocks() []*TerrainBlock {
	if x != nil {
		return x.TerrainBlocks
	}
	return nil
}

var File_scene_proto protoreflect.FileDescriptor

var file_scene_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x73,
	0x63, 0x65, 0x6e, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x01, 0x0a, 0x0c, 0x54, 0x65,
	0x72, 0x72, 0x61, 0x69, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0c,
	0x0a, 0x01, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69,
	0x64, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68,
	0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x72, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4f, 0x0a, 0x0d, 0x54, 0x65, 0x72, 0x72, 0x61, 0x69, 0x6e, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x3e, 0x0a, 0x0d, 0x74, 0x65, 0x72, 0x72, 0x61, 0x69, 0x6e,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73,
	0x63, 0x65, 0x6e, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x65, 0x72, 0x72, 0x61, 0x69,
	0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x0d, 0x74, 0x65, 0x72, 0x72, 0x61, 0x69, 0x6e, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x42, 0x1e, 0x5a, 0x1c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x61, 0x69, 0x6e, 0x6e, 0x39, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_scene_proto_rawDescOnce sync.Once
	file_scene_proto_rawDescData = file_scene_proto_rawDesc
)

func file_scene_proto_rawDescGZIP() []byte {
	file_scene_proto_rawDescOnce.Do(func() {
		file_scene_proto_rawDescData = protoimpl.X.CompressGZIP(file_scene_proto_rawDescData)
	})
	return file_scene_proto_rawDescData
}

var file_scene_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_scene_proto_goTypes = []interface{}{
	(*TerrainBlock)(nil),  // 0: sceneProto.TerrainBlock
	(*TerrainBlocks)(nil), // 1: sceneProto.TerrainBlocks
}
var file_scene_proto_depIdxs = []int32{
	0, // 0: sceneProto.TerrainBlocks.terrainBlocks:type_name -> sceneProto.TerrainBlock
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_scene_proto_init() }
func file_scene_proto_init() {
	if File_scene_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_scene_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TerrainBlock); i {
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
		file_scene_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TerrainBlocks); i {
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
			RawDescriptor: file_scene_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_scene_proto_goTypes,
		DependencyIndexes: file_scene_proto_depIdxs,
		MessageInfos:      file_scene_proto_msgTypes,
	}.Build()
	File_scene_proto = out.File
	file_scene_proto_rawDesc = nil
	file_scene_proto_goTypes = nil
	file_scene_proto_depIdxs = nil
}
