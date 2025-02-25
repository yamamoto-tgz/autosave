# watch-gmail-history

## Environment variables

| Name             | Required | Default |
| ---------------- | -------- | ------- |
| RKTN_PAY_LABEL   | True     | -       |
| RKTN_DEBIT_LABEL | True     | -       |

## Create topic

```
gcloud pubsub topics create watch-gmail-history
```

## Deploy function

```
gcloud functions deploy watch-gmail-history \
    --gen2 \
    --region=us-west1 \
    --runtime=go122 \
    --source=./ \
    --entry-point=watch-gmail-history \
    --trigger-topic=watch-gmail-history \
    --no-allow-unauthenticated \
    --set-env-vars=RKTN_PAY_LABEL=${RKTN_PAY_LABEL},RKTN_DEBIT_LABEL=${RKTN_DEBIT_LABEL}
```

## Create scheduler

```
gcloud scheduler jobs create pubsub watch-gmail-history \
    --location=us-west1 \
    --schedule="0 0 * * 0" \
    --topic=watch-gmail-history \
    --message-body="-" \
    --attributes="" \
    --time-zone="Asia/Tokyo"
```
