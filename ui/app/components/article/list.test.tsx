import { render, screen } from "@testing-library/react";
import Articles from "./list";
import "@testing-library/jest-dom";
import useArticles from "../../data/articles";

jest.mock("../../data/articles");

describe("Articles Component", () => {
  beforeEach(() => {
    jest.clearAllMocks();
  });

  test("renders loading state", () => {
    (useArticles as jest.Mock).mockImplementation(() => ({
      loading: true,
      error: null,
      articles: { articles: [] },
    }));

    render(<Articles feedId="test" />);
    expect(screen.getByText("loading...")).toBeInTheDocument();
  });

  test("renders error state", () => {
    (useArticles as jest.Mock).mockImplementation(() => ({
      loading: false,
      error: "failed to load articles",
      articles: { articles: [] },
    }));

    render(<Articles feedId="test" />);
    expect(screen.getByText("failed to load articles")).toBeInTheDocument();
  });

  test("renders no articles found state", () => {
    (useArticles as jest.Mock).mockImplementation(() => ({
      loading: false,
      error: null,
      articles: { articles: [] },
    }));

    render(<Articles feedId="test" />);
    expect(screen.getByText("no articles found")).toBeInTheDocument();
  });

  test("renders list of articles", () => {
    (useArticles as jest.Mock).mockImplementation(() => ({
      loading: false,
      error: null,
      articles: {
        articles: [{ id: "1", title: "Article 1" }],
        pagination: {},
      },
    }));

    render(<Articles feedId="test" />);
    expect(screen.getByText("Article 1")).toBeInTheDocument();
    expect(screen.getByRole("link", { name: "Article 1" })).toHaveAttribute(
      "href",
      "/articles/1"
    );
  });
});
