package models

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
)

type GameState struct {
	Palabra   string
	Intentos  int16
	Underline []string
}

type Tablero struct {
	State  GameState
	Player Player
}

//Inicializar valores del tablero
func (tablero *Tablero) Init() {
	random := rand.Intn(len(PALABRAS) - 1)
	tablero.State = GameState{
		Palabra:  PALABRAS[random],
		Intentos: 0,
	}

	tablero.Player = Player{}
}

func (tablero Tablero) clearTablero() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func (tablero *Tablero) StartGame() {
	tablero.generateUnderLine()
	for tablero.State.Intentos < int16(len(DIBUJO)) {
		tablero.clearTablero()
		fmt.Print("Adivina la palabra ----->  ")
		for _, letra := range tablero.State.Underline {
			fmt.Printf("%s ", letra)
		}
		fmt.Println()
		tablero.PintarDibujo()
		fmt.Printf("\nIntentos fallidos X: %d\n", tablero.State.Intentos)
		tablero.Player.SetLetra()
		tablero.SetLetraInPalabra()
	}
}

func (tablero *Tablero) generateUnderLine() {
	underLines := strings.Split(tablero.State.Palabra, "")
	for i, letra := range underLines {
		if letra == " " {
			continue
		}
		underLines[i] = "_"
	}

	tablero.State.Underline = underLines
}

// Se verifica que la letra ingresada se encuentra en la palabra
func (tablero *Tablero) SetLetraInPalabra() {
	subString := tablero.State.Palabra
	index := strings.IndexRune(subString, tablero.Player.Letra)

	if index == -1 {
		tablero.State.Intentos += 1
		fmt.Println("No esta la letra")
		return
	}

	//Este ciclo verifica si la letra se encuentra mas de una vez
	//en la palabra
	for len(subString)-1 > 0 {
		tablero.State.Underline[index] = string(tablero.Player.Letra)
		subString = tablero.State.Palabra[index+1:]
		fmt.Println(subString)
		tempIndex := strings.IndexRune(subString, tablero.Player.Letra)
		if tempIndex == -1 {
			break
		}
		index += tempIndex + 1
	}
}

func (tablero *Tablero) PintarDibujo() {
	fmt.Println(DIBUJO[tablero.State.Intentos])
}
