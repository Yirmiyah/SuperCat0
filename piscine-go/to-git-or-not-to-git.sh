curl -s https://zone01normandie.org/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{login:{_eq:\"Santos\"}}){id}}"}' | tr -dc '0-9'






