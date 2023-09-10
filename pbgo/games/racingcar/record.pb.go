// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: games/racingcar/record.proto

package racingcar

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

// *
// The state of the game flow
type GameState int32

const (
	GameState_STATE_IDLE          GameState = 0 // Idle
	GameState_STATE_START         GameState = 1 // Game Starting
	GameState_STATE_COUNTER_START GameState = 2 // Counuter Started
	GameState_STATE_INPUT         GameState = 3 // Wait for input
	GameState_STATE_END           GameState = 7 // Game Ended
	GameState_STATE_CONFIRM       GameState = 8 // Game Ended
)

// Enum value maps for GameState.
var (
	GameState_name = map[int32]string{
		0: "STATE_IDLE",
		1: "STATE_START",
		2: "STATE_COUNTER_START",
		3: "STATE_INPUT",
		7: "STATE_END",
		8: "STATE_CONFIRM",
	}
	GameState_value = map[string]int32{
		"STATE_IDLE":          0,
		"STATE_START":         1,
		"STATE_COUNTER_START": 2,
		"STATE_INPUT":         3,
		"STATE_END":           7,
		"STATE_CONFIRM":       8,
	}
)

func (x GameState) Enum() *GameState {
	p := new(GameState)
	*p = x
	return p
}

func (x GameState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GameState) Descriptor() protoreflect.EnumDescriptor {
	return file_games_racingcar_record_proto_enumTypes[0].Descriptor()
}

func (GameState) Type() protoreflect.EnumType {
	return &file_games_racingcar_record_proto_enumTypes[0]
}

func (x GameState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GameState.Descriptor instead.
func (GameState) EnumDescriptor() ([]byte, []int) {
	return file_games_racingcar_record_proto_rawDescGZIP(), []int{0}
}

// *
// The status fo the current game.
type GameStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionID   string           `protobuf:"bytes,1,opt,name=sessionID,proto3" json:"sessionID,omitempty"`                                                                                    // Unique tableid
	GameRoundID string           `protobuf:"bytes,2,opt,name=gameRoundID,proto3" json:"gameRoundID,omitempty"`                                                                                // Unique session for the game
	TableID     string           `protobuf:"bytes,3,opt,name=tableID,proto3" json:"tableID,omitempty"`                                                                                        // Unique tableid
	State       GameState        `protobuf:"varint,4,opt,name=state,proto3,enum=games.racingcar.GameState" json:"state,omitempty"`                                                            // State of the game
	StartTime   int64            `protobuf:"varint,5,opt,name=startTime,proto3" json:"startTime,omitempty"`                                                                                   // Time when game start
	ShoeID      string           `protobuf:"bytes,6,opt,name=shoeID,proto3" json:"shoeID,omitempty"`                                                                                          // Shoe ID
	Ball1       int32            `protobuf:"varint,10,opt,name=ball1,proto3" json:"ball1,omitempty"`                                                                                          // Value of the ball 1
	Ball2       int32            `protobuf:"varint,11,opt,name=ball2,proto3" json:"ball2,omitempty"`                                                                                          // Value of the ball 2
	Ball3       int32            `protobuf:"varint,12,opt,name=ball3,proto3" json:"ball3,omitempty"`                                                                                          // Value of the ball 3
	Ball4       int32            `protobuf:"varint,13,opt,name=ball4,proto3" json:"ball4,omitempty"`                                                                                          // Value of the ball 4
	Ball5       int32            `protobuf:"varint,14,opt,name=ball5,proto3" json:"ball5,omitempty"`                                                                                          // Value of the ball 5
	Ball6       int32            `protobuf:"varint,15,opt,name=ball6,proto3" json:"ball6,omitempty"`                                                                                          // Value of the ball 6
	Ball7       int32            `protobuf:"varint,16,opt,name=ball7,proto3" json:"ball7,omitempty"`                                                                                          // Value of the ball 7
	Ball8       int32            `protobuf:"varint,17,opt,name=ball8,proto3" json:"ball8,omitempty"`                                                                                          // Value of the ball 8
	Ball9       int32            `protobuf:"varint,18,opt,name=ball9,proto3" json:"ball9,omitempty"`                                                                                          // Value of the ball 9
	Ball10      int32            `protobuf:"varint,19,opt,name=ball10,proto3" json:"ball10,omitempty"`                                                                                        // Value of the ball 10
	Countdown   int32            `protobuf:"varint,20,opt,name=countdown,proto3" json:"countdown,omitempty"`                                                                                  // Coundown value
	Lucky       map[string]int32 `protobuf:"bytes,150,rep,name=lucky,proto3" json:"lucky,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"` // Lucky Number and payout ratio
}

