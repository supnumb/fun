package member


var (
        engine *xorm.Engine
)

func init() {
        var err error
        engine, err = xorm.NewEngine("mysql", "root:www.loclive.cn:1160@/mbox?charset=utf8")
}

// 保存会员到数据库
// 如果会员有会员ID，会员已经存在，则更新该ID会员
// 如果会员不存在，则在数据库增加新会员
func saveMember(m *Member){

        if m.Id>0{
                _,err=engine.ID(m.Id).Update(m)

                if err!=nil{
                        log.Fatal("update member error",err)
                }
        }else{

                saas := c.Request.Header["Saas-Code"]
                coder := new(common.Coder)

                if saas != nil {
                        tm := strings.Split(saas[0], "-")
                        if len(tm) == 2 {
                                coder.SetTenant(tm[0])
                                coder.SetMerchant(tm[1])
                        }
                }
        }
        engine.Sync(m)



}
