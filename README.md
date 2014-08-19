pandora
=======

Export all pandora stations + faves to JSON

To get your username, visit http://www.pandora.com/profile and check the URL.
Also make sure your profile privacy is set to public.

## How it works

It uses the [gopiano](github.com/cellofellow/gopiano) library to connect to
pandora and download the station list.

It then loops through each station and grabs the songs that have been thumbed
up.

## How to use it

    $ pandora [email] [password]
    
    $ pandora stayradiated hunter5
    
        [
            {
                "name": "Pink Floyd Radio",
                "songs": [
                    {
                        "name": "...",
                        "artist": "..."
                    }
                ]
            }
        ]
