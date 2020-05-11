package gadget

import "fmt"

// TapePlayer type
type TapePlayer struct {
	Batteries string
}

// Play plays a song
func (t *TapePlayer) Play(song string) {
	fmt.Println("Playing ", song)
}

// Stop stops the player
func (t *TapePlayer) Stop() {
	fmt.Println("Stopped!")
}

// TapeRecorder type
type TapeRecorder struct {
	Microphones int
}

// Record records a song
func (t *TapeRecorder) Record() {
	fmt.Println("Recording...")
}

// Play plays a song
func (t *TapeRecorder) Play(song string) {
	fmt.Println("Playing ", song)
}

// Stop stops the player
func (t *TapeRecorder) Stop() {
	fmt.Println("Stopped!")
}
