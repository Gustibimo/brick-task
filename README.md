## Onebrick Task

### Setup and Run the application

1. Clone the repository
```bash 
git clone https://github.com/Gustibimo/brick-task.git
````

2. Download the dependencies
```bash
go mod tidy
```

3. Run the application
```bash
go run cmd/main.go
```


### Code Structure

The code is structured in a way that it is easy to understand and maintain. The code is divided into:

1. `handler` package contains the HTTP handlers for the endpoints. The handlers are responsible for parsing the request, calling the service and returning the response.
2. `usecase` package contains the business logic of the application. The usecase is responsible for validating the request, calling the repository and returning the response.
3. `repository` package contains the repository for the application. The repository is responsible for calling the external services and returning the response.
4. `models` package contains the models for the application. The models are responsible for defining the request and response structure for the application.
5. `gateway` package contains the external services for the application. The gateway is responsible for calling the external services and returning the response.

### DB Migration

This application use Goose for DB migration. The migration file is located in `db/postgres` directory. The migration is automatically run when the application is started.

Create migration file can be done by running the following command:
```bash
make migration-create name=[migration_name]
```

### Endpoints
The task is to create 3 endpoints for payments processing. The endpoints are:

1. Validate bank account

```asciidoc
POST /api/v1/payments/validate-bank-account
```

Request:
```asciidoc
{
    "account_number": "1234567890",
    "bank_code": "mandiri"
}
```

Header:
```asciidoc
Content-Type: application/json
```

Response:
```asciidoc
{
    "status": "success",
    "message": "Bank account is valid"
}
```

Error response:
```asciidoc
{
    "status": "failed",
    "message": "Bank account is invalid"
}
```

2. Disbursement

```asciidoc
POST /api/v1/payments/disburse
```

Request:
```asciidoc
{
    "account_number": "123456789
    "bank_code": "mandiri",
    "amount": 100000
}
```

Header:
```asciidoc
Content-Type: application/json
```

Response:
```asciidoc
{
    "status": "success",
    "message": "Disbursement is successful"
}
```

Error response:
```asciidoc
{
    "status": "failed",
    "message": "Disbursement is failed"
}
```

3. Payment callback to update payment status

```asciidoc

POST /api/v1/payments/webhook
```

Request:
```asciidoc
{
    "transaction_id": "1234567890",
    "status": "success"
}
```

Header:
```asciid
Content-Type: application/json
X-Signature-Key: [secret_key]
```

Response:
```asciidoc
{
    "status": "success",
    "message": "Payment status is updated"
}
```

For validating bank account, this service will call bank API to validate the bank account. The bank API is mocked and will return success response for any request. The bank API is mocked using https://65d34599522627d5010875bb.mockapi.io/api/v1/inquiry/2.

For disbursement process, this service will call bank API to process the disbursement. The bank API is mocked and will return success response for any request. The bank API is mocked using https://65d34599522627d5010875bb.mockapi.io/api/v1/doPayment/1.

For payment callback, this service will update the payment status based on the request. The service will validate the request using X-Signature-Key header (currently hardcoded in the service).