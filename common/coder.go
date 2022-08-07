package coder


type Coder struct{

}


/**
 * 生成指定业务对象的一个不重复的编码
 */
func (this *Coder) GenerateCode() string{

	code:=rand.Intn(100)

	fmt.Println(code)

	return string(code)
	
}
