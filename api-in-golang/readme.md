
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
- https://medium.com/@saumya.ranjan/how-to-create-a-rest-api-in-golang-crud-operation-in-golang-a7afd9330a7b
- https://mapereira0101.medium.com/all-you-need-to-build-a-simple-restfull-api-with-golang-part-i-a6ae5ac0a0d8
- https://tutorialedge.net/golang/creating-restful-api-with-golang/
- https://golangdocs.com/golang-mux-router
- https://stackoverflow.com/questions/42091720/api-testing-in-golang
- https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql


## Questions

- gorilla/mux ou gin-gonic/gin ?
gin semble plus populaire et contient plus de fonctionnalités (validation,...)
qui ne serviront peut-être pas. mux est un simple router plus proche de go
(idiomatic) suffisant pour cette petite API

- faire les tests avec httpTest de la librairie standard ?
l'IDE crée des templates de tests auto, idem pour http ?

- maintenance : comment s'assurer que les librairies externes n'auront pas de
régressions ? les inclures au dépôts ?

- besion de mutex ? pour eviter les ecritures/suppresions en bdd (race
condition) ?

## Documentation de l'API

