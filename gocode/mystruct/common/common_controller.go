package common

type CommInterface interface {
	GetConfig() string
}

type Controller struct {
	CommInterface
}

func (c *Controller) GetName() string {

	panic("")

	return "CommonController"
}
