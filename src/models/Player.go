package models

import (
	"bufio"
	"fmt"
	"os"
)

type Player struct {
	Letra rune
}

func (player *Player) SetLetra() {

	reader := bufio.NewReader((os.Stdin))

	fmt.Print("Escribe una letra: ")
	char, _, err := reader.ReadRune()

	if err != nil {
		fmt.Println(err)
	}

	player.Letra = char
}
