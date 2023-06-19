BASE_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
RELEASE_DIR := release
JOB_MONITOR := job-monitor
BUILD_TIME := $(shell date +"%Y-%m-%d %Z %T")

.PHONY: clean buld

clean:
	rm -rf ${BASE_DIR}/${RELEASE_DIR}

build: clean
	@mkdir -p ${BASE_DIR}/${RELEASE_DIR}
	go build -o ${BASE_DIR}/${RELEASE_DIR}/${JOB_MONITOR} ${BASE_DIR}/pkg
