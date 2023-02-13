#! /bin/sh
cat<< EOF >>internal/models/schema/$(date +%s)_$1.sql
-- +goose Up
-- +goose Down
EOF
