#!/usr/bin/env bash

set -eu -o pipefail
declare -a files
files_count=0

function usage() {
    echo "usage: $0 [-e <env-file1>] file1 [file2...fileN]" >&2
    echo "  -h    show this help screen" >&2
    echo "  -e    add environment variables from this env file (can be specified multiple times)'" >&2
    exit 0
}

# Parse command line flags
while true; do
    [[ $# == 0 ]] && break
    case "${1}" in
        -h|--help)
            usage
            ;;
        -e|--env)
            [[ -z "${2}" ]] && usage
            export $(grep -v '^#' ${2} | xargs)
            shift 2
            ;;
        *)
            files[files_count]="${1}"
            files_count=$((  + 1 ))
            shift
            ;;
    esac
done

# Fail if no files/wildcards were provided
if [[ ${files_count} == 0 ]]; then
    echo "Please specify list of files to pre-process" >&2
    exit 1
fi

# Iterate files/wildcards and pre-process them using envsubst
for f in $(ls ${files[*]}); do
    echo "Pre-processing ${f}..." >&2
    cat ${f} | envsubst > ${f}.processed
    mv ${f}.processed ${f}
done

exit 0
