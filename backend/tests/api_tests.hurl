# Boozer API Tests
#
# These tests cover the /items, /item/:item_id, and /user/:user_id endpoints
#
# Prerequisites:
# - Server should be running (e.g., on localhost:8080)
# - Database should be seeded with test data
#
# Run with: hurl --test tests/api_tests.hurl
# Or with custom base URL: hurl --variable base_url=https://your-server.com tests/api_tests.hurl

###############################################################################
# GET /items - Get all items
###############################################################################

GET {{base_url}}/items
HTTP 200
[Asserts]
jsonpath "$" count > 0
jsonpath "$[0].item_id" isInteger
jsonpath "$[0].name" isString
jsonpath "$[0].units" isFloat
jsonpath "$[0].added" isInteger
# Verify we have the expected seed data items
jsonpath "$[?(@.name == 'Früh Kölsch')]" count == 1
jsonpath "$[?(@.name == 'Rothaus Pils')]" count == 1
jsonpath "$[?(@.name == 'Birra Moretti 600ml')]" count == 1

###############################################################################
# GET /item/:item_id - Get specific item (valid ID)
###############################################################################

GET {{base_url}}/item/1
HTTP 200
[Asserts]
jsonpath "$.item_id" == 1
jsonpath "$.name" == "Früh Kölsch"
jsonpath "$.units" == 2.4
jsonpath "$.added" isInteger

###############################################################################
# GET /item/:item_id - Get specific item (another valid ID)
###############################################################################

GET {{base_url}}/item/2
HTTP 200
[Asserts]
jsonpath "$.item_id" == 2
jsonpath "$.name" == "Rothaus Pils"
jsonpath "$.units" == 2.6
jsonpath "$.added" isInteger

###############################################################################
# GET /item/:item_id - Get specific item (last item from seed data)
###############################################################################

GET {{base_url}}/item/58
HTTP 200
[Asserts]
jsonpath "$.item_id" == 58
jsonpath "$.name" == "Birra Moretti 600ml"
jsonpath "$.units" == 3.0
jsonpath "$.added" == 1686604802

###############################################################################
# GET /item/:item_id - Invalid item ID (non-existent)
###############################################################################

GET {{base_url}}/item/9999
HTTP 404

###############################################################################
# GET /item/:item_id - Invalid item ID (non-numeric)
###############################################################################

GET {{base_url}}/item/invalid
#HTTP 400
HTTP 500

###############################################################################
# GET /user/:user_id - Get specific user (valid ID)
###############################################################################

GET {{base_url}}/user/1
HTTP 200
[Asserts]
jsonpath "$.user_id" == 1
jsonpath "$.username" == "Pibble"
jsonpath "$.created" isInteger
# Verify password is NOT returned (UserNoPw struct)
jsonpath "$.password" not exists

###############################################################################
# GET /user/:user_id - Get specific user (another valid ID)
###############################################################################

GET {{base_url}}/user/2
HTTP 200
[Asserts]
jsonpath "$.user_id" == 2
jsonpath "$.username" == "Glorp"
jsonpath "$.created" isInteger
jsonpath "$.password" not exists

###############################################################################
# GET /user/:user_id - Get specific user (third valid ID)
###############################################################################

GET {{base_url}}/user/3
HTTP 200
[Asserts]
jsonpath "$.user_id" == 3
jsonpath "$.username" == "Gleeb"
jsonpath "$.created" isInteger
jsonpath "$.password" not exists

###############################################################################
# GET /user/:user_id - Get specific user (fourth valid ID)
###############################################################################

GET {{base_url}}/user/4
HTTP 200
[Asserts]
jsonpath "$.user_id" == 4
jsonpath "$.username" == "Gnarp"
jsonpath "$.created" isInteger
jsonpath "$.password" not exists

###############################################################################
# GET /user/:user_id - Invalid user ID (non-existent)
###############################################################################

GET {{base_url}}/user/9999
HTTP 404

###############################################################################
# GET /user/:user_id - Invalid user ID (non-numeric)
###############################################################################

GET {{base_url}}/user/invalid
#HTTP 400
HTTP 500
