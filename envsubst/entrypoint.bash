#!/usr/bin/env bash
if [[ $# == 0 ]]; then
    echo "Please specify list of files to pre-process" >&2
    exit 1
fi

set -eu -o pipefail

for f in $(ls $@); do
    cat ${f} | envsubst > ${f}.processed
    mv ${f}.processed ${f}
done

exit 0