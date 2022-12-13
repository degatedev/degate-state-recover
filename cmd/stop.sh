#!/bin/bash
echo "stop StateRecovery"
PID=`ps -ef | grep staterecovery | grep -v grep | awk '{print $2}'`
if [ "" != "$PID" ];then
  echo "kill $PID"
  kill -9 $PID
fi
echo "StateRecovery had been stopped."