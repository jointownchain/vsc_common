package utils

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
	"strings"
	"time"
)

const FormatDateTime = "2006-01-02 15:04:05"
const FormatDate = "2006-01-02"
const FormatDateNoTime = "2006-01-02 00:00:00"
const FormatDateTime2 = "2006/01/02 15:04:05"

func ConvertLocalDateStr(ct *time.Time) string {
	if ct.IsZero() {
		return ""
	}
	return ct.Local().Format(FormatDateTime)
}

//str转time
func ParseTimeByTimeStr(str, Prefix string) (time.Time, error) {
	p := strings.TrimSpace(str)
	if p == "" {
		return time.Time{}, errors.Errorf("%s不能为空", str)
	}

	t, err := time.ParseInLocation(Prefix, str, time.Local)
	if err != nil {
		return time.Time{}, errors.Errorf("%s格式错误", Prefix)
	}
	return t, nil
}

func AddOneDay(date time.Time) time.Time {
	duration, _ := time.ParseDuration("24h")
	return date.Add(duration)
}
func AddAnyDay(date time.Time, dayNum int) time.Time {
	i := 24 * dayNum
	hourStr := strconv.Itoa(i)
	duration, _ := time.ParseDuration(hourStr + "h")
	return date.Add(duration)
}

func TimeToString(time time.Time, format string) string {
	return time.Format(format)
}

func AddHours(date time.Time, hour int64) time.Time {
	duration, _ := time.ParseDuration(strconv.FormatInt(hour, 10) + "h")
	return date.Add(duration)
}

//获取当前时间当前月第一天日期
func GetMonthFirstDay(now time.Time) time.Time {
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()
	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	return firstOfMonth
}

//获取当前日期不含时间time
func GetDayNoTime(now time.Time) time.Time {
	if now.IsZero() {
		return time.Time{}
	}
	format := now.Format(FormatDateNoTime)
	parse, _ := time.Parse(FormatDateNoTime, format)
	return parse
}

func ParseTimeStampToTime(timeStamp int64) time.Time {
	numberLength := len(strconv.FormatInt(timeStamp, 10))
	if numberLength >= 10 {
		timeStamp = timeStamp / Power(10, numberLength-10)
		return time.Unix(timeStamp, 0)
	} else {
		timeStamp = timeStamp * int64((10-numberLength)*10)
		return time.Unix(timeStamp, 0)
	}
}

func Convert2ProtoTime(time time.Time) *timestamppb.Timestamp {
	if time.IsZero() {
		return nil
	}
	reT, _ := ptypes.TimestampProto(time)
	return reT
}
func ConvertPointer2ProtoTime(time *time.Time) *timestamppb.Timestamp {
	if time == nil || time.IsZero() {
		return nil
	}
	reT, _ := ptypes.TimestampProto(*time)
	return reT
}

func ConvertProtoTime2Time(time timestamppb.Timestamp) *time.Time {
	if !time.IsValid() {
		return nil
	}
	timeGo, _ := ptypes.Timestamp(&time)
	return &timeGo
}
func ConvertProtoTimePointer2Time(time *timestamppb.Timestamp) *time.Time {
	if time != nil && !time.IsValid() {
		return nil
	}
	timeGo, _ := ptypes.Timestamp(time)
	return &timeGo
}

func ConvertProtoTime2TimeStr(timeS *timestamppb.Timestamp) string {
	if timeS == nil {
		return ""
	}
	t := time.Unix(timeS.Seconds, int64(timeS.Nanos))
	return t.Format(FormatDateTime)
}

func ConvertProtoTime2DateStr(timeS *timestamppb.Timestamp) string {
	if timeS == nil {
		return ""
	}
	t := time.Unix(timeS.Seconds, int64(timeS.Nanos))
	return t.Format(FormatDate)
}
