// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: recorder/record.proto

package recorder

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
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

// 動作類型
type StepActionType int32

const (
	// 未指定
	StepActionType_STEP_ACTION_TYPE_UNSPECIFIED StepActionType = 0
	// 發牌
	StepActionType_DEAL StepActionType = 1
	// 移除場上
	StepActionType_REMOVE StepActionType = 2
	// 移轉
	StepActionType_TRANSFER StepActionType = 3
	// 狀態改變
	StepActionType_STATUS StepActionType = 4
)

// Enum value maps for StepActionType.
var (
	StepActionType_name = map[int32]string{
		0: "STEP_ACTION_TYPE_UNSPECIFIED",
		1: "DEAL",
		2: "REMOVE",
		3: "TRANSFER",
		4: "STATUS",
	}
	StepActionType_value = map[string]int32{
		"STEP_ACTION_TYPE_UNSPECIFIED": 0,
		"DEAL":                         1,
		"REMOVE":                       2,
		"TRANSFER":                     3,
		"STATUS":                       4,
	}
)

func (x StepActionType) Enum() *StepActionType {
	p := new(StepActionType)
	*p = x
	return p
}

func (x StepActionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StepActionType) Descriptor() protoreflect.EnumDescriptor {
	return file_recorder_record_proto_enumTypes[0].Descriptor()
}

func (StepActionType) Type() protoreflect.EnumType {
	return &file_recorder_record_proto_enumTypes[0]
}

func (x StepActionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StepActionType.Descriptor instead.
func (StepActionType) EnumDescriptor() ([]byte, []int) {
	return file_recorder_record_proto_rawDescGZIP(), []int{0}
}

// 卡牌，代表在此局中使用的各類物品
type Card struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type int64 `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	// 卡牌內容代碼：CardTypeCode。
	Code int64 `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	// 卡牌ID：牌唯一碼，同ID表示同一張實體牌。不使用此欄位填0。
	Id int64 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	// 狀態：翻開、關起、橫放...
	Status bool `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Card) Reset() {
	*x = Card{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recorder_record_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Card) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Card) ProtoMessage() {}

func (x *Card) ProtoReflect() protoreflect.Message {
	mi := &file_recorder_record_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Card.ProtoReflect.Descriptor instead.
func (*Card) Descriptor() ([]byte, []int) {
	return file_recorder_record_proto_rawDescGZIP(), []int{0}
}

func (x *Card) GetType() int64 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Card) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Card) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Card) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

// 卡牌列表
type CardList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*Card `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *CardList) Reset() {
	*x = CardList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recorder_record_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CardList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CardList) ProtoMessage() {}

