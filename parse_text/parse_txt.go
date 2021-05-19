package parse_text

import (
	"go_demo_post/parse_text/rule"
	"os"
)

/*
已一定的规则解析一个文件，并返回结果
 */

type ParseTextService struct {
}

func NewParseText() *ParseTextService {
	return &ParseTextService{}
}


func (service *ParseTextService) Parse(filename string, ruleService rule.IRule) (rule.IResult, error) {
	file ,err:= os.Open(filename)
	if err!=nil{
		return nil,err
	}
	return ruleService.Parse(file)
}