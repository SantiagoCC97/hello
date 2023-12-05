package ejercicios

import (
	"strconv"
)

func ConviertoaNumero(texto string) (int, string) {

	numerito, err := strconv.Atoi(texto)
	if err != nil {
		return 0, "Hubo un error" + err.Error()
	}
	if numerito >= 100 {
		return numerito, "Es mayor a 100"
	} else {
		return numerito, "Es menor a 100"
	}

}
