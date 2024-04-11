// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.3
// source: SC_11003.proto

package protobuf

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

type SC_11003 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                        *uint32             `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Name                      *string             `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
	Level                     *uint32             `protobuf:"varint,3,req,name=level" json:"level,omitempty"`
	Exp                       *uint32             `protobuf:"varint,4,req,name=exp" json:"exp,omitempty"`
	ResourceList              []*RESOURCE         `protobuf:"bytes,5,rep,name=resource_list,json=resourceList" json:"resource_list,omitempty"`
	AttackCount               *uint32             `protobuf:"varint,6,req,name=attack_count,json=attackCount" json:"attack_count,omitempty"`
	WinCount                  *uint32             `protobuf:"varint,7,req,name=win_count,json=winCount" json:"win_count,omitempty"`
	Adv                       *string             `protobuf:"bytes,8,req,name=adv" json:"adv,omitempty"`
	Character                 []uint32            `protobuf:"varint,9,rep,name=character" json:"character,omitempty"`
	ShipBagMax                *uint32             `protobuf:"varint,10,req,name=ship_bag_max,json=shipBagMax" json:"ship_bag_max,omitempty"`
	EquipBagMax               *uint32             `protobuf:"varint,11,req,name=equip_bag_max,json=equipBagMax" json:"equip_bag_max,omitempty"`
	GmFlag                    *uint32             `protobuf:"varint,12,req,name=gm_flag,json=gmFlag" json:"gm_flag,omitempty"`
	Rank                      *uint32             `protobuf:"varint,13,req,name=rank" json:"rank,omitempty"`
	PvpAttackCount            *uint32             `protobuf:"varint,14,req,name=pvp_attack_count,json=pvpAttackCount" json:"pvp_attack_count,omitempty"`
	PvpWinCount               *uint32             `protobuf:"varint,15,req,name=pvp_win_count,json=pvpWinCount" json:"pvp_win_count,omitempty"`
	CollectAttackCount        *uint32             `protobuf:"varint,16,req,name=collect_attack_count,json=collectAttackCount" json:"collect_attack_count,omitempty"`
	GuideIndex                *uint32             `protobuf:"varint,17,req,name=guide_index,json=guideIndex" json:"guide_index,omitempty"`
	BuyOilCount               *uint32             `protobuf:"varint,18,req,name=buy_oil_count,json=buyOilCount" json:"buy_oil_count,omitempty"`
	ChatRoomId                *uint32             `protobuf:"varint,19,req,name=chat_room_id,json=chatRoomId" json:"chat_room_id,omitempty"`
	CardList                  []*CARDINFO         `protobuf:"bytes,20,rep,name=card_list,json=cardList" json:"card_list,omitempty"`
	MaxRank                   *uint32             `protobuf:"varint,21,req,name=max_rank,json=maxRank" json:"max_rank,omitempty"`
	RegisterTime              *uint32             `protobuf:"varint,22,req,name=register_time,json=registerTime" json:"register_time,omitempty"`
	ShipCount                 *uint32             `protobuf:"varint,23,req,name=ship_count,json=shipCount" json:"ship_count,omitempty"`
	AccPayLv                  *uint32             `protobuf:"varint,24,req,name=acc_pay_lv,json=accPayLv" json:"acc_pay_lv,omitempty"`
	StoryList                 []uint32            `protobuf:"varint,25,rep,name=story_list,json=storyList" json:"story_list,omitempty"`
	GuildWaitTime             *uint32             `protobuf:"varint,26,req,name=guild_wait_time,json=guildWaitTime" json:"guild_wait_time,omitempty"`
	ChatMsgBanTime            *uint32             `protobuf:"varint,27,req,name=chat_msg_ban_time,json=chatMsgBanTime" json:"chat_msg_ban_time,omitempty"`
	FlagList                  []uint32            `protobuf:"varint,28,rep,name=flag_list,json=flagList" json:"flag_list,omitempty"`
	CdList                    []*COOLDOWN         `protobuf:"bytes,29,rep,name=cd_list,json=cdList" json:"cd_list,omitempty"`
	CommanderBagMax           *uint32             `protobuf:"varint,30,req,name=commander_bag_max,json=commanderBagMax" json:"commander_bag_max,omitempty"`
	MedalId                   []uint32            `protobuf:"varint,31,rep,name=medal_id,json=medalId" json:"medal_id,omitempty"`
	IconFrameList             []*IDTIMEINFO       `protobuf:"bytes,32,rep,name=icon_frame_list,json=iconFrameList" json:"icon_frame_list,omitempty"`
	ChatFrameList             []*IDTIMEINFO       `protobuf:"bytes,33,rep,name=chat_frame_list,json=chatFrameList" json:"chat_frame_list,omitempty"`
	Display                   *DISPLAYINFO        `protobuf:"bytes,34,opt,name=display" json:"display,omitempty"`
	Rmb                       *uint32             `protobuf:"varint,35,req,name=rmb" json:"rmb,omitempty"`
	Appreciation              *APPRECIATIONINFO   `protobuf:"bytes,36,req,name=appreciation" json:"appreciation,omitempty"`
	ThemeUploadNotAllowedTime *uint32             `protobuf:"varint,37,req,name=theme_upload_not_allowed_time,json=themeUploadNotAllowedTime" json:"theme_upload_not_allowed_time,omitempty"`
	RefundShopInfoList        []*REFUND_SHOPINFO  `protobuf:"bytes,38,rep,name=refund_shop_info_list,json=refundShopInfoList" json:"refund_shop_info_list,omitempty"`
	CartoonReadMark           []uint32            `protobuf:"varint,39,rep,name=cartoon_read_mark,json=cartoonReadMark" json:"cartoon_read_mark,omitempty"`
	CartoonCollectMark        []uint32            `protobuf:"varint,40,rep,name=cartoon_collect_mark,json=cartoonCollectMark" json:"cartoon_collect_mark,omitempty"`
	RandomShipMode            *uint32             `protobuf:"varint,41,req,name=random_ship_mode,json=randomShipMode" json:"random_ship_mode,omitempty"`
	RandomShipList            []uint32            `protobuf:"varint,42,rep,name=random_ship_list,json=randomShipList" json:"random_ship_list,omitempty"`
	MarryShip                 *uint32             `protobuf:"varint,43,req,name=marry_ship,json=marryShip" json:"marry_ship,omitempty"`
	TakingShipList            []*SHIP_TAKING_DATA `protobuf:"bytes,44,rep,name=taking_ship_list,json=takingShipList" json:"taking_ship_list,omitempty"`
	Soundstory                []uint32            `protobuf:"varint,45,rep,name=soundstory" json:"soundstory,omitempty"`
	ChildDisplay              *uint32             `protobuf:"varint,46,req,name=child_display,json=childDisplay" json:"child_display,omitempty"`
}

