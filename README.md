# jsonapi
Mock httpd for testing JSON stuff from other programs I work on.

## Generate a UUID
Uncomment in main.go:

    /*uuid, _ := uuid.NewRandom()  
    fmt.Printf("Generated UUID: %v\n", uuid)
    */

or use your own UUID generator (hint: the program does not validate UUIDs).  
  
## Ingest data
Use example file "data.json" (or modify this file and types in main.go for customization) and ingest with:  
`curl -XPOST -d @data.json localhost:8080/events/ingest/<UUID>` 
  
  
## Query data
`curl localhost:8080/events/<UUID>`

## Customize port
Use option -p to specify a port different from default `8080` .

## Build Docker image
`bash build`  
`docker build -t <repo/image:version> .` 

## Run Docker image
`docker run -p 8080:8080 <repo/image:version>`
