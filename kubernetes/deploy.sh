
kubectl apply -f=environment.yml,pv.yml \
              -f=mysql-deployment.yml,redis-deployment.yml,mongodb-deployment.yml \
              -f=security-deployment.yml,general-deployment.yml,chat-deployment.yml \
              -f=frontend-deployment.yml