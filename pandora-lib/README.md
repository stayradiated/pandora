pandora
=======

Export all your pandora stations and thumbed up tracks as JSON.

**Important Note:** This does require your username and password for Pandora.

## How it works

It uses the [gopiano](//github.com/cellofellow/gopiano) library to connect to
pandora and download the station list. It then loops through each station and
fetches the songs that have been thumbed up. This info is then processed and
printed out as JSON.

## How to install it

```
go get github.com/stayradiated/pandora
```

## How to use it

Checkout https://github.com/stayradiated/pandora for a simple CLI.
