//Aplicacion base para Go,
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	//seteando el log
	AppName := filepath.Base(os.Args[0])
	f, err := os.OpenFile(AppName+".log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Println("Error al crear el archivo de log, se usara solo stdout")
	} else {
		log.SetOutput(f)
	}
	log.SetFlags(log.LstdFlags)
	log.SetPrefix("[" + AppName + "] ")
	log.Printf("Iniciando Base App!\n")
	fmt.Println("Base APP!", os.Args[0])
}
