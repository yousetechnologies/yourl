# yourl

Fast and lightweight URL shortener written in Golang.

A URL shortner takes long URLs and squeezes them into fewer characters to make a link that is easier to share, tweet, or email to friends.

The URL Shortener API allows you to develop applications that interface with this program. You can use simple HTTP methods to create, inspect, and manage yours short URLs from your desktop, mobile, or web application.

# Goals

+ Fully managed by API REST FULL
+ Properly use of redirect codes of the http protocol
+ Persist data in a NoSQL database (read http://kkovacs.eu/cassandra-vs-mongodb-vs-couchdb-vs-redis)
+ ...

# API

## Insert
### Request

Do a POST request to create a new short URL.
```
curl http://_{your-short-domain}_/v1/url -d '_{long-url}_'
```
If successful, this method returns a response body with the following structure:
```json
{
    "id": "http://_{your-short-domain}_/XC4tkM",
    "longUrl": "_{long-url}_"
}
```

## Get

Do a GET request to retrieve informations for a specified short URL.
```
curl http://_{your-short-domain}_/v1/url/_{short-url}_
```
If successful, this method returns an Url resource in the response body.

## List
_still thinking..._

## No XML, Just JSON

## etc

