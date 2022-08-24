# Mage BA Test Case API - Ozgur Gurcan

## LOCAL TEST

### Go Module Init
```bash
go mod init github.com/zgrgrcn/mageBATestCase
```
You don't need to any local database, Thanks to Mongo Atlas!

### Run
```bash
go run main.go
```
Default port is 8080 go to http://localhost:8080/swagger/index.html to see swagger ui

## API Documentation
### SWAGGER
Generate Swagger Documentation 
(Needs to install swag, Docs are already exist, so pass this step...)
```bash
swag init
```
## Postman
There is a postman collection in the root directory of the project. You can import it to your postman and test the API.
File name is Mage.postman_collection.json.


## ONLINE TEST
All api's are online(thanks to github CI/CD) and you can test them with swagger ui or postman collection.
https://shielded-springs-54397.herokuapp.com/swagger/index.html (it is free tier, so don't be surprised if it is slow)


# for security reasons, these api/db/heroku will expire at September 24.