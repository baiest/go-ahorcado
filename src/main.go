package main

import (
	pk "go-ahorcado/src/models"
)

func main() {
	var tablero pk.Tablero
	tablero.Init()
	tablero.StartGame()
}
