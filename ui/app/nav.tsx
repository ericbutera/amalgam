import Link from 'next/link'

export default function Nav() {
    return (
        <nav>
            <ul>
                <li><Link href="/">home</Link></li>
                <li><Link href="/feeds">feeds</Link></li>
            </ul>
        </nav>
    )
}