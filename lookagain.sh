#!/bin/bash

find . \( -type d,f \) -name "*.sh" | sort -r | cut -d '.sh' -f1 | cut -d './' -f2