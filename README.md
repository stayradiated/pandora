Pandora
=======

A simple CLI wrapper for https://github.com/stayradiated/pandora.
Use it to quickly extract a list of your favorite songs from Pandora.

**Important Note**: This does require you to enter your username and password for Pandora.

## Install

```
go install github.com/stayradiated/pandora
```

Or download a binary from the 'releases' tab.

## Usage

Depending on how many stations you have, it could take a while to run. I have
around 750 likes over 95 stations and it takes around 5-10 seconds.

```
pandora -u [username] -p [password]

// example
$ pandora -u john@smith.com -p hunter2
Pink Floyd -- Wish You Were Here

// example as json
$ pandora -u john@smith.com -p hunter2 -json
[{"name": "Pink Floyd Radio","songs": [{"name": "Wish You Were Here","artist": "Pink Floyd"}]}]
```
