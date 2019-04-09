#!/bin/bash

# The paths in the associative array below must be adjusted for the local system

declare -A projects
projects=(
    ["bitflow-docs"]=".."
    ["bitflow-antlr-grammars"]="../../antlr-grammars"
    ["go-bitflow"]="../../go/go-bitflow"
    ["go-bitflow-collector"]="../../go/go-bitflow-collector"
    ["python-bitflow"]="../../python-bitflow"
    ["bitflow4j"]="../../bitflow4j"
)

for project in ${!projects[@]}; do
    sourceFile="${project}.yml"
    targetFile="${projects[$project]}/mkdocs.yml"
    echo "Copying file $sourceFile to $targetFile"
    cp -i "$sourceFile" "$targetFile"
done

