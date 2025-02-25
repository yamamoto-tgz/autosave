# receive-gmail-history

## Environment variables

| Name | Required | Default |
| ---- | -------- | ------- |
| -    | -        | -       |

## Create topic

```
gcloud pubsub topics create receive-gmail-history
gcloud pubsub topics add-iam-policy-binding receive-gmail-history \
    --member "serviceAccount:gmail-api-push@system.gserviceaccount.com" \
    --role "roles/pubsub.publisher"
```

## Deploy function

```
gcloud functions deploy receive-gmail-history \
    --gen2 \
    --region=us-west1 \
    --runtime=go122 \
    --source=./ \
    --entry-point=receive-gmail-history \
    --trigger-topic=receive-gmail-history \
    --no-allow-unauthenticated
```
