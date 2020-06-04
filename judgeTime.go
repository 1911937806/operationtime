package operationtime

import (
	"time"
)

//判断当前时间是否到了某个点
//t int ： 晚上 11点 = 》 22   |    凌晨12 点  =》 0
func  JudgeTime (h  int) bool{
	nh := time.Now().Hour()
	return  nh == h
}