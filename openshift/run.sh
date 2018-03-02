#!/usr/bin/env bash

oc project tomw-gophers-production
oc new-app https://github.com/twhiston/gophers10 --name=gophers --strategy=docker --context-dir=gophers10
oc expose svc/gophers

oc get bc