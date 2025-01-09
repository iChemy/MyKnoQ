#!/bin/bash

pwd
# SVG出力ディレクトリを作成
mkdir -p ./svg

find ./puml -name "*.puml" -print0 | while IFS= read -r -d $'\0' file; do
  java -jar /app/plantuml.jar -tsvg -o "../svg" "$file"
done

ls svg