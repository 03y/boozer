# login to an account (needed for adding new items)
POST https://localhost/api/v1/authenticate
{
    "username": "{{username}}",
    "password": "{{password}}"
}
HTTP 200
[Asserts]
cookie "token" exists

POST https://localhost/api/v1/submit/item
{
    "name": "{{item_name}}",
    "units": {{item_units}}
}
HTTP 201

GET https://localhost/api/v1/items
HTTP 200
[Asserts]
jsonpath "$[0].name" == "{{item_name}}"
jsonpath "$[0].units" == {{item_units}}
