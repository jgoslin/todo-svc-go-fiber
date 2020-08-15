This is a slightly enhanced version of:
** youtube-go-fiber-introduction **
An introduction to Fiber an express-like library for golang
Youtube video: https://youtu.be/MfFi4Gt-tos
Github: https://github.com/EQuimper/youtube-go-fiber-introduction

In this version I have connected to a locally running mongodb server for my datastore. I have used the standard MongoDB driver.
I have also structured the project as below:
1. Database - In this package I have the configuration of the connection to mongo
2. Errors - This is a package containing constants for standard error messages
3. Todo - In this package is all the logic for the TodoAPI. This includes the Todo Model, Repository for storing todos, a todo service this is the API and last the routes configuration.