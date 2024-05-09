# API Documentation

## Errors

All errors will return a JSON object with an appropriate error message and a corresponding HTTP status code.

### Response Format

```json
{
    "ok": false,
    "error": "Error description",
}
```

## User and Auth

### Authenticated
- **Endpoint** : `/api/auth/authenticated`
- **Method** : POST
- **Description** : check if user is authenticated.
- **Parameters** : null
- **Responses** :
```json
{
    "ok": true,
}
```
  - 200 OK : Successful registration.
  - 401 Unauthorized : wrong token

### Registration
- **Endpoint** : `/api/auth/signup`
- **Method** : POST
- **Description** : Allows a user to register by providing an email address and a password.
- **Parameters** :
  - email (string, required) : User's email address.
  - password (string, required) : Password chosen by the user.
- **Responses** :
```json
{
    "ok": true,
    "data": {
        "token": "eg.eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjY1NzQzYWNmZWI0NjU3MTU0Yjg1Y2VjMyIsImlhdCI6MTcwMjExNjA0NywiZXhwIjoxNzAyMjAyNDQ3fQ.hQ2Om2eiNVPquH9npiCC9hOUy3hoizsFVt8QACCPolU",
    }
}
```
  - 200 OK : Successful registration.
  - 400 Bad Request : Data validation error.
  - 500 Internal Server Error: Internal server error.

### Login
- **Endpoint** : `/api/auth/signin`
- **Method** : POST
- **Description** : Allows a user to log in by providing an email address and a password.
- **Parameters** :
  - email (string, required) : User's email address.
  - password (string, required) : User's password.
- **Responses** :
```json
{
    "ok": true,
    "data": {
        "token": "eg.eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjY1NzQzYWNmZWI0NjU3MTU0Yjg1Y2VjMyIsImlhdCI6MTcwMjExNjA0NywiZXhwIjoxNzAyMjAyNDQ3fQ.hQ2Om2eiNVPquH9npiCC9hOUy3hoizsFVt8QACCPolU",
    }
}
```
  - 200 OK : Successful login.
  - 400 Bad Request : Incorrect parameters.
  - 401 Unauthorized : Incorrect email/password
  - 500 Internal Server Error: Internal server error.

### Password Change
- **Endpoint** : `/api/user`
- **Method** : PUT
- **Description** : Allows a user to change their password

- **Header: Authorization (String, required):** JWT Token for authentication.
- **Parameters** :
  - OldPassword (string, required) : Old password.
  - NewPassword (string, required) : New password.
- **Responses** :
```json
{
    "ok": true,
    "data": {}
}
```
  - 200 OK : Successful change.
  - 400 Bad Request : Incorrect parameters.
  - 401 Unauthorized : Incorrect JWT token
  - 500 Internal Server Error: Internal server error.

### Account Information Retrieval
- **Endpoint** : `/api/user`
- **Method** : GET
- **Description** : Allows a user to retrieve their information
- **Header: Authorization (String, required):** JWT Token for authentication.
- **Responses** :
```json
{
    "ok": true,
    "data": {
      "email": "test@gmail.com"
    }
}
```
  - 200 OK : Successful retrieval.
  - 400 Bad Request : Incorrect parameters.
  - 401 Unauthorized : Incorrect JWT token
  - 500 Internal Server Error: Internal server error.

### Account Deletion
- **Endpoint** : `/api/user`
- **Method** : DELETE
- **Description** : Allows a user to delete their account
- **Header: Authorization (String, required):** JWT Token for authentication.
- **Responses** :
```json
{
    "ok": true,
    "data": {}
}
```
  - 200 OK : Successful deletion.
  - 400 Bad Request : Incorrect parameters.
  - 401 Unauthorized : Incorrect JWT token
  - 500 Internal Server Error: Internal server error.
## Search

### Video or Music Search
- **Endpoint** : `/api/media/search`
- **Method** : GET
- **Description** : Allows retrieval of a list of videos or music with a specific title.
- **Header: Authorization (String, required):** JWT Token for authentication.
- **Parameters** :
  - search (String, required) : title of the media.
- **Responses** :
```json
{
    "ok": true,
    "data": [
      {
        "mediaType": "music",
        "title": "sunshine",
        "id": "cdejdnscksjdc"
      },
      {
        "mediaType": "video",
        "title": "sunshine timelapse",
        "id": "assdwfvsrvdcf"
      }
      ...
    ]
}
```
  - 200 OK : Successful request.
  - 400 Bad Request : Incorrect parameters.
  - 401 Unauthorized : Incorrect JWT token.
  - 500 Internal Server Error: Internal server error.

