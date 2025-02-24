# send-line-messages

## Environment variables

| Name         | Required | Default |
| ------------ | -------- | ------- |
| LINE_USER_ID | True     | -       |
| LINE_TOKEN   | True     | -       |

## Create topic

```
gcloud pubsub topics create send-line-messages
```

## Deploy function

```
gcloud functions deploy send-line-messages \
    --gen2 \
    --region=us-west1 \
    --runtime=go122 \
    --source=./ \
    --entry-point=send-line-messages \
    --trigger-topic=send-line-messages \
    --no-allow-unauthenticated \
    --set-env-vars=LINE_USER_ID=${LINE_USER_ID},LINE_TOKEN=${LINE_TOKEN}
```
