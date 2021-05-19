package common

import "fmt"

type StructCommon struct {

}

func (s *StructCommon) String() string {
	return fmt.Sprintf("StructCommon的输出， %%+v=%#v\n", *s)

}
