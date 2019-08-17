package main

import (
	"math/rand"
	"time"

	ws281x "github.com/rpi-ws281x/rpi-ws281x-go"
)

const (
	brightness = 128
	ledCounts  = 300
	gpioPin    = 18
	freq       = 800000
	sleepTime  = 200
)

type ws struct {
	ws2811 *ws281x.WS2811
}

func (ws *ws) init() error {
	err := ws.ws2811.Init()
	if err != nil {
		return err
	}

	return nil
}

func (ws *ws) close() {
	ws.ws2811.Fini()
}

func rgbToColor(r int, g int, b int) uint32 {
	return uint32(uint32(r)<<16 | uint32(g)<<8 | uint32(b))
}

func (ws *ws) randomRGB() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for {
		for i := 0; i < len(ws.ws2811.Leds(0)); i++ {
			ws.ws2811.Leds(0)[i] = rgbToColor(r1.Intn(255), r1.Intn(255), r1.Intn(255))
		}

		ws.ws2811.Render()

		time.Sleep(sleepTime * time.Millisecond)
	}
}

func main() {
	opt := ws281x.DefaultOptions
	opt.Channels[0].Brightness = brightness
	opt.Channels[0].LedCount = ledCounts
	opt.Channels[0].GpioPin = gpioPin
	opt.Frequency = freq

	ws2811, err := ws281x.MakeWS2811(&opt)
	if err != nil {
		panic(err)
	}

	ws := ws{
		ws2811: ws2811,
	}

	err = ws.init()
	if err != nil {
		panic(err)
	}
	defer ws.close()

	ws.randomRGB()
}
