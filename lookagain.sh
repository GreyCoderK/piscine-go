#!/bin/bash

find . \( -type d,f \) -name "*.sh" | ls -r | sed 's/.sh//g'