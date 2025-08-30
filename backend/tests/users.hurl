POST https://localhost/api/v1/signup
{
    "username": "{{username}}",
    "password": "{{password}}"
}
HTTP 201

POST https://localhost/api/v1/authenticate
{
    "username": "{{username}}",
    "password": "{{password}}"
}
HTTP 200
[Asserts]
cookie "token" exists

PUT https://localhost/api/v1/change_password
{
    "old_password": "{{password}}",
    "new_password": "new_password"
}
HTTP 200

POST https://localhost/api/v1/logout
HTTP 200

POST https://localhost/api/v1/authenticate
{
    "username": "{{username}}",
    "password": "{{password}}"
}
HTTP 400

POST https://localhost/api/v1/authenticate
{
    "username": "{{username}}",
    "password": "new_password"
}
HTTP 200

PUT https://localhost/api/v1/change_password
{
    "old_password": "new_password",
    "new_password": "{{password}}"
}
HTTP 200

