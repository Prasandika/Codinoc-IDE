> **Note:**\
> This is not a Personal Project by me [las-nish](https://github.com/las-nish/)\

# Developer Note for Project

## Project Changes

| Type | Details | Status | Changed to |
| -- | -- | -- | -- |
| Database | MariaDB | Changed | PostgreSQL DB |
| Concept | Multi Tenant Database Concept | Removed | Normal Database |
| Service | Domain Name | Removed | Just an IP Address |
| Service | AWS Server | Changed | Vultr Server |
| Software | JetBrains DataGrip | Changed | Visual Studio Code |
| Software | JetBrains GoLand | Changed | Visual Studio Code |

---

## Directory Structure

**Directory structure in `/Server`**

```
adm       -> Administrator section
bgo       -> GO language binary tarball
bin       -> Website binary executable
bpg       -> PostgreSQL binary tarball
cdn       -> Website file storage
cmd       -> GO file of the project
dta       -> Database warehouse
thm       -> Website complete theme
```

---

## Required Technologies

### Languages

| Language | Version | For What | Architecture | Platform |
| -- | -- | -- | -- | -- |
| GO | >= 17.7 | Backend | x86_64 | Linux |
| JS | Undefined | Front End | Undefined | Undefined |
| SCSS | >= 1.49.4 | Front End | Undefined | Undefined |

### Databases

| Database | Version | Architecture | Platform |
| -- | -- | -- | -- |
| PostgreSQL | >= 10.20 | x86_64 | Linux |

### GO Packages

| Library Name | Reason | Why | References |
| -- | -- | -- | -- |
| [Gorilla Handlers](https://github.com/gorilla/handlers) | Optimizing HTTP Server responses with GZIP compression | By sending a compressed response we save bandwidth and download time eventually rendering the page faster | Go Web Development CookBook, Page 48 |
| [Gorilla Mux](https://github.com/gorilla/mux) | Implementing HTTP request routing | Define more than one URL route in web application | Go Web Development CookBook, Page 66 |

### Web Libraries

### Web Frameworks

---

## Setup Development Server

### Directory Structure

**Directory Structure for Server**

```
Server
  |- adm : Administrator Directory
  |- bgo : Go Language Binary Tarball Directory
  |- bin : Web Site's bin Directory
  |- bpg : PostgreSQL Binary Tarball Directory
  |- cdn : Storage Directory
  |- cmd : GO Files Directory
  |- dta : Database Storage Directory
  |- thm : Theme Directory
  |- www : Site Directory
```

Just ignore `Server/bgo` directory, if you installed go compiler using package managers

### Install GO Compiler

[Follow this tutorial to install GO Language in any linux distro](https://golangdocs.com/install-go-linux) or [this one](https://linuxtect.com/how-to-install-go-golang-in-linux/)

### Install PostgreSQL Database

> Note:\
> We use PostgreSQL 10.20 Linux Binary Tarball here and you can download it using [this link](https://sbp.enterprisedb.com/getfile.jsp?fileid=1257992)

- After download, extract it on `/Server/bpg` directory

### Setup GO Language and Packages

```bash
go env -w GO111MODULE=auto

go get github.com/gorilla/mux &&
go get github.com/gorilla/handlers
```

### Setup Local Domain Configuration

Add these content into the `hosts` file using `sudo nano /etc/hosts`

```
# Codinoc IDE

127.0.0.1:8080 codinoc.com
```

## Setup Host Server