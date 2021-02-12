package mensajes

import (
	"fmt"
)

type NodoC struct{
	Siguiente *NodoC
	Anterior *NodoC
	Contenido *Contenido
}

type ListaMensajes struct {
	Inicio *NodoC
	Ultimo *NodoC
	Tamano int
}

//Insertar datos en la lista
func InsertarC(nuevoC *Contenido, lista *ListaMensajes){
	nuevoNodo := NodoC{Contenido: nuevoC}

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
func mostrarM(lista *ListaMensajes){
	aux := lista.Inicio
	for aux != nil {
		fmt.Println(aux.Contenido)
		aux = aux.Siguiente
	}
}

//obtengo el tamaño de la lista
func ObtenerTamanioM(lista *ListaMensajes) int{
	return lista.Tamano
}
