# SI4404_KEL240_FARMTIZEN_Backend

## 1. Introduction
This is the backend of the SI4404_KEL240_FARMTIZEN project using go as backend. This project is a web-based application that is used to manage the farm. This application is made to help farmers to manage their farm easily. This application is made by using the Laravel framework for the frontend. This application is made by the SI4404_KEL240_FARMTIZEN team. This application is made for the final project of the Capstone Project course.

## Architecture
This project is using the DDD (Domain Driven Design) architecture. why we use DDD? with DDD, we can make the code more readable and maintainable.
### Project Structure
```
├───cmd
└───internal
    ├───app
    ├───config
    ├───controller
    ├───domain
    │   ├───entity
    │   └───repository
    ├───dto
    ├───middleware
    └───service
```

### Explanation
- cmd: this folder is used to store the main function of the application
- internal: this folder is used to store the internal code of the application
- internal/app: this folder is used to store the application code
- internal/config: this folder is used to store the configuration of the application
- internal/controller: this folder is used to store the controller of the application
- internal/domain: this folder is used to store the domain of the application
- internal/domain/entity: this folder is used to store the entity of the application aka the model
- internal/domain/repository: this folder is used to store the repository of the application, repository is used to interact with the database
- internal/dto: this folder is used to store the dto of the application, dto is used to store the data that will be sent to the frontend without using model as body request
- internal/middleware: this folder is used to store the middleware of the application
- internal/service: this folder is used to store the service of the application, service is used to store the business logic of the application


