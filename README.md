# GoCook
GoCook is intended to be a simple application to record recipes and build meal plans. 

It might even calculate a shopping list for you and let you collaborate on meal plans with a partner!

This project isn't really anything serious - it exists largely for the author to learn Golang.

# Getting Started
* Install Go
* Clone the project into your GOPATH
* Install go vendor
```
$ go get github.com/kardianos/govendor
```
* Install dependencies
```
$ govendor fetch
```
* Run
```
$ go run src/main.go 
```

# What It Does right Now
* A restful API using `gin-gonic`
  * An attempt at creating a domain orientated project structure, where each logical domain provides its own router and store (handy for splitting the project up in the future when it really takes of, is worth millions and needs to be split into a microservices architecture)
* The start of authentication using oAuth 2
  * Everyone likes JWTs
  * The most popular library has been used `github.com/dgrijalva/jwt-go`
  * Unsupported functionality has been added to support Auth0 (probably works for Okta too): 
    * A mechanism for using JWKS (cached in memory)
    * The construction of RSA256 Public Keys from the modulus and exponent
    * Validation of standard claims `exp`, `iss` and `aud`
  * The above is made available as middleware for gin  
* An abstraction of `http` in `jsonHttp`
  * Making API requests in golang seems to have a fair old chunk of boiler plate
  * `jsonHttp` wraps `http` and abstracts the following:
    * Makes the request (with a sensible timeout)
    * Reads the body into []byte
    * Binds the bodies []byte to a struct
    * "logs" errors and passes them back down to the caller
  * Chances are there is a way better library for doing this but this does give quite a nice interface (even if it lacks referential transparency):
```
error := jsonHttp.Get(url, &structToBind);
```

# Endpoints
```
    GET: /ping <- replies pong
    
    GET: /recipes <- returns all recipes in memory
    GET: /recipes/:id <- returns one recipe
    POST: /recipes <- store or update a recipe in memory
    {
    	"recipe": {
    		"id": null <- provide null to store a new recipe, a UUIDv4 to update one
    		"name": "my amazing recipe"
    		"url": "http://some.recipes.com/amazing"
    	} 	
    }
```  

# Things to Watch Out For
* This has been authored on a Winblowz box
  * I've not checked that the .gitattributes is correct yet - so watch your line endings!
  * I've not checked that when the project is built on a Unix box it actually works
* There are no unit tests
* There are no integration tests
* There are basically no automated tests of any nature  
* There are probably many Golang antipatterns or GOTCHYAs I've not found yet

# Now GoCook!
![alt Silly Gif](https://thumbs.gfycat.com/WarlikeQuarrelsomeBuck-max-1mb.gif)
