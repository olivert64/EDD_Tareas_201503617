package main

import (
	"fmt"
)

type Arbol struct {
	raiz *Nodo
}

type Nodo struct {
	Valor     int
	Izquierdo *Nodo
	Derecho   *Nodo
}

func nuevoArbol() *Arbol{
	return new(Arbol)
}

func (arbol *Arbol) Add(val int){

	if arbol.raiz == nil {
		arbol.raiz = &Nodo{Valor: val}
		return
	}

	arbol.raiz.AgregarNodo(val)
}

func (nodo *Nodo) AgregarNodo(val int){
	if val < nodo.Valor {
		if nodo.Izquierdo == nil{
			nodo.Izquierdo = &Nodo{Valor: val}
		}else{
			nodo.Izquierdo.AgregarNodo(val)
		}
	}else if val > nodo.Valor {
		if nodo.Derecho == nil{
			nodo.Derecho = &Nodo{Valor: val}
		}else{
			nodo.Derecho.AgregarNodo(val)
		}
	}
}

//recorrido preorden raiz-izq-der
func PreOrder(nodo *Nodo){
	if  nodo != nil{
		fmt.Println(nodo.Valor)
		PreOrder(nodo.Izquierdo)
		PreOrder(nodo.Derecho)
	}
}

//recorrido inorden izq-raiz-der
func InOrder(nodo *Nodo){
	if  nodo != nil{
		InOrder(nodo.Izquierdo)
		fmt.Println(nodo.Valor)
		InOrder(nodo.Derecho)
	}
}

//recorrido postorden izq-der-raiz
func PostOrder(nodo *Nodo){
	if  nodo != nil{
		PostOrder(nodo.Izquierdo)
		PostOrder(nodo.Derecho)
		fmt.Println(nodo.Valor)
	}
}

func main(){
	arb := nuevoArbol()

	valores := []int{9,5,4,6,10}

	for _, val := range valores {
		arb.Add(val)
	}

	fmt.Println("-----RECORRIDO PRE ORDEN-----")
	PreOrder(arb.raiz)
	fmt.Println("-----RECORRIDO IN ORDEN-----")
	InOrder(arb.raiz)
	fmt.Println("-----RECORRIDO POST ORDEN-----")
	PostOrder(arb.raiz)

}