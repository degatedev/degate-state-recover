#!/bin/bash
rm -rf StateRecovery.log
touch StateRecovery.log
./staterecovery -conf "config.toml" > StateRecovery.log 2>&1 &
echo "StateRecovery started."