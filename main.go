package main

import (
	"encoding/json"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"fmt"
	"strings"
)

var element []elements

type elements struct {

	AtomicMass string `json:"atomicMass"` //0
	AtomicNumber string `json:"atomicNumber"` //1
	AtomicRadius string `json:"atomicRadius"` //2
	BoilingPoint string `json:"boilingPoint"` //3
	BondingType string `json:"bondingType"` //4
	CpkHexColor string `json:"cpkHexColor"` //5
	Density string `json:"density"` //6
	ElectronAffinity string `json:"electronAffinity"` //7
	Electronegativity string `json:"electronegativity"` //8
	ElectronicConfiguration string `json:"electronicConfiguration"` //9
	GroupBlock string `json:"groupBlock"` //10
	IonRadius string `json:"ionRadius"` //11
	IonizationEnergy string `json:"ionizationEnergy"` //12
	MeltingPoint string `json:"meltingPoint"` //13
	Name string `json:"name"` //14
	OxidationStates string `json:"oxidationStates"` //15
	StandardState string `json:"standardState"` //16
	Symbol string `json:"symbol"` //17
	VanDelWaalsRadius string `json:"vanDelWaalsRadius"` //18
	YearDiscovered string `json:"yearDiscovered"` //19

}

func allElements(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	_ = json.NewEncoder(w).Encode(element)
}

func atomicNumber(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	atomicNumber := vars["atomicNumber"]
	fmt.Println(atomicNumber)
	for _, data := range element {
		if strings.Compare(atomicNumber, data.AtomicNumber) == 0 {
			_ = json.NewEncoder(w).Encode(data)
			break
		}
	}
}


func main() {

	// Reading JSON file
	jsonFile, err := os.Open("data.json")
	if err != nil {
		fmt.Println("Error reading file")
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	_ = json.Unmarshal(byteValue, &element)

	// Routers
	router := mux.NewRouter()
	router.HandleFunc("/", allElements).Methods("GET")
	router.HandleFunc("/atomicNumber/{atomicNumber}", atomicNumber).Methods("GET")

	//CORS Headers
	headersOk := handlers.AllowedHeaders([]string{"Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET"})

	//Starting Server
	err = http.ListenAndServe(":3000", handlers.CORS(originsOk, headersOk, methodsOk)(router))
	if err!=nil {
		log.Fatal(err)
	}
}
