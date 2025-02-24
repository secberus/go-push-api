#
# Copyright 2024 Secberus, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http:#www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
.PHONY: build
.DEFAULT_GOAL := build

SRC_DIR := ../push-api-protos
DST_DIR := .
PROTOVALIDATE_DIR := ../protovalidate/proto/protovalidate

build:
	protoc -I=/usr/local/include -I=$(PROTOVALIDATE_DIR) -I=$(SRC_DIR) --go_out=$(DST_DIR) --go_opt=paths=source_relative --go-grpc_out=$(DST_DIR) --go-grpc_opt=paths=source_relative $(SRC_DIR)/types/v1/*.proto
	protoc -I=/usr/local/include -I=$(PROTOVALIDATE_DIR) -I=$(SRC_DIR) --go_out=$(DST_DIR) --go_opt=paths=source_relative --go-grpc_out=$(DST_DIR) --go-grpc_opt=paths=source_relative $(SRC_DIR)/api/v1/*.proto
	protoc -I=/usr/local/include -I=$(PROTOVALIDATE_DIR) -I=$(SRC_DIR) --go_out=$(DST_DIR) --go_opt=paths=source_relative --go-grpc_out=$(DST_DIR) --go-grpc_opt=paths=source_relative $(SRC_DIR)/service/v1/push/service.proto

