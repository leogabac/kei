# kei

A compilation of miscellaneous automation scripts thrown into a CLI.

This project is for personal use.

## Why?

Over the years, I have compiled many miscellaneous shell scripts that were 
thrown into my dotfiles directory, and then forgotten by god.

They were: undocumented, confusing, and unmaintainable. The perfect recipe for
disaster and duplication.

After getting mad at past me, I decided to compile all the _actual_ useful 
scripts into a small CLI, and take the opportunity to learn Go along the ride.

## Requirements

* the `go` compiler
* `imagemagick`

## Features

* Blur and image
```sh
kei media blur <input> <outut> -r <radius>
```
* Dim an image
```sh
kei media dim <input> <outut> -b <brightness> -d <dimming>
```