func (x *CardList) ProtoReflect() protoreflect.Message {
	mi := &file_recorder_record_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CardList.ProtoReflect.Descriptor instead.
func (*CardList) Descriptor() ([]byte, []int) {
	return file_recorder_record_proto_rawDescGZIP(), []int{1}
}

func (x *CardList) GetList() []*Card {
	if x != nil {
		return x.List
	}
	return nil
}

// 執行動作：對某些資源進行增、刪、移轉
type StepAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 動作代碼：StepActionCode
	Code int64 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// 原資源擁有座位。
	SourceSeat int64 `protobuf:"varint,3,opt,name=source_seat,json=sourceSeat,proto3" json:"source_seat,omitempty"`
	// 資源轉移目標座位。
	TargetSeat []int64 `protobuf:"varint,4,rep,packed,name=target_seat,json=targetSeat,proto3" json:"target_seat,omitempty"`
	// 目標卡牌資源
	// key: ResourceTypeCode, value: 卡牌陣列
	Cards map[int64]*CardList `protobuf:"bytes,6,rep,name=cards,proto3" json:"cards,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// 目標分數資源
	// key: ResourceTypeCode, value: 分數值。
	Scores map[int64]int64 `protobuf:"bytes,7,rep,name=scores,proto3" json:"scores,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// 動作類型
	Type StepActionType `protobuf:"varint,8,opt,name=type,proto3,enum=recorder.StepActionType" json:"type,omitempty"`
	// （可選）附加資料
	Data *anypb.Any `protobuf:"bytes,9,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *StepAction) Reset() {
	*x = StepAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recorder_record_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StepAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StepAction) ProtoMessage() {}

func (x *StepAction) ProtoReflect() protoreflect.Message {
	mi := &file_recorder_record_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StepAction.ProtoReflect.Descriptor instead.
func (*StepAction) Descriptor() ([]byte, []int) {
	return file_recorder_record_proto_rawDescGZIP(), []int{2}
}

func (x *StepAction) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *StepAction) GetSourceSeat() int64 {
	if x != nil {
		return x.SourceSeat
	}
	return 0
}

func (x *StepAction) GetTargetSeat() []int64 {
	if x != nil {
		return x.TargetSeat
	}
	return nil
}

func (x *StepAction) GetCards() map[int64]*CardList {
	if x != nil {
		return x.Cards
	}
	return nil
}

func (x *StepAction) GetScores() map[int64]int64 {
	if x != nil {
		return x.Scores
	}
	return nil
}

func (x *StepAction) GetType() StepActionType {
	if x != nil {
		return x.Type
	}
	return StepActionType_STEP_ACTION_TYPE_UNSPECIFIED
}

func (x *StepAction) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

// 遊戲步驟
type Step struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 步驟代碼：StepCode
	Code int64 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// 玩家、執行者：SeatCode
	Seat int64 `protobuf:"varint,3,opt,name=seat,proto3" json:"seat,omitempty"`
	// 時間戳
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// 執行動作
	Actions []*StepAction `protobuf:"bytes,6,rep,name=actions,proto3" json:"actions,omitempty"`
	// 是否是修正錯誤
	Modify bool `protobuf:"varint,8,opt,name=modify,proto3" json:"modify,omitempty"`
}

func (x *Step) Reset() {
	*x = Step{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recorder_record_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Step) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Step) ProtoMessage() {}

func (x *Step) ProtoReflect() protoreflect.Message {
	mi := &file_recorder_record_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Step.ProtoReflect.Descriptor instead.
func (*Step) Descriptor() ([]byte, []int) {
	return file_recorder_record_proto_rawDescGZIP(), []int{3}
}

func (x *Step) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Step) GetSeat() int64 {
	if x != nil {
		return x.Seat
	}
	return 0
}

func (x *Step) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Step) GetActions() []*StepAction {
	if x != nil {
		return x.Actions
	}
	return nil
}

func (x *Step) GetModify() bool {
	if x != nil {
		return x.Modify
	}
	return false
}

// 座位
type Seats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 座位功能代碼：SeatCode
	Code int64 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// 各類手牌
	// key: ResourceTypeCode, value: 卡牌陣列
	Cards map[int64]*CardList `protobuf:"bytes,6,rep,name=cards,proto3" json:"cards,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// 各類分數
	// key: ResourceTypeCode, value: 分數值。
	Scores map[int64]int64 `protobuf:"bytes,7,rep,name=scores,proto3" json:"scores,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *Seats) Reset() {
	*x = Seats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recorder_record_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Seats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Seats) ProtoMessage() {}

func (x *Seats) ProtoReflect() protoreflect.Message {
	mi := &file_recorder_record_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Seats.ProtoReflect.Descriptor instead.
func (*Seats) Descriptor() ([]byte, []int) {
	return file_recorder_record_proto_rawDescGZIP(), []int{4}
}

func (x *Seats) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Seats) GetCards() map[int64]*CardList {
	if x != nil {
		return x.Cards
	}
	return nil
}

func (x *Seats) GetScores() map[int64]int64 {
	if x != nil {
		return x.Scores
	}
	return nil
}

// 記錄各座位狀態
type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 判讀類型代碼：ResultRecordTypeCode
	Code int64 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// 紀錄者座位代碼
	BySeat string `protobuf:"bytes,3,opt,name=by_seat,json=bySeat,proto3" json:"by_seat,omitempty"`
	// 時間戳
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// 座位組
	// key: SeatCode, value: Seat
	Seats map[int64]*Seats `protobuf:"bytes,6,rep,name=seats,proto3" json:"seats,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recorder_record_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_recorder_record_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_recorder_record_proto_rawDescGZIP(), []int{5}
}

func (x *Result) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Result) GetBySeat() string {
	if x != nil {
		return x.BySeat
	}
	return ""
}

func (x *Result) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Result) GetSeats() map[int64]*Seats {
	if x != nil {
		return x.Seats
	}
	return nil
}

