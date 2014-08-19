pandora
=======

Export all pandora stations + faves to JSON

## How it works

Get full station list (where <USER> is your username):

    http://www.pandora.com/content/stations?startIndex=0&webname=<USER>
    
Get around five faves from a station
    
    http://www.pandora.com/content/station_track_thumbs?stationId=<STATION_ID>&page=true&posFeedbackStartIndex=<START>&posSortAsc=false&posSortBy=date
    
    <STATION_ID> Id of station
    <START> Start at 0 and increment with each call