func (x *SC_11003) Reset() {
	*x = SC_11003{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SC_11003_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SC_11003) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SC_11003) ProtoMessage() {}

func (x *SC_11003) ProtoReflect() protoreflect.Message {
	mi := &file_SC_11003_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SC_11003.ProtoReflect.Descriptor instead.
func (*SC_11003) Descriptor() ([]byte, []int) {
	return file_SC_11003_proto_rawDescGZIP(), []int{0}
}

func (x *SC_11003) GetId() uint32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *SC_11003) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *SC_11003) GetLevel() uint32 {
	if x != nil && x.Level != nil {
		return *x.Level
	}
	return 0
}

func (x *SC_11003) GetExp() uint32 {
	if x != nil && x.Exp != nil {
		return *x.Exp
	}
	return 0
}

func (x *SC_11003) GetResourceList() []*RESOURCE {
	if x != nil {
		return x.ResourceList
	}
	return nil
}

func (x *SC_11003) GetAttackCount() uint32 {
	if x != nil && x.AttackCount != nil {
		return *x.AttackCount
	}
	return 0
}

func (x *SC_11003) GetWinCount() uint32 {
	if x != nil && x.WinCount != nil {
		return *x.WinCount
	}
	return 0
}

func (x *SC_11003) GetAdv() string {
	if x != nil && x.Adv != nil {
		return *x.Adv
	}
	return ""
}

