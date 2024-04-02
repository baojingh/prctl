

push:
	git remote set-url origin  git@github.com:baojingh/prctl.git
	git pull
	git add .
	git commit -m "update"
	git push origin main

	git remote set-url origin git@gitlab.com:localdetector/prctl.git
	git pull
	git add .
	git commit -m "update"
	git push origin main	

