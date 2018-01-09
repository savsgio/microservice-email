package utils

func CheckException(err error) {
	if err != nil {
		panic(err)
	}
}
