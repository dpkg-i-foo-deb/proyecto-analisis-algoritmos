package utilidades

import "log"

func VerificarError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
