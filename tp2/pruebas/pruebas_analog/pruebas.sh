#!/usr/bin/env bash

set -eu

export TIME="Tardó: %Us - Consumió: %M KiB"
PROGRAMA="$1"
RET=0

# Correr con diff y sin Valgrind.
echo "Ejecución de pruebas unitarias:"
echo ""
for x in *.test; do
    b=${x%.test}
    printf "${b} "
    cat ${b}.test
    ($PROGRAMA < ${b}_in > ${b}_actual_out 2> ${b}_actual_err && \
        diff --suppress-common-lines -y -W 60 ${b}_out ${b}_actual_out && \
        diff --suppress-common-lines -y -W 60 ${b}_err ${b}_actual_err && \
        echo "OK") || { RET=$?; echo "ERROR"; }
    echo ""
done

rm *_actual_out *_actual_err

exit $RET
