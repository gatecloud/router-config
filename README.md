# router-config
A tool for routing file configuration


## Usage

### On-premise

Make sure that you have Chrome in your computer before running the `router-config.bat` file  



## How to deploy into a cloud docker manually



1. Compile the program in the local virtual machine within golang environment.  
You can use the command `CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o <name> .` to specify the executable file's name  
2. Create a project folder in the target EC2.For example `/ubuntu/roconfig`  
3. Copy the executable file (e.g. `roconfig`), docker-compose file, Dockerfile and other javascript or css files into the folder which is created in the step 2
4. Modify the permission of the folder and the file. The commands are 
```
sudo chown -R ubuntu:ubuntu roconfig
sudo chmod 400 roconfig
```
5. Build docker image by the command `docker build -t image-roconfig:latest .`  
6. Run the docker image. `sudo docker-compose up -d`



## Other configuration  

In order to run the program smoothly after deployment, we need to update the URL in `/public/xx.js` files  

### PostgreSQL configuration  

1. Install postgreSQL into the target server  
2. Go to `cd /etc/postgresql/10/main/pg_hba.conf` to configure the allowed IPv4 address  
3. Go to `cd /etc/postgresql/10/main/postgresql.conf` to open all listening port  





