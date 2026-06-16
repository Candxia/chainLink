package cmd

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gvalid"
	"strings"
)

func init() {
	gvalid.RegisterRule("cPwd", RulePwd)
	gvalid.RegisterRule("cName", RuleAdmin)
	gvalid.RegisterRule("mustDate", MustDate)
	gvalid.RegisterRule("limitDate", LimitDate)
	gvalid.RegisterRule("limitDateTime", LimitDateTime)
	gvalid.RegisterRule("lenRid", LenRid)
}

// RulePwd 密码规则校验
// 密码长度 6-18 位，必须包含字母和数字
func RulePwd(ctx context.Context, in gvalid.RuleFuncInput) error {
	value := in.Value.String()
	if len(value) < 6 || len(value) > 18 {
		return gerror.New(in.Message)
	}
	hasLetter := false
	hasDigit := false
	for _, c := range value {
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
			hasLetter = true
		}
		if c >= '0' && c <= '9' {
			hasDigit = true
		}
	}
	if !hasLetter || !hasDigit {
		return gerror.New(in.Message)
	}
	return nil
}

// RuleAdmin 管理员用户名规则校验
// 以字母开头，包含字母/数字/下划线，长度 6-32
func RuleAdmin(ctx context.Context, in gvalid.RuleFuncInput) error {
	value := in.Value.String()
	if gregex.IsMatchString(`^[a-zA-Z][\w]{5,31}$`, value) {
		return nil
	}
	return gerror.New(in.Message)
}

// MustDate 必填日期校验，日期不能早于系统设定的截止日期
func MustDate(ctx context.Context, in gvalid.RuleFuncInput) error {
	value := in.Value.String()
	if value == "" {
		return gerror.New(in.Message)
	}
	cur := gtime.NewFromStrFormat(value, "Y-m-d")
	if cur == nil {
		return gerror.New("日期格式无效，请使用 Y-m-d 格式")
	}
	// 默认不允许早于当天
	now := gtime.Now().Format("Ymd")
	if gconv.Int(cur.Format("Ymd")) < gconv.Int(now) {
		return gerror.New(in.Message)
	}
	return nil
}

// LimitDate 日期范围校验（精确到日，格式 Y-m-d）
func LimitDate(ctx context.Context, in gvalid.RuleFuncInput) error {
	return LimitDateFormat(ctx, in, "Y-m-d")
}

// LimitDateTime 日期时间范围校验（精确到秒，格式 Y-m-d H:i:s）
func LimitDateTime(ctx context.Context, in gvalid.RuleFuncInput) error {
	return LimitDateFormat(ctx, in, "Y-m-d H:i:s")
}

// LimitDateFormat 日期范围校验核心逻辑
// 规则格式: "limitDate:startField,offset,unit"
// 例如: "limitDate:start_time,30,d" 表示结束时间距离开始时间不超过30天
func LimitDateFormat(ctx context.Context, in gvalid.RuleFuncInput, format string) error {
	parts := strings.Split(strings.Split(in.Rule, ":")[1], ",")
	if len(parts) < 3 {
		return gerror.New("日期范围规则格式错误")
	}
	var msg string
	switch parts[2] {
	case "y":
		msg = parts[1] + "年"
	case "m":
		msg = parts[1] + "个月"
	case "d":
		msg = parts[1] + "天"
	}

	value := in.Value.String()
	if value == "" {
		return gerror.Newf(in.Message, msg)
	}

	cur := gtime.NewFromStrFormat(value, format)
	if cur == nil {
		return gerror.New("日期格式无效")
	}

	// 获取起始日期字段
	data, ok := in.Data.Val().(map[string]interface{})
	if !ok {
		return gerror.New("数据格式无效")
	}
	startStr, ok := data[parts[0]].(string)
	if !ok || startStr == "" {
		return gerror.Newf(in.Message, msg)
	}
	st := gtime.NewFromStrFormat(startStr, format)
	if st == nil {
		return gerror.New("起始日期格式无效")
	}

	if st.After(cur) {
		return gerror.Newf(in.Message, msg)
	}

	offset := gconv.Int(parts[1])
	switch parts[2] {
	case "y":
		if st.AddDate(offset, 0, 0).Before(cur) {
			return gerror.Newf(in.Message, msg)
		}
	case "m":
		if st.AddDate(0, offset, 0).Before(cur) {
			return gerror.Newf(in.Message, msg)
		}
	case "d":
		if st.AddDate(0, 0, offset).Before(cur) {
			return gerror.Newf(in.Message, msg)
		}
	}
	return nil
}

// LenRid 数组长度限制校验
// 规则格式: "lenRid:maxCount"，限制数组元素个数不超过 maxCount
func LenRid(ctx context.Context, in gvalid.RuleFuncInput) error {
	parts := strings.Split(in.Rule, ":")
	if len(parts) != 2 {
		return gerror.New(in.Rule)
	}
	maxNum := gconv.Int(parts[1])
	value := in.Value.Array()
	if len(value) < maxNum {
		return nil
	}
	return gerror.Newf(in.Message, maxNum)
}
