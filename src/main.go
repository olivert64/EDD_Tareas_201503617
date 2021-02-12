package main

import (
	l "./estructuras"
	o "./mensajes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type Entrada struct{
	Mensajes []struct{
		Origen string
		Destino string
		Msg[] struct{
			Fecha string
			Texto string
		}
	}
}

type Mensajes struct{
	Origen string
	Destino string
	Msg []Contenido
}

type Contenido struct{
	Fecha string
	Texto string
}

var TodosLosMensajes l.ListaD

func getMensajes(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")//para que devuelva la informacion tipo json

	var arrayMsgJson []Mensajes

	auxNodoD := TodosLosMensajes.Inicio

	for auxNodoD != nil{
		auxMensajeJson := Mensajes{}
		auxMensajeJson.Origen = auxNodoD.Mensaje.Origen
		auxMensajeJson.Destino = auxNodoD.Mensaje.Destino

		var arrayContenidoJson []Contenido
		auxNodoMsg := auxNodoD.Mensaje.Contenido.Inicio

		for  auxNodoMsg != nil{
			auxContenidoJson := Contenido{}
			auxContenidoJson.Texto = auxNodoMsg.Contenido.Texto
			auxContenidoJson.Fecha = auxNodoMsg.Contenido.Fecha

			arrayContenidoJson = append(arrayContenidoJson, auxContenidoJson)
			auxNodoMsg = auxNodoMsg.Siguiente
		}

		auxMensajeJson.Msg = arrayContenidoJson

		arrayMsgJson = append(arrayMsgJson, auxMensajeJson)

		auxNodoD = auxNodoD.Siguiente
	}

	json.NewEncoder(w).Encode(arrayMsgJson)
}

func postMensaje(w http.ResponseWriter, r *http.Request){

	var entradaMsg Entrada
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w,"Inserta un Mensaje Valido")
		return
	}


	json.Unmarshal(reqBody, &entradaMsg)



	for i := 0; i < len(entradaMsg.Mensajes); i++{
		auxMensaje := o.Mensaje{}
		auxMensaje.Destino = entradaMsg.Mensajes[i].Destino
		auxMensaje.Origen = entradaMsg.Mensajes[i].Origen


		listaMsgAux := o.ListaMensajes{}
		for j := 0; j < len(entradaMsg.Mensajes[i].Msg); j++{
			auxContenido := o.Contenido{}
			auxContenido.Texto = entradaMsg.Mensajes[i].Msg[j].Texto
			auxContenido.Fecha = entradaMsg.Mensajes[i].Msg[j].Fecha
			o.InsertarC(&auxContenido, &listaMsgAux)
		}
		auxMensaje.Contenido = &listaMsgAux
		l.Insertar(&auxMensaje, &TodosLosMensajes)
	}

	w.Header().Set("Content-Type", "application/json")//para que devuelva la informacion tipo json
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(entradaMsg)
}

//se ejecuta cuando se visite la url principal o inicial
func indexRoute(w http.ResponseWriter, r *http.Request){

	fmt.Fprintf(w, "Bienvenido a mi Tarea")
}

func main() {
	//para iniciar un enrutador
	router := mux.NewRouter().StrictSlash(true)

	//para definir una ruta y decir que haga algo en esea url
	router.HandleFunc("/", indexRoute)

	router.HandleFunc("/getMsg", getMensajes).Methods("GET")

	router.HandleFunc("/postMsg", postMensaje).Methods("POST")


	//crear servidor http, log lo mantiene ejecutando
	log.Fatal(http.ListenAndServe(":3000", router))

}


/*//ruta para optener todas las tareas
func getTareas(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")//para que devuelva la informacion tipo json
	json.NewEncoder(w).Encode(tareas)
}

func crearTarea(w http.ResponseWriter, r *http.Request){
	
	var newTarea Tarea
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w,"Inserta una tarea valida")
		return
		
	}

	json.Unmarshal(reqBody, &newTarea)

	newTarea.ID = len(tareas) + 1

	tareas = append(tareas, newTarea)

	w.Header().Set("Content-Type", "application/json")//para que devuelva la informacion tipo json
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTarea)
}

//metodo para optener una sola tarea
func getTask(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	tareaID, err := strconv.Atoi(vars["id"])

	if err != nil { 
		fmt.Fprintf(w, "ivalid ID xd")
		return
	}

	for _, tarea := range tareas{
		if tarea.ID == tareaID {
			w.Header().Set("Content-Type", "application/json")//para que devuelva la informacion tipo json
			json.NewEncoder(w).Encode(tarea)
		}
	}
	
}

//metodo para eliminar una tarea
func deleteTarea(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	tareaID, err := strconv.Atoi(vars["id"])

	if err != nil { 
		fmt.Fprintf(w, "ivalid ID xd")
		return
	}

	for i, t := range tareas{
		if t.ID == tareaID {
			tareas = append(tareas[:i], tareas[i + 1:]...)
			fmt.Fprintf(w,"La tarea con el ID %v ha sido removida", tareaID)
		}
	}
}

func updTarea(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	tareaID, err := strconv.Atoi(vars["id"])
	var updateTarea Tarea


	if err != nil { 
		fmt.Fprintf(w, "ivalid ID xd")
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w,"Inserta datos validos")
		return
		
	}

	json.Unmarshal(reqBody, &updateTarea)

	for i, t := range tareas{
		if t.ID == tareaID {
			tareas = append(tareas[:i], tareas[i+1:]...)
			updateTarea.ID = tareaID
			tareas = append(tareas, updateTarea)

			fmt.Fprintf(w,"tarea con el ID %v se ha modificado", tareaID)

		} 
	}


}*/


