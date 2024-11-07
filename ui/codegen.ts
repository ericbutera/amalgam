
import type { CodegenConfig } from '@graphql-codegen/cli';

const config: CodegenConfig = {
  overwrite: true,
  schema: "http://localhost:8082/query",
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
