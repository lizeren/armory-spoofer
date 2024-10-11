# Makefile for building the Go `main` package with dbd and imx targets

# Define metadata variables
BUILD_USER ?= $(shell whoami)
BUILD_HOST ?= $(shell hostname)
BUILD_DATE ?= $(shell /bin/date -u "+%Y-%m-%d %H:%M:%S")
BUILD = ${BUILD_USER}@${BUILD_HOST} on ${BUILD_DATE}
REV = $(shell git rev-parse --short HEAD 2> /dev/null || echo "unknown")

SHELL = /bin/bash
BUILD_FLAGS = -trimpath


# Define paths and dependencies
TAMAGO_GO := /home/lizeren/tamago-go-latest/bin/go
GOENV := GO_EXTLINK_ENABLED=0 CGO_ENABLED=0 GOOS=tamago GOARCH=arm GOARM=7 GOROOT=/home/lizeren/tamago-go-latest

# TAMAGO := /home/lizeren/tamago-go-latest/bin/go
TAMAGO_PKG := github.com/usbarmory/tamago@v0.0.0-20240924114619-273d67cd811d
GOMODCACHE := $(shell go env GOMODCACHE)

# Define the Go environment variables for cross-compilation and build flags
APP := imx-usbserial-demo
# GOENV := GO_EXTLINK_ENABLED=0 CGO_ENABLED=0 GOOS=tamago GOARCH=arm GOARM=7
TEXT_START := 0x80010000

# GOFLAGS := -tags linkramsize ${BUILD_FLAGS} -ldflags "-s -w -T $(TEXT_START) -E _rt0_arm_tamago -R 0x1000 -X 'main.Build=${BUILD}' -X 'main.Revision=${REV}'"
GOFLAGS := -tags=linkramsize ${BUILD_FLAGS} -ldflags="-s -w -T $(TEXT_START) -E _rt0_arm_tamago -R 0x1000"

# Application name



# Primary build targets
.PHONY: all clean dcd imx elf

all: $(APP)

imx: $(APP).imx

dcd: $(APP).dcd

elf: $(APP)

#### Utilities ####

# Ensure Go modules are tidy
tidy:
	go mod tidy

# Build the application binary using the Go environment
# $(APP): check_tamago
# 	$(GOENV) $(TAMAGO) build $(GOFLAGS) -o ${APP}
$(APP): check_tamago
	$(GOENV) $(TAMAGO_GO) build $(GOFLAGS) -o ${APP}


# Create the .dcd file by copying the configuration file from the Tamago package
$(APP).dcd: check_tamago
	echo $(GOMODCACHE)
	echo $(TAMAGO_PKG)
	cp -f $(GOMODCACHE)/$(TAMAGO_PKG)/board/usbarmory/mk2/imximage.cfg $(APP).dcd


# Convert the Go binary to a .bin file
$(APP).bin: CROSS_COMPILE=arm-none-eabi-
$(APP).bin: $(APP)
	$(CROSS_COMPILE)objcopy --enable-deterministic-archives \
	    -j .text -j .rodata -j .shstrtab -j .typelink \
	    -j .itablink -j .gopclntab -j .go.buildinfo -j .noptrdata -j .data \
	    -j .bss --set-section-flags .bss=alloc,load,contents \
	    -j .noptrbss --set-section-flags .noptrbss=alloc,load,contents \
	    $(APP) -O binary $(APP).bin


# Create the .bin file using objcopy and build an .imx image from the binary and .dcd file
$(APP).imx: SOURCE_DATE_EPOCH=0
$(APP).imx: $(APP).bin $(APP).dcd
	mkimage -n $(APP).dcd -T imximage -e $(TEXT_START) -d $(APP).bin $(APP).imx
	# Copy entry point from ELF file
	dd if=$(APP) of=$(APP).imx bs=1 count=4 skip=24 seek=4 conv=notrunc
		

#### Dependency Checks ####

# Check if go.mod exists, if not create it and get dependencies
check_tamago:
	@if [ ! -f "go.mod" ]; then \
		echo "Initializing go.mod file..."; \
		go mod init imx-usbserial-demo; \
		go get github.com/usbarmory/imx-usbserial; \
	fi

# Clean up generated files
clean:
	rm -f $(APP)
	@rm -fr $(APP).bin $(APP).imx $(APP).dcd
