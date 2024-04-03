# prctl
A tool to manage components in private repo, currently for JFrog





```bash
prctl login --url xxxx --username xxx --password xxx
prctl logout 

prctl config list
prctl config add    --repo xxx-aa-sss
prctl config remove --repo xxx-aa-sss

prctl download --type debian  --file cc.txt --apt ./detox.list --path /var/canbu/vwnui
prctl upload   --type debian  --path /var/canbu/vwnui


prctl delete   --type debian  --file cc.txt
prctl list     --type debian  
```