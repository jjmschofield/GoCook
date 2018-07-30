# GoCook
GoCook is intended to be a simple application to record recipes and build meal plans. 

It might even calculate a shopping list for you and let you collaborate on meal plans with a partner!

This project isn't really anything serious - it exists largely for the author to learn Golang.

# Getting Started
* Install Go
* Clone the project into
```
$GOPATH/src/github.com/jjmschofield/GoCook
```
* Install go vendor
```
$ go get github.com/kardianos/govendor
```
* Install dependencies
```
$ govendor sync
```
* Run
```
$ go run main.go
```

# API Framework
The project currently makes use of `github.com/gin-gonic/gin` as an API framework, picked largely as it is very minimalist and - providing the best learning opportunities for the author.

# Authentication
GoCook uses OAuth 2.0, provided (at the moment from Auth0) - `security/auth` is the package which you are looking for.

Authentication will validate that a user has a valid session and should have access to the API by using JWT signed with a RSA256 private key.

Switching to an alternate token issuer / IdP shouldn't be too much of a problem should you need to. 

## Authenticating an API endpoint
`security/auth` offers up `IsAuthenticatedMiddleware`. Use it as middleware on any endpoint or router that you want to force a valid user session for.

## JWTs and JWKS and RSA256 oh my!
The most popular JWT library on GitHub has been used to parse JWTs namely `github.com/dgrijalva/jwt-go`

Unfortunately the current stable version (3.2.0) of the library has poor examples/documentation and a few limitations which have been worked around:

* JWKS support
  * Support for retrieving a key from  a JWKS has been added using the `kid` of the JWT
  * Keys are cached (in a pretty naive and non-robust manner)
     * Whenever a `kid` in a token can't be found the cache is resynced           
* RSA256 support
  * Support for creating an RSA 256 public key from the exponent and modulus of the retrieved JSON web key has been added 
* Extended support for standard JWT claims
  * `exp` is ensured to be in the future
  * `iss` is ensured to be whatever is set in the config
  * `aud` is ensured to be whatever is set in the config
  
Don't worry, you get all of this for free when calling `IsAuthenticatedMiddleware`.  

# What It Does right Now
* A restful API using `gin-gonic/gin`
  * An attempt at creating a domain orientated project structure, where each logical domain provides its own router and store (handy for splitting the project up in the future when it really takes of, is worth millions and needs to be split into a microservices architecture)
* The start of authentication using oAuth 2 as gin middleware
  * Everyone likes JWTs
  * Unsupported functionality has been added to support Auth0 (probably works for Okta too)
  * The above is made available as middleware for gin  
* An abstraction of `http` in `net/jsonhttp`
  * Making API requests in golang seems to have a fair old chunk of boiler plate
  * `net/jsonhttp` wraps `http` and abstracts the following:
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
