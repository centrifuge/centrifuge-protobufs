#!/usr/bin/env bash

# push the swagger api json to swagger hub
echo "pushing swagger.json to SwaggerHub"

version=`cat package.json | ${JQ_BIN} '.version'`
echo "found version ${version} from package.json"

curl -i -X POST \
  https://api.swaggerhub.com/apis/centrifuge.io/p2p?version=${version} \
  -H "Authorization: $SWAGGER_API_KEY" \
  -H "Content-Type: application/json" -d @./gen/swagger.json

exit $?