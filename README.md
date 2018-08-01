# GoCook
GoCook is intended to be a simple application to record recipes and build meal plans. 

It might even calculate a shopping list for you and let you collaborate on meal plans with a partner!

This project isn't really anything serious - it exists largely for the author to learn Golang.

# Getting Started
## Building and Running
* Install Go
* Clone the project into
```
$GOPATH/src/github.com/jjmschofield/GoCook
```
* Run
```
$ cd $GOPATH/src/github.com/jjmschofield/GoCook
$ go run main.go
```

## Setting Up The Database
GoCook makes use of Postgres for persistent storage, with database migrations executed by `github.com/golang-migrate/migrate`

* Install Postgres
* Create a database called `gocook`
* Run database migrations

### Running Database Migrations 
* Grab the [pre-built CLI](https://github.com/golang-migrate/migrate/releases) for migrate based on your architecture
```
$ cd db
$ curl -L <pre-built binary> | tar xvz
```
* Run the migrations
```
$ migrate -database postgres://<user>:<password>@localhost:5432/gocook?sslmode=disable --path migrations up
```

#### Notes
* For windows you need build version 17063 or greater to get easy curl/tar 
* You'll also need run a few more commands due to the lack of support for pipes:
```
$ cd db
$ curl -L https://github.com/golang-migrate/migrate/releases/download/v3.4.0/migrate.windows-amd64.exe.tar.gz > migrate.tar.gz
$ tar -xvz -f migrate.tar.gz
$ mv migrate.windows-amd64.exe migrate.exe
$ rm migrate.tar.gz
```
* If you prefer just download the file manually and unpack with 7zip or similar



# Deployment
This project publishes to http://go-cook.herokuapp.com/ on each push to the master branch.

If you want to replicate this, simply fork this repo and point a heroku app and pipeline to your fork. Alternatively use the CLI as described in https://devcenter.heroku.com/articles/deploying-go

# Configuration
Configuration is provided through the popular `github.com/spf13/viper` package.

To add a config value - simply add it to `cook.json`.

To use a config value, simply use the interface supplied by viper eg: `viper.GetString(<your config key>)`

Presently only non-secure configuration values are supported - don't go adding secrets to `cook.json` unless you would like to share them with everyone on GitHub.

At deploy time you will want to copy `cook.json` or create a specific one in a `config/` directory relative to the `main.exe` eg:

```
config/
- cook.json
main.exe
```

# API Framework
The project currently makes use of `github.com/gin-gonic/gin` as an API framework, picked largely as it is very minimalist which provides a good learning opportunity for the author.

# Authentication
GoCook uses OAuth 2.0, provided at the moment from Auth0 - `security/auth` is the package which you are looking for.

Authentication will validate that a user has a valid session and should have access to the API by using JWT signed with a RSA256 private key.

Switching to an alternate token issuer / IdP shouldn't be too much of a problem should you need to. Just configure the issuer, audience and JWKS endpoint in `cook.json` (don't worry, you can publically share this information). 

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
  * `exp` is validated to be in the future
  * `iss` is validated to be whatever is set in the config
  * `aud` is validated to be whatever is set in the config
  
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
* There are no unit tests
* There are no integration tests
* There are basically no automated tests of any nature  
* There are probably many Golang antipatterns or GOTCHYAs I've not found yet

# Now GoCook!
![alt Silly Gif](https://thumbs.gfycat.com/WarlikeQuarrelsomeBuck-max-1mb.gif)
