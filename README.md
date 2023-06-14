# Cassandra CRUD 
## Installation

### Git installation

In a terminal:
```bash
brew install git
```
1.  Copy the text below, replacing "YOUR_EMAIL" with your GitHub account email address.
    
    ```shell
    $ ssh-keygen -t ed25519-sk -C "YOUR_EMAIL"
    ```
    
    **Note :**  If the command fails with the error "invalid format" or "feature not supported," you may be using a hardware security key that does not support the Ed25519 algorithm. In that case, use the following command instead.
    
    ```shell
    $ ssh-keygen -t ecdsa-sk -C "your_email@example.com"
    ```
    
2.  When prompted, press the button on your hardware security key.
    
3.  When prompted to enter a file to save the key, press Enter to accept the default file location.
    
    ```shell
    > Enter a file in which to save the key (/Users/YOU/.ssh/id_ed25519_sk): [Press enter]
    ```
    
4.  When prompted to enter a passphrase, press **Enter**.
    
    ```shell
    > Enter passphrase (empty for no passphrase): [Type a passphrase]
    > Enter same passphrase again: [Type passphrase again]
    ```
    
5.  Add the SSH key to your GitHub account. For more information, refer to « [Adding a new SSH key to your GitHub account](https://docs.github.com/fr/authentication/connecting-to-github-with-ssh/adding-a-new-ssh-key-to-your-github-account) ».


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