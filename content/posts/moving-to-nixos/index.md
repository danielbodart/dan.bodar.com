---
title: "Moving to NixOS"
date: 2026-09-18T09:00:00Z
draft: true
slug: "moving-to-nixos"
categories:
  - linux
  - agents
tags:
  - nixos
  - nix
  - linux
  - containers
---

I switched to [NixOS](https://nixos.org/) two weeks ago, after about five years of being too scared of the learning curve. With Claude alongside me that curve has basically gone, and I can honestly say it's the best computing decision I have ever made.

I've never enjoyed my PC like this. It feels rock solid, and it feels 100% *mine*.

## Building it up in layers

I didn't do it all at once. Each layer took something I used to manage by hand and turned it into config.

**NixOS** gave me an immutable, declarative operating system. Every change is a new generation, and rolling back is instant, either with `nixos-rebuild switch --rollback` or by picking an older generation from the boot menu. The config is a functional program, the module system type checks every option, and I can build or test it before I switch to it (`nixos-rebuild build`, `test` or even `build-vm`) ([NixOS manual](https://nixos.org/manual/nixos/stable/#sec-changing-config)).

**[Home Manager](https://github.com/nix-community/home-manager)** does the same for my user profile: dotfiles, apps, editor plugins, the lot. It's all declared and versioned, and it all rolls back too.

**[sops-nix](https://github.com/Mic92/sops-nix)** handles secrets. They live encrypted in the same repo as everything else and are decrypted into `/run/secrets` when the system activates, so the config is complete without ever putting a plaintext secret in git.

**Next up is [disko](https://github.com/nix-community/disko)**, "declarative disk partitioning and formatting using nix". Once that's in, even the disk layout will be part of the config.

## The 20-minute rebuild

This is where it all paid off. I rebuilt my desktop, in about 20 minutes between meetings, from the same config as my laptop, and it just worked. Every app came back already installed, with every plugin and every config. The only thing I had to do was log back in to Google!

## Things you get for almost nothing

What I love about Nix is how much just falls out of the design for virtually free. When you get the foundations right, good things keep turning up.

Every package lives in its own path in `/nix/store`, so every binary is isolated and you can have different versions of the same library side by side without any issues. Everything is immutable, so nothing can break anything else.

It also means I can make every tool work exactly how I want. Patching source code or patching a binary is the same level of difficulty, and it's 100% repeatable: the patch is part of the config, so it's reapplied automatically every time the package is rebuilt or upgraded.

## Nearly Docker, nearly a VM

Without doing anything special, plain Nix gets you most of the way to Docker. You get isolated binaries, side-by-side versions and immutability without a container in sight.

Add Nix's built-in container support and you get most of the way to a VM, and in some ways it's better than Docker. A Docker image is a stack of layers you have to download, often hundreds of megabytes of them. A Nix container shares the host's `/nix/store`, so every package in the image sits next to every other package on the host. If the image you want is a subset of what you already have installed, it's **zero bytes** to download. Building the image (the bit you'd normally wait minutes for in CI) takes milliseconds when everything is already in the store.

Docker-based sandboxes, by contrast, use a *lot* of disk space per sandbox by default. When I told someone at work about zero-byte containers, their reaction was "OMG, I need this".

Containers share the host kernel, so they won't protect you from a kernel exploit. If you need to close that last gap, [microVMs](https://github.com/microvm-nix/microvm.nix) are just a step away, and you don't have to give up the disk savings to get there. microvm.nix can share the host's `/nix/store` into the guest over virtiofs or 9p ([docs](https://microvm-nix.github.io/microvm.nix/shares.html)), as long as you pick a hypervisor that supports it (QEMU or cloud-hypervisor rather than Firecracker, for example).

## Safe playgrounds for agents

All of this makes working with agents feel *safe*. If an agent trashes something on the host, I roll back. Better still, it doesn't need to touch the host at all: declarative, throwaway containers mean every agent can get its own sandbox built from the same config as everything else.

Because the store is shared, spinning up a throwaway sandbox costs almost nothing. It was cheap enough that I ended up writing [flong](https://github.com/danielbodart/flong), a tiny launcher that starts one in about 100ms.

## Obviously right

Nix is one of those things that is just *obviously right* once you see it. An immutable, versioned, reproducible machine, where every change is a revision and every patch reapplies itself, is how computers should have worked all along.

If, like me, you've been putting it off because of the learning curve, grab an agent and just start. Five years was far too long to wait.