func (x *SC_11003) GetCharacter() []uint32 {
	if x != nil {
		return x.Character
	}
	return nil
}

func (x *SC_11003) GetShipBagMax() uint32 {
	if x != nil && x.ShipBagMax != nil {
		return *x.ShipBagMax
	}
	return 0
}

func (x *SC_11003) GetEquipBagMax() uint32 {
	if x != nil && x.EquipBagMax != nil {
		return *x.EquipBagMax
	}
	return 0
}

func (x *SC_11003) GetGmFlag() uint32 {
	if x != nil && x.GmFlag != nil {
		return *x.GmFlag
	}
	return 0
}

func (x *SC_11003) GetRank() uint32 {
	if x != nil && x.Rank != nil {
		return *x.Rank
	}
	return 0
}

func (x *SC_11003) GetPvpAttackCount() uint32 {
	if x != nil && x.PvpAttackCount != nil {
		return *x.PvpAttackCount
	}
	return 0
}

func (x *SC_11003) GetPvpWinCount() uint32 {
	if x != nil && x.PvpWinCount != nil {
		return *x.PvpWinCount
	}
	return 0
}

func (x *SC_11003) GetCollectAttackCount() uint32 {
	if x != nil && x.CollectAttackCount != nil {
		return *x.CollectAttackCount
	}
	return 0
}

func (x *SC_11003) GetGuideIndex() uint32 {
	if x != nil && x.GuideIndex != nil {
		return *x.GuideIndex
	}
	return 0
}

func (x *SC_11003) GetBuyOilCount() uint32 {
	if x != nil && x.BuyOilCount != nil {
		return *x.BuyOilCount
	}
	return 0
}

func (x *SC_11003) GetChatRoomId() uint32 {
	if x != nil && x.ChatRoomId != nil {
		return *x.ChatRoomId
	}
	return 0
}

func (x *SC_11003) GetCardList() []*CARDINFO {
	if x != nil {
		return x.CardList
	}
	return nil
}

func (x *SC_11003) GetMaxRank() uint32 {
	if x != nil && x.MaxRank != nil {
		return *x.MaxRank
	}
	return 0
}

func (x *SC_11003) GetRegisterTime() uint32 {
	if x != nil && x.RegisterTime != nil {
		return *x.RegisterTime
	}
	return 0
}

func (x *SC_11003) GetShipCount() uint32 {
	if x != nil && x.ShipCount != nil {
		return *x.ShipCount
	}
	return 0
}

func (x *SC_11003) GetAccPayLv() uint32 {
	if x != nil && x.AccPayLv != nil {
		return *x.AccPayLv
	}
	return 0
}

func (x *SC_11003) GetStoryList() []uint32 {
	if x != nil {
		return x.StoryList
	}
	return nil
}

func (x *SC_11003) GetGuildWaitTime() uint32 {
	if x != nil && x.GuildWaitTime != nil {
		return *x.GuildWaitTime
	}
	return 0
}

func (x *SC_11003) GetChatMsgBanTime() uint32 {
	if x != nil && x.ChatMsgBanTime != nil {
		return *x.ChatMsgBanTime
	}
	return 0
}

func (x *SC_11003) GetFlagList() []uint32 {
	if x != nil {
		return x.FlagList
	}
	return nil
}

func (x *SC_11003) GetCdList() []*COOLDOWN {
	if x != nil {
		return x.CdList
	}
	return nil
}

func (x *SC_11003) GetCommanderBagMax() uint32 {
	if x != nil && x.CommanderBagMax != nil {
		return *x.CommanderBagMax
	}
	return 0
}

func (x *SC_11003) GetMedalId() []uint32 {
	if x != nil {
		return x.MedalId
	}
	return nil
}

func (x *SC_11003) GetIconFrameList() []*IDTIMEINFO {
	if x != nil {
		return x.IconFrameList
	}
	return nil
}

func (x *SC_11003) GetChatFrameList() []*IDTIMEINFO {
	if x != nil {
		return x.ChatFrameList
	}
	return nil
}

