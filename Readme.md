Project Structure:
-----------------
- init.sh             : Shell script with installation steps for Go
- trigger            : Go executable file generated from automate.go
- trigger.go         : Go source code file containing the main application logic
- go.mod              : Go module file specifying the project's dependencies
- go.sum              : Go module file containing the expected cryptographic hashes of the content of specific module versions
- deploy.sh           : Shell script executed by trigger.sh to execute the whole CD process
- trigger.sh          : Shell script that triggers pipeline.sh
- go_runner.conf      : Supervisor configuration file for running the Go application

```
    ├── pipeline/
    │   ├── automate/                  # Go application files
    │   │   ├── trigger.go
    │   │   ├── go.mod
    │   │   └── go.sum
    │   ├── scripts/                   # Scripts related to the pipeline
    │   │   ├── init.sh                 # Installation steps for Go
    │   │   ├── deploy.sh             # Pipeline trigger script the main CD script
    │   │   └── trigger.sh              # Script triggered by Go application
    │   └── README.md                  # Documentation for the pipeline
```
Description:
------------
This project is a Go application that runs a web server and executes a pipeline process triggered by an HTTP request. It uses the Gin web framework for handling HTTP requests.

The main logic is implemented in the automate.go file. When the "/task" endpoint is accessed, it executes the trigger.sh script, which in turn executes the pipeline.sh script responsible for the continuous delivery (CD) process.

Installation:
-------------
1. Make sure you have Go installed on your system.
2. Run the init.sh script to install any necessary dependencies for Go.

Running the Application:
------------------------
1. Build the Go application by running the following command:

This will generate the automate executable file.

2. Start the Go application using Supervisor:
- Create or edit the go_runner.conf file in the /etc/supervisor/conf.d/ directory.
- Replace the "command" value with the actual path to the automate executable file.
- Replace the "directory" value with the actual path to the project directory.
- Update other configurations as needed.
- Save the file.

-----------template-go-runner-file found at path /etc/supervisor/conf.d/go_runner.conf -------------------
```
    [program:go_runner]
    command=/path-to-build/automate  #automate  # Replace with the actual path to your Go script runner executable
    directory=/directory-of-path/pipeline  # Replace with the actual path to your Go script runner directory
    autostart=true
    autorestart=true
    startsecs=10
    startretries=3
    user=legal-ai-dev  # Replace with your username
    redirect_stderr=true
    stdout_logfile=/var/log/supervisor/go_runner.log

```
3. Start Supervisor to run the Go application:

sudo supervisorctl reread

sudo supervisorctl update

sudo supervisorctl start go_runner


4. The Go application will be running, and you can access it by making a GET request to the "/task" endpoint.

Logs:
-----
The application logs are redirected to the /var/log/supervisor/go_runner.log file.

NGINX Configuration for Redirecting Traffic:
--------------------------------------------
Assuming you have NGINX installed and running on your server, follow these steps to configure NGINX to redirect traffic to your Go application using a specific URL.

1. Edit the NGINX configuration file. The default location of the configuration file is usually "/etc/nginx/nginx.conf". You can use a different file if you have a custom configuration.

2. Inside the "http" block, add a new "location" block to define the redirection rule. The example below demonstrates how to redirect traffic to the "/task" URL to your Go application running on a specific IP address and port:

   server {
       # Existing server configuration...

       location /task {
           proxy_pass http://ip-of-vm:8080/task;
       }

       # Other location blocks and configurations...
   }

   # Other server blocks and configurations...


Replace "vm-of-ip:8080" with the actual Private IP address and port where your Go application is running.

3. Save the NGINX configuration file.

4. Restart the NGINX server to apply the changes. The command to restart NGINX depends on your operating system. Here are a few examples:

- Ubuntu or Debian:
  ```
  sudo service nginx restart
  ```

or restart the docker-compose file of nginx

Make sure to use the appropriate command based on your system.

5. NGINX is now configured to redirect traffic to your Go application when accessing the specified URL. Test it by making a GET request to the "/task" URL.


