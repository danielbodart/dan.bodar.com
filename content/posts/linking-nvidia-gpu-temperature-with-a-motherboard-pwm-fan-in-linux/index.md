---
title: "Linking Nvidia GPU temperature with a motherboard PWM fan in Linux"
date: 2016-03-03T06:57:48Z
slug: "linking-nvidia-gpu-temperature-with-a-motherboard-pwm-fan-in-linux"
categories:
  - Uncategorized
  - ubuntu
  - memory
  - linux
---

So I have a closed loop water cooler on my GPU and I replaced the stock single speed fan with two Noctua PWM fans (doing push/pull) connected to my motherboard. Even though the fans can be seen by `lm_sensors` and `fancontrol` the Nvidia GPU does not appear as it's proprietary driver.

I knew `nvidia-settings` could be used to query the temperature (and the pump speed if you care) so I wrote a script to tie it all together:

```
#!/usr/bin/env bash

set -e

# This is the path to the PWM controlled fan (use lm_sensors/fancontrol to help you identify this)

fan=/sys/class/hwmon/hwmon1/pwm1

# Read https://www.kernel.org/doc/Documentation/hwmon/ for your PWM chip to find the correct values (I have a nct6792)

automatic=5

manual=1

# Temperature at which to run fan at 100% speed

max=80

# Re-enable automatic fan control on exit

trap "echo ${automatic} > ${fan}_enable; exit" SIGHUP SIGINT SIGTERM ERR EXIT

# Enable manual fan control

echo ${manual} > ${fan}_enable

function temperature() {

        nvidia-settings -q [gpu:0]/gpucoretemp -t

}

function fan_speed() {

        echo Setting FAN Speed to $1%

        echo $(((($1 * 255)) / 100)) > ${fan}

}

while true; do

        temp=`temperature`

        echo GPU Temperature: $temp

        if [ "$temp" -ge "$max" ] ; then

                fan_speed 100

        else

                fan_speed $(($temp + ((100 - $max))))

        fi

        sleep 1

done

```

Then make the script sudoable without password and run on login. Now my fans run at 600rpm at idle and go up to 1100rpm when running a GPU burn in tool.