# Frida Handlers

Handlers for Monodraw functions for [Frida](https://frida.re/).

Currently, all handlers are auto-generated handlers from `frida-trace`, designed to trace all calls to the setter functions on instances of the base class used by all objects that are saved in the output file (`HTModelPersistentObject`).

## Usage
With Frida installed, Monodraw running _as the same architecture as the target Python process_ (so you must be using a native ARM version of Python if using native ARM Monodraw), and System Integrity Protection disabled on your machine, run the following command:
```fish
frida-trace -m '-[HTModelPersistentObject set*]' Monodraw
```