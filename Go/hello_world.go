package main

import (
	"container/list"
	"fmt"
	"reflect"
)

func main() {
	//print("Hola, Go!")

	//Hola Mundo

	/*
	 Esto es un comentario
	*/

	fmt.Println("Hola, Go!")

	//Variables

	var myString string = "Esto es una cadena de texto."
	fmt.Println(myString)

	myString = "Aquí cambió el valor de la cadena de texto."
	fmt.Println(myString)

	//myString = 6 Error

	var myString2 string = "Esto es una cadena de texto"
	fmt.Println(myString2)

	var myInt int = 7
	myInt = myInt + 4
	fmt.Println(myInt)
	fmt.Println(myInt - 1)
	fmt.Println(myInt)

	//fmt.Println(myString + string(myInt))

	fmt.Println(myString, myInt)

	fmt.Println(reflect.TypeOf(myInt))

	var myFloat = 6.5
	fmt.Println(myFloat)
	fmt.Println(reflect.TypeOf(myFloat))

	fmt.Println(float64(myInt) + myFloat)

	var myBool bool = false
	myBool = true
	fmt.Println(myBool)

	//Variable declarada e inicializada de manera abreviada

	myString3 := "Esto es una cadena de texto"
	fmt.Println(myString3)

	//Constantes

	const myConst = "Esto es una constante"
	fmt.Println(myConst)

	//Control de Flujo

	myInt = 10
	myString = "Hola"

	if myInt == 10 && myString == "Hola" {
		fmt.Println("El valor es 10")
	} else if myInt == 11 || myString == "Hola" {
		fmt.Println("El valor es 11")
	} else {
		fmt.Println("El valor no es 10")
	}

	//Arrays

	//var myArray[3] int = [1; 2; 3]
	var myArray [3]int
	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3
	//myArray[3] = 3 Error
	fmt.Println(myArray[2])
	//fmt.Println(myArray[3]) Error

	//Mapas

	myMap := make(map[string]int)

	myMap["Brais"] = 36
	myMap["Crais01"] = 35
	myMap["Hoz98"] = 24
	fmt.Println(myMap)
	fmt.Println(myMap["Brais"])

	myMap2 := map[string]int{"Brais": 36, "Crais01": 35, "Hoz98": 24}
	fmt.Println(myMap2)

	//Listas

	myList := list.New()
	myList.PushBack(1)
	myList.PushBack(2)
	myList.PushBack(3)
	fmt.Println(myList.Back().Value)

	//Bucles

	for index := 0; index < len(myArray); index++ {
		fmt.Println(myArray[index])
	}

	for index, value := range myMap {
		fmt.Println(index, value)
	}

	//Función

	fmt.Println(myFunction())

	//Estructuras

	type MyStruct struct {
		name string
		age  int
	}

	myStruct := MyStruct{"Brais", 36}
	fmt.Println(myStruct)

}

func myFunction() string {
	return "Mi función"
}
