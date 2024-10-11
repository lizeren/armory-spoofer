package main

import (
    "github.com/usbarmory/tamago/soc/nxp/imx6ul"
    usbarmory "github.com/usbarmory/tamago/board/usbarmory/mk2"

    "github.com/usbarmory/imx-usbserial"


)


func init() {

	switch imx6ul.Model() {
	case "i.MX6ULL", "i.MX6ULZ":
		imx6ul.SetARMFreq(imx6ul.FreqMax)
	case "i.MX6UL":
		imx6ul.SetARMFreq(imx6ul.Freq528)
	}

}



func main() {

    

    usbarmory.LED("blue", false)
    usbarmory.LED("white", true)
    // Create a new UART instance
    serial := &usbserial.UART{}

    // Initialize the UART
    serial.Init()
    
    
    // Define the message to send
    message := "Hello from imx-usbserial!"

    for {
        // Send the message
        serial.Write([]byte(message))
    }

}

