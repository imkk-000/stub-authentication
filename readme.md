# Stub Authentication

## Acceptance Test

| Name | Username | Password | Status Code | Response Body |
|---|---|---|---|---|
| login passed | imkk-000 | 1mkkn@ja* | 200 | `{"status": "ok"}` |
| login failed with wrong password | imkk-000 | 1mkkn@ja* | 401 | `{"status": "wrong password"}` |
| login failed with not existing username | fakeusername | noempty | 401 | `{"status": "not existing username"}` |
