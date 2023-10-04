// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: games/racingcar/record.proto

package racingcar

import (
	_ "github.com/sky-platform/skycasino-proto/pbgo/recorder"
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

type ResourceType int32

const (
	// 未指定
	ResourceType_RESOURCE_TYPE_UNSPECIFIED ResourceType = 0
	// 牌靴
	ResourceType_SHOE ResourceType = 1
	// 例牌
	ResourceType_NORMAL ResourceType = 2
	// 非例牌
	ResourceType_EXTRA ResourceType = 3
)

// Enum value maps for ResourceType.
var (
	ResourceType_name = map[int32]string{
		0: "RESOURCE_TYPE_UNSPECIFIED",
		1: "SHOE",
		2: "NORMAL",
		3: "EXTRA",
	}
	ResourceType_value = map[string]int32{
		"RESOURCE_TYPE_UNSPECIFIED": 0,
		"SHOE":                      1,
		"NORMAL":                    2,
		"EXTRA":                     3,
	}
)

func (x ResourceType) Enum() *ResourceType {
	p := new(ResourceType)
	*p = x
	return p
}

func (x ResourceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResourceType) Descriptor() protoreflect.EnumDescriptor {
	return file_games_racingcar_record_proto_enumTypes[0].Descriptor()
}

func (ResourceType) Type() protoreflect.EnumType {
	return &file_games_racingcar_record_proto_enumTypes[0]
}

func (x ResourceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResourceType.Descriptor instead.
func (ResourceType) EnumDescriptor() ([]byte, []int) {
	return file_games_racingcar_record_proto_rawDescGZIP(), []int{0}
}

// 步驟，依據遊戲供應商邏輯，不是每個步驟都會走到
type Step int32

const (
	// 未指定
	Step_STEP_UNSPECIFIED Step = 0
	// 本局開始
	Step_ROUND_START Step = 1
	// 本局完成
	Step_ROUND_FINISHED Step = 2
	// 本局取消
	Step_ROUND_CANCELED Step = 3
	// 派發閒家例牌1
	Step_DEAL_PLAYER_1 Step = 4
	// 派發莊家例牌1
	Step_DEAL_BANKER_1 Step = 5
	// 派發閒家例牌2
	Step_DEAL_PLAYER_2 Step = 6
	// 派發莊家例牌2
	Step_DEAL_BANKER_2 Step = 7
	// 開放下注
	Step_OPEN_FOR_BETTING Step = 8
	// 倒數
	Step_COUNTDOWN Step = 27
	// 停止下注
	Step_CLOSE_THE_BETTING Step = 9
	// 非瞇牌開例牌
	Step_SHOW_NORMAL Step = 12
	// 瞇牌例牌
	Step_PEEK_NORMAL Step = 13
	// 瞇牌公布例牌
	Step_PEEK_SHOW_NORMAL Step = 14
	// 非瞇牌補閒家牌
	Step_DEAL_PLAYER_EXTRA Step = 16
	// 瞇牌閒家補牌
	Step_PEEK_DEAL_PLAYER_EXTRA Step = 17
	// 瞇牌公布閒家補牌
	Step_PEEK_SHOW_PLAYER_EXTRA Step = 18
	// 非瞇牌補莊家牌
	Step_DEAL_BANKER_EXTRA Step = 20
	// 瞇牌莊家補牌
	Step_PEEK_DEAL_BANKER_EXTRA Step = 21
	// 瞇牌公布莊家補牌
	Step_PEEK_SHOW_BANKER_EXTRA Step = 22
	// PITBOSS 修改派發例牌
	Step_PITBOSS_MODIFY_NORMAL Step = 24
	// PITBOSS 修改派發閒家補牌
	Step_PITBOSS_MODIFY_PLAYER_EXTRA Step = 25
	// PITBOSS 修改派發莊家補牌
	Step_PITBOSS_MODIFY_BANKER_EXTRA Step = 26
)

// Enum value maps for Step.
var (
	Step_name = map[int32]string{
		0:  "STEP_UNSPECIFIED",
		1:  "ROUND_START",
		2:  "ROUND_FINISHED",
		3:  "ROUND_CANCELED",
		4:  "DEAL_PLAYER_1",
		5:  "DEAL_BANKER_1",
		6:  "DEAL_PLAYER_2",
		7:  "DEAL_BANKER_2",
		8:  "OPEN_FOR_BETTING",
		27: "COUNTDOWN",
		9:  "CLOSE_THE_BETTING",
		12: "SHOW_NORMAL",
		13: "PEEK_NORMAL",
		14: "PEEK_SHOW_NORMAL",
		16: "DEAL_PLAYER_EXTRA",
		17: "PEEK_DEAL_PLAYER_EXTRA",
		18: "PEEK_SHOW_PLAYER_EXTRA",
		20: "DEAL_BANKER_EXTRA",
		21: "PEEK_DEAL_BANKER_EXTRA",
		22: "PEEK_SHOW_BANKER_EXTRA",
		24: "PITBOSS_MODIFY_NORMAL",
		25: "PITBOSS_MODIFY_PLAYER_EXTRA",
		26: "PITBOSS_MODIFY_BANKER_EXTRA",
	}
	Step_value = map[string]int32{
		"STEP_UNSPECIFIED":            0,
		"ROUND_START":                 1,
		"ROUND_FINISHED":              2,
		"ROUND_CANCELED":              3,
		"DEAL_PLAYER_1":               4,
		"DEAL_BANKER_1":               5,
		"DEAL_PLAYER_2":               6,
		"DEAL_BANKER_2":               7,
		"OPEN_FOR_BETTING":            8,
		"COUNTDOWN":                   27,
		"CLOSE_THE_BETTING":           9,
		"SHOW_NORMAL":                 12,
		"PEEK_NORMAL":                 13,
		"PEEK_SHOW_NORMAL":            14,
		"DEAL_PLAYER_EXTRA":           16,
		"PEEK_DEAL_PLAYER_EXTRA":      17,
		"PEEK_SHOW_PLAYER_EXTRA":      18,
		"DEAL_BANKER_EXTRA":           20,
		"PEEK_DEAL_BANKER_EXTRA":      21,
		"PEEK_SHOW_BANKER_EXTRA":      22,
		"PITBOSS_MODIFY_NORMAL":       24,
		"PITBOSS_MODIFY_PLAYER_EXTRA": 25,
		"PITBOSS_MODIFY_BANKER_EXTRA": 26,
	}
)

func (x Step) Enum() *Step {
	p := new(Step)
	*p = x
	return p
}

func (x Step) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Step) Descriptor() protoreflect.EnumDescriptor {
	return file_games_racingcar_record_proto_enumTypes[1].Descriptor()
}

func (Step) Type() protoreflect.EnumType {
	return &file_games_racingcar_record_proto_enumTypes[1]
}

func (x Step) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Step.Descriptor instead.
func (Step) EnumDescriptor() ([]byte, []int) {
	return file_games_racingcar_record_proto_rawDescGZIP(), []int{1}
}

// 座位功能代碼
type Seat int32

const (
	// 未指定
	Seat_SEAT_UNSPECIFIED Seat = 0
	// Dealer
	Seat_DEALER Seat = 1
	// 閒家
	Seat_PLAYER Seat = 2
	// 莊家
	Seat_BANKER Seat = 3
	// Pitboss
	Seat_PITBOSS Seat = 9
)

// Enum value maps for Seat.
var (
	Seat_name = map[int32]string{
		0: "SEAT_UNSPECIFIED",
		1: "DEALER",
		2: "PLAYER",
		3: "BANKER",
		9: "PITBOSS",
	}
	Seat_value = map[string]int32{
		"SEAT_UNSPECIFIED": 0,
		"DEALER":           1,
		"PLAYER":           2,
		"BANKER":           3,
		"PITBOSS":          9,
	}
)

func (x Seat) Enum() *Seat {
	p := new(Seat)
	*p = x
	return p
}

func (x Seat) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Seat) Descriptor() protoreflect.EnumDescriptor {
	return file_games_racingcar_record_proto_enumTypes[2].Descriptor()
}

