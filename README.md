# Gtk Bindings


## Debian/Ubuntu

```bash
sudo apt-get install libgtk-3-dev install libcairo2-dev libglib2.0-dev
```

## Mac

```bash
brew install gtk+3
PKG_CONFIG_PATH=/opt/X11/lib/pkgconfig:`brew --prefix gtk+3`/lib/pkgconfig go get -u -v github.com/gotk3/gotk3/gdk
```

For gtk4:

```bash
brew install gtk4 gtk+3 gobject-introspection pkg-config
```

Build the binary:

```bash
PKG_CONFIG_PATH=/opt/X11/lib/pkgconfig go build .
```

Run the binary:

```bash
ASSUME_NO_MOVING_GC_UNSAFE_RISK_IT_WITH=go1.18 ./solmate-go-gui
```
