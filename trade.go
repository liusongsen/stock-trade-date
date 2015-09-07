package trade

//股票交易日期(休市/开盘)计算

import (
	"time"
)

//区间
type Interval struct {
	BeginDate time.Time `开始日期`
	EndDate   time.Time `结束日期`
}

//区间列表
type IntervalSlice struct {
	Ivs []Interval
}

//计算
func club(inputTime time.Time) (rtime time.Time) {

	rtime = inputTime
	//法定休假列表
	var ivse IntervalSlice
	//星期六判断
	if inputTime.Weekday() == time.Saturday {
		rtime = inputTime.AddDate(0, 0, -1)
	}
	//星期天判断
	if inputTime.Weekday() == time.Sunday {
		rtime = inputTime.AddDate(0, 0, -2)
	}
	//一月 1.1 1.2
	ivse.Ivs = append(ivse.Ivs, Interval{BeginDate: time.Date(2015, time.January, 1, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2015, time.January, 2, 0, 0, 0, 0, time.UTC)})
	//二月 2.18 -2.20
	ivse.Ivs = append(ivse.Ivs, Interval{BeginDate: time.Date(2015, time.February, 18, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2015, time.February, 20, 0, 0, 0, 0, time.UTC)})
	//二月 2.23 -2.24
	ivse.Ivs = append(ivse.Ivs, Interval{BeginDate: time.Date(2015, time.February, 23, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2015, time.February, 24, 0, 0, 0, 0, time.UTC)})
	//四月 4.6
	ivse.Ivs = append(ivse.Ivs, Interval{BeginDate: time.Date(2015, time.April, 6, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2015, time.April, 6, 0, 0, 0, 0, time.UTC)})
	//五月 5.1
	ivse.Ivs = append(ivse.Ivs, Interval{BeginDate: time.Date(2015, time.May, 1, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2015, time.May, 1, 0, 0, 0, 0, time.UTC)})
	//六月 6.22
	ivse.Ivs = append(ivse.Ivs, Interval{BeginDate: time.Date(2015, time.June, 22, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2015, time.June, 22, 0, 0, 0, 0, time.UTC)})
	//九月 9.3-9.5
	ivse.Ivs = append(ivse.Ivs, Interval{BeginDate: time.Date(2015, time.September, 3, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2015, time.September, 5, 0, 0, 0, 0, time.UTC)})
	//十月 10.1-10.2
	ivse.Ivs = append(ivse.Ivs, Interval{BeginDate: time.Date(2015, time.October, 1, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2015, time.October, 2, 0, 0, 0, 0, time.UTC)})
	//十月 10.5-10.7
	ivse.Ivs = append(ivse.Ivs, Interval{BeginDate: time.Date(2015, time.October, 5, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2015, time.October, 7, 0, 0, 0, 0, time.UTC)})

	//迭代判断当前日期是否在法定日期范围内
	for _, v := range ivse.Ivs {
		if rtime.Before(v.EndDate.AddDate(0, 0, 1)) && rtime.After(v.BeginDate.AddDate(0, 0, -1)) {
			rtime = v.BeginDate.AddDate(0, 0, -1)
			//星期六/星期天判断
			if rtime.Weekday() == time.Saturday || rtime.Weekday() == time.Sunday {
				rtime = club(rtime)
			}
			break
		}
	}
	return
}

//检查今天股市是否开盘
func IsOpen() bool {

	ntime := time.Now()
	rtime := club(ntime)
	return rtime.Equal(ntime)
}

//计算T-1日期
func ClubT1() string {

	rtime := club(time.Now().AddDate(0, 0, -1))
	return rtime.Format("20060102")
}

//计算T-2日期
func ClubT2() string {

	rtime := club(time.Now().AddDate(0, 0, -1))
	rtime = club(rtime.AddDate(0, 0, -1))
	return rtime.Format("20060102")
}
