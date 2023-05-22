package modelos

import (
	"fmt"
	"log"
)

type ListaSimple struct {
	Cabecera *Nodo
	Cantidad int
	Ultimo   *Nodo
	Mapa     map[int]*Nodo
}

func Lista(cantidadNodos int) *ListaSimple {

	primero := &Nodo{}

	l := &ListaSimple{
		Cabecera: primero,
		Cantidad: 1,
		Ultimo:   primero,
		Mapa:     make(map[int]*Nodo),
	}

	l.Mapa[0] = primero

	for i := 0; i < cantidadNodos-1; i++ {
		l.InsertarFinal(0)
	}

	return l
}

func ListaDesdeSlice(datos []int) *ListaSimple {
	l := &ListaSimple{Cabecera: nil, Cantidad: 0, Ultimo: nil, Mapa: make(map[int]*Nodo)}

	for _, valor := range datos {
		l.InsertarFinal(valor)
	}

	return l
}

func (l *ListaSimple) EstaVacia() bool {
	return l.Cantidad == 0
}

func (l *ListaSimple) GetCantidadNodos() int {
	return l.Cantidad
}

func (l *ListaSimple) EliminarLista() {
	l.Cabecera = nil
	l.Cantidad = 0
}

func (l *ListaSimple) InsertarFinal(valor int) {
	nuevo := &Nodo{Valor: valor, Siguiente: nil}

	if l.EstaVacia() {
		l.Cabecera = nuevo
		l.Ultimo = nuevo
		l.Cabecera.Siguiente = l.Ultimo
	} else {

		l.Ultimo.Siguiente = nuevo
		l.Ultimo = nuevo
	}

	l.Cantidad++

	l.Mapa[l.Cantidad-1] = nuevo
}

func (l *ListaSimple) InsertarFinalRecursivo(nodo *Nodo, valor int, posicion int) {
	if nodo == nil && l.Cantidad == posicion {
		nuevo := &Nodo{Valor: valor, Siguiente: nil}
		l.Ultimo.Siguiente = nuevo
		l.Ultimo = nuevo
		l.Cantidad++
		return
	}

	l.InsertarFinalRecursivo(nodo.Siguiente, valor, posicion+1)
}

func (l *ListaSimple) InsertarInicio(valor int) {

	nuevo := &Nodo{Valor: valor, Siguiente: nil}

	if l.EstaVacia() {
		l.Cabecera = nuevo
		l.Ultimo = nuevo
		l.Cabecera.Siguiente = l.Ultimo

	} else {
		nuevo.Siguiente = l.Cabecera

		l.Cabecera = nuevo
	}

	l.Cantidad++

}

func (l *ListaSimple) InsertarPosicion(valor int, posicion int) {
	nuevo := &Nodo{Valor: valor, Siguiente: nil}

	i := 0
	aux := l.Cabecera

	for i != posicion-1 {
		aux = aux.Siguiente
		i++
	}

	nuevo.Siguiente = aux.Siguiente
	aux.Siguiente = nuevo

	l.Cantidad++
}

func (l *ListaSimple) InsertarPosicionRecursivo(nodo *Nodo, valor int, posicion int, i int) {
	if i == posicion {
		nuevo := &Nodo{Valor: valor, Siguiente: nil}
		nuevo.Siguiente = nodo.Siguiente
		nodo.Siguiente = nuevo
		l.Cantidad++
		return
	}

	l.InsertarPosicionRecursivo(nodo.Siguiente, valor, posicion, i+1)
}

func (l *ListaSimple) GetValor(posicion int) int {

	nodo := l.GetNodo(posicion)

	if nodo == nil {
		log.Fatal("Se trató de acceder a una posición que no existe")
	}

	return nodo.Valor
}

func (l *ListaSimple) SetPosicion(posicion int, valor int) {
	i := 0
	aux := l.Cabecera

	for i != posicion {
		aux = aux.Siguiente
		i++
	}

	aux.Valor = valor
}

func (l *ListaSimple) GetNodo(posicion int) *Nodo {
	// Verificar si la posición es válida
	if posicion < 0 || posicion >= l.Cantidad {
		return nil
	}
	// Retornar el nodo del mapa en tiempo constante
	return l.Mapa[posicion]
}

func (l *ListaSimple) EliminarInicio() {
	l.Cabecera = l.Cabecera.Siguiente
	l.Cantidad--

	delete(l.Mapa, 0) // Eliminar el primer nodo del mapa
	// Actualizar los índices de los nodos restantes en el mapa
	for i := 0; i < l.Cantidad; i++ {
		l.Mapa[i] = l.Mapa[i+1]
	}
}

func (l *ListaSimple) EliminarPosicion(posicion int) {
	i := 0
	aux := l.Cabecera

	for i != posicion-1 {
		aux = aux.Siguiente
		i++
	}

	aux.Siguiente = aux.Siguiente.Siguiente
	l.Cantidad--
}

func (l *ListaSimple) EliminarPosicionRecursivo(nodo *Nodo, posicion int, i int) {
	if i == posicion-1 {
		nodo.Siguiente = nodo.Siguiente.Siguiente
		l.Cantidad--
		return
	}

	l.EliminarPosicionRecursivo(nodo.Siguiente, posicion, i+1)
}

func (l *ListaSimple) GetPosicion(referencia *Nodo) int {
	i := 0
	aux := l.Cabecera

	for i < l.Cantidad {

		if aux == referencia {
			return i
		}

		aux = aux.Siguiente
		i++
	}
	return -1
}

func (l *ListaSimple) Buscar(referencia *Nodo) bool {

	i := 0
	aux := l.Cabecera

	for i < l.Cantidad {

		if aux == referencia {
			return true
		}

		aux = aux.Siguiente
		i++
	}

	return false
}

func (l *ListaSimple) BuscarRecursivo(nodo *Nodo, referencia *Nodo, n int) bool {

	if nodo == referencia {
		return true
	}

	if n == l.Cantidad {
		return false
	}

	return l.BuscarRecursivo(nodo.Siguiente, referencia, n+1)

}

func (l *ListaSimple) Imprimir() {
	i := 0
	aux := l.Cabecera

	for i != l.Cantidad {
		fmt.Println(aux.Valor)
		aux = aux.Siguiente
		i++
	}
}
