# hexlink

A URL Shortening service with a hexagonal architecture written in Go

## Examples

### Shortening a URL

```sh
Send a Post Request to /api/shorten With a json body like:
{
"url": "http://www.google.com"
}
Note: The "http://" in the url is important TODO: improve this with http validation/autocorrection

```

### Following a shortened link

```sh
[URL]/[code]
www.hexli.nk/DcJltl3MR # Not a Real URL 
```

TODO:

+ Improve App Scalability
+ Improve ID Generation
+ Improve ID Management
+ Write a front end