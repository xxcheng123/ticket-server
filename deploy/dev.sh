#!/usr/bin

workDir="$PWD"
if [[ $workDir =~ "deploy" ]]; then
  echo "go parent directory.";
  cd ..;
fi
workDir="$PWD"
cd "$workDir/app/user/api" || exit;
goctl api go --api "$workDir/app/user/api/user.api" --dir "$workDir/app/user/api" --style goZero --home "$workDir/deploy/template";

cd "$workDir/app/user/rpc" || exit;
goctl rpc protoc "user.proto"  --go_out=./pb --go-grpc_out=./pb --zrpc_out=. --style goZero --home "$workDir/deploy/template";


cd "$workDir/app/user/rpc/internal/model" || exit;
goctl model mysql ddl --cache true --dir . --src ./user.sql --style goZero --home "$workDir/deploy/template";

go mod tidy;

echo "all success.";