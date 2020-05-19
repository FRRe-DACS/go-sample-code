package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/dimiro1/banner/autoload"
)

type hospital struct {
	ID        string `json:"ID"`
	Nombre    string `json:"Nombre"`
	Direccion string `json:"Direccion"`
	NroCamas  int16  `json:"NroCamas"`
	Email     string `json:"Email"`
}

type allHospitales []hospital

var hospitales = allHospitales{
	{
		ID:        "1",
		Nombre:    "Perrando",
		Direccion: "Av. 9 de Julio 1101",
		NroCamas:  50,
		Email:     "perrando@hospitales.ar",
	},
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func getAllHospitales(w http.ResponseWriter, r *http.Request) {
	log.Print("Devolviendo lista de Hospitales.")
	enableCors(&w)
	json.NewEncoder(w).Encode(hospitales)
}

func findHospitalByID(w http.ResponseWriter, r *http.Request) {
	hopitalID := mux.Vars(r)["id"]
	log.Print("Devolviendo Hospital con ID ", hopitalID)
	enableCors(&w)
	for _, hospital := range hospitales {
		if hospital.ID == hopitalID {
			json.NewEncoder(w).Encode(hospital)
		}
	}
}

func createHopital(w http.ResponseWriter, r *http.Request) {
	var newHospital hospital
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Por favor ingrese bien los datos del hospital")
	}

	json.Unmarshal(reqBody, &newHospital)

	log.Print("Creando Hospital: ", newHospital)
	enableCors(&w)
	hospitales = append(hospitales, newHospital)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newHospital)
}

func deleteHopital(w http.ResponseWriter, r *http.Request) {
	hopitalID := mux.Vars(r)["id"]
	log.Print("Borriando Hospital con ID ", hopitalID)
	enableCors(&w)
	for i, hospital := range hospitales {
		if hospital.ID == hopitalID {
			hospitales = append(hospitales[:i], hospitales[i+1:]...)
			fmt.Fprintf(w, "The event with ID %v has been deleted successfully", hopitalID)
		}
	}
}

func main() {
	log.Print("Starting Hospital RESTful API!")
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/hospital", getAllHospitales).Methods("GET")
	router.HandleFunc("/hospital/{id}", findHospitalByID).Methods("GET")
	router.HandleFunc("/hospital", createHopital).Methods("POST")
	router.HandleFunc("/hospital/{id}", deleteHopital).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
