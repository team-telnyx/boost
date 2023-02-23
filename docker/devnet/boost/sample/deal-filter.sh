#!/bin/bash

STD_IN=$(</dev/stdin)
touch /app/public/deal-filter.log

printf "deal proposed as: $STD_IN\n" >> /app/public/deal-filter.log

exit 0