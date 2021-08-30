#!/bin/sh

u="NGBOX-4c4c4544-0056-4e10-804c-cac04f483533@sirius-api.neuralgalaxy.cn"
if [[ "$u" == "NGBOX-"*"@"*".neuralgalaxy.c"* ]]; then
  echo "ecode: $?"
  exit $?
fi

echo "ecode: 13"
exit 13