package activity

import (
	context2 "context"
	"fmt"
	"github.com/silenceper/wechat/v2/miniprogram/context"
	"github.com/silenceper/wechat/v2/util"
)

const (
	createUrl = "https://api.weixin.qq.com/cgi-bin/message/wxopen/activityid/create?access_token=%s&unionid=%s&openid=%s"
)

type Activity struct {
	*context.Context
}

type RspCreate struct {
	util.CommonError
	ActivityId     string `json:"activity_id"`     // 动态消息的 ID
	ExpirationTime uint64 `json:"expiration_time"` // activity_id 的过期时间戳。默认24小时后过期。
}

func NewActivity(ctx *context.Context) *Activity {
	return &Activity{ctx}
}

func (activity *Activity) Create(unionid, openid string) (RspCreate, error) {
	return activity.CreateContext(context2.Background(), unionid, openid)
}

func (activity *Activity) CreateContext(ctx context2.Context, unionid, openid string) (result RspCreate, err error) {
	var at string
	if at, err = activity.GetAccessToken(); err != nil {
		return
	}

	var response []byte
	if response, err = util.HTTPGetContext(ctx, fmt.Sprintf(createUrl, at, unionid, openid)); err != nil {
		return
	}

	if err = util.DecodeWithError(response, &result, ""); err != nil {
		return
	}

	if result.ErrCode != 0 {
		err = fmt.Errorf("Code2Session error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return
	}

	return
}
