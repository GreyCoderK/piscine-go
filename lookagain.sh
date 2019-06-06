#!/bin/bash

find . \( -type f \) -name "*.sh" | ls -r | sed 's/.sh//g'