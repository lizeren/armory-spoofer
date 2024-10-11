package main

import (
	"log"
	"time"

	"github.com/usbarmory/tamago/board/usbarmory/mk2"
	"github.com/usbarmory/imx-usbserial"
)

func main() {
	// Initialize the board
	mk2.Init()

	// Initialize the serial interface
	serial := &usbserial.UART{}
	err := serial.Init()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize USB
	mk2.USB1.Init()
	mk2.USB1.DeviceMode()

	// Start the USB device
	go mk2.USB1.Start(serial.Device)

	// Wait a bit for the USB connection to be established
	time.Sleep(time.Second)

	// Print "Hello World" every second
	for {
		_, err := serial.Write([]byte("Hello World\r\n"))
		if err != nil {
			log.Printf("Error writing: %v", err)
		}
		time.Sleep(time.Second)
	}
}