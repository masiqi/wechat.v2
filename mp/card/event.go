package card

import (
	"github.com/masiqi/wechat.v2/mp/core"
)

const (
	// 推送到公众号URL上的事件类型
	EventTypeCardPassCheck            core.EventType = "card_pass_check"              // 卡券通过审核
	EventTypeCardNotPassCheck         core.EventType = "card_not_pass_check"          // 卡券未通过审核
	EventTypeUserGetCard              core.EventType = "user_get_card"                // 领取事件推送
	EventTypeUserDelCard              core.EventType = "user_del_card"                // 删除事件推送
	EventTypeUserConsumeCard          core.EventType = "user_consume_card"            // 核销事件推送
	EventTypeUserViewCard             core.EventType = "user_view_card"               // 进入会员卡事件推送
	EventTypeUserEnterSessionFromCard core.EventType = "user_enter_session_from_card" // 从卡券进入公众号会话事件推送
)

// 卡券通过审核, 微信会把这个事件推送到开发者填写的URL
type CardPassCheckEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader

	EventType core.EventType `xml:"Event"  json:"Event"`  // 事件类型, card_pass_check
	CardId    string         `xml:"CardId" json:"CardId"` // 卡券ID
}

func GetCardPassCheckEvent(msg *core.MixedMsg) *CardPassCheckEvent {
	return &CardPassCheckEvent{
		MsgHeader: msg.MsgHeader,
		EventType: msg.EventType,
		CardId:    msg.CardId,
	}
}

// 卡券未通过审核, 微信会把这个事件推送到开发者填写的URL
type CardNotPassCheckEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader

	EventType core.EventType `xml:"Event"  json:"Event"`  // 事件类型, card_not_pass_check
	CardId    string         `xml:"CardId" json:"CardId"` // 卡券ID
}

func GetCardNotPassCheckEvent(msg *core.MixedMsg) *CardNotPassCheckEvent {
	return &CardNotPassCheckEvent{
		MsgHeader: msg.MsgHeader,
		EventType: msg.EventType,
		CardId:    msg.CardId,
	}
}

// 用户在领取卡券时, 微信会把这个事件推送到开发者填写的URL.
type UserGetCardEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader

	EventType       core.EventType `xml:"Event"           json:"Event"`           // 事件类型, user_get_card
	CardId          string         `xml:"CardId"          json:"CardId"`          // 卡券ID
	IsGiveByFriend  int            `xml:"IsGiveByFriend"  json:"IsGiveByFriend"`  // 是否为转赠, 1 代表是, 0 代表否.
	FriendUserName  string         `xml:"FriendUserName"  json:"FriendUserName"`  // 赠送方账号(一个OpenID), "IsGiveByFriend"为1 时填写该参数.
	UserCardCode    string         `xml:"UserCardCode"    json:"UserCardCode"`    // code 序列号. 自定义code 及非自定义code的卡券被领取后都支持事件推送.
	OldUserCardCode string         `xml:"OldUserCardCode" json:"OldUserCardCode"` // 转赠前的code序列号。
	OuterId         int64          `xml:"OuterId"         json:"OuterId"`         // 领取场景值, 用于领取渠道数据统计. 可在生成二维码接口及添加JS API 接口中自定义该字段的整型值.
}

func GetUserGetCardEvent(msg *core.MixedMsg) *UserGetCardEvent {
	return &UserGetCardEvent{
		MsgHeader:       msg.MsgHeader,
		EventType:       msg.EventType,
		CardId:          msg.CardId,
		IsGiveByFriend:  msg.IsGiveByFriend,
		FriendUserName:  msg.FriendUserName,
		UserCardCode:    msg.UserCardCode,
		OldUserCardCode: msg.OldUserCardCode,
		OuterId:         msg.OuterId,
	}
}

// 用户在删除卡券时, 微信会把这个事件推送到开发者填写的URL.
type UserDelCardEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader

	EventType    core.EventType `xml:"Event"        json:"Event"`        // 事件类型, user_del_card
	CardId       string         `xml:"CardId"       json:"CardId"`       // 卡券ID
	UserCardCode string         `xml:"UserCardCode" json:"UserCardCode"` // 商户自定义code 值. 非自定code 推送为空串
}

func GetUserDelCardEvent(msg *core.MixedMsg) *UserDelCardEvent {
	return &UserDelCardEvent{
		MsgHeader:    msg.MsgHeader,
		EventType:    msg.EventType,
		CardId:       msg.CardId,
		UserCardCode: msg.UserCardCode,
	}
}

// 卡券被核销时, 微信会把这个事件推送到开发者填写的URL
type UserConsumeCardEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader

	EventType     core.EventType `xml:"Event"        json:"Event"`          // 事件类型, user_consume_card
	CardId        string         `xml:"CardId"       json:"CardId"`         // 卡券ID
	UserCardCode  string         `xml:"UserCardCode" json:"UserCardCode"`   // 商户自定义code 值. 非自定code 推送为空串
	ConsumeSource string         `xml:"ConsumeSource" json:"ConsumeSource"` // 核销来源。支持开发者统计API核销（FROM_API）、公众平台核销（FROM_MP）、卡券商户助手核销（FROM_MOBILE_HELPER）（核销员微信号）
}

func GetUserConsumeCardEvent(msg *core.MixedMsg) *UserConsumeCardEvent {
	return &UserConsumeCardEvent{
		MsgHeader:     msg.MsgHeader,
		EventType:     msg.EventType,
		CardId:        msg.CardId,
		UserCardCode:  msg.UserCardCode,
		ConsumeSource: msg.ConsumeSource,
	}
}

// 用户在进入会员卡时, 微信会把这个事件推送到开发者填写的URL
type UserViewCardEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader

	EventType    core.EventType `xml:"Event"        json:"Event"`        // 事件类型, user_view_card
	CardId       string         `xml:"CardId"       json:"CardId"`       // 卡券ID
	UserCardCode string         `xml:"UserCardCode" json:"UserCardCode"` // 商户自定义code 值. 非自定code 推送为空串
}

func GetUserViewCardEvent(msg *core.MixedMsg) *UserViewCardEvent {
	return &UserViewCardEvent{
		MsgHeader:    msg.MsgHeader,
		EventType:    msg.EventType,
		CardId:       msg.CardId,
		UserCardCode: msg.UserCardCode,
	}
}

// 从卡券进入公众号会话事件推送
type UserEnterSessionFromCardEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader

	EventType    core.EventType `xml:"Event"        json:"Event"`        // 事件类型, user_view_card
	CardId       string         `xml:"CardId"       json:"CardId"`       // 卡券ID
	UserCardCode string         `xml:"UserCardCode" json:"UserCardCode"` // 商户自定义code 值. 非自定code 推送为空串
}

func GetUserEnterSessionFromCardEvent(msg *core.MixedMsg) *UserEnterSessionFromCardEvent {
	return &UserEnterSessionFromCardEvent{
		MsgHeader:    msg.MsgHeader,
		EventType:    msg.EventType,
		CardId:       msg.CardId,
		UserCardCode: msg.UserCardCode,
	}
}
