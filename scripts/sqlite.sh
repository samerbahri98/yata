#! /usr/bin/env sh

rm -rf $PWD/db;
mkdir $PWD/db;
sqlite3 $PWD/db/tmp.db "" ;
sqlite3 $PWD/db/dev.db "" ;
