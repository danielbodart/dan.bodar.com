---
title: "Rebuilding FFmpeg with NVENC and AAC on Ubuntu"
date: 2016-06-07T21:11:42Z
slug: "rebuilding-ffmpeg-with-nvenc-and-aac-on-ubuntu"
categories:
  - Uncategorized
comments:
  - {"author":"W","email":"DW@DS.COM","url":"http://WD","date":"2016-11-14T12:53:05Z","content":"does not work nvenc doesn't show on the encoder options on obs-studio","parent":0}
  - {"author":"dan","email":"dan@bodar.com","url":"","date":"2016-11-14T19:14:11Z","content":"Have you tried rebuilding obs-studio, maybe it's statically linked","parent":18831}
  - {"author":"Raymond Kimathi","email":"raymond@poravim.co.ke","url":"http://poravim.co.ke","date":"2016-12-15T17:32:15Z","content":"You are my hero bro ! Tried all sorts of tutorials for building stable FFmpeg in $HOME/bin. Your tutorial is the best !! Now my Geforce GM206 GTX 950 Is happy. Thanks !","parent":0}
  - {"author":"Jack Dawson","email":"djo-coeur@live.fr","url":"","date":"2017-03-16T18:34:51Z","content":"Thank you. It works perfectly with Ubuntu 16.04. I replaced the Nvidia SDK by the newer 7.1 and it still worked with my GTX 750ti","parent":0}
  - {"author":"Rene Ramirez","email":"rene.szabo@gmail.com","url":"","date":"2017-04-06T01:37:01Z","content":"Amazing , thank you !","parent":0}
  - {"author":"Zebediah Boss","email":"zebedee.boss@gmail.com","url":"","date":"2018-03-17T23:35:38Z","content":"Absolutely brilliant, my go to site for all buntu installs.  Thank You","parent":0}
  - {"author":"Zebedeeboss","email":"zebedee.boss@gmail.com","url":"","date":"2018-04-04T08:17:02Z","content":"My go to Nvenc install guide - never fails - Thank You","parent":0}
---

This mini guide will show you how to rebuild the exact version of FFmpeg that comes with your version of Ubuntu and just add support for NVidia GPU encoding via the NVENC API and ACC via libfdk\_aac.

```

# Download and unzip the NVIDIA Video Codec SDK from https://developer.nvidia.com/nvidia-video-codec-sdk
wget https://developer.nvidia.com/video-sdk-601
unzip nvidia_video_sdk_6.0.1.zip

# Copy the headers files from the SDK so FFmpeg can find them
sudo cp nvidia_video_sdk_6.0.1/Samples/common/inc/*.h /usr/local/include/

# Make sure "Source code" is enabled in System Settings... -> Software & Updates
# Download the build dependencies for FFmpeg
sudo apt-get build-dep ffmpeg

# Install libfdk_aac
sudo apt-get install libfdk-aac-dev

# Download the source for the exact version of FFmpeg you already have installed (not as root)
apt-get source ffmpeg

# Go into the ffmpeg source you just downloaded
cd ffmpeg-2.8.6

# Find out the exact command the ffmpeg was originally built with
ffmpeg -buildconf

# Copy the single line "configuration:" and pass it to ".configure" but add "--enable-nonfree --enable-nvenc --enable-libfdk-aac" on the end
# Mine looks like this:
./configure --prefix=/usr --extra-version=1ubuntu2 --build-suffix=-ffmpeg --toolchain=hardened --libdir=/usr/lib/x86_64-linux-gnu --incdir=/usr/include/x86_64-linux-gnu --cc=cc --cxx=g++ --enable-gpl --enable-shared --disable-stripping --disable-decoder=libopenjpeg --disable-decoder=libschroedinger --enable-avresample --enable-avisynth --enable-gnutls --enable-ladspa --enable-libass --enable-libbluray --enable-libbs2b --enable-libcaca --enable-libcdio --enable-libflite --enable-libfontconfig --enable-libfreetype --enable-libfribidi --enable-libgme --enable-libgsm --enable-libmodplug --enable-libmp3lame --enable-libopenjpeg --enable-libopus --enable-libpulse --enable-librtmp --enable-libschroedinger --enable-libshine --enable-libsnappy --enable-libsoxr --enable-libspeex --enable-libssh --enable-libtheora --enable-libtwolame --enable-libvorbis --enable-libvpx --enable-libwavpack --enable-libwebp --enable-libx265 --enable-libxvid --enable-libzvbi --enable-openal --enable-opengl --enable-x11grab --enable-libdc1394 --enable-libiec61883 --enable-libzmq --enable-frei0r --enable-libx264 --enable-libopencv --enable-nonfree --enable-nvenc --enable-libfdk-aac

# Now build it
make

# And finally install it over the original
sudo make install

```