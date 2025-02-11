package handler

import (
	"net/http"

	"github.com/kingwrcy/moments/db"
	"github.com/kingwrcy/moments/vo"
	"github.com/labstack/echo/v4"
	"github.com/samber/do/v2"
)

type NoticeHandler struct {
	base BaseHandler
	hc   http.Client
}

func NewNoticeHandler(injector do.Injector) *NoticeHandler {
	return &NoticeHandler{
		base: do.MustInvoke[BaseHandler](injector),
		hc:   http.Client{},
	}
}

type noticeListResp struct {
	List  []db.Notice `json:"list,omitempty"`  //通知列表
	Total int64       `json:"total,omitempty"` //总数
	Page  int         `json:"page,omitempty"`  //页码
	Size  int         `json:"size,omitempty"`  //页大小
}

func (n NoticeHandler) ListNotices(c echo.Context) error {
	var (
		req   vo.ListNoticeReq
		list  []db.Notice
		total int64
	)

	ctx := c.(CustomContext)
	currentUser := ctx.CurrentUser()

	err := c.Bind(&req)
	if err != nil {
		return FailResp(c, ParamError)
	}

	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Size <= 0 {
		req.Size = 10
	}

	tx := n.base.db.Model(&db.Notice{}).
		Where("userId = ?", currentUser.Id)

	if req.Type != 0 {
		tx = tx.Where("type = ?", req.Type)
	}

	if req.IsRead != nil {
		tx = tx.Where("isRead = ?", *req.IsRead)
	}

	tx.Count(&total)
	offset := (req.Page - 1) * req.Size
	tx.Order("createdAt desc").Limit(req.Size).Offset(offset)
	tx.Find(&list)

	return SuccessResp(c, noticeListResp{
		List:  list,
		Total: total,
		Page:  req.Page,
		Size:  req.Size,
	})
}

func (n NoticeHandler) ReadNotices(c echo.Context) error {
	var (
		req vo.ReadNoticeReq
	)

	ctx := c.(CustomContext)
	currentUser := ctx.CurrentUser()

	err := c.Bind(&req)
	if err != nil {
		return FailResp(c, ParamError)
	}

	if len(req.NoticeIds) == 0 {
		return FailResp(c, ParamError)
	}

	n.base.db.Model(&db.Notice{}).
		Where("userId = ?", currentUser.Id).
		Where("id IN ?", req.NoticeIds).
		Update("isRead", true)

	return SuccessResp(c, h{})
}

func (n NoticeHandler) RemoveNotices(c echo.Context) error {
	var (
		req     vo.RemoveNoticeReq
		notices []db.Notice
	)

	ctx := c.(CustomContext)
	currentUser := ctx.CurrentUser()
	err := c.Bind(&req)
	if err != nil {
		return FailResp(c, ParamError)
	}

	if n.base.db.Model(&db.Notice{}).
		Where("userId = ?", currentUser.Id).
		Where("id IN ?", req.NoticeIds).
		Delete(&notices).RowsAffected == 0 {
		return FailRespWithMsg(c, Fail, "删除失败")
	}

	return SuccessResp(c, h{})
}

func (n NoticeHandler) LatestNotice(c echo.Context) error {
	var notice db.Notice

	ctx := c.(CustomContext)
	currentUser := ctx.CurrentUser()

	n.base.db.Model(&db.Notice{}).
		Where("userId = ?", currentUser.Id).
		Order("createdAt desc").First(&notice)

	return SuccessResp(c, notice)
}
