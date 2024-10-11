import Link from 'next/link'

export default function Page({ params }: { params: { id: string } }) {
    return (
        <div>
            <h1>viewing feed by id {params.id}!</h1>
            <Link href="/">home</Link>
        </div>
    );
}