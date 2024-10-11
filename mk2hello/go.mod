module armory-hello-world

go 1.23.0

toolchain go1.23.1

require (
	github.com/usbarmory/imx-usbserial v0.0.0-00010101000000-000000000000
	github.com/usbarmory/tamago v0.0.0-20240924114619-273d67cd811d
)

require golang.org/x/sync v0.8.0 // indirect

replace github.com/usbarmory/imx-usbserial => /home/lizeren/Desktop/armory-ums/imx-usbserial
