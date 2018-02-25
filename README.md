# yetAnotherCRUD application

## Run
#### Clone repository
    $ git clone https://github.com/8tomat8/yetAnotherCRUD.git
    
#### Change dir
    $ cd ./yetAnotherCRUD
    
#### Pull dependencies
    # github.com/golang/dep has to be installed
    $ dep ensure --vendor-only
    
#### Create DB structure
    $ mysql -uroot -h127.0.0.1 -p < migrations/base.sql

#### Build app
    $ go build -o yetAnotherCRUD ./cmd/server/
    
#### List available args
    $ ./yetAnotherCRUD -help
    Usage of ./yetAnotherCRUD:
      -apihost string
        	Host for API listener (default "0.0.0.0")
      -apiport string
            Port for API listener (default "8080")
      -dbhost string
            MySQL host (default "127.0.0.1")
      -dbpass string
            MySQL password (default "root")
      -dbport string
            MySQL port (default "3306")
      -dbuser string
            MySQL user (default "root")
      -shutdownTimeout duration
            Shutdown timeout (default 30s)   
 
#### Run app
    $ ./yetAnotherCRUD
    INFO[0000] Starting listener on 0.0.0.0:8080

## Test API
#### Search users
    $ curl -vvv -X GET \
              'http://127.0.0.1:8080/api/users?sex=female&username=Username&age=50' \
              -H 'cache-control: no-cache'
      Note: Unnecessary use of -X or --request, GET is already inferred.
      *   Trying 127.0.0.1...
      * TCP_NODELAY set
      * Connected to 127.0.0.1 (127.0.0.1) port 8080 (#0)
      > GET /api/users?sex=female&username=Username&age=50 HTTP/1.1
      > Host: 127.0.0.1:8080
      > User-Agent: curl/7.54.0
      > Accept: */*
      > cache-control: no-cache
      >
      < HTTP/1.1 200 OK
      < Date: Sun, 25 Feb 2018 18:39:08 GMT
      < Content-Length: 2
      < Content-Type: application/json; charset=UTF-8
      <
      * Connection #0 to host 127.0.0.1 left intact
    []
      
    OR
    
    $ curl -X GET \
            'http://127.0.0.1:8080/api/users' \
            -H 'cache-control: no-cache'
    [{"UserID":5,"Username":"1111","Password":"sss","Firstname":"ddd","Lastname":"sss","Sex":"male","Birthdate":"15-12-1991"},{"UserID":6,"Username":"1111","Password":"sss","Firstname":"ddd","Lastname":"sss","Sex":"female","Birthdate":"15-12-1991"},{"UserID":1,"Username":"aaa","Password":"sss","Firstname":"ddd","Lastname":"sss","Sex":"male","Birthdate":"15-12-1991"},{"UserID":2,"Username":"aaaAAAAA","Password":"sss","Firstname":"ddd","Lastname":"sss","Sex":"male","Birthdate":"15-12-1991"},{"UserID":4,"Username":"OOOOOOOOO","Password":"sss","Firstname":"ddd","Lastname":"sss","Sex":"male","Birthdate":"15-12-1991"}]
      
#### Create User
    $ curl -X POST \
        http://127.0.0.1:8080/api/users \
        -H 'cache-control: no-cache' \
        -d '{"UserID":1,"Username":"1111","Password":"sss","Firstname":"ddd","Lastname":"sss","Sex":"female","Birthdate":"15-12-1991"}'
    {"UserID":7,"Username":"1111","Password":"sss","Firstname":"ddd","Lastname":"sss","Sex":"female","Birthdate":"15-12-1991"}
    
#### Update User
    $ curl -X PUT \
        http://127.0.0.1:8080/api/users/4 \
        -H 'cache-control: no-cache' \
        -d '{"UserID":1,"Username":"OOOOOOOOO111111","Password":"sss","Firstname":"ddd","Lastname":"sss","Sex":"male","Birthdate":"15-12-1991"}'
    {"UserID":4,"Username":"OOOOOOOOO111111","Password":"sss","Firstname":"ddd","Lastname":"sss","Sex":"male","Birthdate":"15-12-1991"}
    

#### Delete User
    $ curl -v -X DELETE \
        http://127.0.0.1:8080/api/users/4 \
        -H 'cache-control: no-cache'
      *   Trying 127.0.0.1...
      * TCP_NODELAY set
      * Connected to 127.0.0.1 (127.0.0.1) port 8080 (#0)
      > DELETE /api/users/4 HTTP/1.1
      > Host: 127.0.0.1:8080
      > User-Agent: curl/7.54.0
      > Accept: */*
      > cache-control: no-cache
      >
      < HTTP/1.1 204 No Content
      < Content-Type: application/json; charset=UTF-8
      < Date: Sun, 25 Feb 2018 18:49:24 GMT
      <
      * Connection #0 to host 127.0.0.1 left intact
