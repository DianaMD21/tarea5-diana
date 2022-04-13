// main
// Embed demonstrates basic struct embedding.

package main

import (
	"fmt"
	// "time"
	"bufio"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"runtime"

	"github.com/B-Abraham/L5-ISC315/control"
	"github.com/B-Abraham/L5-ISC315/share"
)

// configuration contains the application settings
type configuration struct {
	Server control.Server `json:"Server"`
	User   control.User   `json:"User"`
}

func (c *configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}

var config = &configuration{}

// var config configuration
func init() {
	// Verbose logging with file name and line number
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	// var err error
	file, err := os.OpenFile("biblos.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	log.SetOutput(file)
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func readFile(nameFile string) (tabla []string) {
	file, err := os.Open(nameFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		tabla = append(tabla, line)
	}
	return tabla
}

func main() {
	share.Load("config.json", config)
	data, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Print(err)
	}
	fname := "./data.txt"
	tab := readFile(fname)
	err = json.Unmarshal([]byte(data), &config)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("   Reporte Libros   ")
	fmt.Println()
	control.JLoginGET(config.Server, config.User)

	fmt.Println("Listado de libros con Titulo: ", tab[0])
	list := control.JBook(config.Server, tab[0])
	for i := 0; i < len(list); i++ {
		fmt.Println(list[i].Title)
		fmt.Println(list[i].Author)
		fmt.Println(list[i].Editor)
		fmt.Println(list[i].Language)
		fmt.Println(list[i].Comment)
		fmt.Println(list[i].Year)
		fmt.Println()
	}

	fmt.Println("Listado de libros con Autor: ", tab[1])
	list = control.JAuth(config.Server, tab[1])
	for i := 0; i < len(list); i++ {
		fmt.Println(list[i].Title)
		fmt.Println(list[i].Author)
		fmt.Println(list[i].Editor)
		fmt.Println(list[i].Language)
		fmt.Println(list[i].Comment)
		fmt.Println(list[i].Year)
		fmt.Println()
	}

	fmt.Println("Listado de libros con Editora: ", tab[2])
	list = control.JEdit(config.Server, tab[2])
	for i := 0; i < len(list); i++ {
		fmt.Println(list[i].Title)
		fmt.Println(list[i].Author)
		fmt.Println(list[i].Editor)
		fmt.Println(list[i].Language)
		fmt.Println(list[i].Comment)
		fmt.Println(list[i].Year)
		fmt.Println()
	}

	fmt.Println("Listado de libros con Idioma: ", tab[3])
	list = control.JLang(config.Server, tab[3])
	for i := 0; i < len(list); i++ {
		fmt.Println(list[i].Title)
		fmt.Println(list[i].Author)
		fmt.Println(list[i].Editor)
		fmt.Println(list[i].Language)
		fmt.Println(list[i].Comment)
		fmt.Println(list[i].Year)
		fmt.Println()
	}
}
