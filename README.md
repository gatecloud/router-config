# router-config
A tool for routing file configuration


## Usage

### On-premise

Make sure that you have Chrome in your computer before running the `router-config.bat` file  



## How to deploy into a cloud docker manually


1. Copy project folder from local to virtual machine where there is a golang environment  
2. Execute `CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o roconfig`, the executable file `roconfig` will be generated  
3. Copy `roconfig` to the specific path of the target cloud server, like `~/roconfig`, along with other folder and files, including `/public`, `/templates`, `.env`, `docker-compose.yml` and `Dockerfile`
4. Execute `sudo docker build -t image-roconfig:latest` to build roconfig's docker image. (The cloud server should install docker)   
5. Run `sudo docker-compose up -d` to start docker  



## Other configuration

### PostgreSQL configuration  

1. Install postgreSQL into the target server  
2. Go to `cd /etc/postgresql/10/main/pg_hba.conf` to configure the allowed IPv4 address  
3. Go to `cd /etc/postgresql/10/main/postgresql.conf` to open all listening port  





