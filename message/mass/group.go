package mass

type commonGroupMsgHead struct {
	Filter struct {
		GroupId string `json:"group_id"`
	} `json:"filter"`
	MsgType string `json:"msgtype"`
}

// news ========================================================================

type GroupNewsMsg struct {
	commonGroupMsgHead

	News struct {
		MediaId string `json:"media_id"`
	} `json:"mpnews"`
}

func NewGroupNewsMsg(groupId, mediaId string) *GroupNewsMsg {
	var msg GroupNewsMsg
	msg.Filter.GroupId = groupId
	msg.MsgType = GROUP_MSG_TYPE_NEWS
	msg.News.MediaId = mediaId

	return &msg
}

// text ========================================================================

type GroupTextMsg struct {
	commonGroupMsgHead

	Text struct {
		Content string `json:"content"`
	} `json:"text"`
}

func NewGroupTextMsg(groupId, content string) *GroupTextMsg {
	var msg GroupTextMsg
	msg.Filter.GroupId = groupId
	msg.MsgType = GROUP_MSG_TYPE_TEXT
	msg.Text.Content = content

	return &msg
}

// voice =======================================================================

type GroupVoiceMsg struct {
	commonGroupMsgHead

	Voice struct {
		MediaId string `json:"media_id"`
	} `json:"voice"`
}

func NewGroupVoiceMsg(groupId, mediaId string) *GroupVoiceMsg {
	var msg GroupVoiceMsg
	msg.Filter.GroupId = groupId
	msg.MsgType = GROUP_MSG_TYPE_VOICE
	msg.Voice.MediaId = mediaId

	return &msg
}

// image =======================================================================

type GroupImageMsg struct {
	commonGroupMsgHead

	Image struct {
		MediaId string `json:"media_id"`
	} `json:"image"`
}

func NewGroupImageMsg(groupId, mediaId string) *GroupImageMsg {
	var msg GroupImageMsg
	msg.Filter.GroupId = groupId
	msg.MsgType = GROUP_MSG_TYPE_IMAGE
	msg.Image.MediaId = mediaId

	return &msg
}

// video =======================================================================

type GroupVideoMsg struct {
	commonGroupMsgHead

	Video struct {
		MediaId string `json:"media_id"`
	} `json:"mpvideo"`
}

func NewGroupVideoMsg(groupId, mediaId string) *GroupVideoMsg {
	var msg GroupVideoMsg
	msg.Filter.GroupId = groupId
	msg.MsgType = GROUP_MSG_TYPE_VIDEO
	msg.Video.MediaId = mediaId

	return &msg
}