// 紀錄
type RoundRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 紀錄ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// 遊戲代碼
	GameCode string `protobuf:"bytes,2,opt,name=game_code,json=gameCode,proto3" json:"game_code,omitempty"`
	// 桌代碼
	TableCode string `protobuf:"bytes,3,opt,name=table_code,json=tableCode,proto3" json:"table_code,omitempty"`
	// 局代碼
	RoundCode string `protobuf:"bytes,4,opt,name=round_code,json=roundCode,proto3" json:"round_code,omitempty"`
	// 使用洗牌ID
	ShuffleCode string `protobuf:"bytes,5,opt,name=shuffle_code,json=shuffleCode,proto3" json:"shuffle_code,omitempty"`
	// 使用此洗牌第幾局，從1開始計算
	ShuffleRound int64 `protobuf:"varint,6,opt,name=shuffle_round,json=shuffleRound,proto3" json:"shuffle_round,omitempty"`
	// 遊戲版本
	GameVersion string `protobuf:"bytes,7,opt,name=game_version,json=gameVersion,proto3" json:"game_version,omitempty"`
	// 座位組
	// key: SeatCode, value: Seat
	Seats map[int64]*Seats `protobuf:"bytes,8,rep,name=seats,proto3" json:"seats,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// 玩家入座
	// key: SeatCode, value: 玩家code。
	Players map[int64]string `protobuf:"bytes,9,rep,name=players,proto3" json:"players,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// 歷程
	Process []*Step `protobuf:"bytes,10,rep,name=process,proto3" json:"process,omitempty"`
	// 判讀
	Results []*Result `protobuf:"bytes,11,rep,name=results,proto3" json:"results,omitempty"`
	// 取消
	Cancel bool `protobuf:"varint,12,opt,name=cancel,proto3" json:"cancel,omitempty"`
	// 取消代碼：CancelReasonCode
	CancelCode int64 `protobuf:"varint,13,opt,name=cancel_code,json=cancelCode,proto3" json:"cancel_code,omitempty"`
}

func (x *RoundRecord) Reset() {
	*x = RoundRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recorder_record_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoundRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoundRecord) ProtoMessage() {}

func (x *RoundRecord) ProtoReflect() protoreflect.Message {
	mi := &file_recorder_record_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoundRecord.ProtoReflect.Descriptor instead.
func (*RoundRecord) Descriptor() ([]byte, []int) {
	return file_recorder_record_proto_rawDescGZIP(), []int{6}
}

func (x *RoundRecord) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RoundRecord) GetGameCode() string {
	if x != nil {
		return x.GameCode
	}
	return ""
}

func (x *RoundRecord) GetTableCode() string {
	if x != nil {
		return x.TableCode
	}
	return ""
}

func (x *RoundRecord) GetRoundCode() string {
	if x != nil {
		return x.RoundCode
	}
	return ""
}

func (x *RoundRecord) GetShuffleCode() string {
	if x != nil {
		return x.ShuffleCode
	}
	return ""
}

func (x *RoundRecord) GetShuffleRound() int64 {
	if x != nil {
		return x.ShuffleRound
	}
	return 0
}

func (x *RoundRecord) GetGameVersion() string {
	if x != nil {
		return x.GameVersion
	}
	return ""
}

func (x *RoundRecord) GetSeats() map[int64]*Seats {
	if x != nil {
		return x.Seats
	}
	return nil
}

func (x *RoundRecord) GetPlayers() map[int64]string {
	if x != nil {
		return x.Players
	}
	return nil
}

func (x *RoundRecord) GetProcess() []*Step {
	if x != nil {
		return x.Process
	}
	return nil
}

func (x *RoundRecord) GetResults() []*Result {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *RoundRecord) GetCancel() bool {
	if x != nil {
		return x.Cancel
	}
	return false
}

func (x *RoundRecord) GetCancelCode() int64 {
	if x != nil {
		return x.CancelCode
	}
	return 0
}

var File_recorder_record_proto protoreflect.FileDescriptor

