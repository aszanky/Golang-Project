#!/bin/bash

protoc delivery/contentpb/content.proto --go_out=plugins=grpc:.