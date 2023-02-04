#!/usr/bin/env bash

changes=$(git diff --name-only HEAD^)
filtered=()

for i in $changes
do
  if [[ $i == plugins/*.json ]] && [[ -f $i ]]; then
    filtered+=("$i")
  fi
done

if [[ ${#filtered[@]} ]]; then
  echo "nochanges"
  exit 0
fi

echo "${filtered[@]}"
exit 0