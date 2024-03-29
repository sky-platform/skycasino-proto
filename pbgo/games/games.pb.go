// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: games/games.proto

package games

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

type GameType int32

const (
	// 未指定
	GameType_GAME_TYPE_UNSPECIFIED GameType = 0
	// 百家樂
	GameType_BACCARAT GameType = 16
	// 龍虎鬥
	GameType_THEBIGBATTLE GameType = 17
	// 炸金花
	GameType_THREECARDS GameType = 18
	// 牛牛
	GameType_BULLFIGHT GameType = 19
	// 番攤
	GameType_FANTAN GameType = 20
	// 骰寶
	GameType_SICBO GameType = 32
	// 輪盤
	GameType_ROULETTE GameType = 36
	// 幸運輪
	GameType_LUCKYWHEEL GameType = 40
)

// Enum value maps for GameType.
var (
	GameType_name = map[int32]string{
		0:  "GAME_TYPE_UNSPECIFIED",
		16: "BACCARAT",
		17: "THEBIGBATTLE",
		18: "THREECARDS",
		19: "BULLFIGHT",
		20: "FANTAN",
		32: "SICBO",
		36: "ROULETTE",
		40: "LUCKYWHEEL",
	}
	GameType_value = map[string]int32{
		"GAME_TYPE_UNSPECIFIED": 0,
		"BACCARAT":              16,
		"THEBIGBATTLE":          17,
		"THREECARDS":            18,
		"BULLFIGHT":             19,
		"FANTAN":                20,
		"SICBO":                 32,
		"ROULETTE":              36,
		"LUCKYWHEEL":            40,
	}
)

func (x GameType) Enum() *GameType {
	p := new(GameType)
	*p = x
	return p
}

func (x GameType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GameType) Descriptor() protoreflect.EnumDescriptor {
	return file_games_games_proto_enumTypes[0].Descriptor()
}

func (GameType) Type() protoreflect.EnumType {
	return &file_games_games_proto_enumTypes[0]
}

func (x GameType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GameType.Descriptor instead.
func (GameType) EnumDescriptor() ([]byte, []int) {
	return file_games_games_proto_rawDescGZIP(), []int{0}
}

var File_games_games_proto protoreflect.FileDescriptor

var file_games_games_proto_rawDesc = []byte{
	0x0a, 0x11, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x2a, 0x99, 0x01, 0x0a, 0x08, 0x47,
	0x61, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x47, 0x41, 0x4d, 0x45, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x42, 0x41, 0x43, 0x43, 0x41, 0x52, 0x41, 0x54, 0x10, 0x10,
	0x12, 0x10, 0x0a, 0x0c, 0x54, 0x48, 0x45, 0x42, 0x49, 0x47, 0x42, 0x41, 0x54, 0x54, 0x4c, 0x45,
	0x10, 0x11, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x48, 0x52, 0x45, 0x45, 0x43, 0x41, 0x52, 0x44, 0x53,
	0x10, 0x12, 0x12, 0x0d, 0x0a, 0x09, 0x42, 0x55, 0x4c, 0x4c, 0x46, 0x49, 0x47, 0x48, 0x54, 0x10,
	0x13, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x4e, 0x54, 0x41, 0x4e, 0x10, 0x14, 0x12, 0x09, 0x0a,
	0x05, 0x53, 0x49, 0x43, 0x42, 0x4f, 0x10, 0x20, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x4f, 0x55, 0x4c,
	0x45, 0x54, 0x54, 0x45, 0x10, 0x24, 0x12, 0x0e, 0x0a, 0x0a, 0x4c, 0x55, 0x43, 0x4b, 0x59, 0x57,
	0x48, 0x45, 0x45, 0x4c, 0x10, 0x28, 0x42, 0x7f, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x61,
	0x6d, 0x65, 0x73, 0x42, 0x0a, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6b,
	0x79, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x73, 0x6b, 0x79, 0x63, 0x61,
	0x73, 0x69, 0x6e, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x67, 0x6f, 0x2f,
	0x67, 0x61, 0x6d, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x58, 0x58, 0xaa, 0x02, 0x05, 0x47, 0x61,
	0x6d, 0x65, 0x73, 0xca, 0x02, 0x05, 0x47, 0x61, 0x6d, 0x65, 0x73, 0xe2, 0x02, 0x11, 0x47, 0x61,
	0x6d, 0x65, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x05, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_games_games_proto_rawDescOnce sync.Once
	file_games_games_proto_rawDescData = file_games_games_proto_rawDesc
)

func file_games_games_proto_rawDescGZIP() []byte {
	file_games_games_proto_rawDescOnce.Do(func() {
		file_games_games_proto_rawDescData = protoimpl.X.CompressGZIP(file_games_games_proto_rawDescData)
	})
	return file_games_games_proto_rawDescData
}

var file_games_games_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_games_games_proto_goTypes = []interface{}{
	(GameType)(0), // 0: games.GameType
}
var file_games_games_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_games_games_proto_init() }
func file_games_games_proto_init() {
	if File_games_games_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_games_games_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_games_games_proto_goTypes,
		DependencyIndexes: file_games_games_proto_depIdxs,
		EnumInfos:         file_games_games_proto_enumTypes,
	}.Build()
	File_games_games_proto = out.File
	file_games_games_proto_rawDesc = nil
	file_games_games_proto_goTypes = nil
	file_games_games_proto_depIdxs = nil
}
