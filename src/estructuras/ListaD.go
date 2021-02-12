package estructuras

import (
	"fmt"

	"../mensajes"
)

type Nodo struct {
	Siguiente  *Nodo
	Anterior   *Nodo
	Mensaje    *mensajes.Mensaje
}

type ListaD struct {
	Inicio *Nodo
	Ultimo *Nodo
	Tamano int
}

//Insertar datos en la lista
func Insertar(nuevoMensaje *mensajes.Mensaje, lista *ListaD){
	nuevoNodo := Nodo{Mensaje: nuevoMensaje}

	if lista.Tamano == 0{
		nuevoNodo.Siguiente = nil
		nuevoNodo.Anterior = nil
		lista.Inicio = &nuevoNodo
		lista.Ultimo = &nuevoNodo
		lista.Tamano++
	}else{
		aux := lista.Ultimo
		aux.Siguiente = &nuevoNodo
		nuevoNodo.Anterior = aux
		nuevoNodo.Siguiente = nil
		lista.Ultimo = &nuevoNodo
		lista.Tamano++
	}

}

//Para IMPRIMIR la lista
func mostrar(lista *ListaD){
	aux := lista.Inicio
	for aux != nil {
		fmt.Println(aux.Mensaje)
		aux = aux.Siguiente
	}
}

//obtengo el tamaño de la lista
func ObtenerTamanio(lista *ListaD) int{
	return lista.Tamano
}
