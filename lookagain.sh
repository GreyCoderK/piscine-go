#!/bin/bash

find . \( -type d,f \) -name "*.sh" | sort -r | sed 's/.sh//g' | sed 's/^.*\///g'