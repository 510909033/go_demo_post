package interf

/*
一个缓存服务，定义了Set Get两个通用方法

当它走起来像鸭子，叫起来也像鸭子，那它就是鸭子。





 */

type IKey interface {

}

type IResult interface {

}


type ICacheClient interface {
	Get(key IKey) (IResult, error)
	Set(key IKey, result IResult) (bool, error)
}
