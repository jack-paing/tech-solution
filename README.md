## Question 1 and 2
Question 1 and 2 are inside the folder called algo.

## Question 3
Sample Schema is in db/schema.sql

## Question 4

### Project Structure
/app --> the entry point or initialization for main.go

/configs --> config.yaml file for adding config value

/db --> sql file for migration of db

/internal

1./card --> service and handler are inside
                
2./config --> loading the config file

3./health --> health check 

4./repo --> dabase related layer


/model --> data model

/utils --> common util folder

### How to run the application

#### Pre-requisite
Mysql running with database created

#### Steps to run the application
1.Change necessary database information in config.yaml

2.`make migrate` for creating tables

3.`make run` for running the application

4.`user_service_collection.postman-collection.json` includes CRUD for cards

#### Steps to run tests
1.`make setup-test-db` for setting up test tables

2.`make unit-test` for running unit tests

### API Documentation
`user_service_collection.postman-collection.json` can serve as an API guide for developer. 