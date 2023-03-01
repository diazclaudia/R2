# R2

## Architecture diagram

The Architecture was based in clean architecture, I try to separate the layers for rest and domain, in this case repository does not exists.

![image](https://user-images.githubusercontent.com/16843197/222171331-e250e50f-061f-4e45-be1f-be1312c9a03c.png)

## Sequence diagram

In this diagram we can see the flow of the logic since user to the algorithm.

![image](https://user-images.githubusercontent.com/16843197/222171450-11052bcc-1eb1-4faa-abea-348d55d422af.png)


## Running local

* git clone https://github.com/diazclaudia/R2.git
* open project and put in terminal in the root directory `go mod tidy`
* Run the project
* postman curls:

#### Login
```
curl --location --request POST 'http://localhost:4000/login' \
--header 'user: admin' \
--header 'pass: admin' \
--header 'Cookie: mysession=MTY3NzY4MDIwNnxEdi1CQkFFQ180SUFBUkFCRUFBQUlfLUNBQUVHYzNSeWFXNW5EQVlBQkhWelpYSUdjM1J5YVc1bkRBY0FCV0ZrYldsdXyLaIONWC3LD46t2VFtGYskcFc7ZGUsQc3j5TTpctqyiQ=='
```

#### Logout
```
curl --location --request POST 'http://localhost:4000/logout' \
--header 'user: admin' \
--header 'pass: admin' \
--header 'Cookie: mysession=MTY3NzU1MDUyM3xEdi1CQkFFQ180SUFBUkFCRUFBQUlfLUNBQUVHYzNSeWFXNW5EQVlBQkhWelpYSUdjM1J5YVc1bkRBY0FCV0ZrYldsdXyigBpbxjmFkIbCp2u-aOJIDm1aVCY_Oh6_8o53v0-3fg=='
```

#### Fibonacci
```
curl --location --request POST 'http://localhost:4000/spiral?rows=2&cols=2' \
--header 'Cookie: mysession=MTY3NzU1MDUyM3xEdi1CQkFFQ180SUFBUkFCRUFBQUlfLUNBQUVHYzNSeWFXNW5EQVlBQkhWelpYSUdjM1J5YVc1bkRBY0FCV0ZrYldsdXyigBpbxjmFkIbCp2u-aOJIDm1aVCY_Oh6_8o53v0-3fg=='
```
