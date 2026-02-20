---
title: "Capsper: I turned CapsLock into a ghost writer"
date: 2026-02-19T12:00:00Z
slug: "capsper-push-to-talk-voice-dictation-for-linux"
categories:
  - linux
  - zig
  - voice
  - gpu
tags:
  - capsper
  - whisper
  - dictation
---

Does CapsLock annoy you? Ever wished it actually did something useful instead of SHOUTING AT PEOPLE BY ACCIDENT?

[Capsper](https://github.com/danielbodart/capsper) is push-to-talk voice dictation for Linux. Hold CapsLock, speak, release. Text appears wherever your cursor is. No cloud, no subscription, no Electron app phoning home. Just your GPU doing what GPUs were meant to do.

## The pitch

~2 second latency between speaking and words appearing on screen. Runs entirely on your machine. Works on both X11 and Wayland. Single binary, no dependencies to manage, no Python runtime, no Docker container. You download it, run the installer, and CapsLock becomes useful for the first time in its miserable existence.

## Why does this exist?

I wanted voice dictation on Linux that didn't suck. The options are either cloud-based (privacy, latency, subscriptions) or local Python-based tools that shell out to a dozen different utilities, run slow, and break every time your desktop environment sneezes. I wanted something that just *worked* — fast, efficient, hold a key, talk, release, done.

So I wrote a single [Zig](https://ziglang.org/) binary that handles the entire pipeline: keyboard grab via evdev, audio capture via PipeWire, GPU transcription via [whisper.cpp](https://github.com/ggml-org/whisper.cpp), and text injection via uinput. No xdotool. No ydotool. No keyd. No Python. No external anything.

## How it works

1. A single Zig binary grabs your keyboard via evdev, intercepts CapsLock as push-to-talk
2. Audio is captured directly via PipeWire while the trigger key is held
3. Incremental transcription runs on the GPU with VAD (voice activity detection on CPU) and token accumulation for consistency
4. Transcribed text is injected as keystrokes via uinput into the focused window

That's it. One process, one binary, zero external tools.

## The nerdy bits

This isn't just "whisper.cpp with extra steps." The interesting part is how the streaming works.

Most whisper-based tools either wait for you to finish speaking and then transcribe the whole thing, or they re-transcribe the entire audio buffer every cycle and diff the output — which causes words to flicker and change as context shifts. Capsper does neither.

Instead, it implements [AlignAtt](https://aclanthology.org/2023.findings-emnlp.744/) — a technique from simultaneous speech translation research. Specifically from [SimulStreaming](https://github.com/ufal/SimulStreaming) at Charles University, which won the [IWSLT 2025 Simultaneous Speech Translation Shared Task](https://aclanthology.org/2025.iwslt-1.1/). The core idea: rather than using whisper.cpp's high-level API, manually drive the mel spectrogram → encode → decode pipeline token by token, and introspect the decoder's cross-attention weights at each step. The attention pattern tells you where in the audio the model is "looking" — when attention drifts past the end of available audio or jumps backwards, that's the signal to stop emitting and wait for more speech.

The result is streaming transcription that's both fast and stable. Words appear as you speak them and don't change after they're emitted.

### Token accumulation

Confirmed tokens are committed as a forced decoder prefix for the next cycle, so the model never contradicts itself. It sees its own previous output as given, builds KV cache state, then continues generating from where it left off. No re-decoding, no instability.

When the 30-second sliding window trims old audio, the corresponding tokens get demoted from "forced output" to "conditioning context" rather than being deleted. The model treats them as hints instead of constraints. This two-tier approach prevents the hallucination cascades that plagued earlier attempts.

### Other things going on under the hood

- **CPU-side VAD** — Silero voice activity detection runs on the CPU (~5ms per check) while whisper runs on the GPU. No context switching overhead for the frequent silence checks.
- **Incremental mel spectrograms** — only computes FFT for new audio samples, caching raw mel frames between cycles. The mel filterbank, Hann window, and FFT are all implemented in pure Zig.
- **N-gram repetition guard** — monitors for repeating token patterns up to 64 tokens long. Three consecutive repeats = discard and stop. Catches both "the the the..." loops and longer phrase hallucinations.
- **Domain terms** — feed it a file of jargon (Kubernetes, gRPC, PostgreSQL) and it biases transcription toward your vocabulary without overriding acoustic evidence.

## Performance

On an RTX 5070 Ti with the large-v3-turbo model (q5_0 quantized):

- **~140ms** per transcription cycle
- **~1 second** between word emissions
- **~1.7 GB VRAM** (vs 5-10+ GB for SimulStreaming's approach with a 9B-parameter LLM)
- **Near-zero idle** — no CPU or GPU activity when you're not speaking

## Why Zig?

I wanted deterministic memory management (no GC pauses during transcription), direct C FFI without overhead (whisper.cpp exposes a C API and PipeWire is a C library), and a single binary with minimal runtime dependencies. Zig gives you all of that plus comptime, which turns out to be incredibly useful when you're building mel spectrograms and attention analysis pipelines.

Also it compiles fast. The whole binary builds in about 2 seconds. After years of fighting build times in other languages, this still makes me unreasonably happy.

## Try it

```bash
mkdir capsper && cd capsper
curl -fSL https://github.com/danielbodart/capsper/releases/latest/download/capsper-linux-x86_64.tar.gz | tar -xz
./install.sh
```

You need Linux, an NVIDIA GPU (GTX 1650 or newer), and PipeWire. The installer handles everything else — CUDA libraries (~600 MB), model downloads (~574 MB), permissions, systemd service, microphone detection. Budget about 1.2 GB of disk space total.

Then just:

```bash
systemctl --user start capsper.service
```

Hold CapsLock and speak. Release to stop. That's the whole UX.

The code is on [GitHub](https://github.com/danielbodart/capsper). Issues, PRs, and questions welcome.
