# PersonCRUD_task2 Documentation
This project demonstrates CRUD operation on a resource Person.

```baseUrl: http://five0juri-personcrud-task2.onrender.com```

## Endpoints
Note:
All responses from this API return in the format below, where "data" can either be null or contain a Person object depending on the status or method:
```
{
    "status": int,
    "message": string,
    "data": {}
}
```

### 1. /api/person (Method: POST)
Use this endpoint to create a person object.

Requirement:
Body should contain:
```
  Key       Type        Validation
- name      string  must not be an empty string
- age       int     must be not be less than or equal to zero
- gender    string
```

Sample body:
```
{
    "name": "Oluwafemi Ojuri",
    "age": 5,
    "gender": "Male"
}
```

Sample Request:
```
https://{baseUrl}/api/person
```

Sample response:
```
{
    "status": 200,
    "message": "Person created successfully",
    "data": null
}
```


### 2. /api/person/:name (Method: GET)
Use this endpoint to fetch a movie BY NAME.

Sample Request
```
https://{baseUrl}/api/person/Oluwafemi
```

Sample Response
```
{
    "status": 200,
    "message": "Successful",
    "data": {
        "id":
        "name": "Oluwafemi Ojuri",
        "age": 5,
        "gender": "Male"
    }
}
```

### 3. /api/update/:name (Method: PUT)
Use this endpoint to update a person resource BY NAME

Requirement:
Body should contain:
  Key       Type        Validation
- name      string  must not be an empty string
- age       int     must be not be less than or equal to zero
- gender    string

Sample body:
```
{
    "name": "Oluwafemi Ojuri",
    "age": 15,
    "gender": "Male"
}
```

Sample Request:
```
https://{baseUrl}/api/update/Oluwafemi
```

Sample Response:
```
{
    	"status":  200,
		"message": "Update Successful",
		"data":    null
}
```


### 4. /api/delete/:name (Method: DELETE)
Use this endpoint to delete a person resource BY NAME

Sample Request:
```https://{baseUrl}/api/delete/Oluwafemi```

Sample Response:
```
{
    	"status":  200,
		"message": "Delete Successful",
		"data":    null
}
```