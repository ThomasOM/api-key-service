# api-key-service
Simple Go app using Gin/GORM with API key generation and authentication endpoints 

# Endpoints

| Request       | Relative Path | Description                             | Request Body (JSON) | Response Body (JSON)             | Status Codes |
| ------------- | ------------- | --------------------------------------- | ------------------- | -------------------------------- | ------------ |
| POST          | `/keys`          | Generates key for owner                 | `"owner": "test"`   | `"owner":"test", "key": "mykey"` | 200          |
| GET           | `/keys`          | Gets all keys for owner                 | `"owner": "test"`   | `"owner":"test", "key": "mykey"` | 200          |
| GET           | `/auth`          | Authenticates using raw API key         | `"key": "mykey"`    | N/A                              | 200, 403     |
