# Simple Payment API
This API aims the needs between merchant and customer in order to do transaction. Thus, the functionality of the API includes:

-	Login: do log in customer and if customer does not exist, then got reject. 
-	Payment: customer that has been login can do the payment. No max & min limit amount for transfer. Transfer only do for registered customer. 
-	Logout: do logout for the logged in custome.
-   

All of the activity stored inside transaction table. 

This API were made by using [GoFiber](https://gofiber.io/).

## Requirement
- [Go](https://go.dev/) (>= 1.24)

## Installation
1. Clone the repository:
```sh
git clone https://github.com/FawwazN/mnc_backend.git
```
2. Navigate to the project directory:
```sh
cd mnc-backend
```
3. Setup Environment Variables
Need to create .env file based on .env.example by inputting these required variables:
```
SERVER_PORT=server port to run the API
JWT_SECRET=secret jwt key
CUSTOMER_DB=location of customers storage file in json
MERCHANT_DB=location of merchants storage file in json
TRANSACTION_DB=location of transactions storage file in json
SESSION_DB=location of session storage file in json
```
4. Install some dependencies and run the application
```sh
go mod tidy
go run main.go
```

## Additional information

Documentation about running the HTTP requests can be imported using `./test/mncpay-api.postman_collection.json` on Postman.