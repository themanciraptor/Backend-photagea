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
- I recommend installing mysql workbench and using it for all developer-database interactions.

## 2. Import Development database
From mysql workbench:
- In the navigator sidebar on the right, click data import/restore.
- Select `import from self-contained file`.
- Open "photagea.sql" from the project root.

## Running
- From the project root on the commandline, run: 
    
        go run .

- Obtain authentication token:
    - Post `127.0.0.1:8001/account/signin` with the following body:
`{
    "Email": "dood",
    "Password": "milk"
}`

- The authentication header on all requests should be set to "Bearer \<authentication token>"

## Project TODO's

Obviously, this project isn't production ready yet. I will add steps here as I think/learn of them:
- Enable HTTPS using openssl
- Passwords and secrets currently included as plain-text in the code. This should be stored as secrets and inserted at build time according to whichever environment.
- Endpoints should have documentation for how they are used. Usage could be returned on 400's.