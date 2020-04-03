# time layout
时间模板
==

<pre>
YmdHis := chinese.Now(false, false, false)
time, _ := chinese.Parse("2006年01月02日 15时04分05秒")
fmt.Println(YmdHis, time)

YmdHis2 := dash.Now(false, false, false)
time2, _ := dash.Parse("2006-01-02 15:04:05")
fmt.Println(YmdHis2, time2)

YmdHis3 := slash.Now(false, false, false)
time3, _ := slash.Parse("2006/01/02 15:04:05")
fmt.Println(YmdHis3, time3)

YmdHis4 := nothing.Now(false)
time4, _ := nothing.Parse("20060102150405")
fmt.Println(YmdHis4, time4)
</pre>