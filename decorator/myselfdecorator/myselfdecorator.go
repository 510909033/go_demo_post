package myselfdecorator

import "path/filepath"

/*
导入第三方数据 并创建 账号， 文章
*/

type ThirdDataParam struct {
}
type ThirdDataResult struct {
}

type CreateAccountParam interface {
}
type CreateAccountResult interface {
}

type Handler func(param *ThirdDataParam) (result *ThirdDataResult, err error)

type (
	//模板方法
	Template interface {
		//获取第三方数据
		GetThirdData(param *ThirdDataParam) (result *ThirdDataResult, err error)
		//创建第三方账号
		CreateAccount(param CreateAccountParam) (result CreateAccountResult)
		//创建文章
		CreateArticle() error
	}
)

type Demo struct {
	filepath.WalkFunc
}

func (d *Demo) TimeStats(fn interface{}) (result *ThirdDataResult, err error) {
	return nil, nil
	//return fn(nil)
}

func (d *Demo) TimeStats1(fn Handler) (result *ThirdDataResult, err error) {
	return fn(nil)
}

type CommonDemo struct {
	demo         *Demo
	templateDemo *TemplateDemo
}

func (c *CommonDemo) GetThirdData(endpoints Handler) Handler {
	return func(param *ThirdDataParam) (result *ThirdDataResult, err error) {

		return endpoints(param)
	}

}

func (c *CommonDemo) CreateAccount(param CreateAccountParam) (result CreateAccountResult) {
	panic("implement me")
}

func (c *CommonDemo) CreateArticle() error {
	panic("implement me")
}

func demo() {
	//templateDemo := &TemplateDemo{}
	//demo := &Demo{}
	//demo.TimeStats(templateDemo.Handler)

	var f1 Handler

	commonDemo := &CommonDemo{}
	commonDemo.GetThirdData(f1)
}

type TemplateDemo struct {
}

func (t *TemplateDemo) GetThirdData(param *ThirdDataParam) (result *ThirdDataResult, err error) {

	panic("implement me")
}

func (t *TemplateDemo) CreateAccount(param CreateAccountParam) (result CreateAccountResult) {
	panic("implement me")
}

func (t *TemplateDemo) CreateArticle() error {
	panic("implement me")
}

type S1 struct {
	handler Handler
}
type S2 struct {
	handler Handler
}
type S3 struct {
	handler Handler
}
