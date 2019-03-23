#!/bin/bash
upower -i /org/freedesktop/UPower/devices/battery_BAT0 | grep -o percentage:.* | grep [0-9]*% | grep -o [0-9]*