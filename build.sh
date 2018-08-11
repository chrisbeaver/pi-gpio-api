#!/bin/bash
env GOOS=linux GOARCH=arm GOARM=5 go build
scp pi-gpio-api pi@raspberrypi:/home/pi/share/gpio
rm pi-gpio-api