creation d'un site pour la présentation d'une coproprieté, emplacement, evenements, boutique  
avec une partie administration pour gérer le contact avec les résidents, l'envoi de mails, planification des evenements

## Fonctionnalités

- [ ] page de garde en statique
- [ ] accès administrateur, président, secretaire, trésorier
- [ ] gestion des logements, propriétaire et/ou locataire
- [ ] gestion des commerces
- [ ] envoi de photos
- [ ] envoi de mails groupés
- [ ] partie "blog" avec flux rss
- [ ] comptes-rendu des réunions
- [ ] calendrier des événements
- [ ] planning du gardien

## Stack Technique

~~L'objectif est d'apprendre un framework web fait en go. Le projet est un prétexte, mais c'est un cas plus interressant et concret qu'une application de todos.~~
~~On part sur [Gin](https://gin-gonic.com/docs/quickstart/) uniquement parce que c'est le plus populaire.~~
~~[Buffalo](https://github.com/gobuffalo/buffalo) aurait également été un choix possible.~~

~~On pratique également testify, le paquet de test, utile pour les mockup, et probablement GORM, l'orm le plus utilisé pour les bases de données.~~

*C'est la stack du mainstream, on ignore les problèmes possibles des dépendances*

Ne pas faire un MVC, mais utiliser une organisation par domaine, avec un paquet par domaine. La gestion des paquets dans go *semble* particulière adaptée à ce modèle.

Gin est un outil puissant et largement utilisé, mais il est surtout utilisé pour faire des API. Ce n'est pas ce qu'on recherche là.  
On recherche un projet plus "borné", qui nous donnent des directions, un "vrai" framework web type laravel ou adonis. Dans ce cas on devrait plus s'orienter vers [Buffalo](https://github.com/gobuffalo/buffalo), [Revel](https://github.com/revel/revel), [Beego](https://github.com/beego/beego/).

