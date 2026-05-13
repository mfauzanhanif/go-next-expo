export const metadata = {
  title: 'Aplikasi Web',
  description: 'Frontend Super App',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="id">
      <body>{children}</body>
    </html>
  )
}