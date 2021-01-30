launch compose
```
podman-compose up
```

stop compose
```
podman-compose down
```

enter a container
```
podman exec --user=root --interactive --tty wordpress-testing_db_1 /bin/bash
```

locate db files
```
ls ~/.local/share/containers/storage/volumes/wordpress-testing_db_data/_data/
```


show active pod
```
podman pod ls
POD ID        NAME               STATUS   CREATED        INFRA ID      # OF CONTAINERS
aaaaaaaaaaaa  wordpress-testing  Running  2 minutes ago  bbbbbbbbbbbb  3
```

show pods
```
podman ps
CONTAINER ID  IMAGE                               COMMAND               CREATED             STATUS                 PORTS                                         NAMES
cccccccccccc  docker.io/library/wordpress:latest  apache2-foregroun...  About a minute ago  Up About a minute ago  0.0.0.0:8080->80/tcp, 0.0.0.0:6603->3306/tcp  wordpress-testing_web_1
dddddddddddd  docker.io/library/mariadb:10.5      mysqld                About a minute ago  Up About a minute ago  0.0.0.0:8080->80/tcp, 0.0.0.0:6603->3306/tcp  wordpress-testing_db_1
bbbbbbbbbbbb  k8s.gcr.io/pause:3.2                                      2 minutes ago       Up About a minute ago  0.0.0.0:8080->80/tcp, 0.0.0.0:6603->3306/tcp  aaaaaaaaaaaa-infra
```

transform pod and pods to kube
```
podman generate kube wordpress-testing >> wordpress-testing-kube.yaml
```

launch pod with kube
```
podman play kube wordpress-testing-kube.yaml
```

> and it is fu****ed : "Error establishing a database connection"