func (x *SC_11003) GetDisplay() *DISPLAYINFO {
	if x != nil {
		return x.Display
	}
	return nil
}

func (x *SC_11003) GetRmb() uint32 {
	if x != nil && x.Rmb != nil {
		return *x.Rmb
	}
	return 0
}

func (x *SC_11003) GetAppreciation() *APPRECIATIONINFO {
	if x != nil {
		return x.Appreciation
	}
	return nil
}

func (x *SC_11003) GetThemeUploadNotAllowedTime() uint32 {
	if x != nil && x.ThemeUploadNotAllowedTime != nil {
		return *x.ThemeUploadNotAllowedTime
	}
	return 0
}

func (x *SC_11003) GetRefundShopInfoList() []*REFUND_SHOPINFO {
	if x != nil {
		return x.RefundShopInfoList
	}
	return nil
}

func (x *SC_11003) GetCartoonReadMark() []uint32 {
	if x != nil {
		return x.CartoonReadMark
	}
	return nil
}

func (x *SC_11003) GetCartoonCollectMark() []uint32 {
	if x != nil {
		return x.CartoonCollectMark
	}
	return nil
}

func (x *SC_11003) GetRandomShipMode() uint32 {
	if x != nil && x.RandomShipMode != nil {
		return *x.RandomShipMode
	}
	return 0
}

func (x *SC_11003) GetRandomShipList() []uint32 {
	if x != nil {
		return x.RandomShipList
	}
	return nil
}

func (x *SC_11003) GetMarryShip() uint32 {
	if x != nil && x.MarryShip != nil {
		return *x.MarryShip
	}
	return 0
}

func (x *SC_11003) GetTakingShipList() []*SHIP_TAKING_DATA {
	if x != nil {
		return x.TakingShipList
	}
	return nil
}

func (x *SC_11003) GetSoundstory() []uint32 {
	if x != nil {
		return x.Soundstory
	}
	return nil
}

func (x *SC_11003) GetChildDisplay() uint32 {
	if x != nil && x.ChildDisplay != nil {
		return *x.ChildDisplay
	}
	return 0
}

var File_SC_11003_proto protoreflect.FileDescriptor

