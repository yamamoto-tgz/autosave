# save-expenses

## Environment variables

| Name            | Required | Default          |
| --------------- | -------- | ---------------- |
| SPREADSHEET_ID  | True     | -                |
| RANGE           | True     | -                |
| BUCKET_NAME     | False    | autosave-tgz     |
| CREDENTIAL_JSON | False    | credentials.json |
| TOKEN_JSON      | False    | token.json       |

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
    --no-allow-unauthenticated \
    --set-env-vars=SPREADSHEET_ID=${SPREADSHEET_ID},RANGE=${RANGE}
```
