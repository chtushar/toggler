#!/usr/bin/env sh
set -eu

printf '\n'

RED="$(tput setaf 1 2>/dev/null || printf '')"
BLUE="$(tput setaf 4 2>/dev/null || printf '')"
GREEN="$(tput setaf 2 2>/dev/null || printf '')"
NO_COLOR="$(tput sgr0 2>/dev/null || printf '')"

info() {
  printf '%s\n' "${BLUE}> ${NO_COLOR} $*"
}

error() {
  printf '%s\n' "${RED}x $*${NO_COLOR}" >&2
}

completed() {
  printf '%s\n' "${GREEN}✓ ${NO_COLOR} $*"
}

exists() {
  command -v "$1" 1>/dev/null 2>&1
}

check_dependencies() {
	if ! exists curl; then
		error "curl is not installed."
		exit 1
	fi

	if ! exists docker; then
		error "docker is not installed."
		exit 1
	fi

	if ! exists docker-compose; then
		error "docker-compose is not installed."
		exit 1
	fi
}

check_existing_db_volume() {
	info "checking for an existing docker db volume"
	if docker volume inspect toggler_toggler-data >/dev/null 2>&1; then
		error "toggler-data volume already exists. Please use docker-compose down -v to remove old volumes for a fresh setup of PostgreSQL."
		exit 1
	fi
}


check_dependencies
check_existing_db_volume