import Link from "next/link";
import queryString from "@/app/lib/queryBuilder";

interface Pagination {
  next?: string;
  previous?: string;
}

const getPagination = (cursor: {
  next?: string;
  previous?: string;
}): Pagination => ({
  next: cursor.next ?? undefined,
  previous: cursor.previous ?? undefined,
});

interface Props {
  base: string;
  cursor: {
    next?: string;
    previous?: string;
  };
}

export default function Pagination({ base, cursor }: Props) {
  const p = getPagination(cursor);

  return (
    <>
      {p.previous && (
        <Link
          href={`${base}?${queryString({
            previous: p.previous,
          })}`}
        >
          Previous
        </Link>
      )}

      {p.next && (
        <Link href={`${base}?${queryString({ next: p.next })}`}>Next</Link>
      )}
    </>
  );
}
