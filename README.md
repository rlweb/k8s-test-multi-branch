# Proof of Concept - Kubernetes staging site proxy

A small go test script which would proxy external requests to an internal kubernetes namespace which holds a load balancer service called app. This would be used for staging test sites where we bring up a full cluster of pods per staging branch.

To run use `./setup.sh`. You'll need to have docker and minikube installed locally.

An external staging site URL would be `branch-name-1.staging.test.co.uk` which forwards onto `app.branch-name-1.svc.cluster.local` using kubernetes internal DNS.
