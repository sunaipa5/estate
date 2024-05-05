PROJECT_DIR := .
RELEASE_DIR := release

#Files and folders that should not be copied to the release folder
NOT_COPY_FILES   := makefile set.go .gitattributes .gitignore estate
NOT_COPY_FOLDERS := frontend .git .OLD .release .backup

#Files and folders that should not be deleted in the release folder
NOT_DELETE_FILES := set.go
NOT_DELETE_FOLDERS :=

FRONTEND_FOLDER := frontend/dist

#ANSI Color
GREEN := "\033[0;32m"
MAGENTA := "\033[0;95m"
NC := "\033[0m"

.PHONY: all

all: backup clean copy build

backup:
	@if [ -d ".backup" ]; then \
		echo $(MAGENTA)"Cleaning .backup directory..."$(NC); \
		rm -rf .backup/*; \
	else \
		echo $(MAGENTA)"Creating .backup directory..."$(NC); \
		mkdir -p .backup; \
	fi
	@echo $(GREEN)"Copying files to .backup directory..."$(NC); \
	rsync -av --exclude=".backup" --exclude="$(RELEASE_DIR)" --quiet $(PROJECT_DIR)/ .backup/

copy: $(RELEASE_DIR)
	@echo $(GREEN)"Copying files..."$(NC)
	@rsync -av $(foreach file,$(NOT_COPY_FILES),--exclude=$(file)) $(foreach folder, $(RELEASE_DIR) $(NOT_COPY_FOLDERS),--exclude=$(folder)) $(PROJECT_DIR)/ $(RELEASE_DIR)/
	@echo $(GREEN)"Copying frontend files..."$(NC)
	@mkdir -p $(RELEASE_DIR)/www
	@rsync -av $(PROJECT_DIR)/$(FRONTEND_FOLDER)/ $(RELEASE_DIR)/www/

$(RELEASE_DIR):
	@echo $(GREEN)"Creating release directory..."$(NC)
	@mkdir -p $(RELEASE_DIR)

clean:
	@if [ -d "$(RELEASE_DIR)" ]; then \
		echo $(MAGENTA)"Cleaning release directory..."$(NC); \
		for folder in $(NOT_DELETE_FOLDERS); do \
			find $(RELEASE_DIR)/$$folder -mindepth 1 ! -path "$(RELEASE_DIR)/$$folder/*" -delete; \
		done \
	else \
		echo $(MAGENTA)"Creating release directory..."$(NC); \
		mkdir -p $(RELEASE_DIR); \
	fi
build:
    #OLD BUILD COMMAND: CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo
	@echo $(GREEN)"Building project..."$(NC)
	@cd $(RELEASE_DIR) && CGO_ENABLED=0 GOOS=linux go build
	@echo $(MAGENTA)"Building finish..."$(NC)
 	
