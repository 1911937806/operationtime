package operationtime

import (
	"fmt"
	"testing"
)

func TestJudgeTime(t *testing.T) {
	if  JudgeTime(22)    {
		fmt.Println("到了22点")
	}else {
		fmt.Println("没到22点")
	}
  }