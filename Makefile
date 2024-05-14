STEAMPIPE_INSTALL_DIR ?= ~/.steampipe
BUILD_TAGS = netgo

install:
	go build -o $(STEAMPIPE_INSTALL_DIR)/plugins/hub.steampipe.io/plugins/turbot/microsoft365@latest/steampipe-plugin-microsoft365.plugin -tags "${BUILD_TAGS}" *.go
