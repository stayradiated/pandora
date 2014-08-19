pandora
=======

Export all pandora stations + faves to JSON

To get your username, visit http://www.pandora.com/profile and check the URL

## How it works

Get full station list

    http://www.pandora.com/content/stations?startIndex=0&webname=<USER>
    
    USER: your username
    
Get around five faves from a station
    
    http://www.pandora.com/content/station_track_thumbs?stationId=<STATION_ID>&page=true&posFeedbackStartIndex=<INDEX>&posSortAsc=false&posSortBy=date
    
    STATION_ID: id of station
    INDEX: start at 0 and increment with each call