func (Seat) Type() protoreflect.EnumType {
	return &file_games_racingcar_record_proto_enumTypes[2]
}

func (x Seat) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Seat.Descriptor instead.
func (Seat) EnumDescriptor() ([]byte, []int) {
	return file_games_racingcar_record_proto_rawDescGZIP(), []int{2}
}

// 結果紀錄類型代碼
type ResultRecordType int32

const (
	// 未指定
	ResultRecordType_RESULT_RECORD_TYPE_UNSPECIFIED ResultRecordType = 0
	// 遊戲局結束
	ResultRecordType_ROUND_FINISH ResultRecordType = 1
	// Pitboss 修正
	ResultRecordType_PITBOSS_MODIFY_RESULT ResultRecordType = 2
)

// Enum value maps for ResultRecordType.
var (
	ResultRecordType_name = map[int32]string{
		0: "RESULT_RECORD_TYPE_UNSPECIFIED",
		1: "ROUND_FINISH",
		2: "PITBOSS_MODIFY_RESULT",
	}
	ResultRecordType_value = map[string]int32{
		"RESULT_RECORD_TYPE_UNSPECIFIED": 0,
		"ROUND_FINISH":                   1,
		"PITBOSS_MODIFY_RESULT":          2,
	}
)

func (x ResultRecordType) Enum() *ResultRecordType {
	p := new(ResultRecordType)
	*p = x
	return p
}

