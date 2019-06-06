#!/bin/bash

find . \( -type d,f \) -name "*.sh" | sort -r | cut --output-delimiter='.sh' -f1 | cut --output-delimiter='./' -f2