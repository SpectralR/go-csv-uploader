**Projet 2 — Service de traitement de jobs asynchrones (niveau : conditions réelles cloud-native)***

Contexte : une entreprise reçoit des fichiers (ex : exports CSV) qu'il faut traiter en arrière-plan sans bloquer l'API principale — cas d'usage très courant en poste Go.

**Cahier des charges :**

- Un service API qui reçoit un fichier, le stocke, et pousse un job dans une queue
- Un ou plusieurs workers qui consomment la queue et traitent le fichier (parsing, validation, écriture résultat en base)
- Queue via Redis (liste ou pub/sub) ou RabbitMQ
- Gestion de la concurrence : plusieurs workers doivent pouvoir tourner en parallèle sans traiter deux fois le même job (utilise goroutines + channels)
- Endpoint pour consulter le statut d'un job (pending / processing / done / failed)
- Gestion des erreurs et retry en cas d'échec de traitement
- Logs structurés (JSON) et un endpoint /health
- Déploiement local via docker-compose (API + workers + Redis + Postgres)

**Bonus si tu veux aller plus loin :**

Ajoute un rate limiter sur l'API
Déploie le tout sur un cluster Kubernetes local (minikube/kind) avec un manifest simple