package main

import (
	"machine"
	"time"
)


func main() {
	var led machine.Pin
	led = machine.LED
	println(led)
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		println("running...")
		led.Low()
		time.Sleep(time.Millisecond * 300)
		led.High()
		time.Sleep(time.Millisecond * 300)
	}
}

// const (
// 	D0	= PD0	// RX
// 	D1	= PD1	// TX
// 	D2	= PD2
// 	D3	= PD3
// 	D4	= PD4
// 	D5	= PD5
// 	D6	= PD6
// 	D7	= PD7
// 	D8	= PB0 // 8pin
// 	D9	= PB1
// 	D10	= PB2
// 	D11	= PB3
// 	D12	= PB4
// 	D13	= PB5 // 13pin
// )
