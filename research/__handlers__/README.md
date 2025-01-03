# Frida Handlers

Handlers for Monodraw functions for [Frida](https://frida.re/).

Currently, all handlers are auto-generated handlers from `frida-trace`, designed to trace all calls to the setter functions on instances of the base class used by all objects that are saved in the output file (`HTModelPersistentObject`).

## Usage
- Install [uv](https://docs.astral.sh/uv/getting-started/installation/)
- Duplicate your `Monodraw` application and rename the copy to `Monodraw (Unsigned)`
- Re-sign the duplicate application with your own key
  - `codesign --deep --force --sign - /Applications/Monodraw\ \(Unsigned\).app`
  - Without resigning, disabling System Integrity Protection is required to attach Frida to the app
- Launch your unsigned version of the app
- Run the following command to launch Frida

```fish
uv run --with frida-tools -- frida-trace -m '-[HTModelPersistentObject set*]' Monodraw
```
