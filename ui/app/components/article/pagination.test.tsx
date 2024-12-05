import { render, screen } from "@testing-library/react";
import Pagination from "@/app/components/article/pagination";
import "@testing-library/jest-dom";
import queryString from "@/app/lib/queryBuilder";

jest.mock("@/app/lib/queryBuilder", () => {
  const originalModule = jest.requireActual("@/app/lib/queryBuilder");
  return {
    __esModule: true,
    ...originalModule,
    default: jest.fn(),
  };
});

describe("Pagination Component", () => {
  beforeEach(() => {
    jest.clearAllMocks();
  });

  test("renders 'Previous' and 'Next' links when both cursors are present", () => {
    (queryString as jest.Mock).mockImplementation((params) =>
      new URLSearchParams(params).toString(),
    );

    render(
      <Pagination
        base="/articles"
        cursor={{ previous: "prev_cursor", next: "next_cursor" }}
      />,
    );

    expect(screen.getByText("Previous")).toBeInTheDocument();
    expect(screen.getByText("Next")).toBeInTheDocument();
    expect(screen.getByText("Previous").closest("a")).toHaveAttribute(
      "href",
      "/articles?previous=prev_cursor",
    );
    expect(screen.getByText("Next").closest("a")).toHaveAttribute(
      "href",
      "/articles?next=next_cursor",
    );
  });

  test("renders only 'Previous' link when only previous cursor is present", () => {
    (queryString as jest.Mock).mockImplementation((params) =>
      new URLSearchParams(params).toString(),
    );

    render(
      <Pagination base="/articles" cursor={{ previous: "prev_cursor" }} />,
    );

    expect(screen.getByText("Previous")).toBeInTheDocument();
    expect(screen.queryByText("Next")).not.toBeInTheDocument();
    expect(screen.getByText("Previous").closest("a")).toHaveAttribute(
      "href",
      "/articles?previous=prev_cursor",
    );
  });

  test("renders only 'Next' link when only next cursor is present", () => {
    (queryString as jest.Mock).mockImplementation((params) =>
      new URLSearchParams(params).toString(),
    );

    render(<Pagination base="/articles" cursor={{ next: "next_cursor" }} />);

    expect(screen.getByText("Next")).toBeInTheDocument();
    expect(screen.queryByText("Previous")).not.toBeInTheDocument();
    expect(screen.getByText("Next").closest("a")).toHaveAttribute(
      "href",
      "/articles?next=next_cursor",
    );
  });

  test("renders nothing when both cursors are absent", () => {
    render(<Pagination base="/articles" cursor={{}} />);

    expect(screen.queryByText("Previous")).not.toBeInTheDocument();
    expect(screen.queryByText("Next")).not.toBeInTheDocument();
  });
});
