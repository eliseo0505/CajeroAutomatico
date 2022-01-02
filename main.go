package main

import (
	"fmt"
)

// Denominación y cantidad respectiva
var disponible = map[int]int{
	100: 3,
	50:  6,
	20:  10,
	10:  50,
	1:   50,
}

// Monto Total Disponible
var total = disponible[100]*100 + disponible[50]*50 + disponible[20]*20 + disponible[10]*10 + disponible[1]*1

// Monto a Retirar
var valorARetirar = 1254

func main() {

	fmt.Println("Total Disponible: ", total)
	fmt.Println("Monto solicitado: ", valorARetirar)

	comprobarMonto()

}

func comprobarMonto() {

	if valorARetirar <= total {
		entregarCien()
	} else {
		fmt.Println("Disculpe, no tenemos el monto solicitado.")
		fmt.Println("Por favor, intente con un monto menor.")
		fmt.Println("¡Muchas Gracias!")
	}
}
func entregarCien() int {
	restoDeCien := valorARetirar                   // Se le asigna a una nueva variable el monto a retirar
	cantidadDeBilletes := restoDeCien / 100        // Se calcula la cantidad de billetes necesarios para cubrir el monto solicitado
	montoDisponibleDeCien := disponible[100] * 100 // Se calcula el monto disponible de esta denominación
	fmt.Println("restoDeCien: ", restoDeCien)
	if restoDeCien >= 100 {
		if cantidadDeBilletes >= disponible[100] {
			fmt.Println("Billtes de 100: ", disponible[100])
			//restoDeCien = valorARetirar
			restoDeCien := restoDeCien - montoDisponibleDeCien
			return entregarCincuenta(restoDeCien)
		}

		if cantidadDeBilletes <= disponible[100] {
			fmt.Println("Billetes de 100: ", cantidadDeBilletes)
			//restoDeCien = valorARetirar
			restoDeCien := restoDeCien - cantidadDeBilletes*100
			return entregarCincuenta(restoDeCien)
		}
	} else {
		return entregarCincuenta(restoDeCien)
	}
	return entregarCincuenta(restoDeCien)
}

func entregarCincuenta(restoDeCien int) int {
	restoDeCincuenta := restoDeCien
	cantidadDeBilletes := restoDeCincuenta / 50
	montoDisponibleDeCincuenta := disponible[50] * 50
	if restoDeCincuenta >= 50 {
		if cantidadDeBilletes >= disponible[50] {
			fmt.Println("Billtes de 50: ", disponible[50])
			restoDeCincuenta := restoDeCincuenta - montoDisponibleDeCincuenta
			return entregarVeinte(restoDeCincuenta)
		}

		if cantidadDeBilletes <= disponible[50] {
			fmt.Println("Billetes de 50: ", cantidadDeBilletes)
			restoDeCincuenta := restoDeCincuenta - cantidadDeBilletes*50
			return entregarVeinte(restoDeCincuenta)
		}
	} else {
		return entregarVeinte(restoDeCincuenta)
	}
	return entregarVeinte(restoDeCincuenta)

}

func entregarVeinte(restoDeCincuenta int) int {
	restoDeVeinte := restoDeCincuenta
	cantidadDeBilletes := restoDeVeinte / 20
	montoDisponibleDeVeinte := disponible[20] * 20

	if restoDeVeinte >= 20 {

		if cantidadDeBilletes >= disponible[20] {
			fmt.Println("Billtes de 20: ", disponible[20])
			restoDeVeinte := restoDeVeinte - montoDisponibleDeVeinte
			return entregarDiez(restoDeVeinte)
		}

		if cantidadDeBilletes <= disponible[20] {
			fmt.Println("Billetes de 20: ", cantidadDeBilletes)
			restoDeVeinte := restoDeVeinte - cantidadDeBilletes*20
			return entregarDiez(restoDeVeinte)
		}

	} else {
		return entregarDiez(restoDeVeinte)
	}
	return entregarDiez(restoDeVeinte)
}

func entregarDiez(restoDeVeinte int) int {
	restoDeDiez := restoDeVeinte
	cantidadDeBilletes := restoDeDiez / 10
	montoDisponibleDeDiez := disponible[10] * 10

	if restoDeDiez >= 10 {

		if cantidadDeBilletes >= disponible[10] {
			fmt.Println("Billtes de 10: ", disponible[10])
			restoDeDiez := restoDeDiez - montoDisponibleDeDiez
			return entregarUno(restoDeDiez)
		}

		if cantidadDeBilletes <= disponible[10] {
			fmt.Println("Billetes de 10: ", cantidadDeBilletes)
			restoDeDiez := restoDeDiez - cantidadDeBilletes*10
			return entregarUno(restoDeDiez)
		}

	} else {
		return entregarUno(restoDeDiez)
	}
	return entregarUno(restoDeDiez)

}
func entregarUno(restoDeDiez int) int {
	restoDeUno := restoDeDiez
	cantidadDeBilletes := restoDeUno / 1
	montoDisponibleDeUno := disponible[1] * 1

	if restoDeUno >= 1 {
		if cantidadDeBilletes >= disponible[1] {
			fmt.Println("Billtes de 1: ", disponible[1])
			restoDeUno := restoDeUno - montoDisponibleDeUno
			return restoDeUno
		}

		if cantidadDeBilletes <= disponible[1] {
			fmt.Println("Billetes de 1: ", cantidadDeBilletes)
			restoDeUno := restoDeUno - cantidadDeBilletes*1
			return restoDeUno
		}

	} else {
		return end(0)
	}
	return restoDeUno

}
func end(int) int {
	return 0
}