var file_SC_11003_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x53, 0x43, 0x5f, 0x31, 0x31, 0x30, 0x30, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x1a, 0x0e, 0x52, 0x45, 0x53, 0x4f, 0x55,
	0x52, 0x43, 0x45, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x43, 0x41, 0x52, 0x44, 0x49,
	0x4e, 0x46, 0x4f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x43, 0x4f, 0x4f, 0x4c, 0x44,
	0x4f, 0x57, 0x4e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x49, 0x44, 0x54, 0x49, 0x4d,
	0x45, 0x49, 0x4e, 0x46, 0x4f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x44, 0x49, 0x53,
	0x50, 0x4c, 0x41, 0x59, 0x49, 0x4e, 0x46, 0x4f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16,
	0x41, 0x50, 0x50, 0x52, 0x45, 0x43, 0x49, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x49, 0x4e, 0x46, 0x4f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x52, 0x45, 0x46, 0x55, 0x4e, 0x44, 0x5f, 0x53,
	0x48, 0x4f, 0x50, 0x49, 0x4e, 0x46, 0x4f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x53,
	0x48, 0x49, 0x50, 0x5f, 0x54, 0x41, 0x4b, 0x49, 0x4e, 0x47, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xec, 0x0d, 0x0a, 0x08, 0x53, 0x43, 0x5f, 0x31, 0x31, 0x30,
	0x30, 0x33, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18,
	0x03, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x10, 0x0a, 0x03,
	0x65, 0x78, 0x70, 0x18, 0x04, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x03, 0x65, 0x78, 0x70, 0x12, 0x36,
	0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x2e,
	0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0b, 0x61, 0x74,
	0x74, 0x61, 0x63, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x69, 0x6e,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x08, 0x77, 0x69,
	0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x64, 0x76, 0x18, 0x08, 0x20,
	0x02, 0x28, 0x09, 0x52, 0x03, 0x61, 0x64, 0x76, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x72,
	0x61, 0x63, 0x74, 0x65, 0x72, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x68, 0x61,
	0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0c, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x62,
	0x61, 0x67, 0x5f, 0x6d, 0x61, 0x78, 0x18, 0x0a, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0a, 0x73, 0x68,
	0x69, 0x70, 0x42, 0x61, 0x67, 0x4d, 0x61, 0x78, 0x12, 0x22, 0x0a, 0x0d, 0x65, 0x71, 0x75, 0x69,
	0x70, 0x5f, 0x62, 0x61, 0x67, 0x5f, 0x6d, 0x61, 0x78, 0x18, 0x0b, 0x20, 0x02, 0x28, 0x0d, 0x52,
	0x0b, 0x65, 0x71, 0x75, 0x69, 0x70, 0x42, 0x61, 0x67, 0x4d, 0x61, 0x78, 0x12, 0x17, 0x0a, 0x07,
	0x67, 0x6d, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x0c, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x06, 0x67,
	0x6d, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x0d, 0x20,
	0x02, 0x28, 0x0d, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x12, 0x28, 0x0a, 0x10, 0x70, 0x76, 0x70,
	0x5f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0e, 0x20,
	0x02, 0x28, 0x0d, 0x52, 0x0e, 0x70, 0x76, 0x70, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x70, 0x76, 0x70, 0x5f, 0x77, 0x69, 0x6e, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0f, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0b, 0x70, 0x76, 0x70, 0x57,
	0x69, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x14, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x5f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x10, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x12, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x41, 0x74,
	0x74, 0x61, 0x63, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x67, 0x75, 0x69,
	0x64, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x11, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0a,
	0x67, 0x75, 0x69, 0x64, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x22, 0x0a, 0x0d, 0x62, 0x75,
	0x79, 0x5f, 0x6f, 0x69, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x12, 0x20, 0x02, 0x28,
	0x0d, 0x52, 0x0b, 0x62, 0x75, 0x79, 0x4f, 0x69, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20,
	0x0a, 0x0c, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x13,
	0x20, 0x02, 0x28, 0x0d, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x64,
	0x12, 0x2e, 0x0a, 0x09, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x14, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x2e, 0x43, 0x41,
	0x52, 0x44, 0x49, 0x4e, 0x46, 0x4f, 0x52, 0x08, 0x63, 0x61, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x6d, 0x61, 0x78, 0x5f, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x15, 0x20, 0x02,
	0x28, 0x0d, 0x52, 0x07, 0x6d, 0x61, 0x78, 0x52, 0x61, 0x6e, 0x6b, 0x12, 0x23, 0x0a, 0x0d, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x16, 0x20, 0x02,
	0x28, 0x0d, 0x52, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x17,
	0x20, 0x02, 0x28, 0x0d, 0x52, 0x09, 0x73, 0x68, 0x69, 0x70, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1c, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x5f, 0x70, 0x61, 0x79, 0x5f, 0x6c, 0x76, 0x18, 0x18, 0x20,
	0x02, 0x28, 0x0d, 0x52, 0x08, 0x61, 0x63, 0x63, 0x50, 0x61, 0x79, 0x4c, 0x76, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x19, 0x20, 0x03, 0x28,
	0x0d, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0f,
	0x67, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x77, 0x61, 0x69, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x1a, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0d, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x57, 0x61, 0x69, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x11, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x6d, 0x73, 0x67,
	0x5f, 0x62, 0x61, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x1b, 0x20, 0x02, 0x28, 0x0d, 0x52,
	0x0e, 0x63, 0x68, 0x61, 0x74, 0x4d, 0x73, 0x67, 0x42, 0x61, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x66, 0x6c, 0x61, 0x67, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x1c, 0x20, 0x03,
	0x28, 0x0d, 0x52, 0x08, 0x66, 0x6c, 0x61, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x07,
	0x63, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x1d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x2e, 0x43, 0x4f, 0x4f, 0x4c, 0x44, 0x4f, 0x57, 0x4e,
	0x52, 0x06, 0x63, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x61, 0x67, 0x5f, 0x6d, 0x61, 0x78, 0x18, 0x1e, 0x20,
	0x02, 0x28, 0x0d, 0x52, 0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x42, 0x61,
	0x67, 0x4d, 0x61, 0x78, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x65, 0x64, 0x61, 0x6c, 0x5f, 0x69, 0x64,
	0x18, 0x1f, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x07, 0x6d, 0x65, 0x64, 0x61, 0x6c, 0x49, 0x64, 0x12,
	0x3b, 0x0a, 0x0f, 0x69, 0x63, 0x6f, 0x6e, 0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x20, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62, 0x65, 0x6c, 0x66, 0x61,
	0x73, 0x74, 0x2e, 0x49, 0x44, 0x54, 0x49, 0x4d, 0x45, 0x49, 0x4e, 0x46, 0x4f, 0x52, 0x0d, 0x69,
	0x63, 0x6f, 0x6e, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x0f,
	0x63, 0x68, 0x61, 0x74, 0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x21, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x2e,
	0x49, 0x44, 0x54, 0x49, 0x4d, 0x45, 0x49, 0x4e, 0x46, 0x4f, 0x52, 0x0d, 0x63, 0x68, 0x61, 0x74,
	0x46, 0x72, 0x61, 0x6d, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x07, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x18, 0x22, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x62, 0x65, 0x6c,
	0x66, 0x61, 0x73, 0x74, 0x2e, 0x44, 0x49, 0x53, 0x50, 0x4c, 0x41, 0x59, 0x49, 0x4e, 0x46, 0x4f,
	0x52, 0x07, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x6d, 0x62,
	0x18, 0x23, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x03, 0x72, 0x6d, 0x62, 0x12, 0x3d, 0x0a, 0x0c, 0x61,
	0x70, 0x70, 0x72, 0x65, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x24, 0x20, 0x02, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x2e, 0x41, 0x50, 0x50, 0x52,
	0x45, 0x43, 0x49, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x49, 0x4e, 0x46, 0x4f, 0x52, 0x0c, 0x61, 0x70,
	0x70, 0x72, 0x65, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x1d, 0x74, 0x68,
	0x65, 0x6d, 0x65, 0x5f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x61,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x25, 0x20, 0x02, 0x28,
	0x0d, 0x52, 0x19, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4e, 0x6f,
	0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x4b, 0x0a, 0x15,
	0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x26, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62, 0x65,
	0x6c, 0x66, 0x61, 0x73, 0x74, 0x2e, 0x52, 0x45, 0x46, 0x55, 0x4e, 0x44, 0x5f, 0x53, 0x48, 0x4f,
	0x50, 0x49, 0x4e, 0x46, 0x4f, 0x52, 0x12, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x53, 0x68, 0x6f,
	0x70, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x63, 0x61, 0x72,
	0x74, 0x6f, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x27,
	0x20, 0x03, 0x28, 0x0d, 0x52, 0x0f, 0x63, 0x61, 0x72, 0x74, 0x6f, 0x6f, 0x6e, 0x52, 0x65, 0x61,
	0x64, 0x4d, 0x61, 0x72, 0x6b, 0x12, 0x30, 0x0a, 0x14, 0x63, 0x61, 0x72, 0x74, 0x6f, 0x6f, 0x6e,
	0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x28, 0x20,
	0x03, 0x28, 0x0d, 0x52, 0x12, 0x63, 0x61, 0x72, 0x74, 0x6f, 0x6f, 0x6e, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x12, 0x28, 0x0a, 0x10, 0x72, 0x61, 0x6e, 0x64, 0x6f,
	0x6d, 0x5f, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x29, 0x20, 0x02, 0x28,
	0x0d, 0x52, 0x0e, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x53, 0x68, 0x69, 0x70, 0x4d, 0x6f, 0x64,
	0x65, 0x12, 0x28, 0x0a, 0x10, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x5f, 0x73, 0x68, 0x69, 0x70,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x2a, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0e, 0x72, 0x61, 0x6e,
	0x64, 0x6f, 0x6d, 0x53, 0x68, 0x69, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6d,
	0x61, 0x72, 0x72, 0x79, 0x5f, 0x73, 0x68, 0x69, 0x70, 0x18, 0x2b, 0x20, 0x02, 0x28, 0x0d, 0x52,
	0x09, 0x6d, 0x61, 0x72, 0x72, 0x79, 0x53, 0x68, 0x69, 0x70, 0x12, 0x43, 0x0a, 0x10, 0x74, 0x61,
	0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x2c,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x2e, 0x53,
	0x48, 0x49, 0x50, 0x5f, 0x54, 0x41, 0x4b, 0x49, 0x4e, 0x47, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x52,
	0x0e, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x68, 0x69, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x1e, 0x0a, 0x0a, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x2d, 0x20,
	0x03, 0x28, 0x0d, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12,
	0x23, 0x0a, 0x0d, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x18, 0x2e, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0c, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x44, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66,
}

