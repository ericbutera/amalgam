import Link from 'next/link'

export default function Home() {
  return (
    <div>
      <main>
        <h1>home</h1>
        <div>
          <Link href="/feeds">feeds</Link>
        </div>
      </main>
    </div>
  );
}
