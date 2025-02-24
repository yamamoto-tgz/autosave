# AutoSave

## Create project

```
gcloud projects create autosave-tgz
gcloud config set project autosave-tgz
gcloud billing projects link autosave-tgz --billing-account=`gcloud billing accounts list | tail -1 | cut -d " " -f 1`
```

## Create bucket

```
gcloud storage buckets create gs://autosave-tgz
```

## Deploy send-line-message

[send-line-message](./cloud-functions/send-line-message/README.md)

## Deploy save-gmail-history

[save-gmail-history](./cloud-functions/save-gmail-history/README.md)

## Deploy watch-gmail-history

[watch-gmail-history](./cloud-functions/watch-gmail-history/README.md)

## Deploy save-expenses

[save-expenses](./cloud-functions/save-expenses/README.md)
