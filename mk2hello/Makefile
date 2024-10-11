TAMAGO := /home/lizeren/tamago-go-latest/bin/go

APP := armory-hello-world
GOENV := GO_EXTLINK_ENABLED=0 CGO_ENABLED=0 GOOS=tamago GOARM=7 GOARCH=arm
TEXT_START := 0x80010000
GOFLAGS := -ldflags "-s -w -T $(TEXT_START) -E _rt0_arm_tamago -R 0x1000"

IMX_USB := /home/lizeren/Downloads/imx_usb_loader/imx_usb

all: $(APP)

$(APP): main.go
	$(GOENV) $(TAMAGO) build $(GOFLAGS) -o $@

$(APP).bin: $(APP)
	arm-none-eabi-objcopy -O binary $< $@

$(APP).imx: $(APP).bin
	echo -e "DCD 0\nWRITE_DATA 4 0x30340004 0x4F400005\nWRITE_DATA 4 0x30391000 0x00000002\nJUMP_ADDRESS $(TEXT_START)" > $(APP).conf
	$(IMX_USB) -c $(APP).conf $(APP).bin
	mv $(APP).bin.imx $(APP).imx

imx: $(APP).imx

clean:
	rm -f $(APP) $(APP).bin $(APP).conf $(APP).imx $(APP).bin.imx
