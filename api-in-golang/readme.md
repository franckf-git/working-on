[[_TOC_]]

## Objectifs

Faire un simple server d'api en go, fonctionnel, avec connection à une base de  
données en sqlite3 à l'image [de celui crée en js](https://gitlab.com/franckf/reference-javascript/-/tree/master/full-project-examples/lite-api-crud)

Les fonctionalités à atteindre seront :

- [x] route de création des utilisateurs
- [x] génération de token
- [x] authentification par JWT
- [x] route CRUD pour poster (titre/données/date/idUtilisateur)
- [x] seul l'utilisateur original pourra modifier/supprimer
- [x] validation des entrées pour les mails
- [x] prise en comptes de toutes les erreurs possibles et retours par API avec
      messages explicites
- documentation d'api automatique ?

### Inspirations

On le fera à ma façon et on fera ce qu'il semble le plus naturel. Mais un
peu d'idées est bienvenue :

- https://github.com/gorilla/mux#readme
- https://tutorialedge.net/golang/creating-restful-api-with-golang/
- https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql
- https://golangbyexample.com/json-request-body-golang-http/

## Questions

- [x] gorilla/mux ou gin-gonic/gin ?
      gin semble plus populaire et contient plus de fonctionnalités (validation,...)
      qui ne serviront peut-être pas. mux est un simple router plus proche de go
      (idiomatic) suffisant pour cette petite API

  > gorilla/mux est suffisament complexe pour l'instant

- [x] faire les tests avec httpTest de la librairie standard ?
      l'IDE crée des templates de tests auto, idem pour http ?

  > Tous les tests utilisent la librairie standard pour faire des tests au niveau de l'API

- [ ] maintenance : comment s'assurer que les librairies externes n'auront pas de
      régressions ? les inclures au dépôts ? https://stackoverflow.com/questions/9985559/organizing-a-multiple-file-go-project

- [ ] besoin de mutex ? pour eviter les ecritures/suppressions en bdd (race
      condition) ?

  > Gérer par database/sql et le pool de connection dans go si les requêtes sont bien préparées
  > Tester également `db, err := sql.Open("sqlite3", "file:locked.sqlite?cache=shared&mode=rwc")`

- [x] remplacer `json.Encode` par `json.Marshal` dans les controllers ?
      cela éviterait les appels au Struct mais les données serait moins "stables" ?

  > Marshal est pour les []bytes (chargement en mémoire) - Encode est pour les streams
  > Pour http les deux methods se valent, Marshal est peut-être un peu plus
  > performant, car Encode appelle celui-ci mais au final c'est juste une
  > question de lisibilité et de validation Encode permet DisallowUnknownFields
  > pour retourner une erreur en cas de champs inconnus (plus compliqué avec Marshal)
  >
  > > bref, préférer json.Encode/Decode

## Todos

- [x] harmonisation des requêtes du model (Exec, Prepare, Query, Begin, QueryRow, ...),
      c'est un peu le bazard

- [ ] refactorisation de la partie controller Posts, illisible, beaucoup de
      répétitions et d'erreurs similaires. Besoin de reduction, utiliser les middlewares
      pour gérer certaines choses (formating, ...) un niveau au dessus ?

- [x] utiliser des methodes pour les controleurs des routes Users, pour tester et
      comparer avec Posts

  > les methodes ne peuvent être utilisées que dans le même package
  > difficile dans un projet api multi-package comme celui-ci
  > mais cela reste possible à condition de faire des compromis sur
  > l'emplacement des struct

- [ ] ajouter un système de migration automatique (avec sauvegarde) pour la base de
      données. Par exemple l'ajout d'index :

  ```
  CREATE [UNIQUE] INDEX index_name
  ON table_name(column_list);
  ```

  > nécéssaire suite à l'oubli d'évolution de schéma lors de la création de AddUser

- [x] ajouter un mode debug pour couper si besoin les sorties des logs