func (x *GameStatus) Reset() {
	*x = GameStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_games_racingcar_record_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameStatus) ProtoMessage() {}

func (x *GameStatus) ProtoReflect() protoreflect.Message {
	mi := &file_games_racingcar_record_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameStatus.ProtoReflect.Descriptor instead.
func (*GameStatus) Descriptor() ([]byte, []int) {
	return file_games_racingcar_record_proto_rawDescGZIP(), []int{0}
}

func (x *GameStatus) GetSessionID() string {
	if x != nil {
		return x.SessionID
	}
	return ""
}

func (x *GameStatus) GetGameRoundID() string {
	if x != nil {
		return x.GameRoundID
	}
	return ""
}

func (x *GameStatus) GetTableID() string {
	if x != nil {
		return x.TableID
	}
	return ""
}

func (x *GameStatus) GetState() GameState {
	if x != nil {
		return x.State
	}
	return GameState_STATE_IDLE
}

func (x *GameStatus) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *GameStatus) GetShoeID() string {
	if x != nil {
		return x.ShoeID
	}
	return ""
}

func (x *GameStatus) GetBall1() int32 {
	if x != nil {
		return x.Ball1
	}
	return 0
}

func (x *GameStatus) GetBall2() int32 {
	if x != nil {
		return x.Ball2
	}
	return 0
}

func (x *GameStatus) GetBall3() int32 {
	if x != nil {
		return x.Ball3
	}
	return 0
}

func (x *GameStatus) GetBall4() int32 {
	if x != nil {
		return x.Ball4
	}
	return 0
}

func (x *GameStatus) GetBall5() int32 {
	if x != nil {
		return x.Ball5
	}
	return 0
}

func (x *GameStatus) GetBall6() int32 {
	if x != nil {
		return x.Ball6
	}
	return 0
}

func (x *GameStatus) GetBall7() int32 {
	if x != nil {
		return x.Ball7
	}
	return 0
}

func (x *GameStatus) GetBall8() int32 {
	if x != nil {
		return x.Ball8
	}
	return 0
}

func (x *GameStatus) GetBall9() int32 {
	if x != nil {
		return x.Ball9
	}
	return 0
}

func (x *GameStatus) GetBall10() int32 {
	if x != nil {
		return x.Ball10
	}
	return 0
}

func (x *GameStatus) GetCountdown() int32 {
	if x != nil {
		return x.Countdown
	}
	return 0
}

func (x *GameStatus) GetLucky() map[string]int32 {
	if x != nil {
		return x.Lucky
	}
	return nil
}

var File_games_racingcar_record_proto protoreflect.FileDescriptor

