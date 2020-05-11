package main

import (
	"github.com/a-soliman/head_first_go/gadget"
)

// Player interface
type Player interface {
	Play(string)
	Stop()
}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	var player Player = &gadget.TapePlayer{}
	mixTape := []string{"My Way", "Take me to the Moon", "The way you look", "Strangers in the night"}
	playList(player, mixTape)
	player = &gadget.TapeRecorder{}
	playList(player, mixTape)
}
