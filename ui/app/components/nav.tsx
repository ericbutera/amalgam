import Link from "next/link";

export default function Nav() {
  return (
    <div className="navbar bg-base-100">
      <div className="navbar-start">
        <Link href="/" className="btn btn-ghost text-xl">
          home
        </Link>
        <ul className="menu menu-horizontal px-1">
          <li>
            <Link href="/feeds">feeds</Link>
          </li>
        </ul>
      </div>
    </div>
  );
}
