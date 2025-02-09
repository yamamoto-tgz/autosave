# Local Functions Development

https://cloud.google.com/run/docs/local-dev-functions

* Init
```
cd sandbox
go mod tidy
```

* Run
```
go run main.go
```

* Test (HTTP)
```
curl localhost:8080/hello
```

* Test (PubSub)
```
curl localhost:8080/event \
  -X POST \
  -H "Content-Type: application/json" \
  -H "ce-id: 123451234512345" \
  -H "ce-specversion: 1.0" \
  -H "ce-time: 2020-01-02T12:34:56.789Z" \
  -H "ce-type: google.cloud.pubsub.topic.v1.messagePublished" \
  -H "ce-source: //pubsub.googleapis.com/projects/MY-PROJECT/topics/MY-TOPIC" \
  -d '{
        "message": {
          "data": "SEVMTE8gRVZFTlQ=",
          "attributes": {
             "attr1":"attr1-value"
          }
        },
        "subscription": "projects/MY-PROJECT/subscriptions/MY-SUB"
      }'
```

* Test (Storage)
```
curl localhost:8080/event \
  -X POST \
  -H "Content-Type: application/json" \
  -H "ce-id: 123451234512345" \
  -H "ce-specversion: 1.0" \
  -H "ce-time: 2020-01-02T12:34:56.789Z" \
  -H "ce-type: google.cloud.storage.object.v1.finalized" \
  -H "ce-source: //storage.googleapis.com/projects/_/buckets/MY-BUCKET-NAME" \
  -H "ce-subject: objects/MY_FILE.txt" \
  -d '{
        "bucket": "MY_BUCKET",
        "contentType": "text/plain",
        "kind": "storage#object",
        "md5Hash": "...",
        "metageneration": "1",
        "name": "MY_FILE.txt",
        "size": "352",
        "storageClass": "MULTI_REGIONAL",
        "timeCreated": "2020-04-23T07:38:57.230Z",
        "timeStorageClassUpdated": "2020-04-23T07:38:57.230Z",
        "updated": "2020-04-23T07:38:57.230Z"
      }'
```