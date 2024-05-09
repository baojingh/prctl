

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
# pypi仓库要求组件必须按照如下格式存放在服务端，flask/3.0.3flask-3.0.3-py3-none-any.whl, artifactory中会自动更新.pypi中的meta信息
curl -u sssss:xxxx  -XPUT 'https://xxx.rtf-alm.xxx.cloud/artifactory/xxx-dev-pypi/flask/3.0.3/flask-3.0.3-py3-none-any.whl' -T   flask-3.0.3-py3-none-any.whl

```



# REF
https://stackoverflow.com/questions/56133485/uploading-multiple-files-in-parallel-to-amazon-s3-with-goroutines-channels
