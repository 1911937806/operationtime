package operationtime

import "time"

//转化的时间类型
const (
	timeStringType1 = "2006-01-02 15:04:05"
	timeStringType2 = "2006/01/02 15:04:05"
	timeStringType3 = "2006-01-02"
	timeStringType4 = "15:04:05"
)
//"2006-01-02 15:04:05"
func ParseTimeToString1 (times time.Time) string {
	dayString := times.Format(timeStringType1)
	return  dayString
}
// "2006/01/02 15:04:05"
func ParseTimeToString2 (times time.Time) string {
	dayString := times.Format(timeStringType2)
	return  dayString
}
//"2006-01-02"
func ParseTimeToString3 (times time.Time) string {
	dayString := times.Format(timeStringType3)
	return  dayString
}
//"15:04:05"
func ParseTimeToString4 (times time.Time)string {
	dayString := times.Format(timeStringType4)
	return  dayString
}
