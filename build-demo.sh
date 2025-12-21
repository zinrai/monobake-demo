#!/bin/bash

read TARGET VERSION <<< $(monobake -tag refs/tags/backend/v0.1.0)
[ -n "$TARGET" ] && sudo docker buildx bake --set="${TARGET}.tags=${TARGET}:${VERSION}" "$TARGET"
