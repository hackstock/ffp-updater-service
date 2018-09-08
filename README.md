# About This Service
FFP-Updater-Service is daemon that runs x(configurable) times a day to update points accrued for users of the frequent flyer program and also calculate missing miles for miles flown before joining the frequent flyer program.

# Configurations
To deploy this service, the following environment variables should be defined. Otherwise, the service will panic on startup.
- FFPUPDATER_PORT           : this env var defines the port number on which the service runs. eg. 9000
- FFPUPDATER_ENV            : this env var sets the level of the logger used by the service. eg. development | production
- FFPUPDATER_SYNC_FREQUENCY : this env var defines the number of times the service should run in every 24 hours. eg. 1 | 2 | 3 
- FFPUPDATER_DATABASE_URI   : this env var sets the database connection string for the db from which the service will read flight information.

# How To Deploy With Docker
- Clone this repository into your GOPATH
- CD into the project directory
- Define the env vars listed above in the Dockerfile. 
- RUN docker-compose up
- Have fun!

# How TO Deploy Without Docker (No Recommended)
- Clone this repository into your GOPATH
- CD into the project directory
- Define the env vars listed above in your system's terminal
- Run go build -race .
- Run ./ffp-updater-service
- Have fun!

