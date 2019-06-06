#!/bin/bash

find . \( -type f,d \) -name "*.sh" | ls -r | sed 's/.sh//g'