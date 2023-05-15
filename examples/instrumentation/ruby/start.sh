#! /usr/env/bin bash

unset BUNDLE_PATH
unset BUNDLE_BIN

ruby ./uppercase.rb -s Puma -o 0.0.0.0
