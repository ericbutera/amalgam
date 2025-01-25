// Feed GraphQL error response
export interface ErrorResponse {
  response: {
    errors: {
      message: string;
      path: string[];
    }[];
  };
}
