# Replace this with your image name, i.e. ghcr.io/<your-username>/demo-repo-go:latest
IMAGE_NAME?=ghcr.io/stacklok/demo-repo-go:latest

# Replace this with your GitHub username and PAT.
# This is used to authenticate with GitHub Container Registry (GHCR)
# and push the image to your repository.
# The PAT should have read/write access for packages.
CR_USERNAME?=stacklok
CR_PAT?=ghp_1234567890abcdefghij1234567890abcdefghij

.PHONY: login
login:
	@echo "Logging in to GitHub Container Registry"
	@echo "${CR_PAT}" | docker login ghcr.io -u $(CR_USERNAME) --password-stdin

.PHONY: build-image
build-image:
	@echo "Building a safe image..."
	docker build -t $(IMAGE_NAME) .

.PHONY: build-malicious-image
build-malicious-image:
	@echo "Building a malicious image..."
	@echo "// Maliciously altered on $$(date)" >> main.go
	docker build -t $(IMAGE_NAME) .


.PHONY: push-image
push-image:
	@echo "Pushing image..."
	docker push $(IMAGE_NAME)

.PHONY: keygen
keygen:
	@cosign generate-key-pair


.PHONY: sign-keypair
sign-keypair:
	@cosign sign $(IMAGE_NAME) --key cosign.key

.PHONY: sign-oidc
sign-oidc:
	@cosign sign $(IMAGE_NAME)

.PHONY: build-binary
build-binary:
	@echo "Building a safe binary..."
	go build -o demo-go-binary ./...

.PHONY: build-malicious-binary
build-malicious-binary:
	@echo "Building a malicious binary..."
	@echo "// Maliciously altered on $$(date)" >> main.go
	go build -o demo-go-binary ./...