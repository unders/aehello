VERSION=dev
PROG=hello
#projectID="projectID"

http?=":10080"
httpAddr=$(http)
httpPort=$(shell echo "$(http)" | sed 's/\://g')


.PHONY: help
help:
	@echo ""
	@echo "** $(PROG) [ version: $(VERSION), project: $(projectID) ] **"
	@echo ""
	@echo "Commands:"
	@echo "    make deploy        # deploy and start app version dev"
	@echo ""



.PHONY: deploy
deploy:
	#gcloud app deploy --bucket=hello --version=$(VERSION)
	gcloud app deploy --version=$(VERSION)

