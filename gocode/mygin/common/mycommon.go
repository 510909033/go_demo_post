package common

func Exception(err error) {
	if err != nil {
		panic(err)
	}
}

func EnterSpan() {

}
