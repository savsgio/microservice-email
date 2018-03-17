package utils

type Json map[string]interface{}

func CheckException(err error) {
	if err != nil {
		panic(err)
	}
}
