# Research Tools

Currently, all research is performed via [Frida](https://frida.re/) to enable instrumenting the running Monodraw process.

## Usage

- Install [uv](https://docs.astral.sh/uv/getting-started/installation/) and npm
- Duplicate your `Monodraw` application and rename the copy to `Monodraw (Unsigned)`
- Re-sign the duplicate application with your own key
    - `codesign --deep --force --sign - /Applications/Monodraw\ \(Unsigned\).app`
    - Without resigning, disabling System Integrity Protection is required to attach Frida to the app
- Launch your unsigned version of the app
- Run the following commands to launch Frida

```bash
npm run build
npm run attach
```