## Music

### Get Music Details

- **Endpoint** : `/api/media/music/:id`
- **Method** : GET
- **Description** : Allows retrieval of the details of a music track.
- **Header: Authorization (String, required):** JWT Token for authentication.
- **Responses** :
```json
{
    "ok": true,
    "data": {
      "mediaType": "music",
      "title": "sunshine",
      "author": "OneRepublic",
      "id": "cdejdnscksjdc",
      "likes": 45,
      "liked": true,
      "favorite": false,
      "uri": "spotify:track:cdejdnscksjdc",
      "img_url": "https://spotify.com/img/ucfshcbshe"
    }
}
```
  - 200 OK : Successful request.
  - 400 Bad Request : Incorrect parameters.
  - 404 Not Found : Music not found.
  - 401 Unauthorized : Incorrect JWT token.
  - 500 Internal Server Error: Internal server error.

## Video

### Get Video Details

- **Endpoint** : `/api/media/video/:id`
- **Method** : GET
- **Description** : Allows retrieval of the details of a video.
- **Header: Authorization (String, required):** JWT Token for authentication.
- **Responses** :
```json
{
    "ok": true,
    "data": {
      "mediaType": "video",
      "title": "sunshine timelapse",
      "author": "4K timelapse",
      "id": "assdwfvsrvdcf",
      "likes": 147,
      "liked": true,
      "favorite": false,
      "uri": "https://www.youtube.com/embed/KiIXuKA9JQ0",
      "img_url": "https://yt.com/img/ucfshcbshe"
    }
}
```
  - 200 OK : Successful request.
  - 400 Bad Request : Incorrect parameters.
  - 404 Not Found : Video not found.
  - 401 Unauthorized : Incorrect JWT token.
  - 500 Internal Server Error: Internal server error.

## Interact

### Like
- **Endpoint** : `/api/me/like/:id`
- **Method** : POST
- **Description** : Allows liking a video or music.
- **Header: Authorization (String, required):** JWT Token for authentication.
- **Parameters** :
  - mediaType (String, required) : Type of media (video or music).
- **Responses** :
```json
{
    "ok": true,
    "data": {
      "mediaType": "video",
      "title": "sunshine timelapse",
      "author": "4K timelapse",
      "id": "assdwfvsrvdcf",
      "likes": 147,
      "liked": true,
      "favorite": false,
      "uri": "https://www.youtube.com/embed/KiIXuKA9JQ0",
      "img_url": "https://yt.com/img/ucfshcbshe"
    }
}
```
  - 200 OK : Like successful.
  - 400 Bad Request : Incorrect parameters or missing parameter.
  - 404 Not Found : Video/music not found.
  - 401 Unauthorized : Incorrect JWT token.
  - 500 Internal Server Error: Internal server error.

### Remove Like
- **Endpoint** : `/api/me/like/:id`
- **Method** : DELETE
- **Description** : Allows removing a like from a video or music.
- **Header: Authorization (String, required):** JWT Token for authentication.
- **Parameters** :
  - mediaType (String, required) : Type of media (video or music).
- **Responses** :
```json
{
    "ok": true,
    "data": {
      "mediaType": "video",
      "title": "sunshine timelapse",
      "author": "4K timelapse",
      "id": "assdwfvsrvdcf",
      "likes": 147,
      "liked": false,
      "favorite": false,
      "uri": "https://www.youtube.com/embed/KiIXuKA9JQ0",
      "img_url": "https://yt.com/img/ucfshcbshe"
    }
}
```
  - 200 OK : Like removal successful.
  - 400 Bad Request : Incorrect parameters or missing parameter.
  - 404 Not Found : Video/music not found.
  - 401 Unauthorized : Incorrect JWT token.
  - 500 Internal Server Error: Internal server error.

### Save to Favorites
- **Endpoint** : `/api/me/save/:id`
- **Method** : POST
- **Description** : Allows saving a video or music to favorites.
- **Header: Authorization (String, required):** JWT Token for authentication.
- **Parameters** :
  - mediaType (String, required) : Type of media (video or music).
- **Responses** :
```json
{
    "ok": true,
    "data": {
      "mediaType": "video",
      "title": "sunshine timelapse",
      "author": "4K timelapse",
      "id": "assdwfvsrvdcf",
      "likes": 147,
      "liked": false,
      "favorite": true,
      "uri": "https://www.youtube.com/embed/KiIXuKA9JQ0",
      "img_url": "https://yt.com/img/ucfshcbshe"
    }
}
```
  - 200 OK : Save successful.
  - 400 Bad Request : Incorrect parameters or missing parameter.
  - 404 Not Found : Video/music not found.
  - 401 Unauthorized : Incorrect JWT token.
  - 500 Internal Server Error: Internal server error.

