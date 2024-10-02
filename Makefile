### BEGIN HEADERS ###
# This block is managed by ServiceBot plugin - Make. The content in this block is created using a common
# template and configurations in service.yml.
# Modifications in this block will be overwritten by generated content in the nightly run.
# For more information, please refer to the page:
# https://confluentinc.atlassian.net/wiki/spaces/Foundations/pages/2871328913/Add+Make
SERVICE_NAME := cmf-sdk-go
SERVICE_DEPLOY_NAME := cmf-sdk-go

### END HEADERS ###
### BEGIN MK-INCLUDE UPDATE ###
### This section is managed by service-bot and should not be edited here.
### You can make changes upstream in https://github.com/confluentinc/cc-service-bot

CURL ?= curl
FIND ?= find
TAR ?= tar

# Mount netrc so curl can work from inside a container
DOCKER_NETRC_MOUNT ?= 1

GITHUB_API = api.github.com
GITHUB_MK_INCLUDE_OWNER := confluentinc
GITHUB_MK_INCLUDE_REPO := cc-mk-include
GITHUB_API_CC_MK_INCLUDE := https://$(GITHUB_API)/repos/$(GITHUB_MK_INCLUDE_OWNER)/$(GITHUB_MK_INCLUDE_REPO)
GITHUB_API_CC_MK_INCLUDE_TARBALL := $(GITHUB_API_CC_MK_INCLUDE)/tarball
GITHUB_API_CC_MK_INCLUDE_VERSION ?= $(GITHUB_API_CC_MK_INCLUDE_TARBALL)/$(MK_INCLUDE_VERSION)

MK_INCLUDE_DIR := mk-include
MK_INCLUDE_LOCKFILE := .mk-include-lockfile
MK_INCLUDE_TIMESTAMP_FILE := .mk-include-timestamp
# For optimum performance, you should override MK_INCLUDE_TIMEOUT_MINS above the managed section headers to be
# a little longer than the worst case cold build time for this repo.
MK_INCLUDE_TIMEOUT_MINS ?= 240
# If this latest validated release is breaking you, please file a ticket with DevProd describing the issue, and
# if necessary you can temporarily override MK_INCLUDE_VERSION above the managed section headers until the bad
# release is yanked.
MK_INCLUDE_VERSION ?= v0.1184.0

# Make sure we always have a copy of the latest cc-mk-include release less than $(MK_INCLUDE_TIMEOUT_MINS) old:
# Note: The simply-expanded make variable makes sure this is run once per make invocation.
UPDATE_MK_INCLUDE := $(shell \
	func_fatal() { echo "$$*" >&2; echo output here triggers error below; exit 1; } ; \
	test -z "`git ls-files $(MK_INCLUDE_DIR)`" || { \
		func_fatal 'fatal: checked in $(MK_INCLUDE_DIR)/ directory is preventing make from fetching recent cc-mk-include releases for CI'; \
	} ; \
	trap "rm -f $(MK_INCLUDE_LOCKFILE); exit" 0 2 3 15; \
	waitlock=0; while ! ( set -o noclobber; echo > $(MK_INCLUDE_LOCKFILE) ); do \
	   sleep $$waitlock; waitlock=`expr $$waitlock + 1`; \
	   test 14 -lt $$waitlock && { \
	      echo 'stealing stale lock after 105s' >&2; \
	      break; \
	   } \
	done; \
	test -s $(MK_INCLUDE_TIMESTAMP_FILE) || rm -f $(MK_INCLUDE_TIMESTAMP_FILE); \
	test -z "`$(FIND) $(MK_INCLUDE_TIMESTAMP_FILE) -mmin +$(MK_INCLUDE_TIMEOUT_MINS) 2>&1`" || { \
	   grep -q 'machine $(GITHUB_API)' ~/.netrc 2>/dev/null || \
	      func_fatal 'error: follow https://confluentinc.atlassian.net/l/cp/0WXXRLDh to fix your ~/.netrc'; \
	   $(CURL) --fail --silent --netrc --location "$(GITHUB_API_CC_MK_INCLUDE_VERSION)" --output $(MK_INCLUDE_TIMESTAMP_FILE)T --write-out '$(GITHUB_API_CC_MK_INCLUDE_VERSION): %{errormsg}\n' >&2 \
	   && $(TAR) zxf $(MK_INCLUDE_TIMESTAMP_FILE)T \
	   && rm -rf $(MK_INCLUDE_DIR) \
	   && mv $(GITHUB_MK_INCLUDE_OWNER)-$(GITHUB_MK_INCLUDE_REPO)-* $(MK_INCLUDE_DIR) \
	   && mv -f $(MK_INCLUDE_TIMESTAMP_FILE)T $(MK_INCLUDE_TIMESTAMP_FILE) \
	   && echo 'installed cc-mk-include-$(MK_INCLUDE_VERSION) from $(GITHUB_MK_INCLUDE_REPO)' >&2 \
	   || func_fatal unable to install cc-mk-include-$(MK_INCLUDE_VERSION) from $(GITHUB_MK_INCLUDE_REPO) \
	   ; \
	} || { \
	   rm -f $(MK_INCLUDE_TIMESTAMP_FILE)T; \
	   if test -f $(MK_INCLUDE_TIMESTAMP_FILE); then \
	      touch $(MK_INCLUDE_TIMESTAMP_FILE); \
	      func_fatal 'unable to access $(GITHUB_MK_INCLUDE_REPO) fetch API to check for latest release; next try in $(MK_INCLUDE_TIMEOUT_MINS) minutes'; \
	   else \
	      func_fatal 'unable to access $(GITHUB_MK_INCLUDE_REPO) fetch API to bootstrap mk-include subdirectory'; \
	   fi; \
	} \
)
ifneq ($(UPDATE_MK_INCLUDE),)
    $(error mk-include update failed)
endif

# Export the (empty) .mk-include-check-FORCE target to allow users to trigger the mk-include
# download code above via make but without having to run any of the other targets, e.g. build.
.PHONY: .mk-include-check-FORCE
.mk-include-check-FORCE:
	@echo -n ""
### END MK-INCLUDE UPDATE ###
### BEGIN INCLUDES ###
# This block is managed by ServiceBot plugin - Make. The content in this block is created using a common
# template and configurations in service.yml.
# Modifications in this block will be overwritten by generated content in the nightly run. To include
# additional mk files, please add them before or after this generated block.
# For more information, please refer to the page:
# https://confluentinc.atlassian.net/wiki/spaces/Foundations/pages/2871328913/Add+Make
include ./mk-include/cc-begin.mk
include ./mk-include/cc-semver.mk
include ./mk-include/cc-semaphore.mk
include ./mk-include/cc-testbreak.mk
include ./mk-include/cc-vault.mk
include ./mk-include/cc-end.mk
### END INCLUDES ###