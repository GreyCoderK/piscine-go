#!/bin/bash

ls -CF --format=commas -t | sed 's/ / /g'
