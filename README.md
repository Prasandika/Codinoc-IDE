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

## Server Setup

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

Just ignore `Server/bgo` directory, if you installing GO compiler using package managers

### 1. Base Setup

1. Create temporary global variable in `.bashrc` file, `export c_path=/codinoc_server`
2. Test it using `echo $c_path`
3. Create Main directory using `sudo mkdir $c_path`
4. Add linux user permission to that directory using `sudo chown -R [USER_NAME]:[USER_NAME]`
5. Copy all project files and folders into `$c_path` directory

### 2. Permission Settings

```sh
sudo chmod -R +rwx $c_path/bgo/bin && \
sudo chmod -R +rwx $c_path/bpg/bin
```

### 3. Global Variable Settings

### 4. Custom Tool Installation

```sh
sudo yum install glib* && \
sudo dnf provides */libncurses.so.5 && \
sudo dnf install ncurses-compat-libs
```

### 5. GO Language Installation

1. [Follow this tutorial to install GO Language in any linux distro](https://golangdocs.com/install-go-linux) or [this one](https://linuxtect.com/how-to-install-go-golang-in-linux/)
2. Setup GO Packages used by the project

```sh
go env -w GO111MODULE=auto
```

```
go get github.com/gorilla/mux &&
go get github.com/gorilla/handlers
```

### 6. PostgreSQL Database Installation

Download and extract PostgreSQL archive pn `/Server/bpg` directory

We use PostgreSQL 10.20 Linux Binary Tarball here and you can download it using [this link](https://sbp.enterprisedb.com/getfile.jsp?fileid=1257992)

1. Initialize data directory

```sh
cd $c_path/bpg/bin
```

```sh
./pg_ctl -D $c_path/dta initdb
```

2. Start PostgreSQL server

```sh
./pg_ctl -D $c_path/dta -l logfile start
```

3. Stop PostgreSQL Server

```sh
./pg_ctl -D $c_path/dta -l logfile stop
```

4. Running status of PostgreSQL

```sh
pgrep pg_ctl
```

5. Login to PostgreSQL Server

```sh
./psql [DATABASE-NAME]
```

> Note:\
> On the first login, need to create a Database

```sh
./psql template1
CREATE DATABASE sample_db;
```

### 7. PostgreSQL Databases Setup

TODO

### 8. Local Development Server additional Setup

Add these content into the `hosts` file using `sudo nano /etc/hosts`

```
# Codinoc IDE

127.0.0.1:8080 codinoc.com
```

So after that, we can access to site using `codinoc.com:8080/` in web browser