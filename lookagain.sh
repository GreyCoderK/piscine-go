#!/bin/bash

find . \( -type d,f \) -name "*.sh" | sort -r | sed 's/^*.\/\(.*\).\/sh$/\1/'