# Project: Cassandra CRUD

## End-point: Get All
### Method: GET
>```
>http://127.0.0.1:8080/contact
>```

⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Create
### Method: POST
>```
>http://127.0.0.1:8080/contact
>```
### Body (**raw**)

```json
{
    "id": "aeeabd1c-68f8-41f4-a70c-9d71c824eac5",
    "title": "test",
    "name": "test",
    "address": "test",
    "realAddress": "test",
    "department": "test",
    "country": "test",
    "tel": "test",
    "email": "test"
}
```

### Response: 200
```json
{
    "data": {
        "id": "aeeabd1c-68f8-41f4-a70c-9d71c824eac5",
        "title": "test",
        "name": "test",
        "address": "test",
        "realAddress": "test",
        "department": "test",
        "country": "test",
        "tel": "test",
        "email": "test"
    }
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Delete
### Method: DELETE
>```
>http://127.0.0.1:8080/contact/aeeabd1c-68f8-41f4-a70c-9d71c824eac5
>```
### Response: 200
```json
{
    "data": "aeeabd1c-68f8-41f4-a70c-9d71c824eac5"
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Get By id
### Method: GET
>```
>http://127.0.0.1:8080/contact/2957b80f-b0cd-46d9-9073-858400c37edd
>```
### Response: 200
```json
{
    "data": {
        "id": "2957b80f-b0cd-46d9-9073-858400c37edd",
        "title": "CFA chambre de Métiers et de l'artisanat du nord Lille : S'inscrire, Cursus, Formation ",
        "name": "Chambre de métiers et de l'artisanat Hauts-de-France - LAON",
        "address": "30 rue d’Enfer 02000 LAON France",
        "realAddress": "30 rue d’Enfer",
        "department": "02000 LAON",
        "country": "France",
        "tel": "09 72 72 72 07",
        "email": "contact@cma-hautsdefrance.fr"
    }
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Update
### Method: PUT
>```
>http://127.0.0.1:8080/contact/a60a6819-f365-4e05-98e8-eda8b6b824eb
>```
### Body (**raw**)

```json
{
    "id": "a60a6819-f365-4e05-98e8-eda8b6b824eb",
    "title": "Chambre de métiers et de l'artisanat Hauts-de-France (Centre de formation d’apprentis - Chambre de métiers et de l’artisanat du Nord) ",
    "name": "Chambre de métiers et de l'artisanat Hauts-de-France - ROUVIGNIES",
    "address": "6 rue Edmond Herly 59220 ROUVIGNIES France",
    "realAddress": "6 rue Edmond Herly",
    "department": "59220 ROUVIGNIES",
    "country": "France",
    "tel": "09 72 72 72 07",
    "email": "contact@cma-hautsdefrance.fr"
}
```

### Response: 200
```json
{
    "data": {
        "id": "a60a6819-f365-4e05-98e8-eda8b6b824eb",
        "title": "Chambre de métiers et de l'artisanat Hauts-de-France (Centre de formation d’apprentis - Chambre de métiers et de l’artisanat du Nord) ",
        "name": "Chambre de métiers et de l'artisanat Hauts-de-France - ROUVIGNIES",
        "address": "6 rue Edmond Herly 59220 ROUVIGNIES France",
        "realAddress": "6 rue Edmond Herly",
        "department": "59220 ROUVIGNIES",
        "country": "France",
        "tel": "09 72 72 72 07",
        "email": "contact@cma-hautsdefrance.fr"
    }
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃
_________________________________________________
Powered By: [postman-to-markdown](https://github.com/bautistaj/postman-to-markdown/)
