import type { Metadata } from 'next';
import './globals.css';

export const metadata: Metadata = {
  title: 'hello-word-D',
  description: 'Stored greeting page',
};

export default function RootLayout({ children }: Readonly<{ children: React.ReactNode }>) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  );
}
