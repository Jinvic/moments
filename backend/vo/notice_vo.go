package vo

const (
	NoticeTypeLike    int32 = iota + 1 // 收到点赞
	NoticeTypeComment                  // 收到评论
	NoticeTypeReply                    // 收到回复
)

type ListNoticeReq struct {
	Page   int   `json:"page,omitempty"`   //页码,从1开始
	Size   int   `json:"size,omitempty"`   //页大小,默认10
	Type   int32 `json:"type,omitempty"`   //通知类型,0:全部1:收到点赞2:收到评论3:收到回复
	IsRead *bool `json:"isRead,omitempty"` //是否已读,false:未读true:已读,不传默认全部
}

type ReadNoticeReq struct {
	NoticeIds []int `json:"noticeIds,omitempty"` //通知id
}

type RemoveNoticeReq struct {
	NoticeIds []int `json:"noticeIds,omitempty"` //通知id
}
