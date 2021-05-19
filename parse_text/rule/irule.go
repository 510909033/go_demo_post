package rule

import (
	"os"
)

type IResult interface {
	GetError() []interface{} //错误信息列表
	GetList() interface{}    //处理好的列表 需业务侧断言
}

type IRule interface {
	//解析文件， 如果err!=nil，说明有失败发生
	Parse(f *os.File) (IResult, error)
}
