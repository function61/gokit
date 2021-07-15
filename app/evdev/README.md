Linux [evdev](https://en.wikipedia.org/wiki/Evdev) constants and some utilities.


Direct use of evdev vs. libinput
--------------------------------

In most cases it's a good idea to
[use libinput directly](https://who-t.blogspot.com/2018/07/why-its-not-good-idea-to-handle-evdev.html).

However, there may be some low-level cases where you need to interface with evdev directly.