func (x ResultRecordType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResultRecordType) Descriptor() protoreflect.EnumDescriptor {
	return file_games_racingcar_record_proto_enumTypes[3].Descriptor()
}

func (ResultRecordType) Type() protoreflect.EnumType {
	return &file_games_racingcar_record_proto_enumTypes[3]
}

func (x ResultRecordType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResultRecordType.Descriptor instead.
func (ResultRecordType) EnumDescriptor() ([]byte, []int) {
	return file_games_racingcar_record_proto_rawDescGZIP(), []int{3}
}

// 取消原因
type CancelReason int32

const (
	// 未指定，沒有取消
	CancelReason_CANCEL_REASON_UNSPECIFIED CancelReason = 0
	// 取消說明原因
	CancelReason_NO_REASON CancelReason = 1
)

// Enum value maps for CancelReason.
var (
	CancelReason_name = map[int32]string{
		0: "CANCEL_REASON_UNSPECIFIED",
		1: "NO_REASON",
	}
	CancelReason_value = map[string]int32{
		"CANCEL_REASON_UNSPECIFIED": 0,
		"NO_REASON":                 1,
	}
)

func (x CancelReason) Enum() *CancelReason {
	p := new(CancelReason)
	*p = x
	return p
}

func (x CancelReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CancelReason) Descriptor() protoreflect.EnumDescriptor {
	return file_games_racingcar_record_proto_enumTypes[4].Descriptor()
}

func (CancelReason) Type() protoreflect.EnumType {
	return &file_games_racingcar_record_proto_enumTypes[4]
}

func (x CancelReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CancelReason.Descriptor instead.
func (CancelReason) EnumDescriptor() ([]byte, []int) {
	return file_games_racingcar_record_proto_rawDescGZIP(), []int{4}
}

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
	return file_games_racingcar_record_proto_enumTypes[5].Descriptor()
}

func (GameState) Type() protoreflect.EnumType {
	return &file_games_racingcar_record_proto_enumTypes[5]
}

func (x GameState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GameState.Descriptor instead.
func (GameState) EnumDescriptor() ([]byte, []int) {
	return file_games_racingcar_record_proto_rawDescGZIP(), []int{5}
}

// *
// The status fo the current game.
type GameStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionID string           `protobuf:"bytes,1,opt,name=sessionID,proto3" json:"sessionID,omitempty"`                                                                                    // Unique tableid
	RoundCode string           `protobuf:"bytes,2,opt,name=roundCode,proto3" json:"roundCode,omitempty"`                                                                                    // Unique session for the game
	TableID   string           `protobuf:"bytes,3,opt,name=tableID,proto3" json:"tableID,omitempty"`                                                                                        // Unique tableid
	State     GameState        `protobuf:"varint,4,opt,name=state,proto3,enum=games.racingcar.GameState" json:"state,omitempty"`                                                            // State of the game
	StartTime int64            `protobuf:"varint,5,opt,name=startTime,proto3" json:"startTime,omitempty"`                                                                                   // Time when game start
	ShoeCode  string           `protobuf:"bytes,6,opt,name=shoeCode,proto3" json:"shoeCode,omitempty"`                                                                                      // Shoe ID
	Ball1     int32            `protobuf:"varint,10,opt,name=ball1,proto3" json:"ball1,omitempty"`                                                                                          // Value of the ball 1
	Ball2     int32            `protobuf:"varint,11,opt,name=ball2,proto3" json:"ball2,omitempty"`                                                                                          // Value of the ball 2
	Ball3     int32            `protobuf:"varint,12,opt,name=ball3,proto3" json:"ball3,omitempty"`                                                                                          // Value of the ball 3
	Ball4     int32            `protobuf:"varint,13,opt,name=ball4,proto3" json:"ball4,omitempty"`                                                                                          // Value of the ball 4
	Ball5     int32            `protobuf:"varint,14,opt,name=ball5,proto3" json:"ball5,omitempty"`                                                                                          // Value of the ball 5
	Ball6     int32            `protobuf:"varint,15,opt,name=ball6,proto3" json:"ball6,omitempty"`                                                                                          // Value of the ball 6
	Ball7     int32            `protobuf:"varint,16,opt,name=ball7,proto3" json:"ball7,omitempty"`                                                                                          // Value of the ball 7
	Ball8     int32            `protobuf:"varint,17,opt,name=ball8,proto3" json:"ball8,omitempty"`                                                                                          // Value of the ball 8
	Ball9     int32            `protobuf:"varint,18,opt,name=ball9,proto3" json:"ball9,omitempty"`                                                                                          // Value of the ball 9
	Ball10    int32            `protobuf:"varint,19,opt,name=ball10,proto3" json:"ball10,omitempty"`                                                                                        // Value of the ball 10
	Countdown int32            `protobuf:"varint,20,opt,name=countdown,proto3" json:"countdown,omitempty"`                                                                                  // Coundown value
	Lucky     map[string]int32 `protobuf:"bytes,150,rep,name=lucky,proto3" json:"lucky,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"` // Lucky Number and payout ratio
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

func (x *GameStatus) GetRoundCode() string {
	if x != nil {
		return x.RoundCode
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

func (x *GameStatus) GetShoeCode() string {
	if x != nil {
		return x.ShoeCode
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

type RcBead struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoundCode  string           `protobuf:"bytes,1,opt,name=roundCode,proto3" json:"roundCode,omitempty"`
	ShoeCode   string           `protobuf:"bytes,2,opt,name=shoeCode,proto3" json:"shoeCode,omitempty"`
	Shoe       uint32           `protobuf:"varint,3,opt,name=shoe,proto3" json:"shoe,omitempty"`
	ShoeRound  uint32           `protobuf:"varint,4,opt,name=shoeRound,proto3" json:"shoeRound,omitempty"`
	Result     string           `protobuf:"bytes,5,opt,name=result,proto3" json:"result,omitempty"`
	Balls      []int32          `protobuf:"varint,7,rep,packed,name=balls,proto3" json:"balls,omitempty"`
	Statistics map[string]int32 `protobuf:"bytes,32,rep,name=statistics,proto3" json:"statistics,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *RcBead) Reset() {
	*x = RcBead{}
	if protoimpl.UnsafeEnabled {
		mi := &file_games_racingcar_record_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RcBead) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RcBead) ProtoMessage() {}

func (x *RcBead) ProtoReflect() protoreflect.Message {
	mi := &file_games_racingcar_record_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RcBead.ProtoReflect.Descriptor instead.
func (*RcBead) Descriptor() ([]byte, []int) {
	return file_games_racingcar_record_proto_rawDescGZIP(), []int{1}
}

func (x *RcBead) GetRoundCode() string {
	if x != nil {
		return x.RoundCode
	}
	return ""
}

func (x *RcBead) GetShoeCode() string {
	if x != nil {
		return x.ShoeCode
	}
	return ""
}

func (x *RcBead) GetShoe() uint32 {
	if x != nil {
		return x.Shoe
	}
	return 0
}

func (x *RcBead) GetShoeRound() uint32 {
	if x != nil {
		return x.ShoeRound
	}
	return 0
}

func (x *RcBead) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

func (x *RcBead) GetBalls() []int32 {
	if x != nil {
		return x.Balls
	}
	return nil
}

func (x *RcBead) GetStatistics() map[string]int32 {
	if x != nil {
		return x.Statistics
	}
	return nil
}

var File_games_racingcar_record_proto protoreflect.FileDescriptor

var file_games_racingcar_record_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x2f, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x63, 0x61,
	0x72, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x67, 0x61, 0x6d, 0x65, 0x73, 0x2e, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x63, 0x61, 0x72, 0x1a,
	0x15, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc3, 0x04, 0x0a, 0x0a, 0x47, 0x61, 0x6d, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x30, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x67, 0x61, 0x6d,
	0x65, 0x73, 0x2e, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x63, 0x61, 0x72, 0x2e, 0x47, 0x61, 0x6d,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73,
	0x68, 0x6f, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73,
	0x68, 0x6f, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x61, 0x6c, 0x6c, 0x31,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x62, 0x61, 0x6c, 0x6c, 0x31, 0x12, 0x14, 0x0a,
	0x05, 0x62, 0x61, 0x6c, 0x6c, 0x32, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x62, 0x61,
	0x6c, 0x6c, 0x32, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x61, 0x6c, 0x6c, 0x33, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x62, 0x61, 0x6c, 0x6c, 0x33, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x61, 0x6c,
	0x6c, 0x34, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x62, 0x61, 0x6c, 0x6c, 0x34, 0x12,
	0x14, 0x0a, 0x05, 0x62, 0x61, 0x6c, 0x6c, 0x35, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x62, 0x61, 0x6c, 0x6c, 0x35, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x61, 0x6c, 0x6c, 0x36, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x62, 0x61, 0x6c, 0x6c, 0x36, 0x12, 0x14, 0x0a, 0x05, 0x62,
	0x61, 0x6c, 0x6c, 0x37, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x62, 0x61, 0x6c, 0x6c,
	0x37, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x61, 0x6c, 0x6c, 0x38, 0x18, 0x11, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x62, 0x61, 0x6c, 0x6c, 0x38, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x61, 0x6c, 0x6c, 0x39,
	0x18, 0x12, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x62, 0x61, 0x6c, 0x6c, 0x39, 0x12, 0x16, 0x0a,
	0x06, 0x62, 0x61, 0x6c, 0x6c, 0x31, 0x30, 0x18, 0x13, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x62,
	0x61, 0x6c, 0x6c, 0x31, 0x30, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x64, 0x6f,
	0x77, 0x6e, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x64,
	0x6f, 0x77, 0x6e, 0x12, 0x3d, 0x0a, 0x05, 0x6c, 0x75, 0x63, 0x6b, 0x79, 0x18, 0x96, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x2e, 0x72, 0x61, 0x63, 0x69,
	0x6e, 0x67, 0x63, 0x61, 0x72, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x4c, 0x75, 0x63, 0x6b, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x6c, 0x75, 0x63,
	0x6b, 0x79, 0x1a, 0x38, 0x0a, 0x0a, 0x4c, 0x75, 0x63, 0x6b, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xaa, 0x02, 0x0a,
	0x06, 0x52, 0x63, 0x42, 0x65, 0x61, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x6f, 0x75, 0x6e, 0x64,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x6f, 0x75, 0x6e,
	0x64, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x68, 0x6f, 0x65, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x65, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x68, 0x6f, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x04, 0x73, 0x68, 0x6f, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x65, 0x52, 0x6f, 0x75,
	0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x73, 0x68, 0x6f, 0x65, 0x52, 0x6f,
	0x75, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x62,
	0x61, 0x6c, 0x6c, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x05, 0x52, 0x05, 0x62, 0x61, 0x6c, 0x6c,
	0x73, 0x12, 0x47, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x18,
	0x20, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x2e, 0x72, 0x61,
	0x63, 0x69, 0x6e, 0x67, 0x63, 0x61, 0x72, 0x2e, 0x52, 0x63, 0x42, 0x65, 0x61, 0x64, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x1a, 0x3d, 0x0a, 0x0f, 0x53, 0x74,
	0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x4e, 0x0a, 0x0c, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x19, 0x52, 0x45, 0x53,
	0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x48, 0x4f, 0x45,
	0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x02, 0x12, 0x09,
	0x0a, 0x05, 0x45, 0x58, 0x54, 0x52, 0x41, 0x10, 0x03, 0x2a, 0x90, 0x04, 0x0a, 0x04, 0x53, 0x74,
	0x65, 0x70, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x54, 0x45, 0x50, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x4f, 0x55, 0x4e,
	0x44, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x4f, 0x55,
	0x4e, 0x44, 0x5f, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x45, 0x44, 0x10, 0x02, 0x12, 0x12, 0x0a,
	0x0e, 0x52, 0x4f, 0x55, 0x4e, 0x44, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44, 0x10,
	0x03, 0x12, 0x11, 0x0a, 0x0d, 0x44, 0x45, 0x41, 0x4c, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52,
	0x5f, 0x31, 0x10, 0x04, 0x12, 0x11, 0x0a, 0x0d, 0x44, 0x45, 0x41, 0x4c, 0x5f, 0x42, 0x41, 0x4e,
	0x4b, 0x45, 0x52, 0x5f, 0x31, 0x10, 0x05, 0x12, 0x11, 0x0a, 0x0d, 0x44, 0x45, 0x41, 0x4c, 0x5f,
	0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x32, 0x10, 0x06, 0x12, 0x11, 0x0a, 0x0d, 0x44, 0x45,
	0x41, 0x4c, 0x5f, 0x42, 0x41, 0x4e, 0x4b, 0x45, 0x52, 0x5f, 0x32, 0x10, 0x07, 0x12, 0x14, 0x0a,
	0x10, 0x4f, 0x50, 0x45, 0x4e, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x42, 0x45, 0x54, 0x54, 0x49, 0x4e,
	0x47, 0x10, 0x08, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x44, 0x4f, 0x57, 0x4e,
	0x10, 0x1b, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x5f, 0x54, 0x48, 0x45, 0x5f,
	0x42, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x09, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x48, 0x4f,
	0x57, 0x5f, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x0c, 0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x45,
	0x45, 0x4b, 0x5f, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x0d, 0x12, 0x14, 0x0a, 0x10, 0x50,
	0x45, 0x45, 0x4b, 0x5f, 0x53, 0x48, 0x4f, 0x57, 0x5f, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10,
	0x0e, 0x12, 0x15, 0x0a, 0x11, 0x44, 0x45, 0x41, 0x4c, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52,
	0x5f, 0x45, 0x58, 0x54, 0x52, 0x41, 0x10, 0x10, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x45, 0x45, 0x4b,
	0x5f, 0x44, 0x45, 0x41, 0x4c, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x45, 0x58, 0x54,
	0x52, 0x41, 0x10, 0x11, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x45, 0x45, 0x4b, 0x5f, 0x53, 0x48, 0x4f,
	0x57, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x45, 0x58, 0x54, 0x52, 0x41, 0x10, 0x12,
	0x12, 0x15, 0x0a, 0x11, 0x44, 0x45, 0x41, 0x4c, 0x5f, 0x42, 0x41, 0x4e, 0x4b, 0x45, 0x52, 0x5f,
	0x45, 0x58, 0x54, 0x52, 0x41, 0x10, 0x14, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x45, 0x45, 0x4b, 0x5f,
	0x44, 0x45, 0x41, 0x4c, 0x5f, 0x42, 0x41, 0x4e, 0x4b, 0x45, 0x52, 0x5f, 0x45, 0x58, 0x54, 0x52,
	0x41, 0x10, 0x15, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x45, 0x45, 0x4b, 0x5f, 0x53, 0x48, 0x4f, 0x57,
	0x5f, 0x42, 0x41, 0x4e, 0x4b, 0x45, 0x52, 0x5f, 0x45, 0x58, 0x54, 0x52, 0x41, 0x10, 0x16, 0x12,
	0x19, 0x0a, 0x15, 0x50, 0x49, 0x54, 0x42, 0x4f, 0x53, 0x53, 0x5f, 0x4d, 0x4f, 0x44, 0x49, 0x46,
	0x59, 0x5f, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x18, 0x12, 0x1f, 0x0a, 0x1b, 0x50, 0x49,
	0x54, 0x42, 0x4f, 0x53, 0x53, 0x5f, 0x4d, 0x4f, 0x44, 0x49, 0x46, 0x59, 0x5f, 0x50, 0x4c, 0x41,
	0x59, 0x45, 0x52, 0x5f, 0x45, 0x58, 0x54, 0x52, 0x41, 0x10, 0x19, 0x12, 0x1f, 0x0a, 0x1b, 0x50,
	0x49, 0x54, 0x42, 0x4f, 0x53, 0x53, 0x5f, 0x4d, 0x4f, 0x44, 0x49, 0x46, 0x59, 0x5f, 0x42, 0x41,
	0x4e, 0x4b, 0x45, 0x52, 0x5f, 0x45, 0x58, 0x54, 0x52, 0x41, 0x10, 0x1a, 0x2a, 0x4d, 0x0a, 0x04,
	0x53, 0x65, 0x61, 0x74, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x45, 0x41, 0x54, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45,
	0x41, 0x4c, 0x45, 0x52, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52,
	0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x42, 0x41, 0x4e, 0x4b, 0x45, 0x52, 0x10, 0x03, 0x12, 0x0b,
	0x0a, 0x07, 0x50, 0x49, 0x54, 0x42, 0x4f, 0x53, 0x53, 0x10, 0x09, 0x2a, 0x63, 0x0a, 0x10, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x22, 0x0a, 0x1e, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x4f, 0x55, 0x4e, 0x44, 0x5f, 0x46, 0x49, 0x4e,
	0x49, 0x53, 0x48, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x50, 0x49, 0x54, 0x42, 0x4f, 0x53, 0x53,
	0x5f, 0x4d, 0x4f, 0x44, 0x49, 0x46, 0x59, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x10, 0x02,
	0x2a, 0x3c, 0x0a, 0x0c, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x12, 0x1d, 0x0a, 0x19, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f,
	0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x0d, 0x0a, 0x09, 0x4e, 0x4f, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x10, 0x01, 0x2a, 0x78,
	0x0a, 0x09, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x53,
	0x54, 0x41, 0x54, 0x45, 0x5f, 0x49, 0x44, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x53,
	0x54, 0x41, 0x54, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13,
	0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x45, 0x52, 0x5f, 0x53, 0x54,
	0x41, 0x52, 0x54, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x49,
	0x4e, 0x50, 0x55, 0x54, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f,
	0x45, 0x4e, 0x44, 0x10, 0x07, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x43,
	0x4f, 0x4e, 0x46, 0x49, 0x52, 0x4d, 0x10, 0x08, 0x42, 0xbd, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x2e, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x63, 0x61, 0x72,
	0x42, 0x0b, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6b, 0x79, 0x2d,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x73, 0x6b, 0x79, 0x63, 0x61, 0x73, 0x69,
	0x6e, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x67, 0x6f, 0x2f, 0x67, 0x61,
	0x6d, 0x65, 0x73, 0x2f, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x63, 0x61, 0x72, 0xa2, 0x02, 0x03,
	0x47, 0x52, 0x58, 0xaa, 0x02, 0x0f, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x2e, 0x52, 0x61, 0x63, 0x69,
	0x6e, 0x67, 0x63, 0x61, 0x72, 0xca, 0x02, 0x0f, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x5c, 0x52, 0x61,
	0x63, 0x69, 0x6e, 0x67, 0x63, 0x61, 0x72, 0xe2, 0x02, 0x1b, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x5c,
	0x52, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x63, 0x61, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x3a, 0x3a, 0x52,
	0x61, 0x63, 0x69, 0x6e, 0x67, 0x63, 0x61, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_games_racingcar_record_proto_enumTypes = make([]protoimpl.EnumInfo, 6)
var file_games_racingcar_record_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_games_racingcar_record_proto_goTypes = []interface{}{
	(ResourceType)(0),     // 0: games.racingcar.ResourceType
	(Step)(0),             // 1: games.racingcar.Step
	(Seat)(0),             // 2: games.racingcar.Seat
	(ResultRecordType)(0), // 3: games.racingcar.ResultRecordType
	(CancelReason)(0),     // 4: games.racingcar.CancelReason
	(GameState)(0),        // 5: games.racingcar.GameState
	(*GameStatus)(nil),    // 6: games.racingcar.GameStatus
	(*RcBead)(nil),        // 7: games.racingcar.RcBead
	nil,                   // 8: games.racingcar.GameStatus.LuckyEntry
	nil,                   // 9: games.racingcar.RcBead.StatisticsEntry
}
var file_games_racingcar_record_proto_depIdxs = []int32{
	5, // 0: games.racingcar.GameStatus.state:type_name -> games.racingcar.GameState
	8, // 1: games.racingcar.GameStatus.lucky:type_name -> games.racingcar.GameStatus.LuckyEntry
	9, // 2: games.racingcar.RcBead.statistics:type_name -> games.racingcar.RcBead.StatisticsEntry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
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
		file_games_racingcar_record_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RcBead); i {
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
			NumEnums:      6,
			NumMessages:   4,
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
