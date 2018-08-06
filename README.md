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
* Run the database migrations (more on this later)

# Deployment
This project publishes to http://go-cook.herokuapp.com/ on each push to the master branch.

If you want to replicate this, simply fork this repo and point a Heroku app and pipeline to your fork. Alternatively use the CLI as described in https://devcenter.heroku.com/articles/deploying-go

Note: this isn't presently running via https in production - an SSL cert needs to be applied via Heroku

# Documentation
Swagger Documentation for the endpoints can be found at the following URLs:
* [UI](http://go-cook.herokuapp.com/swagger)
* [JSON](http://go-cook.herokuapp.com/swagger/swagger.json)

To refresh the documentation, pull down a binary for `go-swagger` and execute it eg:

```
$ curl -L https://github.com/go-swagger/go-swagger/releases/download/0.15.0/swagger_windows_amd64.exe > swagger.exe
$ swagger generate spec -o api/public/swagger.json
```

# API Framework
The project currently makes use of `github.com/gin-gonic/gin` as an API framework, picked largely as it is very minimalist which provides a good learning opportunity for the author.

# Authentication
GoCook uses OAuth 2.0, provided at the moment from Auth0 - `common/auth` is the package which you are looking for.

Authentication will validate that a user has a valid session and should have access to the API by using JWT signed with a RSA256 private key.

Switching to an alternate token issuer / IdP shouldn't be too much of a problem should you need to. Just configure the issuer, audience and JWKS endpoint in `cook.json` (don't worry, you can publically share this information). 

## Authenticating an API endpoint
`common/auth` offers up `AuthenticationMiddleware`. Use it as middleware on any endpoint or router that you want to force a valid user session for.

This will also set `token *jwt.Token` and `userId string` to the gin context of the request, to remove any requirement to parse the token again. 

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
  
Don't worry, you get all of this for free when calling `AuthenticationMiddleware`.  

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

## Secrets
Secrets are expected to be provided as params on the CLI. 

This is a simple solution which will be exchanged for either encrypting the strings with a cert, KWS/Key Vault in the future.

# Database
Presently the project is setup to use Postgres as a data store. 

Database management is carried out through database migration files (`github.com/golang-migrate/migrate`) allowing for traceability of changes and super fast provisioning of new stores.

## Running Database Migrations 
* Grab the [pre-built CLI](https://github.com/golang-migrate/migrate/releases) for migrate based on your architecture
```
$ cd db
$ curl -L <pre-built binary> | tar xvz
```
* Run `up` to go forward
```
$ migrate -database postgres://<user>:<password>@localhost:5432/gocook?sslmode=disable --path migrations up
```
* If you want too, `down` to go back
```
$ migrate -database postgres://<user>:<password>@localhost:5432/gocook?sslmode=disable --path migrations down
```

### Notes
* For windows you need build version 17063 or greater to get easy curl/tar 
* You'll also need run a few more commands due to the lack of support for pipes:
```
$ cd sql
$ curl -L https://github.com/golang-migrate/migrate/releases/download/v3.4.0/migrate.windows-amd64.exe.tar.gz > migrate.tar.gz
$ tar -xvz -f migrate.tar.gz
$ mv migrate.windows-amd64.exe migrate.exe
$ rm migrate.tar.gz
```
* If you prefer just download the file manually and unpack with 7zip or similar

## Creating a Migration
* Create a migration:
```
$ cd sql
$ migrate create --ext .sql --dir migrations <name>
``` 
* Add the changes you want in `sql/migrations/<date/time>_<name>.up.sql`
* Add the back out in `sql/migrations/<date/time>_<name>.down.sql`
* A few rules to keep you safe
  * **Do** Use transactions  
  * **Do** test your backout
  * **Don't** modify existing migrations - always create a new one
 
 ## Design Notes
 * We are not using an ORM (though `gorm` looks great)
   * Highly contentious but there are reasons...
   * We take complete control of database migrations
   * We take complete control of the database arch and capabilities
   * We hopefully get a performance boost (we can definitely profile database performance a bit more easily) 
   * We do loose quite a lot of development speed (but some other considerations counter this)
   * We unlock the next point...
 * We don't write SQL in application code - stored procedures all the way
   * We push all storage related logic out of the application, effectively turning postgres into its own kind of API 
   * We get a security boost against injections (but we still need to validate inputs!)
   * We unlock a more sophisticated permissions model (access can be granted to procedures, not just dbs/schemas/tables)
   * We hopefully get a performance boost
   * We can test these using `pgTap`
 * We treat Postgres as a schemaless store with super powers
   * We make heavy use of JSONB columns
   * This allows us to get schemaless models (in a NoSQL kind of way) but retain all the goodness of a relational store 
 * All these points also serve as a really good learning exercise for the author, especially if it needs to change later :)
# What It Does right Now
* A restful API using `gin-gonic/gin`
  * Endpoints for CRUD on basic recipe models  
  * `recipe` package should be splittable 
* JWT based authentication using OAuth2 as gin middleware
  * Everyone likes JWTs
  * Unsupported functionality has been added to support Auth0 (probably works for Okta too)
  * The above is made available as middleware for gin  
* An abstraction of `http` in `common/jsonhttp`
  * Making API requests in golang seems to have a fair old chunk of boiler plate
  * `common/jsonhttp` wraps `http` and abstracts the following:
    * Makes the request (with a sensible timeout)
    * Reads the body into []byte
    * Binds the bodies []byte to a struct
    * "logs" errors and passes them back down to the caller
  * Chances are there is a way better library for doing this but this does give quite a nice interface (even if it lacks referential transparency):
```
error := jsonHttp.Get(url, &structToBind);
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
