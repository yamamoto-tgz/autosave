# save-expenses

## Environment variables

| Name         | Required | Default      |
| ------------ | -------- | ------------ |
| BUCKET_NAME  | False    | autosave-tgz |
| HISTORY_FILE | False    | history.txt  |

## Create topic

```
gcloud pubsub topics create save-expenses
```

## Deploy function

```
gcloud functions deploy save-expenses \
    --gen2 \
    --region=us-west1 \
    --runtime=go122 \
    --source=./ \
    --entry-point=save-expenses \
    --trigger-topic=save-expenses \
    --no-allow-unauthenticated
```
