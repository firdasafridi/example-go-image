# example-go-image
This is for sharing purpose

## Create an image docker
docker build --tag app-golang:1.0 .

## For running the image
- docker images
- docker container create --name app1 -p 8080:8080 app-golang:1.0
- docker container start app1

## Push to image
- docker tag local-image:tagname reponame:tagname
- docker push reponame:tagname

## Sample
- docker tag app-golang:1.0 firdasafridi/go-hello:1.0
- docker push firdasafridi/go-hello:tagname