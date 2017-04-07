VERSION=dev
PROG=world
projectID=projectID

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
	gcloud app deploy --version=$(VERSION)

