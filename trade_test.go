package trade

import (
	"testing"
	"time"
)

//周六判断
func TestClub(t *testing.T) {

	//2013-06-01 该天是星期六
	ntime := time.Date(2013, time.June, 1, 0, 0, 0, 0, time.UTC)
	rtime := club(ntime)
	//开始断言
	if rtime.Format("20060102") != "20130531" {
		t.Log("日期计算错误" + rtime.Format("20060102"))
		t.Fail()
	}
}
