stock-trade-date
================

获取股票交易日期（交易日期判断/T-1交易日期/T-2交易日期）

股票交易日期特别说明:

星期六/星期天股市不开市
国家法定假期股市不开市
有效日期就是把上述两种情况都除去后的有效日期


T-1交易日期说明:
当前股票交易日期的前一天有效日期

T-2交易日期说明:
当前股票交易日期的前两天有效日期

2013年法定假期参考标准：
国务院办公厅关于2013年部分节假日安排的通知
http://www.gov.cn/zwgk/2012-12/10/content_2286598.htm

安装：

go get github.com/liusongsen/stock-trade-date


用法:

import github.com/liusongsen/stock-trade-date

股市今天是否开盘
isopen := trade.IsOpen()

求T-1日期:
time1 := trade.ClubT1()

求T-2日期：
time2 := trade.ClubT2()



[ `go test` | done: 231.117ms ]
	PASS
	ok  	github.com/liusongsen/stock-trade-date	0.003s

