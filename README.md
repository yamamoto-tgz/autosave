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

## Deploy save-gmail-history

[save-gmail-history](./gcp/save-gmail-history/README.md)

## Deploy watch-gmail-history

[watch-gmail-history](./gcp/watch-gmail-history/README.md)
