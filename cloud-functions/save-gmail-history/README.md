# save-gmail-history

## Environment variables

| Name         | Required | Default      |
| ------------ | -------- | ------------ |
| BUCKET_NAME  | False    | autosave-tgz |
| HISTORY_FILE | False    | history.txt  |

## Create topic

```
gcloud pubsub topics create gmail
gcloud pubsub topics add-iam-policy-binding gmail --member "serviceAccount:gmail-api-push@system.gserviceaccount.com" --role "roles/pubsub.publisher"
```

## Deploy function

```
gcloud functions deploy save-gmail-history \
    --gen2 \
    --region=us-west1 \
    --runtime=go122 \
    --source=./ \
    --entry-point=save-gmail-history \
    --trigger-topic=gmail \
    --no-allow-unauthenticated
```
