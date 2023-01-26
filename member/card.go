package member

import "time"

type Card struct{
		Id int
		Name string    //卡名称
		FaceValue float32  // 面额
		LimitedDate time   // 有效期：截止日期

}
