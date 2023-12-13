# Coal
Assuming you have downloaded Go from https://go.dev/doc/install or using your favorite package manager.

### Setup
Since Coal is using Firebase Realtime Database and Firebase Storage you need to setup a file called ```serviceAccount.json``` in the project root folder with this structure:

```
{
  "type": "service_account",
  "project_id": "coal-f8d25",
  "private_key_id": "YOUR-PRIVATE-KEY-ID",
  "private_key": "YOUR-PRIVATE-KEY",
  "client_email": "YOUR-CLIENT-EMAIL",
  "client_id": "YOUR-CLIENT-ID",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/firebase-adminsdk-6c4gj%40coal-f8d25.iam.gserviceaccount.com",
  "universe_domain": "googleapis.com"
}

```

### Start
Finally, to start the Coal API run this command:
```
$ go run cmd/main.go
```

