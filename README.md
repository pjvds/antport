# antport

A go library to communicate via a the ANT usb stick.

## First time use
You need permission to read and write the USB device. To do so add a text file with the following to `/etc/udev/rules.d/99-garmin.rules`

On Ubuntu 10.04 (or lower):

	SUBSYSTEM=="usb", SYSFS{idVendor}=="0fcf", SYSFS{idProduct}=="1008", MODE="666"

On Ubuntu 12.04 (or other distros running newer udev):

	SUBSYSTEM=="usb", ATTR{idVendor}=="0fcf", ATTR{idProduct}=="1008", MODE="666"

