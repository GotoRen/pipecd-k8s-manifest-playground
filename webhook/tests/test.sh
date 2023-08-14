#!/bin/bash

# Health check
curl -i -X GET 'localhost:8081/healthz'

# PipeCD webhook
curl -i -X POST -H "Content-Type: application/json" -H "signatureKey: pipecd-local" -d '@pipecd-webhook-sample.json' localhost:8081/api/hooks/services