- [x] "Don’t Open() and Close() databases frequently. Instead, create one sql.DB object for each distinct datastore you need to access, and keep it until the program is done accessing that datastore. Pass it around as needed, or make it available somehow globally, but keep it open. And don’t Open() and Close() from a short-lived function. Instead, pass the sql.DB into that short-lived function as an argument." dans le main ?  
       ["The returned DB is safe for concurrent use by multiple goroutines and maintains its own pool of idle connections. Thus, the Open function should be called just once. It is rarely necessary to close a DB."](https://pkg.go.dev/database/sql#Open)  
       ["It is rare to Close a DB, as the DB handle is meant to be long-lived and shared between many goroutines."](https://pkg.go.dev/database/sql#DB.Close)

> ouvrir la base de données une fois et l'utiliser tel que, si vraiment necéssaire la fermer à la fin de l'application (dans le main)

- [x] "always prepare queries" mais peut présenter des risques de sécurité si le pool de connection est plein

- [x] "To verify that the data source name is valid, call Ping."

- [x] "Use Exec(), preferably with a prepared statement, to accomplish an INSERT, UPDATE, DELETE, or another statement that doesn’t return rows."

- [x] utiliser une sqlite en mémoire pour les tests ":memory:", "file::memory:?cache=shared" ou "file:test.db?cache=shared&mode=memory"

## Documentation de l'API

### Tous les posts

```
curl http://127.0.0.1:8000/api/v1/posts
```

### Ajouter un post

```
curl --location --request POST 'http://127.0.0.1:8000/api/v1/post' --header 'Content-Type: application/json' --header 'Authorization: Bearer qdyg.7dhq.djqsik' --data-raw '{"title":"from json","datas":"datasfill"}'
```

### Post par id

```
curl http://127.0.0.1:8000/api/v1/post/2
```

### Mettre à jour un post

```
curl --location --request PUT 'http://127.0.0.1:8000/api/v1/post/2' --header 'Content-Type: application/json' --header 'Authorization: Bearer qdyg.7dhq.djqsik' --data-raw '{"title":"from json","datas":"datasfill"}'
```

### Supprimer un post

```
curl --location --request DELETE 'http://127.0.0.1:8000/api/v1/post/2' --header 'Authorization: Bearer qdyg.7dhq.djqsik'
```

### Ajouter un utilisateur

```
curl --location --request POST 'http://127.0.0.1:8000/user' --header 'Content-Type: application/json' --data-raw '{"email":"user1@mail.lan","password":"VERYstrong&Secur3"}'
```

### Ask for JWT

```
curl --location --request POST 'http://127.0.0.1:8000/user/jwt' --header 'Content-Type: application/json' --data-raw '{"email":"user1@mail.lan","password":"VERYstrong&Secur3"}'
```

## Problèmes

L'application est fonctionnelle mais son écriture a relevé certains problèmes de structures et d'organisations. Des lessons à connaitre pour les prochaines écritures.

- On aurait dû commencer par la route utilisateur et la création des JWT avant les routes Posts, il y a aurait eu ainsi moins de modifications 'a la volée' dans les tests et les controlleurs existants. On aurait eu également une meilleure vision de la structure que l'application et cela aurait falliciter le refactoring.

- Trop de tests ou mal organisés, les tests écrits au fur et mesure et jamais vraiment factorisés n'ont pas falliciter la compréhension du code et l'objectif de certains ajouts de fonctionnalités. Les tests manquait parfois de descriptions, difficile donc de déterminer quels tests échouaients. On à tester tous les cas de figures mais certains tests étaient peut-être redondants, voire inutiles.

- Mettre en place rapidement le debug dans les logs aurait faciliter la lecture des tests. Les tests étaient noyés dans les messages d'erreurs des tests d'échecs. Manque donc de lisibilité à la sortie.

- On a pas vraiment appliqué le principe TDD, par difficulté ou par feignantise, les tests étaient crées juste après l'écriture de la fonction pour validation. Pour la plupart des cas ce sont des tests unitaires pas du TDD.
  Seule la partie Test_Fails à vraiment été faites en TDD, et en encore puisque on crée tout les cas de figures à l'avance (tant que l'on les avait en tête).

- Refactoriser au fur et à mesure aurait permis d'éviter d'avoir un code spagetti aussi vite. La refactorisation va maintenant être douloureuse.

- Manque d'utilisations des fonctionnalités de go. Les structs et surtout les pointers ont sous-utilisés. Les struct auraient apportés plus de structures/organisations, à revoir lors de la refactorisation. Il manque également le réflexe d'utiliser les pointers, ce n'est pas encore naturel et pas encore forcément compris.
