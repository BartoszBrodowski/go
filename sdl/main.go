package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

var (
	width     int
	height    int
	numOfAnts int
)

type Ant struct {
	x, y          int32
	width, height int32
	velocityX     int32
	velocityY     int32
}

type Rock struct {
	x, y int32
	size int32
}

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize SDL: %s\n", err)
		return
	}
	defer sdl.Quit()

	flag.IntVar(&width, "width", 1024, "Window width")
	flag.IntVar(&height, "height", 512, "Window height")
	flag.IntVar(&numOfAnts, "count", 10, "Amount of ants")
	flag.Parse()

	screenWidth := int32(width)
	screenHeight := int32(height)

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		return
	}
	defer renderer.Destroy()

	ants := make([]Ant, numOfAnts)
	rocks := generateGrayRocks(5, screenWidth, screenHeight)
	iterationCounter := 0
	var changeVelocityInterval int

	for i := 0; i < numOfAnts; i++ {
		ants[i] = Ant{
			x:         rand.Int31n(screenWidth),
			y:         rand.Int31n(screenHeight),
			width:     4,
			height:    4,
			velocityX: getRandomVelocity(),
			velocityY: getRandomVelocity(),
		}
	}

	for {
		changeVelocityInterval = rand.Intn(21) + 10
		processEvents(&ants)
		update(&ants, &iterationCounter, changeVelocityInterval, rocks)
		render(renderer, &ants, rocks)
		iterationCounter++
	}
}

func processEvents(ants *[]Ant) {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			os.Exit(0)
		}
	}
}

func update(ants *[]Ant, iterationCounter *int, changeVelocityInterval int, rocks []Rock) {
	for i := range *ants {
		ant := &(*ants)[i]
		ant.x += ant.velocityX
		ant.y += ant.velocityY

		// Reverse direction if ant hits screen boundaries
		if ant.x < 0 || ant.x+ant.width > int32(width) {
			ant.velocityX = -ant.velocityX
			// Adjust ant position to stay within screen boundaries
			if ant.x < 0 {
				ant.x = 0
			} else {
				ant.x = int32(width) - ant.width
			}
		}
		if ant.y < 0 || ant.y+ant.height > int32(height) {
			ant.velocityY = -ant.velocityY
			if ant.y < 0 {
				ant.y = 0
			} else {
				ant.y = int32(height) - ant.height
			}
		}

		for _, rock := range rocks {
			if intersects(ant, rock) {
				ant.velocityX = -ant.velocityX
				ant.velocityY = -ant.velocityY
				break
			}
		}

		if *iterationCounter%changeVelocityInterval == 0 {
			changeVelocity(ant)
		}
	}
}

// Check if the ant intersects with the rock
func intersects(ant *Ant, rock Rock) bool {
	antRect := sdl.Rect{X: ant.x, Y: ant.y, W: ant.width, H: ant.height}
	rockRect := sdl.Rect{X: rock.x, Y: rock.y, W: rock.size, H: rock.size}
	return antRect.HasIntersection(&rockRect)
}

func changeVelocity(ant *Ant) {
	if ant.velocityX > 0 {
		ant.velocityX = int32(rand.Intn(4) + 1)
	} else {
		ant.velocityX = int32((rand.Intn(4) + 1) * -1)
	}
	if ant.velocityY > 0 {
		ant.velocityY = int32(rand.Intn(4) + 1)
	} else {
		ant.velocityY = int32((rand.Intn(4) + 1) * -1)
	}
}

func render(renderer *sdl.Renderer, ants *[]Ant, rocks []Rock) {
	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.Clear()

	for _, rock := range rocks {
		renderer.SetDrawColor(150, 150, 150, 255)
		renderer.FillRect(&sdl.Rect{X: rock.x, Y: rock.y, W: rock.size, H: rock.size})
	}

	for i, ant := range *ants {
		red := uint8(i * 10 % 256)
		green := uint8(i * 20 % 256)
		blue := uint8(i * 30 % 256)

		renderer.SetDrawColor(red, green, blue, 255)
		renderer.FillRect(&sdl.Rect{X: ant.x, Y: ant.y, W: ant.width, H: ant.height})
	}

	renderer.Present()

	// How often does the program refresh
	time.Sleep(20 * time.Millisecond)
}

func generateGrayRocks(count int, screenWidth, screenHeight int32) []Rock {
	rocks := make([]Rock, count)

	for i := 0; i < count; i++ {
		size := int32(30)
		x := rand.Int31n((screenWidth - size) % screenWidth)
		y := rand.Int31n((screenHeight - size) % screenHeight)

		rocks[i] = Rock{x: x, y: y, size: size}
	}

	return rocks
}

func getRandomVelocity() int32 {
	randomInt := rand.Intn(2)
	if randomInt == 0 {
		return int32(rand.Intn(5) + 1)
	} else {
		return int32((rand.Intn(5) + 1) * -1)
	}
}