var file_games_racingcar_record_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x2f, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x63, 0x61,
	0x72, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x67, 0x61, 0x6d, 0x65, 0x73, 0x2e, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x63, 0x61, 0x72, 0x22,
	0xc3, 0x04, 0x0a, 0x0a, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c,
	0x0a, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b,
	0x67, 0x61, 0x6d, 0x65, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x67, 0x61, 0x6d, 0x65, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x44, 0x12, 0x18,
	0x0a, 0x07, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x30, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x2e,
	0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x63, 0x61, 0x72, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x68, 0x6f, 0x65,
	0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x65, 0x49, 0x44,
	0x12, 0x14, 0x0a, 0x05, 0x62, 0x61, 0x6c, 0x6c, 0x31, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x62, 0x61, 0x6c, 0x6c, 0x31, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x61, 0x6c, 0x6c, 0x32, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x62, 0x61, 0x6c, 0x6c, 0x32, 0x12, 0x14, 0x0a, 0x05,
	0x62, 0x61, 0x6c, 0x6c, 0x33, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x62, 0x61, 0x6c,
	0x6c, 0x33, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x61, 0x6c, 0x6c, 0x34, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x62, 0x61, 0x6c, 0x6c, 0x34, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x61, 0x6c, 0x6c,
	0x35, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x62, 0x61, 0x6c, 0x6c, 0x35, 0x12, 0x14,
	0x0a, 0x05, 0x62, 0x61, 0x6c, 0x6c, 0x36, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x62,
	0x61, 0x6c, 0x6c, 0x36, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x61, 0x6c, 0x6c, 0x37, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x62, 0x61, 0x6c, 0x6c, 0x37, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x61,
	0x6c, 0x6c, 0x38, 0x18, 0x11, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x62, 0x61, 0x6c, 0x6c, 0x38,
	0x12, 0x14, 0x0a, 0x05, 0x62, 0x61, 0x6c, 0x6c, 0x39, 0x18, 0x12, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x62, 0x61, 0x6c, 0x6c, 0x39, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x61, 0x6c, 0x6c, 0x31, 0x30,
	0x18, 0x13, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x62, 0x61, 0x6c, 0x6c, 0x31, 0x30, 0x12, 0x1c,
	0x0a, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x12, 0x3d, 0x0a, 0x05,
	0x6c, 0x75, 0x63, 0x6b, 0x79, 0x18, 0x96, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x67,
	0x61, 0x6d, 0x65, 0x73, 0x2e, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x63, 0x61, 0x72, 0x2e, 0x47,
	0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x4c, 0x75, 0x63, 0x6b, 0x79, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x6c, 0x75, 0x63, 0x6b, 0x79, 0x1a, 0x38, 0x0a, 0x0a, 0x4c,
	0x75, 0x63, 0x6b, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x78, 0x0a, 0x09, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x49, 0x44, 0x4c, 0x45,
	0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x52,
	0x54, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x4f, 0x55,
	0x4e, 0x54, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b,
	0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x49, 0x4e, 0x50, 0x55, 0x54, 0x10, 0x03, 0x12, 0x0d, 0x0a,
	0x09, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x45, 0x4e, 0x44, 0x10, 0x07, 0x12, 0x11, 0x0a, 0x0d,
	0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x52, 0x4d, 0x10, 0x08, 0x42,
	0xbd, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x2e, 0x72, 0x61,
	0x63, 0x69, 0x6e, 0x67, 0x63, 0x61, 0x72, 0x42, 0x0b, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x6b, 0x79, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f,
	0x73, 0x6b, 0x79, 0x63, 0x61, 0x73, 0x69, 0x6e, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x70, 0x62, 0x67, 0x6f, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x2f, 0x72, 0x61, 0x63, 0x69, 0x6e,
	0x67, 0x63, 0x61, 0x72, 0xa2, 0x02, 0x03, 0x47, 0x52, 0x58, 0xaa, 0x02, 0x0f, 0x47, 0x61, 0x6d,
	0x65, 0x73, 0x2e, 0x52, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x63, 0x61, 0x72, 0xca, 0x02, 0x0f, 0x47,
	0x61, 0x6d, 0x65, 0x73, 0x5c, 0x52, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x63, 0x61, 0x72, 0xe2, 0x02,
	0x1b, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x5c, 0x52, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x63, 0x61, 0x72,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x47,
	0x61, 0x6d, 0x65, 0x73, 0x3a, 0x3a, 0x52, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x63, 0x61, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_games_racingcar_record_proto_rawDescOnce sync.Once
	file_games_racingcar_record_proto_rawDescData = file_games_racingcar_record_proto_rawDesc
)

func file_games_racingcar_record_proto_rawDescGZIP() []byte {
	file_games_racingcar_record_proto_rawDescOnce.Do(func() {
		file_games_racingcar_record_proto_rawDescData = protoimpl.X.CompressGZIP(file_games_racingcar_record_proto_rawDescData)
	})
	return file_games_racingcar_record_proto_rawDescData
}

var file_games_racingcar_record_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_games_racingcar_record_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_games_racingcar_record_proto_goTypes = []interface{}{
	(GameState)(0),     // 0: games.racingcar.GameState
	(*GameStatus)(nil), // 1: games.racingcar.GameStatus
	nil,                // 2: games.racingcar.GameStatus.LuckyEntry
}
var file_games_racingcar_record_proto_depIdxs = []int32{
	0, // 0: games.racingcar.GameStatus.state:type_name -> games.racingcar.GameState
	2, // 1: games.racingcar.GameStatus.lucky:type_name -> games.racingcar.GameStatus.LuckyEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_games_racingcar_record_proto_init() }
func file_games_racingcar_record_proto_init() {
	if File_games_racingcar_record_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_games_racingcar_record_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameStatus); i {
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
			RawDescriptor: file_games_racingcar_record_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_games_racingcar_record_proto_goTypes,
		DependencyIndexes: file_games_racingcar_record_proto_depIdxs,
		EnumInfos:         file_games_racingcar_record_proto_enumTypes,
		MessageInfos:      file_games_racingcar_record_proto_msgTypes,
	}.Build()
	File_games_racingcar_record_proto = out.File
	file_games_racingcar_record_proto_rawDesc = nil
	file_games_racingcar_record_proto_goTypes = nil
	file_games_racingcar_record_proto_depIdxs = nil
}
