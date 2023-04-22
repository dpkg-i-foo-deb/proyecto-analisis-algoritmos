package utilidades

import "os"

func EscribirArchivo(nombre string, contenido []byte) {

	err := os.WriteFile(nombre, contenido, 0644)

	VerificarError(err)
}
