package mmpaymkttransfers

import (
	"github.com/masiqi/wechat.v2/mch/core"
)

// 查询代金券批次信息.
func QueryCouponStock(clt *core.Client, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML("https://api.mch.weixin.qq.com/mmpaymkttransfers/query_coupon_stock", req)
}
