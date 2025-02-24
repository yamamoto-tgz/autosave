# send-line-message

## Environment variables

| Name         | Required | Default |
| ------------ | -------- | ------- |
| LINE_USER_ID | True     | -       |
| LINE_TOKEN   | True     | -       |

## Create topic

```
gcloud pubsub topics create send-line-message
```

## Deploy function

```
gcloud functions deploy send-line-message \
    --gen2 \
    --region=us-west1 \
    --runtime=go122 \
    --source=./ \
    --entry-point=send-line-message \
    --trigger-topic=send-line-message \
    --no-allow-unauthenticated \
    --set-env-vars=LINE_USER_ID=${LINE_USER_ID},LINE_TOKEN=${LINE_TOKEN}
```
