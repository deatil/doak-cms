package time

import (
    "github.com/uniplaces/carbon"

    "github.com/deatil/doak-cms/pkg/config"
)

// 常量
const (
    DefaultFormat       = "2006-01-02 15:04:05"
    DateFormat          = "2006-01-02"
    FormattedDateFormat = "Jan 2, 2006"
    TimeFormat          = "15:04:05"
    HourMinuteFormat    = "15:04"
    HourFormat          = "15"
    DayDateTimeFormat   = "Mon, Jan 2, 2006 3:04 PM"
    CookieFormat        = "Monday, 02-Jan-2006 15:04:05 MST"
    RFC822Format        = "Mon, 02 Jan 06 15:04:05 -0700"
    RFC1036Format       = "Mon, 02 Jan 06 15:04:05 -0700"
    RFC2822Format       = "Mon, 02 Jan 2006 15:04:05 -0700"
    RFC3339Format       = "2006-01-02T15:04:05-07:00"
    RSSFormat           = "Mon, 02 Jan 2006 15:04:05 -0700"
)

// 现在
func Now() *carbon.Carbon {
    location := config.Section("time").Key("loc").MustString("")

    date, _ := carbon.NowInLocation(location)

    return date
}

// 时间戳
// "UTC" / "Asia/Shanghai"
func CreateFromTimestamp(timestamp int64) *carbon.Carbon {
    location := config.Section("time").Key("loc").MustString("")

    date, _ := carbon.CreateFromTimestamp(timestamp, location)

    return date
}

// 时间格式
func CreateFromFormat(layout, value string) *carbon.Carbon {
    location := config.Section("time").Key("loc").MustString("")

    date, _ := carbon.CreateFromFormat(layout, value, location)

    return date
}

// 解析
func Parse(layout, value string) *carbon.Carbon {
    location := config.Section("time").Key("loc").MustString("")

    date, _ := carbon.Parse(layout, value, location)

    return date
}

