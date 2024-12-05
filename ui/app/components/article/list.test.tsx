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

    render(<Articles feedId="test" pagination={{}} />);
    expect(screen.getByText("loading...")).toBeInTheDocument();
  });

  test("renders error state", () => {
    (useArticles as jest.Mock).mockImplementation(() => ({
      loading: false,
      error: "failed to load articles",
      articles: { articles: [] },
    }));

    render(<Articles feedId="test" pagination={{}} />);
    expect(screen.getByText("failed to load articles")).toBeInTheDocument();
  });

  test("renders no articles found state", () => {
    (useArticles as jest.Mock).mockImplementation(() => ({
      loading: false,
      error: null,
      articles: { articles: [] },
    }));

    render(<Articles feedId="test" pagination={{}} />);
    expect(screen.getByText("no articles found")).toBeInTheDocument();
  });

  test("renders list of articles", () => {
    (useArticles as jest.Mock).mockImplementation(() => ({
      loading: false,
      error: null,
      articles: {
        articles: [{ id: "1", title: "Article 1" }],
        cursor: {},
      },
    }));

    render(<Articles feedId="test" pagination={{}} />);
    expect(screen.getByText("Article 1")).toBeInTheDocument();
    expect(screen.getByRole("link", { name: "Article 1" })).toHaveAttribute(
      "href",
      "/articles/1",
    );
  });

  test("applies 'bg-article-new' when article is not viewed", () => {
    (useArticles as jest.Mock).mockImplementation(() => ({
      loading: false,
      error: null,
      articles: {
        articles: [
          { id: "1", title: "Article 1", userArticle: { viewedAt: null } },
        ],
        cursor: {},
      },
    }));

    render(<Articles feedId="test" pagination={{}} />);
    const articleRow = screen
      .getByRole("link", { name: "Article 1" })
      .closest("li");
    expect(articleRow).toHaveClass("bg-article-new");
  });

  test("applies 'bg-article-read' when article is viewed", () => {
    (useArticles as jest.Mock).mockImplementation(() => ({
      loading: false,
      error: null,
      articles: {
        articles: [
          {
            id: "2",
            title: "Article 2",
            userArticle: { viewedAt: "2024-12-01" },
          },
        ],
        cursor: {},
      },
    }));

    render(<Articles feedId="test" pagination={{}} />);

    const articleRow = screen
      .getByRole("link", { name: "Article 2" })
      .closest("li");
    expect(articleRow).toHaveClass("bg-article-read");
  });

  test("renders both viewed and non-viewed articles with correct backgrounds", () => {
    (useArticles as jest.Mock).mockImplementation(() => ({
      loading: false,
      error: null,
      articles: {
        articles: [
          { id: "1", title: "Article 1", userArticle: { viewedAt: null } },
          {
            id: "2",
            title: "Article 2",
            userArticle: { viewedAt: "2024-12-01" },
          },
        ],
        cursor: {},
      },
    }));

    render(<Articles feedId="test" pagination={{}} />);

    const article1Row = screen
      .getByRole("link", { name: "Article 1" })
      .closest("li");
    const article2Row = screen
      .getByRole("link", { name: "Article 2" })
      .closest("li");

    expect(article1Row).toHaveClass("bg-article-new");
    expect(article2Row).toHaveClass("bg-article-read");
  });
});
