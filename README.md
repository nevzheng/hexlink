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

## Docker Commands
```sh
# Creating an Image
docker build -t nevzh/hexlink
docker run nevzh/hexlink
docker push nevzh/hexlink

# Using Docker Compose
docker compose up
```


### Following a shortened link

```sh
[URL]/[code]
www.hexli.nk/DcJltl3MR # Not a Real URL 
```

## OpenAPI Generation
The OpenAPI Specification (OAS) defines a standard, language-agnostic 
Hexlink uses OAS 3.0.
[`oapi-codegen`](https://github.com/deepmap/oapi-codegen) is used to generate types from the provided specification.
We want to use go-kit to implement our server.
However, I couldn't find a publicly available, well documented, and supported code generator for OpenAPI that supports go-kit.

Not all auto-generated types will be used. I expect to refine the process 
and utilization of both go kit and open api.

### Generating Types from specification
1. Ensure that the `hexlink-schema` submodule is the desired version.
2. Install [`oapi-codegen`](https://github.com/deepmap/oapi-codegen)
3. Ensure that go binaries are included in your path
4. `oapi-codegen --config=codegen.cfg.yml hexlink-schema/hexlink.yaml`

TODO:

+ Improve App Scalability
+ Improve ID Generation
+ Improve ID Management
+ Write a front end
