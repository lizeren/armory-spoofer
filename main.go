// https://github.com/usbarmory/armory-ums
//
// Copyright (c) WithSecure Corporation
// https://foundry.withsecure.com
//
// Use of this source code is governed by the license
// that can be found in the LICENSE file.

package main

import (
	"errors"
	"log"
	"time"

	"github.com/usbarmory/tamago/soc/nxp/imx6ul"
	"github.com/usbarmory/tamago/soc/nxp/usb"
	"github.com/usbarmory/tamago/soc/nxp/usdhc"

	usbarmory "github.com/usbarmory/tamago/board/usbarmory/mk2"

)

var cards []*usdhc.USDHC

func init() {
	log.SetFlags(0)

	switch imx6ul.Model() {
	case "i.MX6ULL", "i.MX6ULZ":
		imx6ul.SetARMFreq(imx6ul.FreqMax)
	case "i.MX6UL":
		imx6ul.SetARMFreq(imx6ul.Freq528)
	}
}

func detect(card *usdhc.USDHC) (err error) {
	if card == nil {
		return errors.New("no such device")
	}

	if err = card.Detect(); err != nil {
		return
	}

	info := card.Info()
	capacity := int64(info.BlockSize) * int64(info.Blocks)
	giga := capacity / (1000 * 1000 * 1000)
	gibi := capacity / (1024 * 1024 * 1024)

	log.Printf("imx6_usdhc: %d GB/%d GiB card detected %+v", giga, gibi, info)

	cards = append(cards, card)

	return
}

func main() {
    
	ch, err := usbarmory.DetectDebugAccessory(500 * time.Millisecond)

	value := <- ch
	
	if value == false{
		usbarmory.LED("blue", false)
	} else {
		usbarmory.LED("blue", true)
	}	

	if err != nil{
		usbarmory.LED("white", false)
	} else {
		usbarmory.LED("white", true)
	}	



    // message := "Hello from UART2!\r\n"

	// for i := 0; i < 10; i++ {
	// 	for _, c := range message {
	// 		imx6ul.UART2.Tx(byte(c))
	// 	}
	// }


	err = detect(usbarmory.SD)
	err = detect(usbarmory.MMC)



	device := &usb.Device{
		Setup: setup,
	}
	configureDevice(device)

	iface := buildMassStorageInterface()
	device.Configurations[0].AddInterface(iface)

	log.Printf("Init USB1")
	usbarmory.USB1.Init()
	log.Printf("DeviceMode set for USB1")
	usbarmory.USB1.DeviceMode()
	log.Printf("Reset USB1")
	usbarmory.USB1.Reset()

	log.Printf("USB1 start")
	
	// never returns
	usbarmory.USB1.Start(device)
	log.Printf("It should never return. Bug here.")

}
