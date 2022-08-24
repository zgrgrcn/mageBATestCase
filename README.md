# Mage BA Test Case API - Ozgur Gurcan

## LOCAL TEST

### Go Module Init
```bash
git clone https://github.com/zgrgrcn/mageBATestCase
cd mageBATestCase
go run main.go
```
You don't need to have any local database, Thanks to Mongo Atlas! Create a .env file in the root directory and add the following lines:
```bash
DATABASE_URL=mongodb+srv://
TOKEN_KEY=key
PORT=:8080
```

### Run
```bash
go run main.go
```
The default port is 8080 go to http://localhost:8080/swagger/index.html or https://shielded-springs-54397.herokuapp.com/swagger/index.html to see swagger ui.
#### For /user endpoint you don't need to add any token, but for /leaderboard and /endgame endpoints you need to add a token to the header as Bearer. You can get token from /user/login endpoint.


## ONLINE TEST
All APIs  are online(thanks to GitHub CI/CD) and you can test them with swagger ui or postman collection.
https://shielded-springs-54397.herokuapp.com/swagger/index.html (it a is free tier, so don't be surprised if it is slow)
#### For /user endpoint you don't need to add any token, but for /leaderboard and /endgame endpoints you need to add a token to the header as Bearer. You can get token from /user/login endpoint.



## API Documentation
### SWAGGER
Generate Swagger Documentation 
(Needs to install swag, Docs already exist, so pass this step if you don't want to generate docs again)
```bash
swag init
```
## Postman
There is a postman collection in the root directory of the project. You can import it to your postman and test the API.
The file name is Mage.postman_collection.json.

# for security reasons, these API/DB/Heroku  will expire on the 24 of September.

