# Backend-photagea
golang backend for photagea webapp

# Setup Project

## Prerequisites
Install go: https://golang.org/doc/install

## 1. Install Mysql
- Download and run the installer: https://dev.mysql.com/downloads/installer/
- Leave default selections
- Set the root password, create at least one additional user: dev,
password: developmentpassword. Or edit main.go:29 for your own custom
account.
- I recommend installing mysql workbench and using it for all
developer-database interactions.

## 2. Import Development database
From mysql workbench:
- In the navigator sidebar on the right, click data import/restore.
- Select `import from self-contained file`.
- Open "photagea.sql" from the project root.


