# books

I use existing libs :

 - Ozzo Validation, for input request validation
 - Godotenv, for env loader
 - Testify
 - etc

> I use Redis in assumption that the get book's data by subject is occur frequently, so the processes will be light.
> and I use a simple message queue (in this case I use Nats) to handle pickup schedule's request, so the processes will do asynchronous in assumption there will be many request so the processes allows it to continue running and responding to user requests without having to wait for operations that take a long time to finish executing.

# how to run

- to install some supporting technology stuff (Redis & Nats), i have created "docker-compose.yaml" file. You can adjust the ports, to the ports that you will use on the local device. And configure the .env file to match the port you are using.
- go to the root directory project, than do "docker-compose up" to install the supporting stuff. wait until the process is finished.
- still on root directory, do "go mod tidy" to import all the used packages. wait until the process is finished.
- last but not least, do "go run main.go"

# end points

these are some endpoint that I've made for this service :

> to get books by subject
- curl --location 'http://localhost:8080/api/books?subject=love'

> to make pickup schedule
- curl --location 'http://localhost:8080/api/pickup' \
--header 'Content-Type: application/json' \
--data '{
    "date": "12-05-2023",
    "user": "jody",
    "information": [
        {
            "title": "Wuthering Heights",
            "cover_id": 12818862,
            "edition_count": 1930,
            "authors": [
                {
                    "Name": "Emily Brontë"
                }
            ]
        }
    ]
}'

> to get all pickup schedule
- curl --location 'http://localhost:8080/api/pickup'

- or if you don't want to use curl, you can import the collection from this :
> https://api.postman.com/collections/5535920-38901aad-4e88-4acf-a8d8-e293a94824df?access_key=PMAT-01H096JEP932HYWB331H275EYV

# unit testing
- I've made some unit testing in business process/usecase layer (books & pickup). each unit test has reached 100% coverage
> Running tool: /usr/local/go/bin/go test -timeout 30s -coverprofile=/var/folders/h_/tjhvlj3n32sc9lvvfbm8x9ym0000gn/T/vscode-go5U5I9q/go-code-cover books/src/app/usecase/pickup

>> ok  	books/src/app/usecase/pickup	0.711s	coverage: 100.0% of statements

> Running tool: /usr/local/go/bin/go test -timeout 30s -coverprofile=/var/folders/h_/tjhvlj3n32sc9lvvfbm8x9ym0000gn/T/vscode-go5U5I9q/go-code-cover books/src/app/usecase/books

>> ok  	books/src/app/usecase/books	0.794s	coverage: 100.0% of statements

- you can see the coverage testing by open the project with vscode, choose the testing file, right click on an empty space anywhere in the file, then choose "Go:Toogle Test Coverage in Current Package"

