.PHONY: setup-dev-env
setup-dev-env:
	echo "TODO"

.PHONY: gen
gen:
	openapi-generator generate -i openapi.yaml \
		-g go-gin-server \
		--git-user-id sebastiaopamplona \
		--git-repo-id sherpa2 \
		--additional-properties \
		apiPath=api/openapi
	# remove unwanted files generated by openapi-generator
	rm ./main.go
	rm ./api/openapi/README.md
	rm -rf .openapi-generator
