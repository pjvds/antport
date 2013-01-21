# antport

A go library to communicate via a the ANT usb stick.

## First time use in Ubuntu
You need permission to read and write the USB device. To do so add a text file with the following to `/etc/udev/rules.d/99-garmin.rules`

On Ubuntu 10.04 (or lower):

	SUBSYSTEM=="usb", SYSFS{idVendor}=="0fcf", SYSFS{idProduct}=="1008", MODE="666"

On Ubuntu 12.04 (or other distros running newer udev):

	SUBSYSTEM=="usb", ATTR{idVendor}=="0fcf", ATTR{idProduct}=="1008", MODE="666"

## First time use in Mac OSX
Mac requires libusb to be installed from source code. Make sure you have XCode, Mac OSX SDK and the command line tools installed and run:

	brew install libusb --HEAD

Now when you build the project you can get the following error:

	# github.com/kylelemons/gousb/usb
	iso.go:67:8: struct size calculation error off=72 bytesize=64

[Joseph Poirier](http://code.google.com/p/go/issues/detail?id=3505#c10) proposed a quick fix that should work:

	Just wondering if anyone has looked into this since the last post? I tried to compile Kyle's gousb package on OSX today but ran into this problem; the quick hack was to change the struct's flexible member size to 1 in libusb.h.

You can find `libusb.h` in `/usr/local/include/libusb-1.0/libusb.h`. [Here](https://gist.github.com/4578277#file-libusb-h-L937) is a gist of my modified `libusb.h`. The change has been made at line 937.