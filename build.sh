#!/bin/bash
PROJECT_DIR="."
RELEASE_DIR="release"

#Create database SQL File
CREATE_SQL_FILE=false
DB_USER=dbadmin
DB_NAME=estate
SQL_FILE="estate.sql"

#Files and folders that should not be copied to the release folder
NOT_COPY_FILES=".gitattributes .gitignore estate build.sh options"
NOT_COPY_FOLDERS="frontend .backup .git"


FRONTEND_FOLDER="frontend/dist"

#ANSI Color
GREEN="\033[0;32m"
MAGENTA="\033[0;95m"
NC="\033[0m"

backup() {
    if [ -d ".backup" ]; then
        echo -e "${MAGENTA}Cleaning .backup directory...${NC}"
        rm -rf .backup/*
    else
        echo -e "${MAGENTA}Creating .backup directory...${NC}"
        mkdir -p .backup
    fi
    echo -e "${GREEN}Copying files to .backup directory...${NC}"
    rsync -av --exclude=".backup" --exclude="$RELEASE_DIR" --quiet "$PROJECT_DIR"/ .backup/
}

copy() {
    echo -e "${GREEN}Creating release directory...${NC}"
    mkdir -p "$RELEASE_DIR"
    echo -e "${GREEN}Copying files...${NC}"
    rsync -av $(printf -- '--exclude=%s ' $NOT_COPY_FILES) $(printf -- '--exclude=%s ' $RELEASE_DIR $NOT_COPY_FOLDERS) "$PROJECT_DIR"/ "$RELEASE_DIR"/
    echo -e "${GREEN}Copying frontend files...${NC}"
    mkdir -p "$RELEASE_DIR/www"
    rsync -av "$PROJECT_DIR/$FRONTEND_FOLDER/" "$RELEASE_DIR/www/"
}

build() {
    echo -e "${GREEN}Building project...${NC}"
    cd "$RELEASE_DIR" && CGO_ENABLED=0 GOOS=linux go build
    echo -e "${MAGENTA}Building finish...${NC}"
}

createSql() {
      if [ "$CREATE_SQL_FILE" = true ]; then
        echo -e "${GREEN}Creating SQL file...${NC}"
        mysqldump -u "$DB_USER" -p"$DB_PASS" "$DB_NAME" > "$SQL_FILE"
        echo -e "${MAGENTA}SQL file created: $SQL_FILE${NC}"
    else
        echo -e "${MAGENTA}SQL file creation skipped.${NC}"
    fi
}

main() {
    backup
    copy
    build
    createSql
}

main "$@"

