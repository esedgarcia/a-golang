// hola.go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Escribe tu nombre: ")
    nombre, _ := reader.ReadString('\n')
    fmt.Printf("Hola, %s!", nombre)
}
