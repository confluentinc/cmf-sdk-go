## 0.x.x


Introduces support for the following APIs:

### Flink Environment 

- GET `/environments`: Used to list all the enviroments
- POST `environments`: Used to update/create environment
- DELETE `/environments/{envName}`: Used to delete the environment
- GET `/environments/{envName}`: Used to describe the environment
- GET `/environments/{envName}/applications`: Used to list all the applications
- POST `/environments/{envName}/applications`: Used to create/update an application
- DELETE `/environments/{envName}/applications/{appName}`: Used to delete an application
- GET `/environments/{envName}/applications/{appName}`: Used to describe an application
- POST `/environments/{envName}/applications/{appName}/start`: Used to start an application
- POST `/environments/{envName}/applications/{appName}/suspend`: Used to suspend an application