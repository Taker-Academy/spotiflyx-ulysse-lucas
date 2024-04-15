# Documentation de l'API

## Erreurs

Toutes les erreurs renverront un objet JSON avec un message d'erreur approprié et un code d'état HTTP correspondant.

### Format de réponse

```json
{
    "ok": false,
    "error": "Description de l'erreur",
}
```

## User

### Inscription
- **Endpoint** : `/api/user/signup`
- **Méthode** : POST
- **Description** : Permet à un utilisateur de s'inscrire en fournissant une adresse e-mail et un mot de passe.
- **Paramètres** :
  - email (string) : Adresse e-mail de l'utilisateur.
  - password (string) : Mot de passe choisi par l'utilisateur.
- **Réponses** :
```json
{
    "ok": true,
    "data": {
        "token": "eg.eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjY1NzQzYWNmZWI0NjU3MTU0Yjg1Y2VjMyIsImlhdCI6MTcwMjExNjA0NywiZXhwIjoxNzAyMjAyNDQ3fQ.hQ2Om2eiNVPquH9npiCC9hOUy3hoizsFVt8QACCPolU",
    }
}
```
  - 200 OK : Inscription réussie.
  - 400 Bad Request : Erreur de validation des données.
  - 500 Internal Server Error: Erreur interne du serveur.

### Connexion
- **Endpoint** : `/api/user/signin`
- **Méthode** : POST
- **Description** : Permet à un utilisateur de se connecter en fournissant une adresse e-mail et un mot de passe.
- **Paramètres** :
  - email (string) : Adresse e-mail de l'utilisateur.
  - password (string) : Mot de passe de l'utilisateur.
- **Réponses** :
```json
{
    "ok": true,
    "data": {
        "token": "eg.eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjY1NzQzYWNmZWI0NjU3MTU0Yjg1Y2VjMyIsImlhdCI6MTcwMjExNjA0NywiZXhwIjoxNzAyMjAyNDQ3fQ.hQ2Om2eiNVPquH9npiCC9hOUy3hoizsFVt8QACCPolU",
    }
}
```
  - 200 OK : connexion réussie.
  - 400 Bad Request : mauvais parametres.
  - 401 Unauthorized : mauvais email/mot de passe
  - 500 Internal Server Error: Erreur interne du serveur.

### modification du mot de passe
- **Endpoint** : `/api/user`
- **Méthode** : PUT
- **Description** : Permet à un utilisateur de modifier son mot de passe

- **Header: Authorization (String, required):** Token JWT pour l'authentification.
- **Paramètres** :
  - oldPassword (string) : ancien mot de passe.
  - newPassword (string) : nouveau mot de passe.
- **Réponses** :
```json
{
    "ok": true,
    "data": {}
}
```
  - 200 OK : modification réussi.
  - 400 Bad Request : mauvais parametres.
  - 401 Unauthorized : mauvais token jwt
  - 500 Internal Server Error: Erreur interne du serveur.

### suppression du compte
- **Endpoint** : `/api/user`
- **Méthode** : PUT
- **Description** : Permet à un utilisateur de supprimer son compte
- **Header: Authorization (String, required):** Token JWT pour l'authentification.
- **Réponses** :
```json
{
    "ok": true,
    "data": {}
}
```
  - 200 OK : modification réussi.
  - 400 Bad Request : mauvais parametres.
  - 401 Unauthorized : mauvais token jwt
  - 500 Internal Server Error: Erreur interne du serveur.