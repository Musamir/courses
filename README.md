# courses

<p align="center" width="100%">
    <img src="https://user-images.githubusercontent.com/43841786/126863750-3e165586-6a7f-4daa-a1aa-7aa2563df036.png" alt="welcome-PNG32" border="0">
</p>

# Courses - Course management system

This is a simple course management system.

## Contents
- [Courses](#Courses)
  - [Contents](#contents)
  - [Quick start](#quick-start)
  - [Swagger](#swagger)


## Quick start

1. You need to install [Go](https://golang.org/) compiler(version 1.13+ is required) and set your Go environment.

2. Set your environment variables Courses_DATABASE_URL(postgres url), PORT(port for http) or specify them in the configuration file config/conf.json

3. Open your command line in the program folder and use the below Go command to run the program.

```sh
$ go run cmd/app/main.go
```

## Swagger
Run the program, and browse to [http://localhost:8085/swagger/index.html](http://localhost:8085/swagger/index.html). You will see Swagger 2.0 Api documents as shown below:
![swagger screen](https://user-images.githubusercontent.com/43841786/126864073-ed84bf65-feec-4799-b970-8fb2317c316d.png)