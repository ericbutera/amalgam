
import type { CodegenConfig } from '@graphql-codegen/cli';

const schemaURL = process.env.GRAPH_HOST || 'http://localhost:8082/query';

const config: CodegenConfig = {
  overwrite: true,
  schema: schemaURL,
  documents: [
    "app/graphql/queries.graphql",
    "app/graphql/mutations.graphql",
  ],
  generates: {
    "app/generated/graphql.ts": {
      plugins: [
        'typescript',
        'typescript-operations',
        'typescript-graphql-request'
      ]
    },
    "./app/generated/graphql.schema.json": {
      plugins: ["introspection"]
    }
  }
};

export default config;
