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

    go install github.com/stayradiated/pandora

## How to use it

Depending on how many stations you have, it could take a while to run. I have
around 750 likes over 95 stations and it takes around 3.5 seconds.

    $ pandora [email] [password]
    
    $ pandora stayradiated hunter5
    
        [
            {
                "name": "Pink Floyd Radio",
                "songs": [
                    {
                        "name": "Wish You Were Here",
                        "artist": "Pink Floyd"
                    },
                    ...
                ]
            },
            ...
        ]
