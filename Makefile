empty :=
space := $(empty) $(empty)
PACKAGE := github.com/envoyproxy/protoc-gen-validate

BUF := $(if $(shell which buf),buf,$(error "No buf in PATH, visit https://docs.buf.build/installation to install"))

# protoc-gen-go parameters for properly generating the import path for PGV
VALIDATE_IMPORT := Mvalidate/validate.proto=${PACKAGE}/validate
GO_IMPORT_SPACES := ${VALIDATE_IMPORT},\
	Mgoogle/protobuf/any.proto=google.golang.org/protobuf/types/known/anypb,\
	Mgoogle/protobuf/duration.proto=google.golang.org/protobuf/types/known/durationpb,\
	Mgoogle/protobuf/struct.proto=google.golang.org/protobuf/types/known/structpb,\
	Mgoogle/protobuf/timestamp.proto=google.golang.org/protobuf/types/known/timestamppb,\
	Mgoogle/protobuf/wrappers.proto=google.golang.org/protobuf/types/known/wrapperspb,\
	Mgoogle/protobuf/descriptor.proto=google.golang.org/protobuf/types/descriptorpb
GO_IMPORT:=$(subst $(space),,$(GO_IMPORT_SPACES))

.DEFAULT_GOAL := help

.PHONY: help
help: Makefile
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: validate/validate.pb.go ## generates the PGV binary and installs it into $$GOPATH/bin
	go install .

.PHONY: bazel
bazel: ## generate the PGV plugin with Bazel
	bazel build //test/...

.PHONY: build_generation_tests
build_generation_tests:
	bazel build //test/generation/...

.PHONY: gazelle
gazelle: ## runs gazelle against the codebase to generate Bazel BUILD files
	bazel run //:gazelle -- update-repos -from_file=go.mod -prune -to_macro=dependencies.bzl%go_third_party
	bazel run //:gazelle

.PHONY: lint
lint: bin/golint bin/shadow ## lints the package for common code smells
	test -z "$(shell gofmt -d -s ./*.go)" || (gofmt -d -s ./*.go && exit 1)
	# golint -set_exit_status
	# check for variable shadowing
	go vet -vettool=$(shell pwd)/bin/shadow ./...
	# lints the python code for style enforcement
	flake8 --config=python/setup.cfg python/protoc_gen_validate/validator.py
	isort --check-only python/protoc_gen_validate/validator.py

bin/shadow:
	GOBIN=$(shell pwd)/bin go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow

bin/golint:
	GOBIN=$(shell pwd)/bin go install golang.org/x/lint/golint

bin/protoc-gen-go:
	GOBIN=$(shell pwd)/bin go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1

bin/harness:
	cd test && go build -o ../bin/harness ./harness/executor

.PHONY: harness
harness: testcases test/harness/go/harness.pb.go test/harness/go/main/go-harness test/harness/cc/cc-harness bin/harness ## runs the test harness, validating a series of test cases in all supported languages
	./bin/harness -go -cc

.PHONY: bazel-tests
bazel-tests: ## runs all tests with Bazel
	bazel test //test/... --test_output=errors

.PHONY: example-workspace
example-workspace: ## run all tests in the example workspace
	cd example-workspace && bazel test //... --test_output=errors

.PHONY: testcases
testcases: bin/protoc-gen-go ## generate the test harness case protos
	$(BUF) generate test/cases --template=buf.gen.test.yaml
	

validate/validate.pb.go: bin/protoc-gen-go validate/validate.proto
	$(BUF) generate validate --template=buf.gen.validate.yaml --output=validate

test/harness/go/harness.pb.go: bin/protoc-gen-go test/harness/harness.proto
	# generates the test harness protos
	cd test/harness && protoc -I . \
		--plugin=protoc-gen-go=${GOPATH}/bin/protoc-gen-go \
		--go_out="module=${PACKAGE}/test/harness/go,${GO_IMPORT}:./go" harness.proto

test/harness/go/main/go-harness:
	# generates the go-specific test harness
	cd test && go build -o ./harness/go/main/go-harness ./harness/go/main

test/harness/cc/cc-harness: test/harness/cc/harness.cc
	# generates the C++-specific test harness
	# use bazel which knows how to pull in the C++ common proto libraries
	bazel build //test/harness/cc:cc-harness
	cp bazel-bin/test/harness/cc/cc-harness $@
	chmod 0755 $@

test/harness/java/java-harness:
	# generates the Java-specific test harness
	mvn -q -f java/pom.xml clean package -DskipTests

.PHONY: prepare-python-release
prepare-python-release:
	cp validate/validate.proto python/
	cp LICENSE python/

.PHONY: python-release
python-release: prepare-python-release
	rm -rf python/dist
	python3.8 -m build --no-isolation --sdist python
	# the below command should be identical to `python3.8 -m build --wheel`
	# however that returns mysterious `error: could not create 'build': File exists`.
	# setuptools copies source and data files to a temporary build directory,
	# but why there's a collision or why setuptools stopped respecting the `build_lib` flag is unclear.
	# As a workaround, we build a source distribution and then separately build a wheel from it.
	python3.8 -m pip wheel --wheel-dir python/dist --no-deps python/dist/*
	python3.8 -m twine upload --verbose --skip-existing --repository ${PYPI_REPO} --username "__token__" --password ${PGV_PYPI_TOKEN} python/dist/*

.PHONY: check-generated
check-generated: ## run during CI; this checks that the checked-in generated code matches the generated version.
	for f in validate/validate.pb.go ; do \
	  mv $$f $$f.original ; \
	  make $$f ; \
	  mv $$f $$f.generated ; \
	  cp $$f.original $$f ; \
	  diff $$f.original $$f.generated ; \
	done

.PHONY: ci
ci: lint bazel testcases bazel-tests build_generation_tests example-workspace check-generated

.PHONY: clean
clean: ## clean up generated files
	(which bazel && bazel clean) || true
	rm -f \
		bin/protoc-gen-go \
		bin/harness \
		test/harness/cc/cc-harness \
		test/harness/go/main/go-harness \
		test/harness/go/harness.pb.go
	rm -rf \
		test/gen 
	rm -rf \
		python/dist \
		python/*.egg-info
