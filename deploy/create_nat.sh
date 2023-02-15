#!/bin/bash

SUBNET_ID="subnet-0f290ea8bd4f975f7"
ALLOCATION_ID="eipalloc-09b060101d2026dd0"
ROUTE_TABLE_ID="rtb-031b05efcb74ce8aa"
NAT_NAME="nat-lambda"
NATID=`aws ec2 describe-nat-gateways --filter "Name=tag:Name,Values=$NAT_NAME" --filter 'Name=state,Values=available' | grep NatGatewayId | awk -F \" '{ print $4 }'`

if [ "$NATID" == "" ]; then

  aws ec2 create-nat-gateway --subnet-id $SUBNET_ID --allocation-id $ALLOCATION_ID --tag-specifications "ResourceType=natgateway,Tags=[{Key=Name,Value=$NAT_NAME}]"

  for j in {0..2}
  do
    NATID=`aws ec2 describe-nat-gateways --filter "Name=tag:Name,Values=$NAT_NAME" --filter 'Name=state,Values=available' | grep NatGatewayId | awk -F \" '{ print $4 }'`
    if [ "$NATID" == "" ]; then
      sleep 20s
    else
      break
    fi
  done

  # Abandon
  if [ "$NATID" == "" ]; then
    echo "Wait for nat gateway timeout. Failed to create the nat gateway"
    exit 1
  fi
fi

aws ec2 create-route --route-table-id $ROUTE_TABLE_ID --destination-cidr-block 0.0.0.0/0 --nat-gateway-id $NATID
