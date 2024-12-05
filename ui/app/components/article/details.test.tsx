import { render, screen } from "@testing-library/react";
import ArticleDetails from "./details";
import "@testing-library/jest-dom";
import useArticle from "../../data/article";

jest.mock("../../data/article");

const mockUseArticle = useArticle as jest.Mock;

describe("ArticleDetails Component", () => {
  beforeEach(() => {
    jest.clearAllMocks();
  });

  test("renders loading state", () => {
    mockUseArticle.mockImplementation(() => ({
      loading: true,
      error: null,
      article: null,
    }));

    render(<ArticleDetails id="1" />);
    expect(screen.getByText("loading...")).toBeInTheDocument();
  });

  test("renders error state", () => {
    mockUseArticle.mockImplementation(() => ({
      loading: false,
      error: "Failed to load",
      article: null,
    }));

    render(<ArticleDetails id="1" />);
    expect(screen.getByText("failed to load article")).toBeInTheDocument();
  });

  test("renders 'Article not found' state", () => {
    mockUseArticle.mockImplementation(() => ({
      loading: false,
      error: null,
      article: null,
    }));

    render(<ArticleDetails id="1" />);
    expect(screen.getByText("article not found")).toBeInTheDocument();
  });

  test("renders article details", () => {
    mockUseArticle.mockImplementation(() => ({
      loading: false,
      error: null,
      article: {
        id: "1",
        title: "Sample Article",
        content: "This is a test article.",
        feedId: "feed1",
      },
    }));

    render(<ArticleDetails id="1" />);
    expect(screen.getByText("Sample Article")).toBeInTheDocument();
    expect(screen.getByText("This is a test article.")).toBeInTheDocument();
    expect(screen.getByRole("link", { name: "Back to Feed" })).toHaveAttribute(
      "href",
      "/feeds/feed1/articles",
    );
  });
});
