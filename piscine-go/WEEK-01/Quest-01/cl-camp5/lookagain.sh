#!/bin/bash

find . -type f -name "*.sh" | awk -F "/" '{print $NF}' | sed 's/\.sh$//'
