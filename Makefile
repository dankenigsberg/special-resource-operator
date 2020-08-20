SPECIALRESOURCE  ?= nvidia-gpu
OPERATOR_SDK     ?= operator-sdk-v0.12.0-x86_64-linux-gnu


REGISTRY         ?= quay.io
ORG              ?= openshift-psap
TAG              ?= $(shell git branch | grep \* | cut -d ' ' -f2)
IMAGE            ?= $(REGISTRY)/$(ORG)/special-resource-operator:$(TAG)
NAMESPACE        ?= openshift-sro
PULLPOLICY       ?= IfNotPresent
TEMPLATE_CMD      = sed 's+REPLACE_IMAGE+$(IMAGE)+g; s+REPLACE_NAMESPACE+$(NAMESPACE)+g; s+Always+$(PULLPOLICY)+; s+REPLACE_SPECIALRESOURCE+$(SPECIALRESOURCE)+'

DEPLOY_NAMESPACE  = namespace.yaml
DEPLOY_OBJECTS    = service_account.yaml role.yaml role_binding.yaml operator.yaml
DEPLOY_CRD        = crds/sro.openshift.io_specialresources_crd.yaml 

PACKAGE           = github.com/openshift-psap/special-resource-operator
MAIN_PACKAGE      = $(PACKAGE)/cmd/manager
DOCKERFILE        = Dockerfile
ENVVAR            = GOOS=linux CGO_ENABLED=0
GOOS              = linux
GO111MODULE       = auto
GO_BUILD_RECIPE   = GO111MODULE=$(GO111MODULE) GOOS=$(GOOS) go build -mod=vendor -o $(BIN) $(MAIN_PACKAGE)

BIN=$(lastword $(subst /, ,$(PACKAGE)))

GOFMT_CHECK=$(shell find . -not \( \( -wholename './.*' -o -wholename '*/vendor/*' \) -prune \) -name '*.go' | sort -u | xargs gofmt -s -l)

all: build

build:
	$(GO_BUILD_RECIPE)

test: verify
	go test ./cmd/... ./pkg/... -coverprofile cover.out

test-e2e: 
	$(eval TMP := $(shell mktemp -d)/test-init.yaml)
	@$(TEMPLATE_CMD) deploy/service_account.yaml > $(TMP)
	echo -e "\n---\n" >> $(TMP)
	@$(TEMPLATE_CMD) deploy/role.yaml >> $(TMP)
	echo -e "\n---\n" >> $(TMP)
	@$(TEMPLATE_CMD) deploy/role_binding.yaml >> $(TMP)
	echo -e "\n---\n" >> $(TMP)
	@$(TEMPLATE_CMD) deploy/operator.yaml >> $(TMP)

	go test -v ./test/e2e/... -root $(PWD) -kubeconfig=$(KUBECONFIG) -tags e2e  -globalMan deploy/$(DEPLOY_CRD) -namespacedMan $(TMP)

$(DEPLOY_CRD):
	@$(TEMPLATE_CMD) deploy/$@ | kubectl apply -f -

deploy-crd: $(DEPLOY_CRD) 
	@sleep 1 

$(DEPLOY_NAMESPACE): deploy-crd 
	@$(TEMPLATE_CMD) deploy/$@ | kubectl apply -f -


deploy-objects: deploy-crd
	@for obj in $(DEPLOY_OBJECTS); do               \
		$(TEMPLATE_CMD) deploy/$$obj | kubectl apply -f - ; \
	done 

include recipes/$(SPECIALRESOURCE)/config/Makefile

specialresource: $(SPECIALRESOURCE) 

update-api:
	$(OPERATOR_SDK) generate k8s
	$(OPERATOR_SDK) generate openapi

deploy: $(DEPLOY_NAMESPACE) deploy-objects 

undeploy: 
	@for obj in $(DEPLOY_CRD) $(DEPLOY_OBJECTS) $(DEPLOY_NAMESPACE); do  \
		$(TEMPLATE_CMD) deploy/$$obj | kubectl delete -f - 2>/dev/null ; \
	done | true

verify:	verify-gofmt

verify-gofmt:
ifeq (, $(GOFMT_CHECK))
	@echo "verify-gofmt: OK"
else
	@echo "verify-gofmt: ERROR: gofmt failed on the following files:"
	@echo "$(GOFMT_CHECK)"
	@echo ""
	@echo "For details, run: gofmt -d -s $(GOFMT_CHECK)"
	@echo ""
	@exit 1
endif

clean:
	go clean
	rm -f $(BIN)

local-image:
	@rm -f special-resource-operator
	podman build --no-cache -t $(IMAGE) -f $(DOCKERFILE) .

local-image-push:
	podman push $(IMAGE) 

.PHONY: all build generate verify verify-gofmt clean test test-e2e local-image local-image-push $(DEPLOY_CRDS) grafana
