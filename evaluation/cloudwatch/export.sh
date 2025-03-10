#!/bin/bash
LAMBDA="/aws/lambda/CacheNode"
FILE="aws/lambda/CacheNode"
#2023-06-08
start='2023-06-08 00:00:00'
end='2023-06-09 00:00:00'

# Convert date into seconds (Format is %s)
startTime=$(date  -j -f "%Y-%m-%d %H:%M:%S" "$start" +%s)000
endTime=$(date  -j -f "%Y-%m-%d %H:%M:%S" "$end" +%s)000


for i in {0..5}
do
  # Wait for the end the last task
  for j in {0..15}
  do
    RUNNING=`aws logs describe-export-tasks --status-code "RUNNING" | grep taskId | awk -F \" '{ print $4 }'`
    if [ "$RUNNING" != "" ]; then
      sleep 2s
    else
      break
    fi
  done

  # Abandon
  if [ "$RUNNING" != "" ]; then
    echo "Detect running task and wait timeout, killing task \"$RUNNING\"..."
    aws logs --profile CWLExportUser cancel-export-task --task-id \"$RUNNING\"
    echo "Done"
  fi

  aws logs create-export-task --log-group-name $LAMBDA$i --from ${startTime} --to ${endTime} --destination "sion-datapool" --destination-prefix $FILE$PREFIX$i
  sleep 2s
done
