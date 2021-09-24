[[_TOC_]]

## Objectifs

Faire un simple server d'api en go, fonctionnel, avec connection à une base de  
données en sqlite3 à l'image [de celui crée en js](https://gitlab.com/franckf/reference-javascript/-/tree/master/full-project-examples/lite-api-crud)

Les fonctionalités à atteindre seront :

- route de création des utilisateurs
- génération de token
- authentification par JWT
- route CRUD pour poster (titre/données/date/idUtilisateur)
- seul l'utilisateur original pourra modifier/supprimer
- validation des entrées pour les mails
- prise en comptes de toutes les erreurs possibles et retours par API avec
  messages explicites
- documentation d'api automatique ?

### Inspirations

On le fera à ma façon et on fera s=ce qu'il semble le plus naturel. Mais un
peu d'idées est bienvenue :

- https://github.com/gorilla/mux#readme
- https://tutorialedge.net/golang/creating-restful-api-with-golang/
- https://golangdocs.com/golang-mux-router
- https://stackoverflow.com/questions/42091720/api-testing-in-golang
- https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql
- https://golangbyexample.com/json-request-body-golang-http/

## Questions

- gorilla/mux ou gin-gonic/gin ?
  gin semble plus populaire et contient plus de fonctionnalités (validation,...)
  qui ne serviront peut-être pas. mux est un simple router plus proche de go
  (idiomatic) suffisant pour cette petite API

- faire les tests avec httpTest de la librairie standard ?
  l'IDE crée des templates de tests auto, idem pour http ?

  > Tous les tests utilisent la librairie standard pour faire des tests au niveau de l'API

- maintenance : comment s'assurer que les librairies externes n'auront pas de
  régressions ? les inclures au dépôts ?

- besoin de mutex ? pour eviter les ecritures/suppressions en bdd (race
  condition) ?

- remplacer `json.Encode` par `json.Marshal` dans les controllers ?
  cela éviterait les appels au Struct mais les données serait moins "stables" ?
  a tester sur Users ?

## Todos

- harmonisation des requêtes du model (Exec, Prepare, Query, Begin, QueryRow, ...),
  c'est un peu le bazard

- refactorisation de la partie controller Posts, illisible, beaucoup de
  répétitions et d'erreurs similaires. Besoin de reduction, utiliser les middlewares
  pour gérer certaines choses (formating, ...) un niveau au dessus ?

- utiliser des methodes pour les controleurs des routes Users, pour tester et
  comparer avec Posts

- ajouter un système de migration automatique (avec sauvegarde) pour la base de
  données. Par exemple l'ajout d'index :
  ```
  CREATE [UNIQUE] INDEX index_name
  ON table_name(column_list);
  ```

## Documentation de l'API

### Tous les posts

```
curl http://127.0.0.1:8000/api/v1/posts
```

### Ajouter un post

```
curl --location --request POST 'http://127.0.0.1:8000/api/v1/post' --header 'Content-Type: application/json' --data-raw '{"title":"from json","datas":"datasfill","idUser":5}'
```

### Post par id

```
curl http://127.0.0.1:8000/api/v1/post/2
```

### Mettre à jour un post

```
curl --location --request PUT 'http://127.0.0.1:8000/api/v1/post/2' --header 'Content-Type: application/json' --data-raw '{"title":"from json","datas":"datasfill","idUser":5}'
```

### Supprimer un post

```
curl --location --request DELETE 'http://127.0.0.1:8000/api/v1/post/2'
```
