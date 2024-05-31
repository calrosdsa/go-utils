push:
	git add .
	git commit -m "changes"
	git push -u origin main
	git tag $v
	git push origin $v