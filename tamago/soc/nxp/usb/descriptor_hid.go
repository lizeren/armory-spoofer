// descriptor_hid.go: USB HID Descriptor Support
// https://github.com/usbarmory/tamago
//
// Copyright (c) WithSecure Corporation
// https://foundry.withsecure.com
//
// Use of this source code is governed by the license
// that can be found in the LICENSE file.

package usb

import (
	"bytes"
	"encoding/binary"
)

// HID descriptor constants
const (
	// HID Class Specification release number in binary-coded decimal (1.11)
	HID_BCD  = 0x0111
	HID_TYPE = 0x21 // HID Descriptor type

	// HID descriptor lengths
	HID_DESCRIPTOR_LENGTH = 9 // Typically 9 bytes for HID descriptor
)

// HIDDescriptor represents the HID descriptor, which provides information
// about the HID device to the host.
type HIDDescriptor struct {
	Length            uint8  // Size of the HID descriptor (in bytes)
	DescriptorType    uint8  // Descriptor type (HID)
	BcdHID            uint16 // HID Class Specification release number
	CountryCode       uint8  // Country code of the localized hardware
	NumDescriptors    uint8  // Number of subordinate descriptors
	ReportDescriptor  uint8  // Descriptor type (Report)
	DescriptorLength  uint16 // Length of the report descriptor
}

// SetDefaults initializes the HID descriptor with default values.
func (d *HIDDescriptor) SetDefaults() {
	d.Length = HID_DESCRIPTOR_LENGTH          // HID Descriptor size
	d.DescriptorType = HID_TYPE               // HID Descriptor type
	d.BcdHID = HID_BCD                        // HID specification version (1.11)
	d.CountryCode = 0                         // No specific localization
	d.NumDescriptors = 1                      // Number of subordinate descriptors
	d.ReportDescriptor = 0x22                 // Report Descriptor type
	d.DescriptorLength = uint16(len(HIDReportDescriptor)) // Length of the report descriptor
}

// Bytes converts the HID descriptor structure to a byte slice.
func (d *HIDDescriptor) Bytes() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, d)
	return buf.Bytes()
}

// HIDReportDescriptor is the report descriptor for the HID device.
// This defines how the host should interpret the data that the HID device sends.
// You can customize this descriptor based on the type of HID device (e.g., keyboard, mouse).
var HIDReportDescriptor = []byte{
	0x05, 0x01,       // Usage Page (Generic Desktop)
	0x09, 0x06,       // Usage (Keyboard)
	0xA1, 0x01,       // Collection (Application)
	0x05, 0x07,       // Usage Page (Key Codes)
	0x19, 0xE0,       // Usage Minimum (224)
	0x29, 0xE7,       // Usage Maximum (231)
	0x15, 0x00,       // Logical Minimum (0)
	0x25, 0x01,       // Logical Maximum (1)
	0x75, 0x01,       // Report Size (1)
	0x95, 0x08,       // Report Count (8)
	0x81, 0x02,       // Input (Data, Variable, Absolute)
	0x95, 0x01,       // Report Count (1)
	0x75, 0x08,       // Report Size (8)
	0x81, 0x03,       // Input (Constant) - Reserved byte
	0x95, 0x05,       // Report Count (5)
	0x75, 0x01,       // Report Size (1)
	0x05, 0x08,       // Usage Page (LEDs)
	0x19, 0x01,       // Usage Minimum (Num Lock)
	0x29, 0x05,       // Usage Maximum (Kana)
	0x91, 0x02,       // Output (Data, Variable, Absolute)
	0x95, 0x01,       // Report Count (1)
	0x75, 0x03,       // Report Size (3)
	0x91, 0x03,       // Output (Constant) - Padding
	0x95, 0x06,       // Report Count (6)
	0x75, 0x08,       // Report Size (8)
	0x15, 0x00,       // Logical Minimum (0)
	0x25, 0x65,       // Logical Maximum (101)
	0x05, 0x07,       // Usage Page (Key Codes)
	0x19, 0x00,       // Usage Minimum (0)
	0x29, 0x65,       // Usage Maximum (101)
	0x81, 0x00,       // Input (Data, Array)
	0xC0,             // End Collection
}
