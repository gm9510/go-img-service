curl http://localhost:3000/files | json_pp

#v1: curl -X POST http://localhost:3000/files | json_pp
#v2: curl -X POST http://localhost:3000/files -H 'Content-Type: application/json'  -d '{"name":"my_login","ext":"my_password", "size":4}' | json_pp
curl -X POST http://localhost:3000/files -F "file=@download.pdf" -H 'Content-Type: multipart/form-data'

curl -X GET http://localhost:3000/files/1 | json_pp

curl -X PUT http://localhost:3000/files/1 | json_pp

curl -X PATCH http://localhost:3000/files/1 | json_pp

curl -X DELETE http://localhost:3000/files/1 | json_pp
