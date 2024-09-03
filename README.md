# Testing task "REST API for the News service"

 #go #fiber #reform #mysql #rest-api #json-api #docker #api-test #hurl

## Excercise

Implement JSON REST API service with the following routes:

* `POST /edit/:Id` * Update a news record by its Id.
* `GET /list` - Retrieve a list of news records.
* `POST /add` - Add a news record.
* `POST /add-category/:NewsId/:CategoryId` - Add a category to a news record.
* `DELETE /:NewsId` - Delete a news record.

For database storage, you can use either MySQL or PostgreSQL.

The server is built using Fiber, and Reform is used for database interactions.

The connection to the database should utilize a connection pool. All settings should be configured through environment variables and/or Viper.

### DB schema

```SQL
CREATE TABLE `News` (
  `Id` bigint NOT NULL,
  `Title` tinytext NOT NULL,
  `Content` longtext NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `NewsCategories` (
  `NewsId` bigint NOT NULL,
  `CategoryId` bigint NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

ALTER TABLE `News`
  ADD PRIMARY KEY (`Id`);

ALTER TABLE `NewsCategories`
  ADD PRIMARY KEY (`NewsId`,`CategoryId`);

ALTER TABLE `News`
  MODIFY `Id` bigint NOT NULL AUTO_INCREMENT;
```

It is important to note that the association between news and categories is managed in a separate table.

### Input Data Format for the Edit Endpoint

```json
{
  "Id": 64,
  "Title": "Lorem ipsum",
  "Content": "Dolor sit amet <b>foo</b>",
  "Categories": [1,2,3]
}
```

If a field is not provided in the input, that field should not be updated.

### Output Data Format for the List Endpoint

```json
{
    "Success": true,
    "News": [
      {
        "Id": 64,
        "Title": "Lorem ipsum",
        "Content": "Dolor sit amet <b>foo</b>",
        "Categories": [1,2,3]
      },
      {
        "Id": 1,
        "Title": "first",
        "Content": "tratata",
        "Categories": [1]
      }
    ]
}
```

## Requirements and Recommendations

* If you are familiar with Docker, we would like to see the service containerized.
* Authorization via the Authorization header and well-organized code structure and routing by groups/folders.
* Field validation during the editing process.
* Pagination in the list endpoint.
* Proper logging using any popular logger (e.g., logrus).
* Robust error handling.
