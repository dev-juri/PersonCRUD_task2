# PersonCRUD_task2 Documentation
This project demonstrates CRUD operation on a resource Person.
A NoSQL database was used, you can view the collection model along side the UML class diagram for the application [here](https://lucid.app/documents/view/260db3b1-1d9b-4835-9f19-b8181d2b3cb8)

```baseUrl: https://five0juri-personcrud-task2.onrender.com```

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
    "name": "Oluwafemi",
    "age": 5,
    "gender": "Male"
}
```

Sample Request:
```{baseUrl}/api/person```

Sample response:
```
{
    "status": 200,
    "message": "Person created successfully",
    "data": {
        "person": {
            "id": "64ff73ac452fe02b75c391dc",
            "name": "Oluwafemi",
            "age": 5,
            "gender": "Male"
        }
    }
}
```


### 2. /api/person/:name (Method: GET)
Use this endpoint to fetch a movie BY NAME.

Sample Request
```{baseUrl}/api/person/Oluwafemi```

Sample Response
```
{
    "status": 200,
    "message": "Successful",
    "data": {
        "person": {
            "id": "64ff73ac452fe02b75c391dc",
            "name": "Oluwafemi",
            "age": 5,
            "gender": "Male"
        }
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
    "name": "Oluwafemi",
    "age": 15,
    "gender": "Male"
}
```

Sample Request:
```{baseUrl}/api/update/Oluwafemi```

Sample Response:
```
{
    "status":  200,
	"message": "Update Successful",
	"data":    {
        "person": {
            "id": "64ff73ac452fe02b75c391dc",
            "name": "Oluwafemi",
            "age": 15,
            "gender": "Male"
        }
    }
}
```


### 4. /api/delete/:name (Method: DELETE)
Use this endpoint to delete a person resource BY NAME

Sample Request:
```{baseUrl}/api/delete/Oluwafemi```

Sample Response:
```
{
    "status":  200,
	"message": "Record for person with name Oluwafemi deleted successfully",
	"data":    null
}
```

### Testing 
[Test the API on Postman [here](https://api.getpostman.com/collections/14969266-d38c4aed-1403-4004-9ae1-b9d7f87f5177)](https://elements.getpostman.com/redirect?entityId=14969266-d38c4aed-1403-4004-9ae1-b9d7f87f5177&entityType=collection)https://elements.getpostman.com/redirect?entityId=14969266-d38c4aed-1403-4004-9ae1-b9d7f87f5177&entityType=collection)
