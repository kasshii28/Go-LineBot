InitToPush:
	git init
	git add -A
	git commit -m "first commit"
	git branch -M main
	git remote add origin git@github.com:kasshii28/Go-LineBot.git
	git push -u origin main