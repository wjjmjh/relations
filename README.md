### Local Development with Neo4j database

```
docker pull neo4j:latest

docker run \
  -p 7474:7474 -p 7687:7687 \
  -v $HOME/neo4j/data:/data \
  -v $HOME/neo4j/logs:/logs \
  -v $HOME/neo4j/import:/var/lib/neo4j/import \
  -v $HOME/neo4j/plugins:/plugins \
  --env NEO4J_AUTH=neo4j/neo4j-password \
  neo4j:latest
```

^^ with username: `neo4j` and password: `neo4j-password`
