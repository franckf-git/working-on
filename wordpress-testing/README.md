podman-compose up --down

podman-compose down

podman exec --user=root --interactive --tty wordpress-testing_db_1 /bin/bash

ls ~/.local/share/containers/storage/volumes/wordpress-testing_db_data/_data/
