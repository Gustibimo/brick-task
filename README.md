## Onebrick Task

The task is to create 3 endpoints for payments processing. The endpoints are:

1. Validate bank account

```asciidoc
POST /api/v1/validate-bank-account
```

Request:
```asciidoc
{
    "account_number": "1234567890",
    "bank_code": "mandiri"
}
```

Response:
```asciidoc
{
    "status": "success",
    "message": "Bank account is valid"
}
```