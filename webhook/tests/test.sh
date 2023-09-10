#!/bin/bash

# Health check
curl -i -X GET 'localhost:8081/healthz'

# PipeCD webhook
curl -i -X POST -H "Content-Type: application/json" -H "signatureKey: pipecd-playground" -d '@pipecd-webhook-sample.json' http://localhost:8081/api/hooks/services
