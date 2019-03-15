package main

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#123456",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		println(color + "=>" + hex)
	}
}
