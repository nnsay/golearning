#!/bin/sh
#
# Example external authenticator program for use with `ext_auth`.
#

# read u p
u="NGBOX-4c4c4544-0056-4e10-804c-cac04f483533@sirius-api.neuralgalaxy.cn"
p="SlZOTEg1My0uSlZOTEg1My5DTkZDUDAwMEFLMDA0TC4K"

# echo "USER: $u" >&2

if [ "$u" == "cicd" ]; then
  # if [ "$p" == "cicdpwd" ]; then
  #   echo '{"labels": {"group": ["CICD"]}}'
  #   exit 0
  # fi
  if [ "$p" == "$CICD_SECRET" ]; then
    echo '{"labels": {"group": ["CICD"]}}'
    exit 0
  fi
  exit 1
fi

if [[ "$u" == "NGBOX-"*"@"*".neuralgalaxy.c"* ]]; then
  python /Users/wangjian/Workspace/practice/golearning/learn/test/dockerauth/extensions/boxlogin.py "$u" "$p"
  exit $?
fi

exit 1
