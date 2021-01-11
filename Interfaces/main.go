package main

type englishBot struct{}
type spanishBot struct{}

func main() {

}

func (eb englishBot) getGretting() string {
	// VERY custom logic for generating an english greeting
	return "Hi There!"
}

func (sb spanishBot) getGretting() string {
	return "Hola!"
}
