# Orca multitask
##### Multitask calculation with Orca (quantum package): https://www.orcasoftware.de/

## Install
Download the sources and compile with the command:
```
go build cmd/main.go
```
## Setting
Customize the file **configs/config.yml**
**orca_path** - path to orca executable
**work_path** - shared folder for your calculations

If there are no other folders in the working folder, the program will try to run the calculation from the same directory (**one task**). If there are other folders in the working folder, then the program will try to run the calculation from each folder, but no further than the first level (**many tasks**).

Create a .env file. and set it up, you will need mail to send messages (look **.env.example**). Currently only **gmail.com** is supported.

## Run

Run the compiled file. In case of errors, you can see the logs in the folder **logs**.