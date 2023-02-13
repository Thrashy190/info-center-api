package utils

import (
	"log"

	"github.com/fatih/color"
)

func Warning(text string){
	color.Set(color.FgRed)
	log.Fatalf(text)
	color.Unset()
}

func Succes(text string){
	color.Set(color.FgGreen)
	log.Println(text)
	color.Unset()
}

func Process(text string){
	color.Set(color.FgBlue)
	log.Println(text)
	color.Unset()
}