### Remove from Favorites
- **Endpoint** : `/api/me/save/:id`
- **Method** : DELETE
- **Description** : Allows removing a video or music from favorites.
- **Header: Authorization (String, required):** JWT Token for authentication.
- **Parameters** :
  - mediaType (String, required) : Type of media (video or music).
- **Responses** :
```json
{
    "ok": true,
    "data": {
      "mediaType": "video",
      "title": "sunshine timelapse",
      "author": "4K timelapse",
      "id": "assdwfvsrvdcf",
      "likes": 147,
      "favorite": false,
      "liked": false,
      "uri": "https://www.youtube.com/embed/KiIXuKA9JQ0",
      "img_url": "https://yt.com/img/ucfshcbshe"
    }
}
```
  - 200 OK : Removal successful.
  - 400 Bad Request : Incorrect parameters or missing parameter.
  - 404 Not Found : Video/music not found.
  - 401 Unauthorized : Incorrect JWT token.
  - 500 Internal Server Error: Internal server error.

### Get Favorites
- **Endpoint** : `/api/me/favorites`
- **Method** : GET
- **Description** : Allows retrieving favorite videos and music.
- **Header: Authorization (String, required):** JWT Token for authentication.
- **Responses** :
```json
{
    "ok": true,
    "data": [
      {
        "mediaType": "music",
        "title": "sunshine",
        "id": "cdejdnscksjdc"
      },
      {
        "mediaType": "video",
        "title": "sunshine timelapse",
        "id": "assdwfvsrvdcf"
      }
      ...
    ]
}
```
  - 200 OK : Retrieval successful.
  - 400 Bad Request : Incorrect parameters or missing parameter.
  - 401 Unauthorized : Incorrect JWT token.
  - 500 Internal Server Error: Internal server error.

## Create Media
- **Endpoint** : `/api/media/create`
- **Method** : POST
- **Description** : Allows creating a video or music.
- **Header: Authorization (String, required):** JWT Token for authentication.
- **Parameters** :
  - mediaType (string, [music, video], required) : type of media.
  - title (string) : set custom name for the media.
  - url (string, required) : Link to retrieve the content from the corresponding API
- **Responses** :
```json
{
    "ok": true,
    "data": {
      "mediaType": "video",
      "title": "sunshine timelapse",
      "author": "4K timelapse",
      "id": "assdwfvsrvdcf",
      "likes": 0,
      "favorite": false,
      "liked": false,
      "url": "https://www.youtube.com/embed/KiIXuKA9JQ0",
      "imgUrl": "https://yt.com/img/ucfshcbshe"
    }
}
```
  - 201 OK : media created.
  - 400 Bad Request : Incorrect parameters or missing parameter.
  - 401 Unauthorized : Incorrect JWT token.
  - 422 Unprocessable Entity: Invalid url.
  - 500 Internal Server Error: Internal server error.

## get recent media
- **Endpoint** : `/api/media/latest`
- **Method** : GET
- **Description** : get the 3 most recent media in each category (video and music).
- **Header: Authorization (String, required):** JWT Token for authentication.
- **Responses** :
```json
{
    "ok": true,
    "data": {
      "music": [
        {
        "mediaType": "music",
        "title": "sunshine",
        "imgUrl": "https://img.com",
        "id": "cdejdnscksjdc"
        },
        {
          "mediaType": "music",
          "title": "waterstyle",
        "imgUrl": "https://img.com",
          "id": "cdejfugjjfddc"
        },
        {
          "mediaType": "music",
          "title": "loving me",
        "imgUrl": "https://img.com",
          "id": "cfyukolkjsjdc"
        }
      ],
      "video": [
        {
        "mediaType": "video",
        "title": "sunshine timelapse",
        "imgUrl": "https://img.com",
        "id": "assdwfvsrvdcf"
        },
        {
        "mediaType": "video",
        "title": "new fortnite battle pass !!!",
        "imgUrl": "https://img.com",
        "id": "zeffefvsrvdck"
        },
        {
        "mediaType": "video",
        "title": "The history of Mr Beast",
        "imgUrl": "https://img.com",
        "id": "ascvbbcsrvdcj"
        }
      ]
    }
}
```
  - 200 OK : request successful.
  - 400 Bad Request : Incorrect parameters or missing parameter.
  - 401 Unauthorized : Incorrect JWT token.
  - 500 Internal Server Error: Internal server error.