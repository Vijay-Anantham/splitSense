SPLIT WISE

Add aplit expensees among friends and groups at ease/ Track spending over months weeks and cultivate a healthy spending behaviour.

### How to run

-> Intall golang <br>
-> get all the dependencies `go mod tidy` <br>
-> Get the db ready `docker-compose up -d` <br>
-> run backend `cd ./backend` and `go run ./cmd/main.go` <br>

### Schema Design MongoDb

--->  User
{
    id: string
    name: string
    groups: [] -> Usercan be part of only so many groups and not otherwise
}

---> Groups
{
    id: string
    name: string
    state: string
}

--> spending
{
    id: string
    name: string
    category: string
    group: string
}


