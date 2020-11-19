# MarbleFileServer
**Version 1.0.0**

A File server built with golang and HTML.

Connect to a servers ip via a browser to upload and download files anywhere, supports postgre SQL to allow for secure logins.

### Built With

* [Golang](https://golang.org/)

## Getting Started

MarbleFileServer is a simple file server software that runs on Linux but may also support Linux. It only requires a few changes in order to get it to work.

### Prerequisites

Currently runs on Linux distros but should run on windows if you build it from source.

### Installation

#### From release
Download the latest release and extract the contents.

```sh
sudo tar -xvf MarbleFileServer_v0.1.tar.gz
```

Edit the config.yaml file to add your drives you wish for MarbleFileServer to have access tp. Add the postgresql username, password and databasename. Users will need to be added manually within the server it also uses bcrypt to store read the passwords.

Simply run the server in the dir

```sh
./MarbleFileServer
```

Open a browser and go to the servers local ip address. port 8080 will need to be open if you want to access the server outside of the network.

#### From Source
clone the master branch to build it yourself on your chosen operating system that supports golang.

```sh
git clone https://github.com/ChrisPHP/MarbleFileServer.git
```

need to simply unzip the contents and use golang to compile the source code.

```sh
go build
```

### Contributors

- Chris Page <17636418@students.lincoln.ac.uk>

### License
Distributed under the MIT License. See `LICENSE` for more information.

