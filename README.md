# AutoSave

家計簿自動保存システム

## Functions

### sendLineMessages

Deploy Command:

```
gcloud functions deploy sendLineMessages --gen2 --region=asia-northeast1 --runtime=nodejs22 --source=./sendLineMessages --entry-point=sendLineMessages --trigger-http --no-allow-unauthenticated --set-env-vars=LINE_TOKEN=${LINE_TOKEN},LINE_USER_ID=${LINE_USER_ID}
```

Environment Variables:

-   LINE_TOKEN
-   LINE_USER_ID
