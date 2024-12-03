import { render, screen } from "@testing-library/react";
import "@testing-library/jest-dom";
import useArticle from "@/app/data/article";
import { getGraph } from "@/app/lib/fetch";
import Page from "./page";

jest.mock("@/app/data/article");
jest.mock("@/app/lib/fetch");

describe("Page Component", () => {
  beforeEach(() => {
    jest.clearAllMocks();
  });

  test("renders loading state", () => {
    (useArticle as jest.Mock).mockImplementation(() => ({
      loading: true,
      error: null,
      article: null,
    }));
    render(<Page params={{ id: "test-article" }} />);
    expect(screen.getByText("Loading...")).toBeInTheDocument();
  });

  test("renders error state", () => {
    (useArticle as jest.Mock).mockImplementation(() => ({
      loading: false,
      error: new Error("Something went wrong"),
      article: null,
    }));
    render(<Page params={{ id: "test-article" }} />);
    expect(screen.getByText("An error has occurred.")).toBeInTheDocument();
  });

  test("renders 'Article not found' state", () => {
    (useArticle as jest.Mock).mockImplementation(() => ({
      loading: false,
      error: null,
      article: null,
    }));
    render(<Page params={{ id: "test-article" }} />);
    expect(screen.getByText("Article not found.")).toBeInTheDocument();
  });

  test("renders article and marks it as read", () => {
    const mockMarkArticleRead = jest.fn();
    (getGraph as jest.Mock).mockReturnValue({
      MarkArticleRead: mockMarkArticleRead,
    });

    const article = {
      id: "test-article",
      title: "Test Article",
      userArticle: { viewedAt: null },
    };

    (useArticle as jest.Mock).mockImplementation(() => ({
      loading: false,
      error: null,
      article,
    }));

    render(<Page params={{ id: "test-article" }} />);
    expect(mockMarkArticleRead).toHaveBeenCalledWith({ id: "test-article" });
    expect(screen.getByText("Test Article")).toBeInTheDocument();
  });

  test("does not mark article as read if already viewed", () => {
    const mockMarkArticleRead = jest.fn();
    (getGraph as jest.Mock).mockReturnValue({
      MarkArticleRead: mockMarkArticleRead,
    });

    const article = {
      id: "test-article",
      title: "Test Article",
      userArticle: { viewedAt: "2024-12-01" },
    };

    (useArticle as jest.Mock).mockImplementation(() => ({
      loading: false,
      error: null,
      article,
    }));

    render(<Page params={{ id: "test-article" }} />);
    expect(mockMarkArticleRead).not.toHaveBeenCalled();
  });
});
