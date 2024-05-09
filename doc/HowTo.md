

```bash

prctl login --url baidu.com --username bob --password pd
```

```bash
prctl logout

```


```bash
prctl download --input components.txt --output ./outputs 
```

```bash
apt-get install --no-install-recommends -y --download-only cron

# apt-get download  cron
```

```bash
# download all indirect dependencies
apt-get install --no-install-recommends -y --download-only xxx=1.2


# just download the package in current path , no dependencies
apt-get download xxx=1.2
```


# REF
https://stackoverflow.com/questions/56133485/uploading-multiple-files-in-parallel-to-amazon-s3-with-goroutines-channels

- golang change directory 
  https://stackoverflow.com/questions/52435908/how-to-change-the-shells-current-working-directory-in-go

  