var (
	file_SC_11003_proto_rawDescOnce sync.Once
	file_SC_11003_proto_rawDescData = file_SC_11003_proto_rawDesc
)

func file_SC_11003_proto_rawDescGZIP() []byte {
	file_SC_11003_proto_rawDescOnce.Do(func() {
		file_SC_11003_proto_rawDescData = protoimpl.X.CompressGZIP(file_SC_11003_proto_rawDescData)
	})
	return file_SC_11003_proto_rawDescData
}

var file_SC_11003_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SC_11003_proto_goTypes = []interface{}{
	(*SC_11003)(nil),         // 0: belfast.SC_11003
	(*RESOURCE)(nil),         // 1: belfast.RESOURCE
	(*CARDINFO)(nil),         // 2: belfast.CARDINFO
	(*COOLDOWN)(nil),         // 3: belfast.COOLDOWN
	(*IDTIMEINFO)(nil),       // 4: belfast.IDTIMEINFO
	(*DISPLAYINFO)(nil),      // 5: belfast.DISPLAYINFO
	(*APPRECIATIONINFO)(nil), // 6: belfast.APPRECIATIONINFO
	(*REFUND_SHOPINFO)(nil),  // 7: belfast.REFUND_SHOPINFO
	(*SHIP_TAKING_DATA)(nil), // 8: belfast.SHIP_TAKING_DATA
}
var file_SC_11003_proto_depIdxs = []int32{
	1, // 0: belfast.SC_11003.resource_list:type_name -> belfast.RESOURCE
	2, // 1: belfast.SC_11003.card_list:type_name -> belfast.CARDINFO
	3, // 2: belfast.SC_11003.cd_list:type_name -> belfast.COOLDOWN
	4, // 3: belfast.SC_11003.icon_frame_list:type_name -> belfast.IDTIMEINFO
	4, // 4: belfast.SC_11003.chat_frame_list:type_name -> belfast.IDTIMEINFO
	5, // 5: belfast.SC_11003.display:type_name -> belfast.DISPLAYINFO
	6, // 6: belfast.SC_11003.appreciation:type_name -> belfast.APPRECIATIONINFO
	7, // 7: belfast.SC_11003.refund_shop_info_list:type_name -> belfast.REFUND_SHOPINFO
	8, // 8: belfast.SC_11003.taking_ship_list:type_name -> belfast.SHIP_TAKING_DATA
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_SC_11003_proto_init() }
func file_SC_11003_proto_init() {
	if File_SC_11003_proto != nil {
		return
	}
	file_RESOURCE_proto_init()
	file_CARDINFO_proto_init()
	file_COOLDOWN_proto_init()
	file_IDTIMEINFO_proto_init()
	file_DISPLAYINFO_proto_init()
	file_APPRECIATIONINFO_proto_init()
	file_REFUND_SHOPINFO_proto_init()
	file_SHIP_TAKING_DATA_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SC_11003_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SC_11003); i {
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
			RawDescriptor: file_SC_11003_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SC_11003_proto_goTypes,
		DependencyIndexes: file_SC_11003_proto_depIdxs,
		MessageInfos:      file_SC_11003_proto_msgTypes,
	}.Build()
	File_SC_11003_proto = out.File
	file_SC_11003_proto_rawDesc = nil
	file_SC_11003_proto_goTypes = nil
	file_SC_11003_proto_depIdxs = nil
}
