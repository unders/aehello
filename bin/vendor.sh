#!/usr/bin/env bash

set -e

readonly PACKAGES=(
	"github.com/pkg/errors"
	"cloud.google.com/go/logging"
)

main() {
	local cmd=$1

	if [[ "$cmd" = list ]]; then
		list
	elif [[ "$cmd" = update ]]; then
		update
	elif [[ "$cmd" = deps ]]; then
		deps
	elif [[ "$cmd" = pkg ]]; then
		pkg
	else
		echo "Usage:"
		echo ""
		echo "./bin/vendor list                       - show deps not vendored"
		echo "./bin/vendor deps                       - show vendored deps"
		echo "./bin/vendor update                     - update all vendored deps"
	fi
}

update() {
	rm -rf vendor
	mkdir vendor

	for pkg in "${PACKAGES[@]}"
	do
		echo "vendoring pkg: $pkg"
		vendor "$pkg"
	done
}

deps() {
	echo "Vendored dependencies:"
	echo ""
	for pkg in "${PACKAGES[@]}"
	do
		echo "    $pkg"
	done
}

list() {
	local pkgs=$(go list ./... | grep -v vendor | xargs go list \
				  -f '{{join .Imports "\n"}} {{join .TestImports "\n"}}' | \
				  xargs go list -f '{{if not .Standard}}{{.ImportPath}}{{end}}')

	echo "Vendor these packages:"
	echo ""
	for pkg in $pkgs
	do
		if [[ "$pkg" != *github.com/unders/aehello* ]]; then
			echo "    $pkg"
		fi
	done
}

vendor() {
	local pkg=$1
	local home=$(pwd)

	local tmpdir=$(mktemp -d -t tmp.XXXXXXXXXX)

	if [[ $tmpdir == "/var/folders"* ]]; then

		if [[ ${#tmpdir} -le 40 ]]; then
			echo "tmpdir to short..."
			return
		fi

		set_gopath $tmpdir

		go get -u "$pkg"

		prune "$GOPATH/src"
		cd "$GOPATH/src"
		cp -R . "$home/vendor/"
		rm -rf "$tmpdir"

		cd "$home"
		return
	fi

	echo "Could not create tmpdir..."
}


set_gopath() {
	local tmpdir=$1

	mkdir "$tmpdir"/bin
	mkdir "$tmpdir"/pkg
	mkdir "$tmpdir"/src

	export GOPATH="$tmpdir"
}

prune() {
	local path=$1

	find "$path" -type d -name ".git" -prune -exec rm -rf '{}' '+'
	find "$path" -type d -name ".github" -prune -exec rm -rf '{}' '+'
	find "$path" -type f -name "*_test.go" -delete
	find "$path" -type f ! -name "*.go" \
		  -and ! -name "*.proto" \
		  -and ! -name "*.s" \
		  -and ! -name "*.c" \
		  -and ! -name "LICENSE" \
		  -and ! -name "LICENCE" \
		  -and ! -name "UNLICENSE" \
		  -and ! -name "COPYING" \
		  -and ! -name "COPYRIGHT" \
		  -and ! -name "PATENTS" \
		  -delete
}

main $@
