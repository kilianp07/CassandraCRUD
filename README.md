# Cassandra CRUD 
## Installation

### Using Docker
1. Install Docker
    * Download Docker Desktop for Mac ARM64 from [Docker Hub](https://desktop.docker.com/mac/main/arm64/Docker.dmg?utm_source=docker&utm_medium=webreferral&utm_campaign=docs-driven-download-mac-arm64)
    * Double-click `Docker.dmg` to open the installer, then drag the Docker icon to the Applications folder.
    * Double-click Docker.app in the Applications folder to start Docker.
    * The Docker menu (whale menu) displays the Docker Subscription Service Agreement window. Select Accept to continue.
    * From the installation window, select: \
    Use recommended settings (Requires password). This let’s Docker Desktop automatically set the necessary configuration settings.

2. Run the container
```bash
git clone https://github.com/kilianp07/CassandraCRUD.git && cd CassandraCRUD && docker-compose up 
```

The API may crash because the database is not ready yet. \
You need to fill the database with the following tool : [CassandraSeeder](https://github.com/kilianp07/CassandraSeeder.git) \
Once the database is filled, you can restart the API with the following command : 
```bash
docker-compose restart
```

The API is now available on `http://localhost:8080`

## Usage
See [Postman doc](./postman.md) to test the API. \
You also can downloand the [Postman collection](./postman.json) and import it to your Postman.