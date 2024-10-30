### Ok: Ping-pong
GET {{url_base}}/ping
Origin: http://any-origin.io
Content-type: any-content-type
HTTP 200
content-type: application/json
content-length: 18
{"Message":"pong"}

### Fail: Cannot get list because unauthorized
GET {{url_base}}/list
HTTP 401
[Asserts]
body == "Unauthorized"

### Ok: Authorize to get access token
POST {{url_base}}/login
HTTP 200
[Captures]
token: jsonpath "$['token']"

### Ok: Get empty list when authorized
GET {{url_base}}/list
Authorization: Bearer {{token}}
HTTP 200
content-type: application/json
content-length: 17
{"Success":false}

### Ok: Add 1st news record
POST {{url_base}}/add
Authorization: Bearer {{token}}
{"Title":"News title: America", "Content":"News content: America"}
HTTP 200
content-type: application/json
[Asserts]
jsonpath "$.Success" == true
jsonpath "$.News.Title" == "News title: America"
jsonpath "$.News.Content" == "News content: America"
[Captures]
news-id-a: jsonpath "$['News']['ID']"

### Ok: Add 2nd news record
POST {{url_base}}/add
Authorization: Bearer {{token}}
{"Title":"News title: Belgium", "Content":"News content: Belgium"}
HTTP 200
content-type: application/json
[Asserts]
jsonpath "$.Success" == true
jsonpath "$.News.Title" == "News title: Belgium"
jsonpath "$.News.Content" == "News content: Belgium"
[Captures]
news-id-b: jsonpath "$['News']['ID']"

### Ok: Add 3rd news record
POST {{url_base}}/add
Authorization: Bearer {{token}}
{"Title":"Cyprus", "Content":"Cyprus", "Categories":[100,200,300]}
HTTP 200
content-type: application/json
[Asserts]
jsonpath "$.Success" == true
jsonpath "$.News.Title" == "Cyprus"
jsonpath "$.News.Content" == "Cyprus"
jsonpath "$.Categories[0]" == 100
jsonpath "$.Categories[1]" == 200
jsonpath "$.Categories[2]" == 300
[Captures]
news-id-c: jsonpath "$['News']['ID']"

### Ok: Edit 3rd news record
POST {{url_base}}/edit/{{news-id-c}}
Authorization: Bearer {{token}}
{"ID":{{news-id-c}},"Title":"News title: Cyprus","Categories":[300,400]}
HTTP 200
content-type: application/json
[Asserts]
jsonpath "$.Success" == true
jsonpath "$.News.ID" == {{news-id-c}}
jsonpath "$.News.Title" == "News title: Cyprus"
jsonpath "$.News.Content" == "Cyprus"
jsonpath "$.Categories" count == 2
jsonpath "$.Categories[0]" == 300
jsonpath "$.Categories[1]" == 400

### Fail: Cannot edit, because of id parsing error
POST {{url_base}}/edit/abc
Authorization: Bearer {{token}}
{"ID":1,"Title":"any title"}
HTTP 404

### Fail: Cannot edit, because different news IDs in route and in body
POST {{url_base}}/edit/{{news-id-a}}
Authorization: Bearer {{token}}
{"ID":{{news-id-b}},"Title":"any title"}
HTTP 200
content-type: application/json
[Asserts]
jsonpath "$.Success" == false
jsonpath "$.Message" == "News ID in route (ID: {{news-id-a}}) does not match News record ID in request body (ID: {{news-id-b}})"

### Fail: Cannot edit, because news record absent
POST {{url_base}}/edit/9999999
Authorization: Bearer {{token}}
{"ID":9999999,"Title":"any title"}
HTTP 200
content-type: application/json
[Asserts]
jsonpath "$.Success" == false
jsonpath "$.Message" == "Cannot find News record (ID: 9999999)"

### Ok: Add category
POST {{url_base}}/add-category/{{news-id-c}}/500
Authorization: Bearer {{token}}
HTTP 200
content-type: application/json
[Asserts]
jsonpath "$.Success" == true
jsonpath "$.Message" == "Category (ID: 500) assigned to news record (ID: {{news-id-c}})"

### Fail: Cannot add category, because of NewsId parsing error
POST {{url_base}}/add-category/abc/600
Authorization: Bearer {{token}}
HTTP 404

### Fail: Cannot add category, because of CatId parsing error
POST {{url_base}}/add-category/{{news-id-c}}/abc
Authorization: Bearer {{token}}
HTTP 404

### Fail: Cannot add category, because news record is absent
POST {{url_base}}/add-category/999999/600
Authorization: Bearer {{token}}
HTTP 200
content-type: application/json
[Asserts]
jsonpath "$.Success" == false
jsonpath "$.Message" == "Cannot find News record (ID: 999999)"

### Ok: List records with custom pagination
GET {{url_base}}/list?page=1&per-page=2
Authorization: Bearer {{token}}
HTTP 200
content-type: application/json
[Asserts]
jsonpath "$.News" count == 2
{"Success":true,"News":[{"ID":{{news-id-a}},"Title":"News title: America","Content":"News content: America","Categories":null},{"ID":{{news-id-b}},"Title":"News title: Belgium","Content":"News content: Belgium","Categories":null}]}

### Ok: List records with custom pagination
GET {{url_base}}/list?page=2&per-page=2
Authorization: Bearer {{token}}
HTTP 200
content-type: application/json
[Asserts]
jsonpath "$.News" count == 1
{"Success":true,"News":[{"ID":{{news-id-c}},"Title":"News title: Cyprus","Content":"Cyprus","Categories":[300,400,500]}]}

### Ok: List records with default pagination
GET {{url_base}}/list
Authorization: Bearer {{token}}
HTTP 200
content-type: application/json
[Asserts]
jsonpath "$.News" count == 3

### Ok: Delete news record
DELETE {{url_base}}/{{news-id-a}}
Authorization: Bearer {{token}}
HTTP 200
content-type: application/json
[Asserts]
jsonpath "$.Success" == true
jsonpath "$.Message" == "News record (ID: {{news-id-a}}) is deleted"

### Ok: Delete news record
DELETE {{url_base}}/{{news-id-b}}
Authorization: Bearer {{token}}
HTTP 200
content-type: application/json
[Asserts]
jsonpath "$.Success" == true
jsonpath "$.Message" == "News record (ID: {{news-id-b}}) is deleted"

### Ok: Delete news record
DELETE {{url_base}}/{{news-id-c}}
Authorization: Bearer {{token}}
HTTP 200
content-type: application/json
[Asserts]
jsonpath "$.Success" == true
jsonpath "$.Message" == "News record (ID: {{news-id-c}}) is deleted"