var file_recorder_record_proto_rawDesc = []byte{
	0x0a, 0x15, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a,
	0x04, 0x43, 0x61, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2e, 0x0a, 0x08, 0x43, 0x61, 0x72, 0x64, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x22, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x52,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0xba, 0x03, 0x0a, 0x0a, 0x53, 0x74, 0x65, 0x70, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x73, 0x65, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x65, 0x61, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x61, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0a,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x65, 0x61, 0x74, 0x12, 0x35, 0x0a, 0x05, 0x63, 0x61,
	0x72, 0x64, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x65, 0x70, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x43, 0x61, 0x72, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x63, 0x61, 0x72, 0x64,
	0x73, 0x12, 0x38, 0x0a, 0x06, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x65,
	0x70, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x06, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x65, 0x70, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x1a, 0x4c, 0x0a, 0x0a, 0x43, 0x61, 0x72, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x61,
	0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x39, 0x0a, 0x0b, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x4a, 0x04, 0x08, 0x02,
	0x10, 0x03, 0x22, 0xc2, 0x01, 0x0a, 0x04, 0x53, 0x74, 0x65, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x65, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73,
	0x65, 0x61, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x2e, 0x0a,
	0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x65, 0x70, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x6d,
	0x6f, 0x64, 0x69, 0x66, 0x79, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x4a, 0x04, 0x08, 0x04, 0x10,
	0x05, 0x4a, 0x04, 0x08, 0x07, 0x10, 0x08, 0x22, 0x91, 0x02, 0x0a, 0x05, 0x53, 0x65, 0x61, 0x74,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x63, 0x61, 0x72, 0x64, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e,
	0x53, 0x65, 0x61, 0x74, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x05, 0x63, 0x61, 0x72, 0x64, 0x73, 0x12, 0x33, 0x0a, 0x06, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x2e, 0x53, 0x65, 0x61, 0x74, 0x73, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x1a, 0x4c, 0x0a, 0x0a,
	0x43, 0x61, 0x72, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x39, 0x0a, 0x0b, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x06, 0x22, 0xf9, 0x01, 0x0a, 0x06,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x79,
	0x5f, 0x73, 0x65, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x79, 0x53,
	0x65, 0x61, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x31, 0x0a,
	0x05, 0x73, 0x65, 0x61, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x53,
	0x65, 0x61, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x73, 0x65, 0x61, 0x74, 0x73,
	0x1a, 0x49, 0x0a, 0x0a, 0x53, 0x65, 0x61, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x25, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x61, 0x74, 0x73,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x4a, 0x04, 0x08, 0x02, 0x10,
	0x03, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x05, 0x22, 0xef, 0x04, 0x0a, 0x0b, 0x52, 0x6f, 0x75, 0x6e,
	0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x61, 0x6d, 0x65, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x61, 0x6d, 0x65,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x68, 0x75, 0x66, 0x66, 0x6c, 0x65, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x68, 0x75, 0x66, 0x66, 0x6c,
	0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x68, 0x75, 0x66, 0x66, 0x6c, 0x65,
	0x5f, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x73, 0x68,
	0x75, 0x66, 0x66, 0x6c, 0x65, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x67, 0x61,
	0x6d, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x67, 0x61, 0x6d, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a,
	0x05, 0x73, 0x65, 0x61, 0x74, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x2e, 0x53, 0x65, 0x61, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05,
	0x73, 0x65, 0x61, 0x74, 0x73, 0x12, 0x3c, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73,
	0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x2e, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x73, 0x12, 0x28, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x18, 0x0a,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e,
	0x53, 0x74, 0x65, 0x70, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x12, 0x2a, 0x0a,
	0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x63, 0x61, 0x6e, 0x63, 0x65,
	0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x43, 0x6f,
	0x64, 0x65, 0x1a, 0x49, 0x0a, 0x0a, 0x53, 0x65, 0x61, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x25, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x61,
	0x74, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3a, 0x0a,
	0x0c, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x62, 0x0a, 0x0e, 0x53, 0x74, 0x65,
	0x70, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x1c, 0x53,
	0x54, 0x45, 0x50, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a,
	0x04, 0x44, 0x45, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x45, 0x4d, 0x4f, 0x56,
	0x45, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x10,
	0x03, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x10, 0x04, 0x42, 0x90, 0x01,
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x0b,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x33, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6b, 0x79, 0x2d, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x73, 0x6b, 0x79, 0x63, 0x61, 0x73, 0x69, 0x6e, 0x6f,
	0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0xa2, 0x02, 0x03, 0x52, 0x58, 0x58, 0xaa, 0x02, 0x08, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0xca, 0x02, 0x08, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0xe2, 0x02,
	0x14, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_recorder_record_proto_rawDescOnce sync.Once
	file_recorder_record_proto_rawDescData = file_recorder_record_proto_rawDesc
)

func file_recorder_record_proto_rawDescGZIP() []byte {
	file_recorder_record_proto_rawDescOnce.Do(func() {
		file_recorder_record_proto_rawDescData = protoimpl.X.CompressGZIP(file_recorder_record_proto_rawDescData)
	})
	return file_recorder_record_proto_rawDescData
}

