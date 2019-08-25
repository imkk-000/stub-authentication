# Stub Authentication

## Acceptance Test

| Name | Username | Password | Status Code | Response Body |
|---|---|---|---|---|
| login passed | imkk-000 | 1mkkn@ja* | 200 | `{"status": "ok"}` |
| login failed with wrong password | imkk-000 | wrongP@ssw0rd | 401 | `{"status": "wrong password"}` |
| login failed with not existing username | fakeusername | noempty | 401 | `{"status": "not existing username"}` |

## API Provider

### Post: /login

#### Request

```text
Request Type: Post Form
Request Body: username=imkk-000&password=1mkkn@ja*
```

#### Response

```text
Response Type: JSON
Response Body: {"status": "ok"}
```
