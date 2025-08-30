POST https://localhost/api/v1/signup
{
    "username": "{{random_string}}",
    "password": "password123"
}
HTTP 201

POST https://localhost/api/v1/authenticate
{
    "username": "{{random_string}}",
    "password": "password123"
}
HTTP 200
[Asserts]
cookie "token" exists

PUT https://localhost/api/v1/change_password
{
    "old_password": "password123",
    "new_password": "new_password"
}
HTTP 200

POST https://localhost/api/v1/logout
HTTP 200

POST https://localhost/api/v1/authenticate
{
    "username": "{{random_string}}",
    "password": "password123"
}
HTTP 400

POST https://localhost/api/v1/authenticate
{
    "username": "{{random_string}}",
    "password": "new_password"
}
HTTP 200