var file_recorder_record_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_recorder_record_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_recorder_record_proto_goTypes = []interface{}{
	(StepActionType)(0),           // 0: recorder.StepActionType
	(*Card)(nil),                  // 1: recorder.Card
	(*CardList)(nil),              // 2: recorder.CardList
	(*StepAction)(nil),            // 3: recorder.StepAction
	(*Step)(nil),                  // 4: recorder.Step
	(*Seats)(nil),                 // 5: recorder.Seats
	(*Result)(nil),                // 6: recorder.Result
	(*RoundRecord)(nil),           // 7: recorder.RoundRecord
	nil,                           // 8: recorder.StepAction.CardsEntry
	nil,                           // 9: recorder.StepAction.ScoresEntry
	nil,                           // 10: recorder.Seats.CardsEntry
	nil,                           // 11: recorder.Seats.ScoresEntry
	nil,                           // 12: recorder.Result.SeatsEntry
	nil,                           // 13: recorder.RoundRecord.SeatsEntry
	nil,                           // 14: recorder.RoundRecord.PlayersEntry
	(*anypb.Any)(nil),             // 15: google.protobuf.Any
	(*timestamppb.Timestamp)(nil), // 16: google.protobuf.Timestamp
}
var file_recorder_record_proto_depIdxs = []int32{
	1,  // 0: recorder.CardList.list:type_name -> recorder.Card
	8,  // 1: recorder.StepAction.cards:type_name -> recorder.StepAction.CardsEntry
	9,  // 2: recorder.StepAction.scores:type_name -> recorder.StepAction.ScoresEntry
	0,  // 3: recorder.StepAction.type:type_name -> recorder.StepActionType
	15, // 4: recorder.StepAction.data:type_name -> google.protobuf.Any
	16, // 5: recorder.Step.timestamp:type_name -> google.protobuf.Timestamp
	3,  // 6: recorder.Step.actions:type_name -> recorder.StepAction
	10, // 7: recorder.Seats.cards:type_name -> recorder.Seats.CardsEntry
	11, // 8: recorder.Seats.scores:type_name -> recorder.Seats.ScoresEntry
	16, // 9: recorder.Result.timestamp:type_name -> google.protobuf.Timestamp
	12, // 10: recorder.Result.seats:type_name -> recorder.Result.SeatsEntry
	13, // 11: recorder.RoundRecord.seats:type_name -> recorder.RoundRecord.SeatsEntry
	14, // 12: recorder.RoundRecord.players:type_name -> recorder.RoundRecord.PlayersEntry
	4,  // 13: recorder.RoundRecord.process:type_name -> recorder.Step
	6,  // 14: recorder.RoundRecord.results:type_name -> recorder.Result
	2,  // 15: recorder.StepAction.CardsEntry.value:type_name -> recorder.CardList
	2,  // 16: recorder.Seats.CardsEntry.value:type_name -> recorder.CardList
	5,  // 17: recorder.Result.SeatsEntry.value:type_name -> recorder.Seats
	5,  // 18: recorder.RoundRecord.SeatsEntry.value:type_name -> recorder.Seats
	19, // [19:19] is the sub-list for method output_type
	19, // [19:19] is the sub-list for method input_type
	19, // [19:19] is the sub-list for extension type_name
	19, // [19:19] is the sub-list for extension extendee
	0,  // [0:19] is the sub-list for field type_name
}

func init() { file_recorder_record_proto_init() }
func file_recorder_record_proto_init() {
	if File_recorder_record_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_recorder_record_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Card); i {
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
		file_recorder_record_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CardList); i {
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
		file_recorder_record_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StepAction); i {
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
		file_recorder_record_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Step); i {
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
		file_recorder_record_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Seats); i {
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
		file_recorder_record_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_recorder_record_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoundRecord); i {
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
			RawDescriptor: file_recorder_record_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_recorder_record_proto_goTypes,
		DependencyIndexes: file_recorder_record_proto_depIdxs,
		EnumInfos:         file_recorder_record_proto_enumTypes,
		MessageInfos:      file_recorder_record_proto_msgTypes,
	}.Build()
	File_recorder_record_proto = out.File
	file_recorder_record_proto_rawDesc = nil
	file_recorder_record_proto_goTypes = nil
	file_recorder_record_proto_depIdxs = nil
}
