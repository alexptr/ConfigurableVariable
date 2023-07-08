package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"net/http"

	"github.com/spf13/viper"
)

var MYGLOBALVARIABLE string

func main() {

	systemVar := os.Getenv("CONFIG_VALUE")
	if systemVar != "" {
		fmt.Println("Found systemVar: ", systemVar)
		MYGLOBALVARIABLE = systemVar
	} else {
		fmt.Println("No systemVar found, searching local config file...")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		if err := viper.ReadInConfig(); err != nil {
			fmt.Printf("Error reading config file: %s\n", err)
			return
		}
		myVariable := viper.GetString("my_variable")
		if myVariable == "" {
			fmt.Println("no my_variable in local config file")
			return
		}
		fmt.Println("MyVariable: ", myVariable)
		MYGLOBALVARIABLE = myVariable
	}

	handleRequest()

}

func getVariable(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(MYGLOBALVARIABLE)
}

func home(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("home page")
}

func handleRequest() {
	http.HandleFunc("/", home)
	http.HandleFunc("/getvariable", getVariable)
	log.Fatal(http.ListenAndServe(":10000", nil))
}
