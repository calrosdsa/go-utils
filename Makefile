push:
	git tag $v
	git add .
	git commit -m "changes"
	git push -u origin main
	git push origin $v