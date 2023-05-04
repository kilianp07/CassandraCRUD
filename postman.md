# Project: Cassandra CRUD
# 📁 Collection: Restaurant 


## End-point: Delete Restaurant
### Method: DELETE
>```
>http://127.0.0.1:8080/restaurant/{{$restaurant_id}}
>```
### Response: 200
```json
{
    "Message": "Restaurant deleted successfully"
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Get By ID
### Method: GET
>```
>http://127.0.0.1:8080/restaurant/{{$restaurant_id}}
>```
### Response: 200
```json
{
    "data": {
        "address": {
            "id": "7363d9e5-ca88-47dd-bb17-70190749cf2a",
            "building": "37-02",
            "coord": {
                "type": "",
                "coordinates": [
                    0,
                    0
                ]
            },
            "street": "Main Street",
            "zipcode": "11354",
            "restaurant_id": "35a69bd6-569c-472d-9e28-977b6f402aad"
        },
        "borough": "Brooklyn",
        "cuisine": "Italian",
        "grades": null,
        "name": "New Restaurant Name",
        "id": "35a69bd6-569c-472d-9e28-977b6f402aad"
    }
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Get all Restaurants
### Method: GET
>```
>http://127.0.0.1:8080/restaurant/all
>```

⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: New Restaurant
### Method: POST
>```
>http://127.0.0.1:8080/restaurant
>```
### Body (**raw**)

```json
{
    "borough": "Brooklyn",
    "cuisine": "Italian",
    "name": "Test Restaurant"
}

```

### Response: 201
```json
{
    "restaurant_id": "45cf9764-5204-473e-a43d-8e527e983e77",
    "success": "Restaurant created successfully"
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Update Restaurant
### Method: PUT
>```
>http://127.0.0.1:8080/restaurant/{{$restaurant_id}}
>```
### Body (**raw**)

```json

{
    "name": "New Restaurant Name",
    "borough": "Brooklyn",
    "cuisine": "Italian",
    "restaurant_id": "{{$restaurant_id}}"
}
```

### Response: 200
```json
{
    "Message": "Restaurant updated successfully",
    "data": {
        "address": {
            "building": "",
            "coord": {
                "type": "",
                "coordinates": [
                    0,
                    0
                ]
            },
            "street": "",
            "zipcode": ""
        },
        "borough": "Brooklyn",
        "cuisine": "Italian",
        "grades": null,
        "name": "New Restaurant Name",
        "restaurant_id": "35a69bd6-569c-472d-9e28-977b6f402aad"
    }
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃
# 📁 Collection: Address 


## End-point: Get All
### Method: GET
>```
>http://127.0.0.1:8080/address/all
>```

⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Delete
### Method: DELETE
>```
>http://127.0.0.1:8080/address/{{$address_id}}
>```
### Response: 200
```json
{
    "success": "Address deleted successfully"
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃
_________________________________________________
Powered By: [postman-to-markdown](https://github.com/bautistaj/postman-to-markdown/)
