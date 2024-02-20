## Onebrick Task

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