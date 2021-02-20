ansible permet de recupérer des informations sur le système. cela permettrait de faire un outil de monitoring simple sans agent et sans installation de serveur sur la machine qui "surveille"

- une machine "surveillante" (qui peut être un poste perso) lance à intervals réguliers un playbook ansible pour récupérer des infos sur des systèmes
- les informations sont stockées dans une bdd sqlite3 (plus simple à gérer)
- un serveur web servira les informations stockées, avec une librairie js pour afficher les graphiques
- en cas d'alertes - dépassements de seuils on pourra envoyer un mail ou notification de bureau

> ansible peut être remplacé par l'envoi de commandes ssh type uptime, free, cpuinfo, ...

peut être écrit en n'importe quel language node, go, python
