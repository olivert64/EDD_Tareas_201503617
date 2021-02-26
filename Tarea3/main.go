package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
)

type nodo struct {
	nombre, apellido, apodo, favoritos string
	Siguiente, Anterior                *nodo
}

type lista struct {
	cabeza *nodo
	cola   *nodo
}

func (this *lista) Insertar(nuevo *nodo) {
	if this.cabeza == nil {
		this.cabeza = nuevo
		this.cola = nuevo
	} else {
		this.cola.Siguiente = nuevo
		nuevo.Anterior = this.cola
		this.cola = nuevo
	}
}

func validarGrafico(lista *lista) bool{
	error := false
		error = escribirArchivoDot(lista)
		compilarDot("./grafico/archivo.dot", "./grafico/imagen.png")
	return error
}

func compilarDot(origen string, destino string){
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpng", origen).Output()
	mode := int(0777)
	ioutil.WriteFile(destino, cmd, os.FileMode(mode))
}

func escribirArchivoDot(lista *lista) bool{
	archivo, err := os.Create("./grafico/archivo.dot")
	if err != nil {
		fmt.Println("Error al crear Archivo")
		return false
	}

	defer archivo.Close()

	var cadena string = ""

			cadena += "digraph Graphic{\n"
			NodoTemp := lista.cabeza
			i := 0
			//Escribir apuntador inicial

			for NodoTemp != nil {
				cadena += "		nodo_"+strconv.Itoa(i)+"[\n"
				cadena += "		shape=\"tab\"\n"
				cadena += "		style=\"filled\"\n"
				cadena += "		fillcolor=\"orange\"\n"
				cadena += "		fontsize=20\n"
				cadena += "		label=\""+ NodoTemp.nombre+"\n"
				cadena += "    "+ NodoTemp.apellido+"\n"
				cadena += "      "+ NodoTemp.apodo+"\n"
				cadena += "    "+ NodoTemp.favoritos+"\"\n"
				cadena += "		];\n"

				if NodoTemp.Anterior != nil {
					cadena += "		nodo_"+strconv.Itoa(i)+"->"
					cadena += "		nodo_"+strconv.Itoa(i-1)+";\n"
				}
				if NodoTemp.Siguiente != nil{
					cadena += "		nodo_"+strconv.Itoa(i)+"->"
					cadena += "		nodo_"+strconv.Itoa(i+1)+";\n"
				}
				i++
				NodoTemp = NodoTemp.Siguiente
			}

	cadena +=" }\n"
	cadena += "}"

	archivo.Sync()
	buffer := bufio.NewWriter(archivo)
	buffer.WriteString(cadena)

	if err != nil {
		panic(err)
	}

	buffer.Flush()
	return true
}


func main() {
	li := lista{nil, nil}
	a := nodo{"Marvin", "Martinez", "Marvin25ronal", "Jugar apex", nil, nil}
	b := nodo{"Yaiza", "Pineda", "Bambi", "Patinar", nil, nil}
	c := nodo{"Jonathan", "Lopez", "Pancho", "Comer", nil, nil}
	d := nodo{"usuario1", "bla", "bla", "Jugar apex", nil, nil}
	e := nodo{"usuario2", "bla", "bla", "Jugar apex", nil, nil}
	f := nodo{"usuario3", "sale edd", "vamos con todo", "100 en la fase 1", nil, nil}
	li.Insertar(&a)
	li.Insertar(&b)
	li.Insertar(&c)
	li.Insertar(&d)
	li.Insertar(&e)
	li.Insertar(&f)

	validarGrafico(&li)

}
