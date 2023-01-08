# KrakenAPI

A program made to archive informations from Kraken API and presenting them

## Functionnalities

The app will have to :
- Store server status and timing
- Trading pairs & their informations
- Display and store the data in files and on a Postgres DB
- Write an API route to download the file(s)
- Write an API route to see the informations stored in the DB on a web page

The users will be able :
- To archive all cryptos available on Kraken
- To sort informations for all pairs
- To see the data on a web page
- To download the files every hour

## Go

 We'll have to use those notions :
 - Webserver HTTP
 - Goroutine & Waitgroup
 - Channels

## Docker
This whole application has to be Dockerised so that its deployment can be easy