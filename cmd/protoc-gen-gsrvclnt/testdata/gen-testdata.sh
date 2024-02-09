#!/bin/bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

go install $SCRIPT_DIR/..

protoc --proto_path=$SCRIPT_DIR --go_out=$SCRIPT_DIR --go_opt=paths=source_relative --gsrvclnt_out=$SCRIPT_DIR --gsrvclnt_opt=paths=source_relative --go-grpc_out=$SCRIPT_DIR --go-grpc_opt=paths=source_relative $SCRIPT_DIR/route_guide/route_guide.proto
